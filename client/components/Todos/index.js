import React from 'react';
import PropTypes from 'prop-types';
import { graphql } from 'react-apollo';

import history from '../../history';
import todoListQuery from '../../graphql/queries/todos';
import { createTodoMutation, updateTodoMutation } from '../../graphql/mutations/todos';

@graphql(createTodoMutation, {
    props: ({ mutate }) => ({
        createTodo: content => mutate({ variables: { content } }),
    }),
    options: {
        update: (store, { data }) => {
            const storeData = store.readQuery({ query: todoListQuery });
            storeData.todoList.push(data.createTodo);
            store.writeQuery({ query: todoListQuery, data: storeData });
        },
    },
})
@graphql(updateTodoMutation, {
    props: ({ mutate }) => ({
        updateTodo: (id, isCompleted) => mutate({ variables: { id, isCompleted } }),
    }),
})
@graphql(todoListQuery)
export default class Todos extends React.Component {
    static propTypes = {
        createTodo: PropTypes.func.isRequired,
        updateTodo: PropTypes.func.isRequired,
        data: PropTypes.shape({
            todoList: PropTypes.array,
            loading: PropTypes.bool.isRequired,
        }).isRequired,
    };

    state = {
        content: '',
    };

    handleClick = () => {
        const { content } = this.state;
        const { createTodo } = this.props;

        if (content) createTodo(content);
    };

    logout = () => {
        localStorage.removeItem('token');
        history.push('/login');
    };

    render() {
        const { todoList, loading } = this.props.data;
        const { updateTodo } = this.props;

        const renderTodo = todo => (
            <div key={todo.id} className="todos">
                <li>{todo.content}</li>
                <input
                    type="checkbox"
                    checked={todo.isCompleted}
                    onChange={() => updateTodo(todo.id, !todo.isCompleted)}
                />
            </div>
        );

        if (loading) {
            return 'loading';
        }

        return (
            <ul>
                {todoList && todoList.map(todo => renderTodo(todo))}
                <input type="text" onChange={e => this.setState({ content: e.target.value })} />
                <button onClick={this.handleClick}>create todo</button>
                <button onClick={this.logout}>logout</button>
            </ul>
        );
    }
}
