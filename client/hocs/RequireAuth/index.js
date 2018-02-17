import React from 'react';
import Auth from '../../utils/Auth';
import history from '../../history';

export default function (ComposedComponent) {
    class RequireAuth extends React.Component {
        componentWillMount() {
            if (!Auth.isUserAuthenticated()) {
                history.push('/login');
            }
        }

        componentWillUpdate() {
            if (!Auth.isUserAuthenticated()) {
                history.push('/login');
            }
        }

        render() {
            return <ComposedComponent {...this.props} />;
        }
    }

    return RequireAuth;
}
