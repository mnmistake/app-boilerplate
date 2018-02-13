import React from 'react';
import ReactDOM from 'react-dom';
import { ApolloProvider } from 'react-apollo';

// import './sass/style.scss';
// import './sass/_fonts.scss';

import Router from './Router';
import client from './apollo';

ReactDOM.render(
    <ApolloProvider client={client}>
        <Router />
    </ApolloProvider>,
    document.getElementById('app'),
);
