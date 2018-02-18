import gql from 'graphql-tag';

const todoListQuery = gql`
    query {
        todoList {
            id
            content
            isCompleted
        }
    }
`;

const todoQuery = gql`
    query todoQuery($id: Int!) {
        todo(id: $id) {
            id
            content
            isCompleted
        }
    }
`;

export {
    todoListQuery,
    todoQuery,
};
