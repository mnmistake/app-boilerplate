import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';

import history from './history';

import Authentication from './components/Authentication';
import Todos from './components/Todos';
import Todo from './components/Todo';
import Navbar from './components/Navbar';

import RequireAuth from './hocs/RequireAuth';
import PersistLogin from './hocs/PersistLogin';

const AppRouter = () => (
    <Router history={history}>
        <div>
            <Navbar />
            <Switch>
                <Route exact path="/" component={RequireAuth(Todos)} />
                <Route path="/todo/:id" component={RequireAuth(Todo)} />
                <Route path="/login" component={PersistLogin(Authentication)} />
                <Route path="/register" component={PersistLogin(Authentication, true)} />
            </Switch>
        </div>
    </Router>
);


export default AppRouter;
