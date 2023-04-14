import axios from 'axios';

import { useState, useEffect } from 'react';
import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import 'bootstrap/dist/css/bootstrap.min.css';
import "./prod.css"
import { useNavigate } from 'react-router-dom';


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
      axios.get(loginurl + "/api/account/authcheck", {}).then(function(resp) {
         if (resp.data.status) {
            setIsloggedin(true)
         } else {
            setIsloggedin(false)
         }
         // console.log(resp.data.status)
      }).catch(function() {
         setIsloggedin(false)
      })
   }

   const getproducts = () => {
      axios.get(produrl + "/api/products/all", {}).then(function(resp) {

         setData(resp.data.products)
      })
   }
   const navigate = useNavigate()
   const logout = () => {
      axios.post(loginurl + "/api/account/logout", {}).then(function (resp) {
         if (resp.data.status) {
            setIsloggedin(false)
            navigate("/login")
         } else {
            setIsloggedin(true)
         }
         console.log(resp.data.status)
      }).catch(function () {
         setIsloggedin(false)
         navigate("/login")
      })
   }
   return (
      <div >
         <nav class="navbar navbar-dark bg-dark">
            <div className="p-2">
               <a class="navbar-brand"> CloudStore </a>
               <Button onClick={logout} class="btn btn-light pull-right"> Logout </Button>
            </div>
         </nav>
         <ul hidden={!isloggedin} className="list-group">
            {
               data.map(user => {
                  return (
                     <Card style={{ minWidth: '13rem', margin: '1rem' }}>

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

         </ul>
         <div>{order}</div>
         <div hidden={isloggedin}> Go login dumass</div>
      </div >
   );
}

export default App;
