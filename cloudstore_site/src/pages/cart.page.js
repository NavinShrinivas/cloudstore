
import { HomeNavbar } from '../components'
import { Container, Row, Col, Image, Button } from 'react-bootstrap';
import { useSelector, useDispatch } from 'react-redux'
import { login, logout, cart } from '../redux/features/userSlice'
import axios from "axios";
import { useEffect } from 'react';

function CartPage() {
    const cart = useSelector((state) => state.user.cart)
    const newArr = [...cart];

    const senddetails = () => {

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
                <Button onClick={() => { }}>Make order</Button>
            </center>
        </>
    );
}
export default CartPage;