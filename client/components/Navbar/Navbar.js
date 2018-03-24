// @flow
import React, { PureComponent } from 'react';
import { graphql } from 'react-apollo';

import * as styles from './Navbar.scss';
import type { UserType } from '../../types/User.types';
import userQuery from '../../graphql/queries/User.graphql';
import setUserMutation from '../../graphql/mutations/User.graphql';
import Auth from '../../utils/Auth';
import history from '../../history';

type Props = {
    user: UserType,
    clearUser: () => void,
};

@graphql(userQuery, {
    props: ({ data: { user } }: Object<UserType>) => ({
        user,
    }),
})
@graphql(setUserMutation, {
    props: ({ mutate }) => ({
        clearUser: () => mutate({ variables: { username: null, id: null } }),
    }),
})
export default class Navbar extends PureComponent<Props> {
    logout = () => {
        this.props.clearUser();
        Auth.removeToken();
        history.push('/login');
    };

    render() {
        const { user } = this.props;
        return (
            <div className={styles.navbar}>
                <div className="container">
                    <h1>{Auth.doesTokenExist() && user && user.username}</h1>
                    {Auth.doesTokenExist() && <button onClick={this.logout}>Log out</button>}
                </div>
            </div>
        );
    }
}
