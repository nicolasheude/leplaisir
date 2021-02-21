import React from "react";
import ReactDOM from 'react-dom';
import './index.css';
import AppRouter from './AppRouter';
import 'bulma/css/bulma.css'

ReactDOM.render(
  <React.StrictMode>
    <AppRouter />
  </React.StrictMode>,
  document.getElementById('root')
);