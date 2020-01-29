import { useQuery } from "@apollo/react-hooks";
import gql from "graphql-tag";
import * as React from "react";
import { Card, Container, Loader, Menu } from "semantic-ui-react";
import { XFileCard, XFileUploadModal, XNoFilesFound } from "../components/file";
import { XErrorMessage } from "../components/messages";
import { File } from "../graphql/models";

export const MULTI_FILE_QUERY = gql`
  {
    files {
      id
      name
      contentType
      size
      creationTime
      lastModifiedTime

      links {
        id
        alias
        clicks
        expirationTime
      }
    }
  }
`;

type MultiFile = {
  files: File[];
};

const XMultiFileView = () => {
  const { called, loading, error, data } = useQuery<MultiFile>(
    MULTI_FILE_QUERY,
    {
      pollInterval: 5000
    }
  );

  const showCards = () => {
    if (!data || !data.files || data.files.length < 1) {
      return <XNoFilesFound />;
    }
    return (
      <Card.Group centered itemsPerRow={4}>
        {data.files.map(file => (
          <XFileCard key={file.id} {...file} />
        ))}
      </Card.Group>
    );
  };

  return (
    <Container style={{ padding: "10px" }}>
      <Menu secondary>
        <Menu.Item position="right">
          <XFileUploadModal
            button={{ positive: true, circular: true, icon: "plus" }}
          />
        </Menu.Item>
      </Menu>
      <Container fluid style={{ padding: "20px" }}>
        <Loader disabled={!called || !loading} />
        <XErrorMessage title="Error Loading Files" err={error} />
        {showCards()}
      </Container>
    </Container>
  );
};

// XTargetCardGroup.propTypes = {
//     targets: PropTypes.arrayOf(PropTypes.shape({
//         id: PropTypes.number.isRequired,
//         name: PropTypes.string.isRequired,
//         tags: PropTypes.arrayOf(PropTypes.string),
//     })).isRequired,
// };

export default XMultiFileView;