import React from 'react';
import { graphql } from 'react-apollo';
import gql from 'graphql-tag';

const TodoList = ({ data }) => {
    const { todoList } = data;
    return (
        <ul>
            {todoList && todoList.map(todo =>
                <li key={todo.id}>{todo.content}</li>
            )}
        </ul>
    )
}

export default graphql(gql`
    query {
        todoList {
            id,
            content,
            isCompleted,
        }
    }
`)(TodoList)