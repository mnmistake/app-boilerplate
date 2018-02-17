import React from 'react';
import PropTypes from 'prop-types';
import { graphql } from 'react-apollo';

import todosQuery from '../../queries/todos';
import createTodoMutation from '../../mutations/todos';

@graphql(createTodoMutation, {
    name: 'createTodo',
    options: {
        update: (store, { data }) => {
            const storeData = store.readQuery({ query: todosQuery });
            storeData.todoList.push(data.createTodo);
            store.writeQuery({ query: todosQuery, data: storeData });
        },
    },
})
@graphql(todosQuery)
export default class Todos extends React.Component {
    static propTypes = {
        createTodo: PropTypes.func.isRequired,
        data: PropTypes.shape({
            todoList: PropTypes.array,
            loading: PropTypes.bool.isRequired,
        }).isRequired,
    };

    state = {
        content: '',
    };

    createTodo() {
        if (this.state.content) {
            this.props.createTodo({
                variables: {
                    content: this.state.content,
                },
            });
        }
    }

    render() {
        const { todoList: todos, loading } = this.props.data;
        if (loading) {
            return 'loading';
        }

        return (
            <ul>
                {todos && todos.map(todo => <li key={todo.id}>{todo.content}</li>)}
                <input type="text" onChange={e => this.setState({ content: e.target.value })} />
                <button onClick={() => this.createTodo()}>create todo</button>
            </ul>
        );
    }
}
