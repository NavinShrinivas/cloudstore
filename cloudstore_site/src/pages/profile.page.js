
import { HomeNavbar } from '../components'
import { Container, Row, Col, Image, Button } from 'react-bootstrap';
import { useSelector, useDispatch } from 'react-redux'

import axios from "axios";
import { useEffect } from 'react';

function ProfilePage() {
    const user = useSelector((state) => state.user.value)

    console.log(user)
    useEffect(() => {

    }, [])
    return (
        <>
            <HomeNavbar />
            <div className="vh-100 d-flex justify-content-center align-items-center text-white" style={{ backgroundColor: "#070e18" }}>
                <div id="loginCard" className="text-black card shadow position-absolute top-50 start-50 translate-middle ml-5" style={{ width: "30%", minWidth: "300px", borderRadius: "30px" }}>
                    <Container style={{ padding: 20 }}>
                        <Row >


                            <h1 style={{
                                display: 'flex', justifyContent: 'center'
                            }}>{user.username}</h1>
                            <p style={{
                                display: 'flex', justifyContent: 'center'
                            }}>{user.email}</p>


                        </Row>
                        <Row>
                            <Col sm={12}>
                                <Row>
                                    <h6>Phone number: {user.phone}</h6>
                                    <h6>Name: {user.name}</h6>
                                    <h6>Type: {user.usertype}</h6>

                                </Row>

                            </Col>
                        </Row>
                        <Row>
                            {/* <Button variant="primary">Edit Profile</Button> */}
                        </Row>
                    </Container>
                </div>
            </div>
        </>
    );
}
export default ProfilePage;