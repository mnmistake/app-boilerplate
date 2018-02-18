import gql from 'graphql-tag';

export default gql`
    query {
        user @client {
            username
            id
        }
    }
`;
