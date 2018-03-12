import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';
import { Link } from 'react-router-dom';
import { graphql } from 'react-apollo';

import * as styles from './Sheets.scss';
import { sheetsQuery } from '../../graphql/queries/sheets';
import { createTodoMutation, updateTodoMutation } from '../../graphql/mutations/todos';

@graphql(createTodoMutation, {
    props: ({ mutate }) => ({
        createTodo: content => mutate({ variables: { content } }),
    }),
    options: {
        update: (store, { data }) => {
            const storeData = store.readQuery({ query: sheetsQuery });
            storeData.todoList.push(data.createTodo);
            store.writeQuery({ query: sheetsQuery, data: storeData });
        },
    },
})
@graphql(updateTodoMutation, {
    props: ({ mutate }) => ({
        updateTodo: (id, isCompleted) => mutate({ variables: { id, isCompleted } }),
    }),
})
@graphql(sheetsQuery)
export default class Sheets extends React.Component {
    static propTypes = {
        data: PropTypes.shape({
            sheets: PropTypes.array,
            loading: PropTypes.bool.isRequired,
        }).isRequired,
    };

    render() {
        const { sheets, loading } = this.props.data;

        if (loading) return 'loading';

        const LastSheet = () => (
            <Link to={`/sheet/create`} className="gridItem">
                Create your sheet...
            </Link>
        );

        const Sheet = ({ id, name, createdAt, user: { username }, isLastSheet }) => (
            <React.Fragment>
                <Link to={`/sheet/${id}`} className="gridItem">
                    <h1>{name}</h1>
                    <p>{username}</p>
                    Created {moment(createdAt).fromNow()}
                </Link>
                {isLastSheet && <LastSheet />}
            </React.Fragment>
        );


        return (
            <div className="gridWrapper">
                {sheets && sheets.map((sheet, idx) => {
                    const isLastSheet = sheets.length - 1 === idx;

                    return (
                        <Sheet key={sheet.id} {...sheet} isLastSheet={isLastSheet} />
                    );
                })}
            </div>
        );
    }
}
