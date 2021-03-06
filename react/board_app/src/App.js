import './css/App.css';
import React from 'react';
import { BrowserRouter, Route } from 'react-router-dom';
import Home from './routes/Home';
import Board from './routes/Board';
import Navigator from './components/Navigator';
import Article from './components/Article';

function App() {
  return (
    <BrowserRouter>
      <Navigator />
      <Route path="/" exact={true} component={Home} />
      <Route path="/board/:id" exact={true} component={Board} />
      <Route path="/board/:id/article/:aid" component={Article} />
    </BrowserRouter>
  );
}

export default App;
 