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
                <Route exact path="/create" component={Createevent}/>
                <Route exact path="/book" component={Booking}/>
                <Route exact path="/list" component={Listevents}/>
                <Route exact path="/home" component={User}/>
                <Route exact path="/" component={User}/>
                <Route exact path="/dashboard" component={Dashboard}/>
            </div>
        )
    }
}
 
export default Main;