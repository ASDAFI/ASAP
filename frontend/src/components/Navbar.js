import axios from 'axios';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom'
import React, { useEffect , useState } from "react";

const Navbar = () =>  {
//   console.log("check token " + localStorage.getItem('token'));


//   const logout =  async () =>  {
//       const headers = {
//         'Authorization': `Bearer ${localStorage.getItem('token')}`,
//     };
//     axios.get('http://localhost:8000/logout', { headers })
//     .then(res => {
//       localStorage.removeItem('token');
//         console.log("logged out. the token is " + localStorage.getItem('token'));
//         if ( localStorage.getItem('token') === null ) {
//           console.log("successfully logged out");
//         }
//         window.location.reload();
//     }).catch(err => {console.log(err)});
//   } 
    


//   if (Logged) {
//     console.log("token is " + localStorage.getItem('token'));

//     return(
//       <nav className="navbar navbar-expand-lg navbar-light fixed-top">
//           <div className="container">
//             <Link className="navbar-brand" to={'/'}>
//               Home
//             </Link>
//             <div className="collapse navbar-collapse" id="navbarTogglerDemo02">
//                     <div>
//                     <ul className="navbar-nav ml-auto">
//       <li className="nav-item">
//         <Link to="/sign-in" className="nav-link" onClick={logout}>Logout</Link>
//       </li>
//     </ul>
//                     </div>
//             </div>
//           </div>
//         </nav>

//     )
//   } 
//   if (!Logged){
//     return(
//       <nav className="navbar navbar-expand-lg navbar-light fixed-top">
//       <div className="container">
//         <Link className="navbar-brand" to={'/'}>
//           Home
//         </Link>
//         <div className="collapse navbar-collapse" id="navbarTogglerDemo02">
//                 <div>
//                 <ul className="navbar-nav ml-auto">
//     <li className="nav-item">
//       <Link className="nav-link" to={'/sign-in'} >
//         Sign in
//       </Link>
//     </li>
//     <li className="nav-item">
//       <Link className="nav-link" to={'/sign-up'}>
//         Sign up
//       </Link>
//     </li>
//   </ul>
//                 </div>
//         </div>
//       </div>
//     </nav>
//     )


//   }

};


export default Navbar;