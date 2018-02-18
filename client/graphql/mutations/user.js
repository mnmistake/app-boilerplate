import gql from 'graphql-tag';

export default gql`
    mutation setUser($username: String!, $id: Int!) {
        setUser(username: $username, id: $id) @client {
            username
            id
        }
    }
`;
