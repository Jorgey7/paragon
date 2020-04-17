package worker

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/kcarretto/paragon/ent"
	"github.com/kcarretto/paragon/graphql"
	"github.com/kcarretto/paragon/pkg/cdn"
	"github.com/kcarretto/paragon/pkg/event"
	"github.com/kcarretto/paragon/pkg/script"
	libassert "github.com/kcarretto/paragon/pkg/script/stdlib/assert"
	libassets "github.com/kcarretto/paragon/pkg/script/stdlib/assets"
	libcdn "github.com/kcarretto/paragon/pkg/script/stdlib/cdn"
	libcrypto "github.com/kcarretto/paragon/pkg/script/stdlib/crypto"
	libenv "github.com/kcarretto/paragon/pkg/script/stdlib/env"
	libfile "github.com/kcarretto/paragon/pkg/script/stdlib/file"
	libssh "github.com/kcarretto/paragon/pkg/script/stdlib/ssh"
	libsys "github.com/kcarretto/paragon/pkg/script/stdlib/sys"

	"go.starlark.net/starlark"
)

const ServiceTag = "svc-pg-worker"

type Worker struct {
	cdn.Uploader
	cdn.Downloader
	Graph  graphql.Client
	Config string
}

func (w *Worker) HandleTaskQueued(ctx context.Context, info event.TaskQueued) {
	log.Printf("Handling new task queued event")

	target := info.Target
	if target == nil || target.ID == 0 {
		log.Printf("[DBG] task queued event with invalid target")
		return
	}

	task := info.Task
	if task == nil || task.ID == 0 {
		log.Printf("[DBG] task queued event with invalid task")
		return
	}

	found := false
	for _, tag := range info.Tags {
		if tag == nil {
			continue
		}
		if tag.Name != ServiceTag {
			continue
		}

		found = true
		break
	}
	if !found && info.Tags != nil {
		log.Printf("[DBG] task queued event with service tags that were not for the worker")
		return
	}

	if target.PrimaryIP == "" {
		log.Printf("[DBG] task queued event with invalid target ip")
		return
	}

	w.ExecTargetTask(ctx, task, target, info.Credentials)
}

func (w *Worker) ExecTargetTask(ctx context.Context, task *ent.Task, target *ent.Target, credentials []*ent.Credential) {
	_, err := w.Graph.ClaimTask(ctx, task.ID)
	if err != nil {
		log.Printf("[ERR] Failed to claim task: %v", err)
	}

	output := &taskOutput{
		ID:    task.ID,
		Ctx:   ctx,
		Graph: w.Graph,
	}
	output.Start()
	var execErr error
	defer func() {
		output.Stop(execErr)
	}()

	log.Printf("[DBG] Executing new task (%d) on %s (%d credentials)",
		task.ID,
		target.PrimaryIP,
		len(credentials),
	)

	env := libenv.Environment{
		PrimaryIP:       target.PrimaryIP,
		OperatingSystem: target.OS.String(),
	}

	/* Build Assets Bundle */
	assets := libassets.Environment{
		Downloader: w,
	}
	code := script.New(
		fmt.Sprintf("%d", task.ID),
		strings.NewReader(task.Content),
		script.WithOutput(output),
		libfile.Include(),
		libassert.Include(),
		libsys.Include(), // TODO: Deprecate after multi-file is finished
		assets.Include(),
		env.Include(),
	)
	if _, err := code.Call("init", starlark.Tuple{}); err != nil {
		execErr = fmt.Errorf("failed to initialize assets: %w", err)
		return
	}

	bundle := libassets.TarGZBundler{}
	if err := bundle.Bundle(append(
		assets.Files,
		libassets.NamedReader{
			Name:   "scripts/main.rg",
			Reader: strings.NewReader(task.Content),
		},
	)...); err != nil {
		execErr = fmt.Errorf("failed to bundle assets: %w", err)
		return
	}

	/* Use Config to execute task */
	sshConnector := &SSHConnector{
		Credentials: credentials,
	}
	defer sshConnector.Close()

	sshEnv := &libssh.Environment{
		RemoteHost: target.PrimaryIP,
		Connector:  sshConnector,
	}
	defer sshEnv.Close()

	cdnEnv := &libcdn.Environment{
		Uploader:   w,
		Downloader: w,
	}

	var configScript = strings.NewReader(DefaultConfig)
	if w.Config != "" {
		configScript = strings.NewReader(w.Config)
	}

	config := script.New(
		"worker_config.rg",
		configScript,
		script.WithOutput(output),
		libfile.Include(),
		libassert.Include(),
		libcrypto.Include(),
		env.Include(),
		cdnEnv.Include(),
		sshEnv.Include(),
	)

	var res starlark.Value = starlark.None
	defer func() {
		if _, err := config.Call("worker_exit", starlark.Tuple{res}); err != nil {
			log.Printf("[ERR] worker_exit failed: %v", err)
		}
	}()

	res, execErr = config.Call("worker_run", starlark.Tuple{starlark.String(bundle.Buffer.Bytes())})
}
