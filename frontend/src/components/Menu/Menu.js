import React, { Component } from 'react';
import counterburgersymbol from './counterburgersymbol.png';
import cbsymbol from './cbsymbol.jpg';
import burgerdetails from './burgerdetails.png';
import {Link} from 'react-router-dom';
import axios from 'axios';
import './Menu.css';
// import starter from './starter.png'
// import starter5 from './starter5.png'
// import starter6 from './starter6.png'
 import starter7 from './starter7.png'
// import starter8 from './starter8.png'
// import burger from './burger.png'
// import burger1 from './burger1.png'
// import burger2 from './burger2.png'
// import burger3 from './burger3.png'
// import burger4 from './burger4.png'
// import burger5 from './burger5.png'
// import drinks from './drinks.png'
// import drinks1 from './drinks1.png'
// import drinks2 from './drinks2.png'
// import drinks3 from './drinks3.png'
// import drinks4 from './drinks4.png'
import { Card } from 'antd';
const { Meta } = Card;

class Menu extends Component {
    
    
    constructor(props){
        super(props);
        this.state = {
            allmenu: [],  
            starters: [],
            burgers: [],
            beverages: [],
        }
        //this.handleLogout = this.handleLogout.bind(this);
        //this.submitBooking = this.submitBooking.bind(this);
    }  

    

    componentWillMount()
    {
        axios.get("http://localhost:8000/menu")
                    .then((response) => {
                        console.log("Response data", response.data)
                        this.setState({
                            allmenu : response.data
                        })
                });
                console.log("Checking menu details", this.state.allmenu)
    }



    render() {

        let wholemenu = this.state.allmenu.map((wholemenu,j) => {
        return(
                <Card
                hoverable
                style={{ width: 400 }}
                cover={<img src = {cbsymbol} height="320" width="670" alt=""></img>}
                className="MenuCards"
                >
                <br></br>
                <Meta
                title={wholemenu.ItemType}
                description={wholemenu.ItemName}
                />
                <br></br>
               <p>Description: {wholemenu.Description}</p>
               <p>Price : {wholemenu.Price}$</p>
               <br></br>
               <button onClick={()=>this.addToCart(wholemenu.ItemId)} className="btn btn-danger cartButton ">Add to Cart</button>
                </Card>
        )
        })

        return (
            <div>
                <div className="counterburgersymbol">
                    <img src = {counterburgersymbol} height="100" width="200" alt=""></img>
                    <div className="container MenuOustide">
                    <div className="storedetails">
                        <b>THE COUNTER</b>
                        <br></br>
                       <Link to="/location">Change Location</Link>
                       <br></br>
                       <p>Phone: (408) 423-9200</p>
                       <p>Pickup Hours: Open today 11am-10pm </p>
                       <p>Accepted Cards: Mastercard, Visa, American Express, Discover</p>
                    </div>
                    <div className="burgerdetails">
                    <img src = {burgerdetails} height="250" width="670" alt=""></img>
                    <p>Selections vary by location and may have limited availability.<br></br>
                        <a className="allergy" href="/nutrition">Nutritional information<br></br>
                        Allergen information</a>
                    </p>
                    </div>
                    <div className="container Menu">
                        {wholemenu}
                        <br></br>
                    </div>
                </div>
                </div>
            </div>
        );
    }
}

export default Menu;