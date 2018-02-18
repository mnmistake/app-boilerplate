import React from 'react';
import Auth from '../../utils/Auth';
import history from '../../history';

export default function (ComposedComponent, isRegister = false) {
    class PersistLogin extends React.Component {
        componentWillMount() {
            if (Auth.doesTokenExist()) {
                history.push('/');
            }
        }

        componentWillUpdate() {
            if (Auth.doesTokenExist()) {
                history.push('/');
            }
        }

        render() {
            return <ComposedComponent {...this.props} isRegister={isRegister} />;
        }
    }

    return PersistLogin;
}
