import React, { Fragment } from 'react';
import { Router, Route, Switch } from 'react-router-dom';

import history from './history';

import Authentication from './components/Authentication';
import Sheets from './components/Sheets';
import CreateSheet from './components/CreateSheet';
import Sheet from './components/Sheet';
import Navbar from './components/Navbar';

import RequireAuth from './hocs/RequireAuth';
import PersistLogin from './hocs/PersistLogin';

const AppRouter = () => (
    <Router history={history}>
        <Fragment>
            <Navbar history={history} />
            <Switch>
                <Route exact path="/" component={RequireAuth(Sheets)} />
                <Route path="/sheet/:id" component={RequireAuth(Sheet)} />
                <Route path="/create" component={RequireAuth(CreateSheet)} />
                <Route path="/login" component={PersistLogin(Authentication)} />
                <Route path="/register" component={PersistLogin(Authentication, true)} />
            </Switch>
        </Fragment>
    </Router>
);


export default AppRouter;
