
import { HomeNavbar } from '../components'
import { Container, Row, Col, Image, Button } from 'react-bootstrap';
import { useSelector, useDispatch } from 'react-redux'
import { login, logout, cart } from '../redux/features/userSlice'

import axios from "axios";
import { useEffect } from 'react';

function CartPage() {
    const cart = useSelector((state) => state.cart.value)
    useEffect(() => {

    }, [])
    return (
        <>
            {/* <HomeNavbar />
            {value.map(prod => {

                return (
                    <Container>
                        <Row>
                            <Col>ID:{prod.ID}</Col>
                            <Col>{prod.name}</Col>
                        </Row>
                        <Row>
                            <Col>{prod.seller_username}</Col>
                            <Col>{prod.price} $</Col>
                        </Row>

                        <hr />
                    </Container>
                )
            }

            )} */}
        </>
    );
}
export default CartPage;