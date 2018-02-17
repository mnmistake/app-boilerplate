import React from 'react';
import Auth from '../../utils/Auth';
import history from '../../history';

export default function (ComposedComponent, isRegister = false) {
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
            return <ComposedComponent {...this.props} isRegister={isRegister} />;
        }
    }

    return IsAuthenticated;
}
