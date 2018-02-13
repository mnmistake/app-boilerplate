import React from 'react';

export default class Login extends React.Component {
    state = {
        username: '',
        password: '',
    }

    render() {
        return [
            <input type="text" />,
            <input type="password" />,
        ];
    }
}
