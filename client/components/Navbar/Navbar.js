// @flow
import React, { PureComponent } from 'react';
import { graphql, withApollo } from 'react-apollo';
import { Link } from 'react-router-dom';

import * as styles from 'components/Navbar/Navbar.scss';
import type { UserType } from 'types/User.types';
import userQuery from 'graphql/queries/User.graphql';
import Auth from 'utils/Auth';

type Props = {
    user: UserType,
    client: Object,
};

@graphql(userQuery, {
    props: ({ data: { user } }: Object<UserType>) => ({
        user,
    }),
})
@withApollo
export default class Navbar extends PureComponent<Props> {
    logout = () => {
        this.props.client.resetStore();
        Auth.removeToken();
        this.props.history.push('/login');
    };

    render() {
        const { user } = this.props;
        return (
            <div className={styles.navbar}>
                <div className="container">
                    <h1>{Auth.doesTokenExist() && user && user.username}</h1>
                    {Auth.doesTokenExist() &&
                        <div className={styles.actions}>
                            <Link to="/create" className="button primary">
                                <span>New sheet</span>
                            </Link>
                            <button className="button default" onClick={this.logout}>
                                Log out
                            </button>
                        </div>
                    }
                </div>
            </div>
        );
    }
}
