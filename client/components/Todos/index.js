import React from 'react';
import { graphql, compose } from 'react-apollo';

import TodoList from './TodoList';
import { todosQuery } from '../../queries/todos';
import { createTodoMutation } from '../../mutations/todos';

const Todos = (props) => <TodoList {...props} />

export default compose(
    graphql(todosQuery),
    graphql(createTodoMutation, { 
        name: 'createTodo',
        options: {            
            update: (store, { data: createTodoMutation }) => {
                const data = store.readQuery({ query: todosQuery });                
                data.todoList.push(createTodoMutation.createTodo);                
                store.writeQuery({ query: todosQuery, data });
            }        
        }
    }),
)(Todos);