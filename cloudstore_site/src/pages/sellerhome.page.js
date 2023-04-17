// import { useState } from 'react'
// import { useSelector } from 'react-redux'

import Nav from 'react-bootstrap/Nav';
import { Form, Button } from 'react-bootstrap';
import Navbar from 'react-bootstrap/Navbar';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { HomeNavbar, ProductSellerCard } from '../components'
import axios from 'axios';
import { useState, useEffect } from 'react';


function SellerHomePage() {

   // const [user, setUser] = useState(useSelector((state) => state.user.value))
   const [products, setProducts] = useState([])
   const [query, setQuery] = useState({
      name: "",
      price: 0,
      limit: 0,
      manufacturer: "",
   });

   const handleChange = () => (e) => {
      const name = e.target.name;
      const value = e.target.value;
      setQuery((prevState) => ({
         ...prevState,
         [name]: value
      }));
   }

   const handleSubmit = async (e) => {
      e.preventDefault()
      const formData = new FormData();
      Object.entries(query).forEach(([key, value]) => {
         formData.append(key, value);
      });
      const jsonObject = {};

      for (const [key, value] of formData) {
         jsonObject[key] = value;
      }


      console.log(query)
      query.limit = parseInt(query.limit)
      query.price = parseInt(query.price)
      
      await axios.post("/api/products/create", query).then(resp => console.log(resp))
      window.location.reload()

   };


   const getproducts = async () => {
      await axios.get("/api/products/seller", {}).then(function(resp) {
         console.log(resp)
         setProducts(resp.data.items)
         // for (let i = 0; i < resp.data.products.length; i++) {
         //    sellers.add(resp.data.products[i].seller_username)
         //    manufacturers.add(resp.data.products[i].manufacturer)
         // }
      })
   }
   useEffect(() => {
      getproducts()
   }, [])

   return (
      <>
         <div className="home-page" style={{ backgroundColor: '#f5f5f5', overflow: 'hidden' }}>
            <HomeNavbar />
            <center>
               <div className="mx-4" style={{ marginTop: '4rem' }}>
                  <div className="row">
                     <div className="col-12">
                        <div className="row">
                           {
                              products.map(product => {
                                 return (
                                    <ProductSellerCard key={product.id} product={product} />
                                 )
                              })
                           }
                        </div>
                     </div>
                  </div>
               </div>
            </center>

            <Container style={{ margin: "20px" }}>
               <Form onSubmit={handleSubmit}>
                  <h2>Add new products</h2>

                  <Row>

                     <Col>
                        <Form.Group controlId="formBasicName">
                           <Form.Label>Name</Form.Label>
                           <Form.Control required value={query.name} onChange={handleChange()} type="text" name="name" />
                        </Form.Group>
                     </Col>
                     <Col>
                        <Form.Group controlId="formBasicName">
                           <Form.Label>price</Form.Label>
                           <Form.Control required value={query.price} onChange={handleChange()} type="number" name="price" />
                        </Form.Group>
                     </Col>
                  </Row>
                  <br />
                  <Row>
                     <Col>

                        <Form.Group controlId="formBasicEmail">
                           <Form.Label> Limit </Form.Label>
                           <Form.Control required value={query.limit} onChange={handleChange()} type="number" name="limit" />
                        </Form.Group>
                     </Col>
                     <Col>
                        <Form.Group controlId="formBasicPassword">
                           <Form.Label>Manufacturer</Form.Label>
                           <Form.Control required value={query.manufacturer} onChange={handleChange()} type="text" name="manufacturer" />
                        </Form.Group>
                     </Col>

                  </Row>
                  <br />


                  <center>
                     <Button variant="primary" type="submit">Add product</Button>
                  </center>
               </Form>
            </Container>
         </div >
      </>
   );
}

export default SellerHomePage;
