import { useState } from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';
import { useNavigate } from "react-router-dom";
import { useSelector, useDispatch } from 'react-redux'
import { updateCart, clearCart } from '../redux/features/userSlice'

function OrderSidebar(props) {
    const navigate = useNavigate();
    const dispatch = useDispatch()
    var value = props.order

    // console.log(value)
    const updateorder = () => {
        dispatch(updateCart(value))
        navigate("/cart")
    }

    return (
        <div className="position-fixed filter-sidebar m-2 p-3 card" style={{ height: '85vh', maxHeight: '85vh', width: '23%', overflowY: 'scroll' }}>
            {
                value.map(prod => {

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

                            <hr />
                        </Container>
                    )
                }

                )
            }
            <center>
                <Button onClick={updateorder} variant="primary">Cart</Button>
            </center>
        </div >

    )

}

export default OrderSidebar;