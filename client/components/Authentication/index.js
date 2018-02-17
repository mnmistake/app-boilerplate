import React from 'react';
import PropTypes from 'prop-types';
import { graphql } from 'react-apollo';

import Auth from '../../utils/Auth';
import { loginMutation, registerMutation } from '../../graphql/mutations/auth';
import history from '../../history';

@graphql(loginMutation, {
    props: ({ mutate }) => ({
        login: ({ username, password }) => mutate({ variables: { username, password } }),
    }),
})
@graphql(registerMutation, {
    props: ({ mutate }) => ({
        register: ({ username, password }) => mutate({ variables: { username, password } }),
    }),
})
export default class Login extends React.Component {
    static propTypes = {
        login: PropTypes.func.isRequired,
        register: PropTypes.func.isRequired,
        isRegister: PropTypes.bool.isRequired,
    };

    state = {
        username: '',
        password: '',
        errors: [],
    };

    handleSubmit = async (e) => {
        e.preventDefault();
        const { username, password } = this.state;
        const { login, register, isRegister } = this.props;

        try {
            const res = isRegister ? await register({ username, password }) : await login({ username, password });

            if (res.data.loginUser) {
                const { token } = res.data.loginUser;
                Auth.setToken(token);
                history.push('/');
            }
        } catch (error) {
            const errors = error.graphQLErrors.map(err => err.message);
            this.setState({ errors });
        }
    };

    render() {
        const { errors } = this.state;
        const { isRegister } = this.props;

        return (
            <React.Fragment>
                <form onSubmit={e => this.handleSubmit(e)}>
                    <input type="text" onChange={e => this.setState({ username: e.target.value })} />
                    <input type="password" onChange={e => this.setState({ password: e.target.value })} />
                    <button>{isRegister ? 'Register' : 'Login'}</button>
                </form>
                {!!errors.length && errors.map(err => <li key={err}>{err}</li>)}
            </React.Fragment>
        );
    }
}