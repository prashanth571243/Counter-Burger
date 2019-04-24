import React, { Component } from 'react';
import counterburgersymbol from './counterburgersymbol.png';
import {Link} from 'react-router-dom';
import axios from 'axios';
import './Payments.css';
import { CountryDropdown} from 'react-country-region-selector';

class Payments extends Component {

    constructor(props){
        super(props);
        this.state = {
            orderId : '',
            price: '',
            cardNumber: '',
            expirationDate: '',
            cvv: '',
            zipcode: '',
            country: ''
        }
        //this.handleLogout = this.handleLogout.bind(this);
        this.onChange = this.onChange.bind(this);
    }

    onChange(e){
        this.setState({[e.target.name]:e.target.value})
    }

    render() {
        return (
            <div>
                <div className="counterburgersymbol">
                    <img src = {counterburgersymbol} height="100" width="200" alt=""></img>
                </div>
                <div className="container paymentcontainer">
                <br></br>
                    <b className="BillDetails">
                        Bill Details <br></br>
                    </b>
                    <b>
                        Order Id : {this.state.orderId} <br></br>
                        Amount to be Paid : {this.state.price} <br></br>
                    </b>
                    <br></br>
                    <p className="paymentMethod">
                        Payment Method - 
                    </p>
                    <p>
                        Enter Credit Card Details:
                    </p>
                    <form onSubmit>
                        Enter Card Number:
                        <div class="form-group">
                            <input onChange = {this.onChange} type="text" class="form-control" name="cardNumber" value={this.state.cardNumber} placeholder="Card Number" required/>
                        </div>
                        Enter Expiration Date:
                        <div class="form-group">
                            <input onChange = {this.onChange} type="date" class="form-control" name="expirationDate" value={this.state.expirationDate} placeholder="Expiration Date" required/>
                        </div>
                        Enter Security Code:
                        <div class="form-group">
                            <input onChange = {this.onChange} type="number" class="form-control" value={this.state.cvv} name="cvv" placeholder="CVV" required/>
                        </div>
                        Enter Zip/Postal Code:
                        <div class="form-group">
                            <input onChange = {this.onChange} type="number" class="form-control" value={this.state.zipcode} name="zipcode" placeholder="Zip Code" required/>
                        </div>
                        Select Country:
                        <div class="form-group">
                        <CountryDropdown
                            value={this.state.country}
                            onChange={this.onChange}
                            name="country" 
                            class="form-control"
                        />
                        </div>
                        <div>
                            <button class="btn btn-success btn-lg btn-block">Place Order</button>  
                        </div>
                    </form> 

                </div>
            </div>
        );
    }
}

export default Payments; 
