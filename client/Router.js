import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';

import history from './history';

import Authentication from './components/Authentication';
import Todos from './components/Todos';
import RequireAuth from './hocs/RequireAuth';
import IsAuthorized from './hocs/IsAuthorized';

const AppRouter = () => (
    <Router history={history}>
        <div>
            <Switch>
                <Route path="/login" component={IsAuthorized(Authentication)} />
                <Route path="/register" component={IsAuthorized(Authentication, true)} />
                <Route exact path="/" component={RequireAuth(Todos)} />
            </Switch>
        </div>
    </Router>
);


export default AppRouter;
