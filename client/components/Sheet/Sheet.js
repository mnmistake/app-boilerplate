import React from 'react';
import { graphql } from 'react-apollo';

import Segment from '../Segment';
import { sheetQuery } from '../../graphql/queries/sheets';

const Sheet = ({ data: { sheet, loading } }) => {
    if (loading) {
        return 'loading...';
    }

    const { segments } = sheet;

    return (
        <React.Fragment>
            <h1>{sheet.name}</h1>
            <div className="gridWrapper">
                {segments && segments.map(segment => <Segment {...segment} key={segment.id} />)}
            </div>
        </React.Fragment>
    );
};

export default graphql(sheetQuery, {
    options: ({ match: { params: { id } } }) => ({
        variables: { id },
    }),
})(Sheet);
