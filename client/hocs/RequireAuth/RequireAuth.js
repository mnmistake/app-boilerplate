import React, { Component } from 'react';
import { graphql } from 'react-apollo';
import PropTypes from 'prop-types';
import jwtDecode from 'jwt-decode';

import setUserMutation from 'graphql/mutations/User.graphql';
import Auth from 'utils/Auth';

export default function (ComposedComponent) {
    @graphql(setUserMutation, {
        props: ({ mutate }) => ({
            setUser: ({ username, id }) => mutate({ variables: { username, id } }),
        }),
    })
    class RequireAuth extends Component {
        static propTypes = {
            setUser: PropTypes.func.isRequired,
        };

        componentDidMount() {
            if (!Auth.doesTokenExist()) {
                this.props.history.push('/login');
            }
            this.setUser();
        }

        componentDidUpdate() {
            if (!Auth.doesTokenExist()) {
                this.props.history.push('/login');
            }
            this.setUser();
        }

        setUser() {
            try {
                const decoded = jwtDecode(Auth.getToken());
                const { setUser } = this.props;

                if (decoded) {
                    const { username, id } = decoded;
                    setUser({ username, id });
                }
            } catch (err) {
                // invalid token error, abort and send back to login
                Auth.removeToken();
                this.props.history.push('/login');
            }
        }

        render() {
            return <ComposedComponent {...this.props} />;
        }
    }

    return RequireAuth;
}
