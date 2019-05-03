import React, {Component} from 'react';
import "./location.css";
import {Link} from 'react-router-dom';
import axios from 'axios';



class Location extends Component {
    constructor(props){
        super(props);
        this.state = {
            location : "",
            locationInfo: []
        }
        this.onChangeLocation = this.onChangeLocation.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
        this.settingLocation = this.settingLocation.bind(this);
    }

    onChangeLocation (e) {
        e.preventDefault();
        this.setState({[e.target.name] : e.target.value})
    }
    async onSubmit(e){
        e.preventDefault();
         const locationData = this.state.location
        try{
            const connectionReqResponse = await axios.get(`http://localhost:8000/location/getLocation/${locationData}`)
           console.log("checkpoint")
            console.log(connectionReqResponse.data)
            this.setState({
                locationInfo : this.state.locationInfo.concat(connectionReqResponse.data)
            })
            console.log(this.state.locationInfo)
           
        } catch(err) {
            
        }
    
    }

    settingLocation(locationName) {
        console.log("this")
        localStorage.setItem("locationName",locationName)
        this.props.history.push("/menu")

    }

    render(){
        var details = this.state.locationInfo.map((value,i) => {
            return(
                
                <div >
                   <table>
                   <th><button  type="button" onClick = {() => this.settingLocation(value.locationName)} className="locationButton">{value.locationName}</button></th>
                    {/* <th>{value.locationName}</th> */}
                    <tr>{value.addressLine1}</tr>

                   </table>
                </div>      
            )

        })
       return(
           <div>
               
            <div className="location"> 
            <div className="locationNav">
         <span className="locationSpan"> <a href="/" id="A_4"></a>
         </span>
         <div id="menu-outer-location">
  <div className="tableLocation">
    <ul id="horizontal-list-location">
        <li><Link to="/home"><font color = "black" face="Impact" size="4">HOME</font></Link></li>
      <li><Link to="/login"><font color = "black" face="Impact" size="4">VIEW MENUS</font></Link></li>
      <li><Link to="/signup"><font color = "black" face="Impact" size="4">CREATE ACCOUNT</font></Link></li>
      <li><Link to="/login"><font color = "black" face="Impact" size="4 ">SIGN IN</font></Link></li>
    </ul>
  </div>
</div>  
 </div>
  </div>
  <div className="locationBox"> 
    <h1 id="H1_3">
        Find Locations Near You
    </h1>
    <h2 id="H2_7">
        Enter Zip Code:
    </h2>
    <form action="/location" id="FORM_8" onSubmit = {this.onSubmit}  >
    <input type="text" id="INPUT_9" name="location" placeholder="Street address or zip code" onChange = {this.onChangeLocation} />
<br></br>
        <input type="submit" id="BUTTON_14" />
    </form>
    <br></br>

    {details}
    
   </div>
 
           </div>
       )

    }

}
export default Location;
