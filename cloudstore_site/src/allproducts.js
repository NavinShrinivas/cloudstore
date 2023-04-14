import axios from 'axios';

import { useState, useEffect } from 'react';
import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import 'bootstrap/dist/css/bootstrap.min.css';
import "./prod.css"


function App(props) {
   const loginurl = "http://localhost:5001"
   const produrl = "http://localhost:5002"
   const [isloggedin, setIsloggedin] = useState(true)
   const [data, setData] = useState([])
   // var data = []
   const [order, setOrder] = useState([])



   axios.defaults.withCredentials = true //NOTE : This is very important to be able to set cookies 
   //Do check if setting default here applies to all axios resquests from here on

   useEffect(() => {
      handleLogin()
      getproducts()
   }, [])


   const AddtoOrder = (event) => {
      order.push(event.target.value)
      // li_order.push(event.target.value)
      console.log(order)
   }


   const handleLogin = () => {
      axios.get(loginurl + "/api/account/authcheck", {}).then(function (resp) {
         if (resp.data.status) {
            setIsloggedin(true)
         } else {
            setIsloggedin(false)
         }
         // console.log(resp.data.status)
      }).catch(function () {
         setIsloggedin(false)
      })
   }

   const getproducts = () => {
      axios.get(produrl + "/api/products/info", {}).then(function (resp) {

         setData(resp.data.items)
      })
   }
   return (
      <div >
         <div hidden={!isloggedin} className="f-container">
            {
               data.map(user => {
                  return (
                     <Card style={{ width: '15rem', minWidth: '13rem', margin: '4rem' }}>

                        <Card.Body>
                           <Card.Title>{user.name}</Card.Title>
                           <Card.Text>
                              {user.age}
                           </Card.Text>
                           <Card.Text>
                              {user.manufacturer}
                           </Card.Text>
                           <Card.Text>
                              {user.price} $ / {user.limit} Items left
                           </Card.Text>
                           <Button key={user.name} value={user.name} onClick={AddtoOrder} variant="primary" >Order</Button>
                        </Card.Body>
                     </Card>
                  )
               })}

         </div>
         <div>{order}</div>
         <div hidden={isloggedin}> Go login dumass</div>
      </div >
   );
}

export default App;
