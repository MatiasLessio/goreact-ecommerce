import React from 'react';
import {BrowserRouter as Router} from "react-router-dom";
import {Paginas} from './Paginas'
import {Header} from './header/Header';
import './App.css';

function App() {
  return (
    <div>
      <Router>
        <div><Header/></div>
        <div><Paginas/></div>
      </Router>
    </div>
  );
}

export default App;
