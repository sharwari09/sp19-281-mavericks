import React, {Component} from 'react';
import {Route} from 'react-router-dom';
import Booking from "./Booking";

class Main extends Component {
    state = {  }
    render(){
        return(
            <div>
              
               
                {/* <Route path="/home" component={Home}/> */}
                <Route path="/book" component={Booking}/>
                
            </div>
        )
    }
}
 
export default Main;