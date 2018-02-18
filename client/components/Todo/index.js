import React from 'react';
import { graphql } from 'react-apollo';

import { todoQuery } from '../../graphql/queries/todos';

const Todo = ({ data }) => {
    const { loading, todo } = data;

    if (loading) {
        return 'loading...';
    }

    return todo && <h1>{todo.content} - {todo.id}</h1>;
};

export default graphql(todoQuery, {
    options: ({ match }) => ({
        variables: {
            id: match.params.id,
        },
    }),
})(Todo);
