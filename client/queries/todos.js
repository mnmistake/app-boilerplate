import gql from 'graphql-tag';

const todosQuery = gql`
query {
    todoList {
        id,
        content,
        isCompleted,
    }
}
`;

export {
    todosQuery,
}