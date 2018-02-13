import gql from 'graphql-tag';

const registerMutation = gql`
    mutation registerUser($username: String!, $password: String!) {
        registerUser(username: $username, password: $password) {
            token
        }
    }
`;

export default registerMutation;
