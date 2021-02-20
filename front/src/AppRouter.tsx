// import Router from "react-dom-router";

import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import React from "react";

import IndexComponent from './IndexComponent'
import NameForm from './Form/app'
import SignIn from './SignIn/app'
import SignUpConfirm from './SignUpConfirm/app'

export default function AppRouter() {
     return (
       <Router>
         <div>
           <nav>
             <ul>
               <li>
                 <Link to="/">Home</Link>
               </li>
           <li>
             <Link to="/sign-up">Sign Up</Link>
               </li>
              <li>
                <Link to="/sign-in">Sign In</Link>
              </li>
            </ul>
          </nav>

           <Route path="/" exact component={IndexComponent} />
           <Route path="/sign-up" component={NameForm} />
           <Route path="/sign-up-confirm" component={SignUpConfirm} />
           <Route path="/sign-in" component={SignIn} />
         </div>
       </Router>
     );
  }