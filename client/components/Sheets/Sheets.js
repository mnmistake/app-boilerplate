// @flow
import React, { Fragment } from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';
import { Link } from 'react-router-dom';
import { graphql } from 'react-apollo';
import classNames from 'classnames';

import * as styles from './Sheets.scss';
import { sheetsQuery } from '../../graphql/queries/Sheets.graphql';

import Avatar from '../Avatar';
import Spinner from '../Spinner';

type Props = {
  data: {
      sheets: Array<Object>,
      loading: boolean,
  }
};

@graphql(sheetsQuery)
export default class Sheets extends React.Component<Props> {
    static propTypes = {
        data: PropTypes.shape({
            sheets: PropTypes.array,
            loading: PropTypes.bool.isRequired,
        }).isRequired,
    };

    render() {
        const { sheets, loading } = this.props.data;

        if (loading) {
            return <Spinner />;
        }
        
        return (
            <div className={classNames('container', styles.sheetsWrapper)}>
                {sheets && sheets.map((sheet, idx) => {
                    const isLastSheet = sheets.length - 1 === idx;
                    const { id, name, createdAt, user: { username } } = sheet;

                    return (
                        <Fragment>
                            <Link to={`/sheet/${id}`} className={styles.sheet} title={name}>
                                <div className={styles.sheetDetails}>
                                    <h1>{name}</h1>
                                    <p className="note">Created {moment(createdAt).fromNow()}</p>
                                </div>
                                <Avatar username={username} />
                            </Link>
                            {isLastSheet &&
                                <Link to={`/create`} className={styles.sheet}>
                                    Create your sheet...
                                </Link>
                            }
                        </Fragment>
                    )
                })}
            </div>
        );
    }
}
