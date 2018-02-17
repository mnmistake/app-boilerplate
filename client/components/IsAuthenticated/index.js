import React from 'react';
import Auth from '../../middleware/Auth';
import history from '../../history';

export default function (ComposedComponent) {
    class IsAuthenticated extends React.Component {
        componentWillMount() {
            if (Auth.isUserAuthenticated()) {
                history.push('/');
            }
        }

        componentWillUpdate() {
            if (Auth.isUserAuthenticated()) {
                history.push('/');
            }
        }

        render() {
            return <ComposedComponent {...this.props} />;
        }
    }

    return IsAuthenticated;
}
