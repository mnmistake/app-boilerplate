import React from 'react';
import { graphql } from 'react-apollo';
import classNames from 'classnames';

import * as styles from './Sheet.scss';
import Spinner from '../Spinner';
import Segment from '../Segment';
import { sheetQuery } from '../../graphql/queries/Sheets.graphql';

const Sheet = ({ data: { sheet, loading } }) => {
    if (loading) {
        return <Spinner />;
    }

    const { segments } = sheet;
    console.log(segments);

    return (
        <div className={classNames('container', styles.sheetWrapper)}>
            <h1>{sheet.name}</h1>
            <div className="segmentsWrapper">
                {segments && segments.map(segment => <Segment {...segment} key={segment.id} isCreator={false} />)}
            </div>
        </div>
    );
};

export default graphql(sheetQuery, {
    options: ({ match: { params: { id } } }) => ({
        variables: { id },
    }),
})(Sheet);
