import * as React from "react";
import { Card } from "semantic-ui-react";
import { Job } from "../../graphql/models";
import { XTags } from "../tag";
import { XTaskSummary } from "../task";

const XJobCard = (j: Job) => (
  <Card fluid>
    <Card.Content>
      <Card.Header href={"/jobs/" + j.id}>{j.name}</Card.Header>

      <XTaskSummary tasks={j.tasks} />
    </Card.Content>
    <Card.Content extra>
      <XTags tags={j.tags} defaultText="None" />
    </Card.Content>
  </Card>
);

export default XJobCard;
