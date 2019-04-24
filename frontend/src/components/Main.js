import React, {Component} from 'react';
import {Route} from 'react-router-dom';
import SignUp from './User/SignUp'
import Login from './User/Login'
import LandingPage from './LandingPage/LandingPage'



class Main extends Component {
    render(){
        return(
            <div>
             <Route path="/signup" component={SignUp}/>
             <Route path="/login" component={Login}/>
             <Route path="/home" component={LandingPage}/>
            </div>
        )

    }
}
export default Main;
