import React, { Component } from 'react';
import Auth from '../../utils/Auth';

export default function (ComposedComponent, isRegister = false) {
    class PersistLogin extends Component {
        componentDidMount() {
            if (Auth.doesTokenExist()) {
                this.props.history.push('/');
            }
        }

        componentDidUpdate() {
            if (Auth.doesTokenExist()) {
                this.props.history.push('/');
            }
        }

        render() {
            return <ComposedComponent {...this.props} isRegister={isRegister} />;
        }
    }

    return PersistLogin;
}
