import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';

import history from './history';

import Login from './components/Login';
import Register from './components/Register';
import Todos from './components/Todos';

const AppRouter = () => (
    <Router history={history}>
        <div>
            <Switch>
                <Route path="/login" component={Login} />
                <Route path="/register" component={Register} />
                <Route exact path="/" component={Todos} />
            </Switch>
        </div>
    </Router>
);


export default AppRouter;