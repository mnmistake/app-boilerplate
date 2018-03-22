// @flow
import { withClientState } from 'apollo-link-state';
import type { InMemoryCache } from 'apollo-cache-inmemory';
import userQuery from './graphql/queries/User.graphql';

type User = {
    __typename: string,
    username: ?string,
    id: ?number,
};

type State = {|
  user: User,
|};

const initialState: State = {
    user: {
        __typename: 'CurrentUser',
        username: null,
        id: null,
    },
};

export default (cache: InMemoryCache) =>
    withClientState({
        cache,
        defaults: initialState,
        resolvers: {
            Mutation: {
                setUser: (_, { username, id }: { username: boolean, id: number }, { cache }: InMemoryCache): void => {
                    const prevState = cache.readQuery({ query: userQuery });
                    const data: State = {
                        ...prevState,
                        user: {
                            ...prevState.user,
                            username,
                            id,
                        },
                    };

                    cache.writeData({ data });
                    return null;
                },
            },
        },
    });
