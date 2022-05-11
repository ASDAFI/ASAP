import React from 'react'
import '../node_modules/bootstrap/dist/css/bootstrap.min.css'
import './App.css'
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom'
import Login from './components/login.component'
import SignUp from './components/signup.component'
import Home from './components/home'
import Navbar from './components/Navbar'
import { useEffect , useState } from "react";
import axios from 'axios';

import EventBus from "./common/EventBus";
function App() {
    localStorage.removeItem('token');
    localStorage.removeItem('name');

    const [currentUser, setCurrentUser] = useState(undefined);
    useEffect(() => {
      const user = localStorage.getItem("name");
  
      if (user) {
        setCurrentUser(user);
      }
  
      EventBus.on("logout", () => {
        logout();
      });
  
      return () => {
        EventBus.remove("logout");
      };
    }, []);
  
    const logout = () => {
      const headers = {
        'Authorization': `Bearer ${localStorage.getItem('token')}`,
    };
    axios.get('http://localhost:8000/logout', { headers })
    .then(res => {
      localStorage.removeItem('token');
        console.log("logged out. the token is " + localStorage.getItem('token'));
        if ( localStorage.getItem('token') === null ) {
          console.log("successfully logged out");
        }

    }).catch(err => {console.log(err)});
    localStorage.removeItem('token');
    localStorage.removeItem('name');
    setCurrentUser(undefined)

    };

    

  return (
    <Router>
      <div className="App">
      <nav className="navbar navbar-expand-lg navbar-light fixed-top">
      <div className="container">
        <Link className="navbar-brand" to={'/'}>
          Home
        </Link>
        <div className="collapse navbar-collapse" id="navbarTogglerDemo02">
                <div>
                <ul className="navbar-nav ml-auto">
                {currentUser ? (
          <div className="navbar-nav ml-auto">
            <li className="nav-item">
              <Link to={"/profile"} className="nav-link">
                {currentUser.username}
              </Link>
            </li>
            <li className="nav-item">
              <a href="/sign-in" className="nav-link" onClick={logout}>
                LogOut
              </a>
            </li>
          </div>
        ) : (
          <div className="navbar-nav ml-auto">
            <li className="nav-item">
              <Link to={"/sign-in"} className="nav-link">
                Login
              </Link>
            </li>
          </div>
        )}
  </ul>
                </div>
        </div>
      </div>
    </nav>
        <div className="outer">
          <div className="inner">
            <Routes>
              <Route exact path="/" element={<Home  />} />
              <Route path="/sign-in" element={<Login />} />
            </Routes>
          </div>
        </div>
      </div>
    </Router>
  )
}

export default App