import { useState } from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';
function OrderSidebar(props) {

    var value = props.order

    // console.log(value)

    return (
        <div className="position-fixed filter-sidebar m-2 p-3 card" style={{ height: '85vh', maxHeight: '85vh', width: '23%', overflowY: 'scroll' }}>
            {
                value.map(prod => {

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

                )
            }
            <center>
                <a href="/cart"><Button variant="primary">Cart</Button></a>
            </center>
        </div >

    )

}

export default OrderSidebar;