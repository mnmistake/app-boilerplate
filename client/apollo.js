import { ApolloClient } from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';
import { setContext } from 'apollo-link-context';
import { ApolloLink } from 'apollo-link';
import { onError } from 'apollo-link-error';

import history from './history';
import stateLink from './state';
import Auth from './utils/Auth';

const cache = new InMemoryCache();

const httpLink = new HttpLink({
    uri: '/graphql',
});

const authLink = setContext((req, { headers }) => {
    const token = Auth.getToken();
    return {
        headers: {
            ...headers,
            ...token && { Authorization: `Bearer ${token}` },
        },
    };
});

const errorLink = onError(({ graphQLErrors, networkError }) => {
    if (graphQLErrors)
        graphQLErrors.map(({ message }) => {
            if (message === 'Invalid token') {
                Auth.removeToken();
                history.push('/login');
            }
        });
    if (networkError) console.error(`[Network error]: ${networkError}`);
});

const client = new ApolloClient({
    link: ApolloLink.from([
        errorLink,
        authLink,
        stateLink(cache),
        httpLink,
    ]),
    cache,
});

export default client;
