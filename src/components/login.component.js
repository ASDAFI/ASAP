import React, { Component, useState } from "react";
import {Navigate} from "react-router-dom";

export default class Login extends Component {
    
    constructor(props) {
        super(props);
        this.state = {
          redirect:false ,
          credentials: {username: '', password: ''} 
        };
      }
  
    login = event => {
        console.log('submitted!');
    }
    

    inputChanged = event => {
        const cred = this.state.credentials;
        cred[event.target.name] = event.target.value;
        this.setState({credentials: cred});
      }
      onSubmit =  event => {
        event.preventDefault();
        // console.log(this.state.credentials.username , this.state.credentials.password );
        const u  = this.state.credentials.username;
        const p  =  this.state.credentials.password
        const response = fetch('http://localhost:8000/login' ,{
            method: 'POST',
            Headers :{
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json',
            } ,
            body: JSON.stringify(
                { username : u , password : p }
            )
        }).then(response =>  response.json() ).then( data => {
            console.log(data);
            if(data.token){
                console.log('logged in');
                localStorage.setItem('token' , data.token);
                this.setState({redirect: true});
            }
            else{
                console.log('not logged in');
                this.setState({redirect: false});
            }
        }).catch(err => {console.log(err)});
    }
  
    render() {
        if (this.state.redirect) {
            return < Navigate to="/"/>;
        }
        return (
            <form  style={{ }} onSubmit={this.onSubmit}>
                <h3>Log in</h3>
                <div className="form-group" >
                    <label>Username</label>
                    <input type="username" className="form-control" placeholder="Enter username"  name="username" value={this.state.credentials.username}  onChange={this.inputChanged} />
                </div>

                <div className="form-group">
                    <label>Password</label>
                    <input type="password" className="form-control" placeholder="Enter password" id = "password" name="password" value={this.state.credentials.password} onChange={this.inputChanged}   />
                </div>

                <div className="form-group">
                    <div className="custom-control custom-checkbox">
                        <input type="checkbox" className="custom-control-input" id="customCheck1" />
                        <label className="custom-control-label" htmlFor="customCheck1">Remember me</label>
                    </div>
                </div>
                <button type="submit" onClick={this.login} className="btn btn-dark btn-lg btn-block"  >Sign in</button>
            </form>
        );
    }
}
