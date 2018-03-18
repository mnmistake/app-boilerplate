import React from 'react';
import PropTypes from 'prop-types';
import { graphql } from 'react-apollo';

import { registerUser, loginUser } from '../../graphql/mutations/auth.graphql';
import Auth from '../../utils/Auth';
import history from '../../history';

console.log(registerUser);

@graphql(loginUser, {
    props: ({ mutate }) => ({
        login: ({ username, password }) => mutate({ variables: { username, password } }),
    }),
})
@graphql(registerUser, {
    props: ({ mutate }) => ({
        register: ({ username, password }) => mutate({ variables: { username, password } }),
    }),
})
export default class Authentication extends React.Component {
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

            if (res.data) {
                const { token } = isRegister ? res.data.registerUser : res.data.loginUser;
                Auth.setToken(token);
                history.push('/');
            }
        } catch (error) {
            console.error(error);
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
                {!!errors.length && errors.map((err, idx) => <li key={`${err}__${idx}`}>{err}</li>)}
            </React.Fragment>
        );
    }
}
