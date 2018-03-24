// @flow
import React, { Component, Fragment } from 'react';
import { graphql } from 'react-apollo';

import { registerUser, loginUser } from '../../graphql/mutations/Authentication.graphql';
import Auth from '../../utils/Auth';
import history from '../../history';

type User = {
    username: string,
    password: string,
};

type Props = {
    login: User => Object,
    register: User => Object,
    isRegister: boolean,
};

type State = {
  ...$Exact<User>,
  errors: Array<string>,
};

@graphql(loginUser, {
    props: ({ mutate }) => ({
        login: ({ username, password }: User) => mutate({ variables: { username, password } }),
    }),
})
@graphql(registerUser, {
    props: ({ mutate }) => ({
        register: ({ username, password }: User) => mutate({ variables: { username, password } }),
    }),
})
export default class Authentication extends Component<Props, State> {
    state = {
        username: '',
        password: '',
        errors: [],
    };

    handleSubmit = async (e: Event) => {
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
            <Fragment>
                <form onSubmit={e => this.handleSubmit(e)}>
                    <input type="text" onChange={e => this.setState({ username: e.target.value })} />
                    <input type="password" onChange={e => this.setState({ password: e.target.value })} />
                    <button>{isRegister ? 'Register' : 'Login'}</button>
                </form>
                {!!errors.length && errors.map((err, idx) => <li key={`${err}__${idx}`}>{err}</li>)}
            </Fragment>
        );
    }
}
