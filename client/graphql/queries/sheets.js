import gql from 'graphql-tag';

const sheetsQuery = gql`
    query {
        sheets {
            id
            name    
            createdAt
            user {
                id
                username
            }
        }
    }
`;

const sheetQuery = gql`
    query sheetQuery($id: Int!) {
        sheet(id: $id) {
            name
            segments {
                id
                label
                content
                createdAt
            }
        }
    }
`;

export {
    sheetsQuery,
    sheetQuery,
};
