import gql from 'graphql-tag';

const createTodoMutation = gql`
    mutation createTodo($content: String!) {
        createTodo(content: $content) {
            id,
            content,
            isCompleted,
        }
    }
`;

const updateTodoMutation = gql`
    mutation updateTodo($id: Int!, $isCompleted: Boolean!) {
        updateTodo(id: $id, isCompleted: $isCompleted) {
            id,
            content,
            isCompleted,
        }
    }
`;

export {
    createTodoMutation,
    updateTodoMutation,
};
