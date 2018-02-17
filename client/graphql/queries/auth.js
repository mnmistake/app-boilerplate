import gql from 'graphql-tag';

const getUserStatusQuery = gql`
    query {
        getUserStatus {
            isAuthorized
        }
    }
`;

export default getUserStatusQuery;
