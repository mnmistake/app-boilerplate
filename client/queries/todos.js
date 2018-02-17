import gql from 'graphql-tag';

const todoListQuery = gql`
    query {
        todoList {
            id,
            content,
            isCompleted,
        }
    }
`;

export default todoListQuery;
