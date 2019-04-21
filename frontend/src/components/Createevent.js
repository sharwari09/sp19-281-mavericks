import React, { Component } from 'react';
import axios from 'axios';
import { eventURL } from '../config/environment';
import {Link} from 'react-router-dom';
var swal = require('sweetalert')

class Createevent extends Component {
    constructor(props){
       
        super(props);
          
        this.state = {
            eventName : "",
            location : "",
            date : ""
        }
    
    this.submitEvent = this.submitEvent.bind(this);
    this.eventNameHandler = this.eventNameHandler.bind(this);
    this.locationChangeHandler = this.locationChangeHandler.bind(this);
    this.DateHandler = this.DateHandler.bind(this);
    }
    eventNameHandler = (e) => {
        this.setState({
            email : e.target.value
        })
    }
    locationChangeHandler = (e) => {
        this.setState({
            email : e.target.value
        })
    }
    DateHandler = (e) => {
        this.setState({
            email : e.target.value
        })
    }
    handleSignout = () => {
       
        localStorage.removeItem('firstname');
        localStorage.removeItem('id');
    }
    submitEvent = (e) => {
        var headers = new Headers();
        e.preventDefault();
        var data = {
            bucketname : "eventbrite",
            org_id : localStorage.getItem("id"),
            eventName : this.state.eventName, 
            location : this.state.location,
            date : this.state.date
        }
        console.log("data : ",  data);
        //axios.defaults.withCredentials = true;
        axios.post(eventURL + 'events', data, { headers: { 'Content-Type': 'application/json'}})
            .then(response => { 
            console.log("response :", response)
            if(response.status == 200)
            {
                swal("Event created successfully!", "", "success");
            }
        })
        .catch(error => {
            console.log(error)
        });
    }
    render() { 
        var firstname = localStorage.getItem("firstname")
        return (
            <div>
            <div className=" ht nav-height">
            <div class="navbar-header">
            <Link to="/">
            <h1 style={{'margin-left':'20px', color:'rgb(27, 167, 231)'}}>eventbrite</h1></Link>
            </div>
            <nav class="navbar nav">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        
            <ul class="nav navbar-nav mr-4">
                
                <Link to="#" class="nav-link dropdown-toggle" id="navbarDropdownMenuLink" data-toggle="dropdown" aria-haspopup="true">
                <h3><span class="glyphicon glyphicon-user"></span><font color="black">{firstname}</font></h3>    
                </Link>
                <div className=" dropdown-menu item" aria-labelledby="navbarDropdownMenuLink" >
                <li><Link to="/home" class="head1" onClick = {this.handleSignout}><span class="glyphicon glyphicon-log-out"></span>Sign out</Link></li>
                </div>
                </ul>
                </nav>
                </div>
            <div className="back">
            <div className="main_cont">
            <div className="main-div3">
            <div class="login-form signupform"><br/>
            <h1 className="hidden-xs">Create Event</h1>
            <form onSubmit={this.submitEvent.bind(this)}>
                <div class="form-group">
                    <input onChange = {this.eventNameHandler} type="text" name="eventName" class="form-control1" placeholder="Event Name" required/>
                </div>
                <div class="form-group">
                    <input onChange = {this.locationChangeHandler} type="text" name="location" class="form-control1"  placeholder="Location" required/>
                </div>
                <div class="form-group">
                    <input onChange = {this.DateHandler} type="date" name="date" class="form-control1"  placeholder="Date" required/>
                </div>
                <button onClick = {this.submitEvent} type="submit" className="btn btn-primary1">Create Event</button>
            </form>
            </div>
            </div>
            </div>
            </div>
            
            </div>
          );
    }
}
 
export default Createevent;