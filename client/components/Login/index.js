import React from 'react';
import { graphql } from 'react-apollo';

import Auth from '../../middleware/Auth';
import { loginMutation } from '../../mutations/auth';
import history from '../../history';

@graphql(loginMutation, {
    name: 'loginUser',
})
export default class Login extends React.Component {
    state = {
        username: '',
        password: '',
    };

    handleSubmit = async (e) => {
        e.preventDefault();
        const { username, password } = this.state;
        const { loginUser } = this.props;

        const res = await loginUser({
            variables: {
                username,
                password,
            },
        });
        const { token } = res.data.loginUser;
        Auth.setToken(token);
        history.push('/');
    };

    render() {
        return (
            <form onSubmit={e => this.handleSubmit(e)}>
                <input type="text" onChange={e => this.setState({ username: e.target.value })} />
                <input type="password" onChange={e => this.setState({ password: e.target.value })} />
                <button>Login</button>
            </form>
        );
    }
}
