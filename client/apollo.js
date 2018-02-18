import { ApolloClient } from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';
import { setContext } from 'apollo-link-context';
import { ApolloLink } from 'apollo-link';

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

const client = new ApolloClient({
    link: ApolloLink.from([
        authLink,
        stateLink(cache),
        httpLink,
    ]),
    cache,
});

export default client;
