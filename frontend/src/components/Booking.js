
import React, { Component } from 'react';
import axios from 'axios';
import { goURL } from '../config/environment';

class Booking extends Component {
    constructor(props){
       
    super(props);
      
    this.state = {
        price : "",
        quantity : ""
    }
this.ticketChangeHandler = this.ticketChangeHandler.bind(this);
this.submitBooking = this.submitBooking.bind(this);

}
ticketChangeHandler = (e) => {
    this.setState({
        quantity : e.target.value
    })
}

    submitBooking = (e) => {
        var headers = new Headers();
        e.preventDefault();
        var price = this.state.quantity * 40;
        console.log("price" + price);
        // this.setState({
        //     price : price,
        // });
        var data = {
            EventName : "San Jose Career Fair",
            Price : price,
            EventID: "0123012218",
            UserID: "1923840290"
        }
        console.log("data : ",  data);
        //set the with credentials to true
        axios.defaults.withCredentials = true;
        axios.post(goURL + 'book', data, { headers: { 'Content-Type': 'application/json'}})
            .then(response => { 
            console.log("response :", response)
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
            <div class="login-form1 signupform">
            <div className="headerback"></div>
            <div>
                <h1 className="header">Career Fair</h1>
                <h3 className="title">Organizer     :</h3>
                <p style={{'margin-right':'-20px','margin-top' : '-37px'}}>San Jose State University</p>
                <h3 className="title">Location :</h3>
                <p style={{'margin-left':'-70px','margin-top' : '-37px'}}>San Jose, CA</p>
                <h3 className="title">Date and Time :</h3>
                <p style={{'margin-left':'90px','margin-top' : '-37px'}}> July 27, 2019, 9:00 AM – 7:00 PM PDT</p>
                <h3 className="title">Price :</h3>
                <p style={{'margin-left':'-150px','margin-top' : '-37px'}}>$40</p>
                <h3 className="title">Quantity   :</h3>
                <div style={{'margin-left':'-115px','margin-top' : '-37px'}}>
                        <select class="form-control1"  onChange = {this.ticketChangeHandler} required>
                        <option value="Select">Select</option>
                            <option value="1">1</option>
                            <option value="2">2</option>
                            <option value="3">3</option>
                            <option value="4">4</option>
                            <option value="5">5</option>
                        </select>

                    </div>
                <button onClick = {this.submitBooking} className="btn btn-primary1" type="submit">Register</button>
            </div>
            </div>
            </div>
            </div>
            </div>

            </div>
         );
    }
}
 
export default Booking;