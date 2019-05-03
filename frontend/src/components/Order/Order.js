import React, { Component } from 'react';
import {NavLink} from 'react-router-dom';
import axios from 'axios';
import {Link} from 'react-router-dom';
import {Redirect} from 'react-router';
var result=[]
var hostname=""

class Order extends Component {
    constructor(props) {
        super(props);
        this.state = { 
            listed:[],
            amount:0.0
         }
         this.handleDelete=this.handleDelete.bind(this)
         this.handleModify=this.handleModify.bind(this)
    }

    componentDidMount(){
        console.log("cmpwilmount")
        var userId="9"
        axios.get(`http://localhost:8000/orders/${userId}`)
        .then((response) => { 

            result=response.data
            console.log("Result(response.data):"+response.data," Length"+response.data.length)
            console.log("Result(Stringify)"+JSON.stringify(result))

            this.setState({
                listed : this.state.listed.concat(response.data)
            })
            console.log("listed"+this.state.listed)
         });
    }

    handleDelete(e1,e2){
       console.log("Delete handler being called with item id and order id",e1,e2) 
       var orderId=e2; 
       const data={
        ItemId:e1
       }
       axios.post(`http://localhost:8000/order/item/${orderId}`,data)
       .then((response) => { 

           result=response.data
           console.log("Result(response.data):"+response.data," Length"+response.data.length)
           console.log("Result(Stringify)"+JSON.stringify(result))

           this.setState({
               listed : this.state.listed.concat(response.data)
           })
           console.log("listed"+this.state.listed)
        }); 
    }

    handleModify(e){
        console.log("Modify handler being called"+e)
    }

  

    render() {
      //  const data = this.state.listed;
      const templates = this.state.listed;
      console.log("length",this.state.listed.length)
      const fullrecord = this.state.listed;
              // const App = () => (
              //   <div>
              //     {
              //       Object.keys(templates).map(template_name => {
              //         return (
              //           <div>
              //             <div>{template_name.userId}</div>
              //             {
              //               templates[template_name].items.map(item => {
              //                 return(
              //                 <div>{item.itemId}  <span>{item.itemName}</span>  <span>{item.description}</span>
              //                 <span>{item.price}</span>
              //                 </div>
                                
                                
              //                   )
              //               })
              //             }
              //           </div>
              //         )
              //       })
              //     }
              //   </div>
              // );



        var amount=0.0
        var details
        // if(this.state.listed.length>0){
        details= this.state.listed.map((item) => {

            // amount=amount+item.totalAmount
            return(
            <div class="ml-4 w-75 mt-3 p-3 bg-light">

                  {
                    item.items.map(detail_item => {
                      return(
                        <div ><h6>1x {detail_item.itemName} <span className="ml-5" >${detail_item.price}</span> </h6>  <span>{detail_item.description}</span>
                        <button onClick={()=>this.handleDelete(detail_item.itemId,item.orderId)} className="btn-danger ml-4">Delete</button>
                        <hr></hr>
                        </div>)
                    })
                  }
               
                <tr><h4><span>Total Amount:</span><span className="ml-2">${item.totalAmount}</span></h4></tr>
                <button className="btn-primary mt-3 p-3 align-center"> Proceed to Checkout</button>
            </div>
        )}) 
        // } else{
        //   details="You have not placed any order yet!"
        // }

  
        return(
            <div class="ml-5">
            <h2>Order Contents:</h2>
              {details}
            </div>
        )
    }
}

export default Order;