import React from 'react';
import { graphql, compose } from 'react-apollo';

import TodoList from './TodoList';
import todosQuery from '../../queries/todos';
import createTodoMutation from '../../mutations/todos';

const Todos = props => <TodoList {...props} />;

export default compose(
    graphql(todosQuery),
    graphql(createTodoMutation, {
        name: 'createTodo',
        options: {
            update: (store, { data }) => {
                const storeData = store.readQuery({ query: todosQuery });
                storeData.todoList.push(data.createTodo);
                store.writeQuery({ query: todosQuery, storeData });
            },
        },
    }),
)(Todos);
