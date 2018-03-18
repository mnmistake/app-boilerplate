import React from 'react';
import { graphql } from 'react-apollo';
import classNames from 'classnames';

import * as styles from './Sheet.scss';
import { sheetQuery } from '../../graphql/queries/Sheets.graphql';

import Spinner from '../Spinner';
import Segment from '../Segment';
import Avatar from '../Avatar';

const Sheet = ({ data: { sheet, loading } }) => {
    if (loading) {
        return <Spinner />;
    }

    const { segments, user: { username } } = sheet;

    const renderSegments = () =>
        segments.map(segment => <Segment {...segment} key={segment.id} isCreator={false} />);

    return (
        <div className={classNames('container', styles.sheetWrapper)}>
            <div className={styles.header}>
                <div>
                    <h1>{sheet.name}</h1>
                </div>
                <div>
                    <Avatar username={username} size="50px" />
                </div>
            </div>
            <div className="segmentsWrapper">
                {segments.length ? renderSegments() : <h1>No segments</h1>}
            </div>
        </div>
    );
};

export default graphql(sheetQuery, {
    options: ({ match: { params: { id } } }) => ({
        variables: { id },
    }),
})(Sheet);
