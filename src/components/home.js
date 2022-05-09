import React, { useEffect , useState } from "react";
import axios from 'axios';

function Home({name}) {
    console.log("check token " + localStorage.getItem('token'));
    // console.log(localStorage.getItem('token'));

    return (
        <div>
            <h3> {name ? 'Welcome ' + name : 'You are not logged in'} </h3>
        </div>
    );
}
export default Home;