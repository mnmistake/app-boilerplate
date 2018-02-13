import React from 'react';
import { graphql } from 'react-apollo';
import registerMutation from '../../mutations/register';
// import history from '../../history';

@graphql(registerMutation, {
    name: 'registerUser',
})
export default class Register extends React.Component {
    state = {
        username: '',
        password: '',
    };

    handleSubmit = async (e) => {
        e.preventDefault();
        const { username, password } = this.state;
        const { registerUser } = this.props;

        const res = await registerUser({
            variables: {
                username,
                password,
            },
        });
        const { token } = res.data.registerUser;
        console.log(token);
    }

    render() {
        return (
            <form onSubmit={e => this.handleSubmit(e)}>
                <input type="text" onChange={e => this.setState({ username: e.target.value })} />
                <input type="password" onChange={e => this.setState({ password: e.target.value })} />
                <button>Register</button>
            </form>
        );
    }
}