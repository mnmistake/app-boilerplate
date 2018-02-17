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

export default createTodoMutation;
