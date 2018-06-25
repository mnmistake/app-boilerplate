// @flow
import React, { Component } from 'react';
import { graphql } from 'react-apollo';
import classNames from 'classnames';
import { Link } from 'react-router-dom'

import * as styles from 'components/Authentication/Authentication.scss';
import Field from 'components/Field';
import { registerUser, loginUser } from 'graphql/mutations/Authentication.graphql';
import Auth from 'utils/Auth';

type User = {
    username: string,
    password: string,
};

type Props = {
    login: User => Object,
    register: User => Object,
    isRegister: boolean,
};

type State = User & {
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
        const { login, register, isRegister, history } = this.props;

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
        const pageLabel = isRegister ? 'Register' : 'Login';

        return (
            <div className={classNames('container', styles.wrapper)}>
                <h1>{pageLabel}</h1>
                <form onSubmit={e => this.handleSubmit(e)} className={styles.form}>
                    <Field
                        required
                        type="text"
                        placeholder="Username"
                        onChange={e => this.setState({ username: e.target.value })}
                    />
                    <Field
                        required
                        type="password"
                        placeholder="Password"
                        onChange={e => this.setState({ password: e.target.value })}
                    />
                    <button className="primary">{pageLabel}</button>
                    <Link to={isRegister ? 'login' : 'register'} className="button default">
                        <span>{isRegister ? 'Login' : 'Register'}</span>
                    </Link>
                </form>
                {!!errors.length && errors.map((err, idx) => <li key={`${err}__${idx}`}>{err}</li>)}
            </div>
        );
    }
}
