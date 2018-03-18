import React from 'react';
import { graphql } from 'react-apollo';
import PropTypes from 'prop-types';
import jwtDecode from 'jwt-decode';

import setUserMutation from '../../graphql/mutations/User.graphql';
import Auth from '../../utils/Auth';
import history from '../../history';

export default function (ComposedComponent) {
    @graphql(setUserMutation, {
        props: ({ mutate }) => ({
            setUser: ({ username, id }) => mutate({ variables: { username, id } }),
        }),
    })
    class RequireAuth extends React.Component {
        static propTypes = {
            setUser: PropTypes.func.isRequired,
        };

        componentWillMount() {
            if (!Auth.doesTokenExist()) {
                history.push('/login');
            }
            this.setUser();
        }

        componentWillUpdate() {
            if (!Auth.doesTokenExist()) {
                history.push('/login');
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
                history.push('/login');
            }
        }

        render() {
            return <ComposedComponent {...this.props} />;
        }
    }

    return RequireAuth;
}
