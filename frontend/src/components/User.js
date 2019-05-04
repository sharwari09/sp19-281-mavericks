import React, { Component } from 'react';
import axios from 'axios';
import { userURL } from '../config/environment';
var swal = require('sweetalert')
class User extends Component {
    constructor(props){
       
        super(props);
          
        this.state = {
            email : "",
            password : ""
        }
    
    this.submitLogin = this.submitLogin.bind(this);
    this.emailChangeHandler = this.emailChangeHandler.bind(this);
    this.passwordChangeHandler = this.passwordChangeHandler.bind(this);
    }
    emailChangeHandler = (e) => {
        this.setState({
            email : e.target.value
        })
    }
    passwordChangeHandler = (e) => {
        this.setState({
            password : e.target.value
        })
    }

    componentDidMount(){
        axios.get(`${userURL}users`).then((response)=>{
            console.log(response.status);
            console.log(response.data);
        })
    }
    submitLogin = (e) => {
        var headers = new Headers();
        e.preventDefault();
        var data = {
            Email : this.state.email, 
            Password : this.state.password
        }
        console.log("data : ",  data);
        //axios.defaults.withCredentials = true;
        axios.post(userURL + 'users/signin', data, { headers: { 'Content-Type': 'application/json'}})
            .then(response => { 
            console.log("response :", response)
            if(response.status == 200)
            {
                localStorage.setItem("firstname", response.data.firstname)
                localStorage.setItem("id", response.data.id)
                this.props.history.push("/list");
            }
                
                
                //swal("User logged in Successfully!", "", "success");
        })
        .catch(error => {
            console.log(error)
        });
    }
    render() { 
        return ( 
            <div>
            <div className="nav-height">
                {/* <div class="container-fluid"> */}
                    <div class="navbar-header">
                    <h1 style={{'margin-left':'20px', color:'rgb(27, 167, 231)'}}>eventbrite</h1>
                    </div>
                    <nav class="navbar nav"> </nav>
            </div>
            <div className="back">
            <div className="main_cont">
            <div className="main-div3">
            <div class="login-form signupform"><br/>
            <h1 className="hidden-xs">Log In</h1>
            <form onSubmit={this.submitLogin.bind(this)}>
                <div class="form-group">
                    <input onChange = {this.emailChangeHandler} type="email" name="email" class="form-control1" placeholder="Email Address" required/>
                </div>
                <div class="form-group">
                    <input onChange = {this.passwordChangeHandler} type="password" name="password" class="form-control1"  placeholder="Password" required/>
                </div>
                <button onClick = {this.submitLogin} type="submit" className="btn btn-primary1">Log In</button>
            </form>
            </div>
            </div>
            </div>
            </div>
            
            </div>
         );
    }
}
 
export default User;