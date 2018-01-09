import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';

import history from './history';

import App from './components/App/App';

const AppRouter = () => (
    <Router history={history}>                                
        <div>
            <Switch>
                <Route path="/" component={App}/>
            </Switch>
        </div>        
    </Router>
);


export default AppRouter;