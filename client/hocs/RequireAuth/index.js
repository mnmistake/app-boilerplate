import React from 'react';
import { graphql } from 'react-apollo';
import history from '../../history';

import getUserStatusQuery from '../../graphql/queries/auth';

export default function (ComposedComponent) {
    @graphql(getUserStatusQuery)
    class RequireAuth extends React.Component {
        componentWillUpdate(nextProps) {
            const { isAuthorized } = nextProps.data.getUserStatus ? nextProps.data.getUserStatus : false;

            if (!isAuthorized) {
                history.push('/login');
            }
        }

        render() {
            return <ComposedComponent {...this.props} />;
        }
    }

    return RequireAuth;
}
