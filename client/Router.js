import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';

import history from './history';

import Login from './components/Login';
import Register from './components/Register';
import Todos from './components/Todos';
import RequireAuth from './components/RequireAuth';
import IsAuthenticated from './components/IsAuthenticated';

const AppRouter = () => (
    <Router history={history}>
        <div>
            <Switch>
                <Route path="/login" component={IsAuthenticated(Login)} />
                <Route path="/register" component={IsAuthenticated(Register)} />
                <Route exact path="/" component={RequireAuth(Todos)} />
            </Switch>
        </div>
    </Router>
);


export default AppRouter;
