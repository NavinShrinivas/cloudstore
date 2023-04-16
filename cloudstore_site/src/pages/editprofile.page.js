import { Form, Button } from 'react-bootstrap';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { useState } from 'react';
import axios from 'axios';

import { LogoutButton } from "../components";

import CloudstoreLogo from '../static/cloudstore-logo.png';

function EditProfile() {
    const [query, setQuery] = useState({
        username: "",
        password: "",
        newusername: "",
        newpassword: "",
        newname: "",
        newemail: "",
        newphone: "",
        newaddress: "",
        newusertype: ""
    });

    const handleChange = () => (e) => {
        const name = e.target.name;
        const value = e.target.value;
        setQuery((prevState) => ({
            ...prevState,
            [name]: value
        }));
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        const formData = new FormData();
        Object.entries(query).forEach(([key, value]) => {
            formData.append(key, value);
        });
        const jsonObject = {};

        for (const [key, value] of formData) {
            jsonObject[key] = value;
        }


        // console.log(query)
        axios.put("/api/account/update", query).then(resp => console.log(resp))


    };

    return (
        <>
            <div className="vh-100 d-flex justify-content-center align-items-center text-white" style={{ backgroundColor: "#070e18" }}>

                <Navbar collapseOnSelect expand="lg" bg="black" variant="dark" fixed="top">
                    <Container fluid>
                        <Navbar.Brand href="/" className="d-flex align-items-center">
                            <Navbar.Collapse id="responsive-navbar-nav">
                                <img
                                    src={CloudstoreLogo}
                                    height="30"
                                    className="mx-3 d-inline-block align-top d-none d-lg-block"
                                    alt="Cloudstore logo"
                                />
                            </Navbar.Collapse>
                            <div className="d-none d-lg-block">
                                <h3 className="m-0">Cloudstore</h3>
                            </div>
                        </Navbar.Brand>

                        <Navbar.Collapse id="responsive-navbar-nav">
                            <Nav className="me-auto">
                            </Nav>
                            <Nav.Link href="/profile">
                                <div className="btn btn-outline-light mx-2" >Profile</div>
                            </Nav.Link>
                            <Nav.Link href="/">
                                <div className="btn btn-outline-light mx-2" >Home</div>
                            </Nav.Link>
                            <div className="d-flex justify-content-around">
                                <LogoutButton />
                            </div>
                        </Navbar.Collapse>
                    </Container>
                </Navbar>
                <Container style={{ marginTop: "50px" }}>
                    <Form onSubmit={handleSubmit}>
                        <h2>Enter current details</h2>

                        <Row>
                            <Col>
                                <Form.Group controlId="formBasicUserName">
                                    <Form.Label>Enter UserName</Form.Label>
                                    <Form.Control required value={query.username} onChange={handleChange()} type="text" name="username" />
                                </Form.Group>
                            </Col>

                            <Col>
                                <Form.Group controlId="formBasicName">
                                    <Form.Label>Enter Password</Form.Label>
                                    <Form.Control required value={query.password} onChange={handleChange()} type="text" name="password" />
                                </Form.Group>
                            </Col>
                        </Row>
                        <hr />
                        <h2>Enter new details</h2>

                        <Row>
                            <Col>
                                <Form.Group controlId="formBasicUserName">
                                    <Form.Label>Enter new UserName</Form.Label>
                                    <Form.Control required value={query.newusername} onChange={handleChange()} type="text" name="newusername" />
                                </Form.Group>
                            </Col>
                            <Col>
                                <Form.Group controlId="formBasicName">
                                    <Form.Label>Enter new Name</Form.Label>
                                    <Form.Control required value={query.newname} onChange={handleChange()} type="text" name="newname" />
                                </Form.Group>
                            </Col>
                            <Col>
                                <Form.Group controlId="formBasicName">
                                    <Form.Label>Enter new Number</Form.Label>
                                    <Form.Control required value={query.newphone} onChange={handleChange()} type="text" name="newphone" />
                                </Form.Group>
                            </Col>
                        </Row>
                        <br />
                        <Row>
                            <Col>

                                <Form.Group controlId="formBasicEmail">
                                    <Form.Label> Enter new Email</Form.Label>
                                    <Form.Control required value={query.newemail} onChange={handleChange()} type="email" name="newemail" />
                                </Form.Group>
                            </Col>
                            <Col>
                                <Form.Group controlId="formBasicPassword">
                                    <Form.Label> Enter new Password</Form.Label>
                                    <Form.Control required value={query.newpassword} onChange={handleChange()} type="password" name="newpassword" />
                                </Form.Group>
                            </Col>
                            <Col>
                                <Form.Group controlId="formBasictype">
                                    <Form.Label> Enter new UserType</Form.Label>
                                    <Form.Control required value={query.newusertype} onChange={handleChange()} type="password" name="newusertype" />
                                </Form.Group>
                            </Col>
                        </Row>
                        <br />

                        <Form.Group controlId="formBasicAddress">
                            <Form.Label>Address</Form.Label>
                            <Form.Control required value={query.newaddress} onChange={handleChange()} as="textarea" rows={3} name="newaddress" />
                        </Form.Group>
                        <hr />

                        <center>
                            <Button variant="primary" type="submit">
                                Save Changes
                            </Button>
                        </center>
                    </Form>
                </Container>
            </div>

        </>
    );
}

export default EditProfile