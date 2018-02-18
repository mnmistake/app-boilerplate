import React from 'react';
import PropTypes from 'prop-types';
import { graphql } from 'react-apollo';

import userQuery from '../../graphql/queries/user';
import setUserMutation from '../../graphql/mutations/user';
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
            <React.Fragment>
                <h1>{user && user.username}</h1>
                <button onClick={this.logout}>Log out</button>
            </React.Fragment>
        );
    }
}
