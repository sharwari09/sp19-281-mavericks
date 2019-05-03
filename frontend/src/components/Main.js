import React, {Component} from 'react';
import {Route} from 'react-router-dom';
import Booking from "./Booking";
import User from "./User";
import Listevents from './Listevents';
import Createevent from "./Createevent";
import Dashboard from "./dashboard.js";

class Main extends Component {
    state = {  }
    render(){
        return(
            <div>
                <Route path="/create" component={Createevent}/>
                <Route path="/book" component={Booking}/>
                <Route path="/list" component={Listevents}/>
                <Route path="/home" component={User}/>
                <Route path="/dashboard" component={Dashboard}/>
            </div>
        )
    }
}
 
export default Main;