import React, { Component } from 'react';
import axios from 'axios';
import { goURL } from '../config/environment';

class Listevents extends Component {
    
    constructor(props){
        super(props);

        this.state = {
            results:[]
        }
    }
    componentDidMount(){

        console.log("inside componentdidmount")
        // let ID = this.props.user._id;
        //axios.defaults.withCredentials = true;
         axios.get(goURL + 'events')
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
    render() { 

        let event = this.state.results.map(item =>{
            return(
                <div className="row events">
                    <div className="col-md-1">
                    {/* <img className="clogo" src={item.CompanyLogo}/>&nbsp; */}
                     </div>
                     <div className="col-md-11" >
                     <h3 className="title1">{item.eventName}</h3>
                    <h4 className="title1">Organizer:{item.organizer}</h4>
                    <h4 className="title1">Location:{item.location}</h4>
                    <h4 className="title1">Date: {new Date(item.date).toDateString()}</h4>
        
                    </div>

                </div>
            )
        })
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
                <br/>
                    {event}
                </div>
            </div>
         );
    }
}
 
export default Listevents;