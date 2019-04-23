import React, { Component } from 'react';
import axios from 'axios';
import { eventURL } from '../config/environment';
import {Link} from 'react-router-dom';
import {Redirect} from 'react-router';
class Listevents extends Component {
    
    constructor(props){
        super(props);

        this.state = {
            results:[],
            redirectVar :null
        }
        this.showBooking = this.showBooking.bind(this);
    }
    handleSignout = () => {
       
        localStorage.removeItem('firstname');
        localStorage.removeItem('id');
    }
    
    componentWillMount(){
        console.log("inside componentdidmount")
        var firstname = localStorage.getItem("firstname")
        //axios.defaults.withCredentials = true;
         axios.get(eventURL + 'events')
             .then((response) => {
                console.log("response", response)
                if(response.status == 200)
                {
                    console.log("events count ", response.data.count)
                    console.log("event 1 ", response.data.events[0])
                    this.setState({
                        results : this.state.results.concat(response.data.events)
                    });
                }

        });
    }

    showBooking = (ID ,e) => {
        var headers = new Headers();
        e.preventDefault();
        localStorage.setItem("eid", ID)
        this.setState({
            redirectVar : <Redirect to= "/book"/>
        })

        //axios.defaults.withCredentials = true;
        // axios.get(eventURL + 'events/' + ID)
        //     .then(response => {
        //         console.log("Response :",response);
        //         if(response.status === 200){
        //             console.log("successfully rendered to booking page");
        //             this.setState({
        //                 redirectVar : <Redirect to= "/book"/>
        //         })
                    
        //         }
        //     });
    }
    render() { 
        var firstname = localStorage.getItem("firstname")
        if (!localStorage.getItem("id")) {
            this.setState({
                redirectVar : <Redirect to="/home" />
            })
            
        }
        let event = this.state.results.map(item =>{
            return(
                <div className="row events">
                
                    <div className="col-md-1">

                     </div>
                     <div className="col-md-11" >
                     <h3 className="title1"><Link to="/book" onClick = {this.showBooking.bind(this, item.eventId)}>{item.eventName}</Link></h3>
                    <h4 className="title1">Location:{item.location}</h4>
                    <h4 className="title1">Date: {new Date(item.date).toDateString()}</h4>
        
                    </div>

                </div>
            )
        })
        return ( 
            <div>
            {this.state.redirectVar}
                {/* <div className="nav-height"> */}
                {/* <div class="container-fluid"> */}
                    {/* <div class="navbar-header">
                    <h1 style={{'margin-left':'20px', color:'rgb(27, 167, 231)'}}>eventbrite</h1>
                    </div>
                    <nav class="navbar nav">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; */}
                    {/* <nav class="navbar nav">{email} </nav> */}
                    {/* <ul class="nav navbar-nav mr-4">
                    <Link to="#" class="nav-link dropdown-toggle" id="navbarDropdownMenuLink" data-toggle="dropdown" aria-haspopup="true">
                <h3><span class="glyphicon glyphicon-user"></span><font color="black">{email}</font></h3>    
                </Link>
                <div className=" dropdown-menu item" aria-labelledby="navbarDropdownMenuLink" >

                <li><Link to="/" class="head1" onClick = {this.handleSignout}><span class="glyphicon glyphicon-log-out"></span>Sign out</Link></li>
                </div>
                    </ul>
                    </nav> */}
                {/* </div> */}
            <div className=" ht nav-height">
            <div class="navbar-header">
            <Link to="/">
            <h1 style={{'margin-left':'20px', color:'rgb(27, 167, 231)'}}>eventbrite</h1></Link>
            </div>
            <nav class="navbar nav">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            <ul><a href="/list" className="buttons">List Events</a></ul>
            <ul><a href="/create" className="buttons">Create Event</a></ul>
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
                <br/>
                    {event}
                </div>
            </div>
         );
    }
}
 
export default Listevents;