import React from 'react';
import history from '../../history';

import Auth from '../../utils/Auth';

export default function (ComposedComponent) {
    class RequireAuth extends React.Component {
        componentWillMount() {
            if (!Auth.doesTokenExist()) {
                history.push('/login');
            }
        }

        componentWillUpdate() {
            if (!Auth.doesTokenExist()) {
                history.push('/login');
            }
        }

        render() {
            return <ComposedComponent {...this.props} />;
        }
    }

    return RequireAuth;
}
