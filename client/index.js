import React from 'react';
import ReactDOM from 'react-dom';
import { ApolloProvider } from 'react-apollo';
//import { Provider } from 'react-redux';

//import './sass/style.scss';
//import './sass/_fonts.scss';

import Router from './Router';
import Client from './apollo';


ReactDOM.render(
    <ApolloProvider client={Client}>
        <Router />
    </ApolloProvider>,
    document.getElementById('app'),
);