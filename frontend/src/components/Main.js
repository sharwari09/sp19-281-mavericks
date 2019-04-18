import React, {Component} from 'react';
import {Route} from 'react-router-dom';
import Booking from "./Booking";
import User from "./User";
import Listevents from './Listevents';

class Main extends Component {
    state = {  }
    render(){
        return(
            <div>
              
               
                {/* <Route path="/home" component={Home}/> */}
                <Route path="/book" component={Booking}/>
                <Route path="/list" component={Listevents}/>
                <Route path="/" component={User}/>
                
            </div>
        )
    }
}
 
export default Main;