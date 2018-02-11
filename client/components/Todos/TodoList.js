import React from 'react';
import PropTypes from 'prop-types';

export default class TodoList extends React.Component {
    state = {
        content: '',
    }

    static propTypes = {
        createTodo: PropTypes.func.isRequired,
        data: PropTypes.shape(PropTypes.shape({
            todoList: PropTypes.shape({}).isRequired,
            loading: PropTypes.bool.isRequired,
        })).isRequired,
    }

    createTodo() {
        if (this.state.content) {
            this.props.createTodo({
                variables: {
                    content: this.state.content,
                }
            })
        }
    }

    render() {    
        const { createTodo } = this.props;
        const { todoList: todos, loading } = this.props.data;
        if (loading) {
            return 'loading'
        }

        return (
            <ul>
                {todos && todos.map(todo =>
                    <li key={todo.id}>{todo.content}</li>
                )}
                <input type="text" onChange={e => this.setState({ content: e.target.value })} />
                <button onClick={() => this.createTodo()}>create todo</button>
            </ul>
        )
    }
}