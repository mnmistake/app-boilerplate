// @flow
import { withClientState } from 'apollo-link-state';
import type { InMemoryCache } from 'apollo-cache-inmemory';
import userQuery from 'graphql/queries/User.graphql';
import type { UserType } from 'types/User.types';

type State = {|
  user: UserType,
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
                setUser: (_, { username, id }: { username: string, id: number }, { cache }: InMemoryCache): null => {
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
