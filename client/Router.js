import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';

import history from './history';

import Todos from './components/Todos';

const AppRouter = () => (
    <Router history={history}>                                
        <div>
            <Switch>
                <Route path="/" component={Todos} />
            </Switch>
        </div>        
    </Router>
);


export default AppRouter;