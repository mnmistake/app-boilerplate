import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';

import history from './history';

import Authentication from './components/Authentication';
import Todos from './components/Todos';
import RequireAuth from './hocs/RequireAuth';
import IsAuthenticated from './hocs/IsAuthenticated';

const AppRouter = () => (
    <Router history={history}>
        <div>
            <Switch>
                <Route path="/login" component={IsAuthenticated(Authentication)} />
                <Route path="/register" component={IsAuthenticated(Authentication, true)} />
                <Route exact path="/" component={RequireAuth(Todos)} />
            </Switch>
        </div>
    </Router>
);


export default AppRouter;
