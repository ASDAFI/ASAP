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

function App() {
    localStorage.removeItem('token');
    const [name, setName] = useState('');
    const[isLogged, setIsLogged] = useState(false);

  useEffect(() => {
    (
        async () => {
            
            const headers = {
                'Authorization': `Bearer ${localStorage.getItem('token')}`,
            };
            axios.get('http://localhost:8000/user', { headers })
            .then(res => {
                if(res.status === 200){
                    setIsLogged(true);
                   }else{setIsLogged(false)}
                setName(res.data.first_name);
            }).catch(err => {console.log(err)});
        }
    )();
});

  return (
    <Router>
      <div className="App">
       <Navbar Logged={isLogged} />
        <div className="outer">
          <div className="inner">
            <Routes>
              <Route exact path="/" element={<Home name={name} />} />
              <Route path="/sign-in" element={<Login />} />
              <Route path="/sign-up" element={<SignUp />} />
            </Routes>
          </div>
        </div>
      </div>
    </Router>
  )
}

export default App