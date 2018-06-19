// @flow
import React, { Component } from 'react';
import { graphql } from 'react-apollo';
import classNames from 'classnames';

import type { UserType } from '../../types/User.types';
import type { SegmentsType } from '../../types/Segment.types';
import * as styles from './Sheet.scss';
import { sheetQuery, sheetsQuery } from '../../graphql/queries/Sheets.graphql';
import { deleteSheet } from '../../graphql/mutations/Sheets.graphql';
import userQuery from '../../graphql/queries/User.graphql';
import Spinner from '../Spinner';
import Segment from '../Segment';
import Avatar from '../Avatar';

type Props = {
    data: {
        loading: boolean,
        sheet: {
            name: string,
            segments: SegmentsType,
            user: UserType,
        },
    }
};

@graphql(sheetQuery, {
    options: ({ match: { params: { id } } }) => ({
        variables: { id },
    }),
})
@graphql(userQuery, {
    props: ({ data: { user } }: Object<UserType>) => ({
        user,
    }),
})
@graphql(deleteSheet, {
    props: ({ mutate }) => ({
        deleteSheet: (id: number) => (
            mutate({ variables: { id } })
        ),
    }),
    options: {
        // Optimistically update the sheets to avoid re-fetching
        update: (proxy, { data: { deleteSheet }}) => {
            const data = proxy.readQuery({ query: sheetsQuery });
            const filteredSheets = data.sheets.filter(s => s.id !== deleteSheet.id);
            proxy.writeQuery({
                query: sheetsQuery,
                data: {
                    ...data,
                    sheets: filteredSheets,
                },
            });
        },
    },
})
export default class Sheet extends Component<Props> {
    deleteSheet = id => {
        this.props.deleteSheet(id);
        this.props.history.push('/');
    };

    renderSegments = segments =>
        segments.map(segment => <Segment {...segment} key={segment.id} isCreator={false} />);

    render() {
        const { data: { sheet, loading }, user: currentUser } = this.props;


        if (loading) {
            return <Spinner />;
        }

        return (
            <div className={classNames('container', styles.sheetWrapper)}>
                <div className={styles.header}>
                    <div>
                        <h1>{sheet.name}</h1>
                        {sheet.user.id === currentUser.id &&
                            <button onClick={() => this.deleteSheet(sheet.id)}>Delete</button>
                        }
                    </div>
                    <div>
                        <Avatar username={sheet.user.username} size="50px" />
                    </div>
                </div>
                <div className="segmentsWrapper">
                    {sheet.segments.length ? this.renderSegments(sheet.segments) : <h1>No segments</h1>}
                </div>
            </div>
        );
    }
}
