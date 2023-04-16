
import { HomeNavbar } from '../components'
import { Container, Row, Col, Image, Button } from 'react-bootstrap';
import { useSelector, useDispatch } from 'react-redux'
import { login, logout, cart } from '../redux/features/userSlice'
import axios from "axios";

import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom';

function OrderPage() {
    const [orders, setorders] = useState([])
    let id = []
    const user = useSelector((state) => state.user.value)

    const navigate = useNavigate();

    const createOrder = async () => {
        axios.defaults.withCredentials = true
        // console.log(user)
        await axios.get("api/orders/all", {
            Name: user.name,
            Username: user.username,
            UserType: user.usertype
        }).then(async (resp) => setorders(resp.data.orders))
        console.log(orders)

    }
    useEffect(() => {
        createOrder()

    }, [])
    return (
        <>
            <HomeNavbar />
            <div style={{ marginTop: "100px" }}></div>

            <Container >
                <Row>
                    <Col><h4>Order Id</h4></Col>
                </Row>
            </Container>
            {orders.map(prod => {
                return (
                    <>
                        <Container key={prod.orderid}>

                            <Row>
                                <Col>{prod.orderid}</Col>
                                <Col>{prod.userid}</Col>
                                <Col>{prod.items.map(items => {
                                    return (<p>Product-{items.productid}</p>)

                                })}</Col>
                            </Row>


                            <hr />
                        </Container>
                    </>)
            }

            )
            }
            {/* <center>
                <Button onClick={createOrder}>Make order</Button>
            </center> */}
        </>
    );
}
export default OrderPage;