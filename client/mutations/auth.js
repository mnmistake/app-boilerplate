import gql from 'graphql-tag';

const registerMutation = gql`
    mutation registerUser($username: String!, $password: String!) {
        registerUser(username: $username, password: $password) {
            token
        }
    }
`;

const loginMutation = gql`
    mutation loginUser($username: String!, $password: String!) {
        loginUser(username: $username, password: $password) {
            token
        }
    }
`;

export {
    registerMutation,
    loginMutation,
};
