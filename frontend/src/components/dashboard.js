import React, { Component } from 'react';
import axios from 'axios';
import { dashboardURL,bucket } from '../config/environment';
import './../css/eventCard.css';
import {Link} from 'react-router-dom';
import {Redirect} from 'react-router';


function PostedEventCard(props){
    const {posted} = props;
console.log(posted);
    return ( 
        <div className="row row-style row-booking">
             <div style={{padding:'8px 4px 4px 8px'}}>
                <p className="clearfix" >Id: {posted.eventId}</p>
                <p className="clearfix" >Event Name : {posted.eventName}</p>
                <p className="clearfix" >Location : {posted.location}</p>
                <p className="clearfix" >Date : {posted.date}</p>
                <p className="clearfix" >Number of Bookings : {posted.numberOfBookings}</p>
             </div>
        </div>
);
}

function BookedEventCard(props){
    const {booked} = props;
    return ( 
        <div className="row row-style row-booking">
        <div style={{padding:'8px 4px 4px 8px'}}>
           <p className="clearfix" >Organizer Id: {booked.orgId}</p>
           <p className="clearfix" >Event Name : {booked.eventName}</p>
           <p className="clearfix" >Location : {booked.location}</p>
           <p className="clearfix" >Date : {booked.date}</p>
        </div>
   </div>
);
}


class Dashboard extends Component {
    
    constructor(props){
        super(props);

        this.state = {
            postedEvents:[],
            bookedEvents:[],
            user_uuid:localStorage.getItem('id'),
            firstname:localStorage.getItem('firstname'),
        }
        
    }
    
    handleSignout = () => {
        localStorage.removeItem('firstname');
        localStorage.removeItem('id');
    }
    
    componentDidMount(){
        
        var user_uuid = this.state.user_uuid
        //axios.defaults.withCredentials = true;
        const payload = {
            "bucket" : bucket,
            "user_uuid" : this.state.user_uuid
        }
        console.log("inside component did mount")
         axios.post(dashboardURL,payload)
             .then((response) => {
                console.log("response", JSON.stringify(response))
                if(response.status == 200){
                    this.setState({ 
                        bookedEvents:response.data.bookedEvents,
                        postedEvents:response.data.postedEvents
                    });
                }

        });
    }

    

    render() { 
        
        var redirectVar = null;
        if (!localStorage.getItem("id")) {
            redirectVar = <Redirect to="/home" />
        }

        const {postedEvents,bookedEvents} = this.state;
        const {firstname,user_uuid} = this.state;

        
        return ( 
            <div>
            {redirectVar}
                
            <div className=" ht nav-height">
            <div class="navbar-header">
            <Link to="/">
            <h1 style={{'margin-left':'20px', color:'rgb(27, 167, 231)'}}>Eventbrite</h1></Link>
            </div>
            <nav class="navbar nav">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            <ul><Link to="/list" className="buttons">List Events</Link></ul>
            <ul><Link to="/create" className="buttons">Create Event</Link></ul>
            <ul><Link to="/dashboard" className="buttons">Dashboard</Link></ul>
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

                <div className="back container">

                <div className="row w-100 justify-content-md-center" >
                    <div className="col-md-4" style={{border:'1px solid #00FFFF'}}>
                        <p>Posted Events</p>
                    </div>

                    <div className="col-md-4" style={{border:'1px solid #00FFFF'}}>
                        <p>Booked Events</p>
                    </div>
                </div>
                    <div className="row w-100 justify-content-md-center" >
                        
                        <div className="col-md-4" style={{border:'1px solid #00FFFF'}}>
                            {
                                postedEvents.map((posted,index)=>{
                                    return(
                                    <PostedEventCard posted={posted} key={index} />
                                    );
                                })
                            }
                        </div>

                        

                        <div className="col-md-4" style={{border:'1px solid #00FFFF'}} >
                        {
                                bookedEvents.map((booked,index)=>{
                                    return(
                                    <BookedEventCard booked={booked} key={index} />
                                    );
                                })
                                
                            }
                        </div>

                        
                    </div>
                </div>
            </div>
         );
    }
}
 
export default Dashboard;