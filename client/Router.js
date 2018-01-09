import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';

import history from './history';

import TodoList from './components/TodoList/TodoList';

const AppRouter = () => (
    <Router history={history}>                                
        <div>
            <Switch>
                <Route path="/" component={TodoList}/>
            </Switch>
        </div>        
    </Router>
);


export default AppRouter;