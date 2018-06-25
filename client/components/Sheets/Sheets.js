// @flow
import React, { Fragment } from 'react';
import moment from 'moment';
import { Link } from 'react-router-dom';
import { graphql } from 'react-apollo';

import * as styles from 'components/Sheets/Sheets.scss';
import { sheetsQuery } from 'graphql/queries/Sheets.graphql';
import Avatar from 'components/Avatar';
import Spinner from 'components/Spinner';

type Props = {
  data: {
      sheets: Array<Object>,
      loading: boolean,
  }
};

const Sheets = ({ data: { sheets, loading } }: Props) => {
    if (loading) {
        return <Spinner />;
    }

    return (
        <div className={styles.sheetsWrapper}>
            {sheets && sheets.map(sheet => {
                const { id, name, createdAt, user: { username } } = sheet;
                const createdAtTimestamp = moment(moment(createdAt).toDate(), "YYYY-MM-DD HH:mm:ss").fromNow();

                return (
                    <Fragment key={id}>
                        <Link to={`/sheet/${id}`} className={styles.sheet} title={name}>
                            <div className={styles.sheetDetails}>
                                <h1>{name}</h1>
                                <p className="note">Created {createdAtTimestamp}</p>
                            </div>
                            <Avatar username={username} />
                        </Link>
                    </Fragment>
                )
            })}
        </div>
    );
};

export default graphql(sheetsQuery)(Sheets);
