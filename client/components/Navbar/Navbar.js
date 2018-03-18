import React from 'react';
import PropTypes from 'prop-types';
import { graphql } from 'react-apollo';

import * as styles from './Navbar.scss';
import userQuery from '../../graphql/queries/User.graphql';
import setUserMutation from '../../graphql/mutations/User.graphql';
import Auth from '../../utils/Auth';
import history from '../../history';

@graphql(userQuery, {
    props: ({ data: { user } }) => ({
        user,
    }),
})
@graphql(setUserMutation, {
    props: ({ mutate }) => ({
        clearUser: () => mutate({ variables: { username: null, id: null } }),
    }),
})
export default class Navbar extends React.Component {
    static propTypes = {
        user: PropTypes.shape({}).isRequired,
        clearUser: PropTypes.func.isRequired,
    };

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
                    <h1>{user && user.username}</h1>
                    <button onClick={this.logout}>Log out</button>
                </div>
            </div>
        );
    }
}
