
import { HomeNavbar } from '../components'
import { Container, Row, Col, Image, Button } from 'react-bootstrap';
import { useSelector, useDispatch } from 'react-redux'
import { login, logout, cart } from '../redux/features/userSlice'
import axios from "axios";
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

function CartPage() {
   const cart = useSelector((state) => state.user.cart)
   const newArr = [...cart];
   const navigate = useNavigate();

   const createOrder = async () => {
      axios.defaults.withCredentials = true //NOTE : This is very important to be able to set cookies 
      let id = [];
      for (let i = 0; i < newArr.length; i++) {
         id = [...id, newArr[i].ID]
      }
      console.log(id)
      await axios.post("/api/orders/create", {
         ids: id,
      }).then(async (res) => {
         if (res.data.status) {
            alert("Order created successfully!")
            navigate("/")
         } else {
            alert("Order not created :(")
            navigate("/")
         }
      }).catch(err => {
         alert("Order not created :(")
         navigate("/")
      })
   }
   return (
      <>
         <HomeNavbar />
         <div style={{ marginTop: "100px" }}></div>
         {newArr.map(prod => {
            // const deleteelement = (id) => {
            //     newArr.splice(id, 1)
            // }

            return (
               <Container key={prod.ID}>
                  <Row>
                     <Col>ID:{prod.ID}</Col>
                     <Col>{prod.name}</Col>
                  </Row>
                  <Row>
                     <Col>{prod.seller_username}</Col>
                     <Col>{prod.price} $</Col>
                  </Row>
                  {/* <Button onClick={() => { deleteelement(2) }}>X</Button> */}
                  <hr />
               </Container>
            )
         }

         )}
         <center>
            <Button onClick={createOrder}>Make order</Button>
         </center>
      </>
   );
}
export default CartPage;
