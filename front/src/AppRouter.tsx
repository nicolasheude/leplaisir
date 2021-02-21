// import Router from "react-dom-router";

import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import React from "react";

import IndexComponent from './IndexComponent'
import NameForm from './Form/app'
import SignIn from './SignIn/app'
import SignUpConfirm from './SignUpConfirm/app'
import Tableau from './Form/Table/Table'

export default function AppRouter() {
     return (
         <div>
            <Router>
              <nav className="navbar" role="navigation" aria-label="main navigation">
                <div className="navbar-item">
                  <img src="/logo.png" width="112" height="28" />
                </div>
                <div className="navbar-start">
                  <Link to="/sign-in" className="navbar-item">Admin</Link>
                  <Link to="/sign-up" className="navbar-item">Formulaire</Link>
                  <Link to="/dash" className="navbar-item">Planning</Link>
                </div>
              </nav>

              <Route path="/" exact component={IndexComponent} />
              <Route path="/sign-up" component={NameForm} />
              <Route path="/sign-in" component={SignIn} />
              <Route path="/manager" component={Tableau} />
            </Router>

         </div>
     );
  }