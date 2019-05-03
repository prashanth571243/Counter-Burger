import React, {Component} from 'react';
import {Route} from 'react-router-dom';
import Location from './Location/Location'




class Main extends Component {
    render(){
        return(
            <div>
             <Route path="/location" component={Location}/>

            </div>
        )

    }
}
export default Main;