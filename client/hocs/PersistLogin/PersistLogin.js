import React from 'react';
import { graphql } from 'react-apollo';
import Auth from '../../utils/Auth';
import history from '../../history';

import getUserStatusQuery from '../../graphql/queries/auth';

export default function (ComposedComponent, isRegister = false) {
    @graphql(getUserStatusQuery, {
        skip: () => !Auth.isUserAuthenticated(), // skip the getStatus query if there is no JWT to verify
    })
    class PersistLogin extends React.Component {
        componentWillUpdate(nextProps) {
            const { isAuthorized } = nextProps.data.getUserStatus ? nextProps.data.getUserStatus : false;
            if (isAuthorized) {
                history.push('/');
            }
        }

        render() {
            return <ComposedComponent {...this.props} isRegister={isRegister} />;
        }
    }

    return PersistLogin;
}
