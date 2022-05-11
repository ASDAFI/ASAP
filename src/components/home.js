import React, { useEffect , useState } from "react";
import axios from 'axios';
import Sidebar from './Sidebar'

function Home() {
    const[name  , setName] = useState(undefined);
    console.log("check token " + localStorage.getItem('token'));
    // console.log(localStorage.getItem('token'));
  
    useEffect(() => {
        (
            async () => {
                
                const headers = {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`,
                };
                axios.get('http://localhost:8000/user', { headers })
                .then(res => {
                    
                    localStorage.setItem('name',res.data.first_name);
                    setName(res.data.first_name);
                }).catch(err => {console.log(err)});
            }
        )();
    });
    
    return (
        
        <div className="home">
            <Sidebar />
            <h3> {name ? 'Welcome ' + name : 'You are not logged in'} </h3>
            
        </div>
    );
}
export default Home;