import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import Container from 'react-bootstrap/Container';

import CloudstoreLogo from '../static/cloudstore-logo.png';

function LandNavbar() {
    return (
        <Navbar collapseOnSelect expand="lg" bg="black" variant="dark" fixed="top">
            <Container fluid>
                <Navbar.Brand href="#home" className="d-flex align-items-center">
                    <Navbar.Collapse id="responsive-navbar-nav">
                        <img
                            src={CloudstoreLogo}
                            height="30"
                            className="d-inline-block align-top d-none d-lg-block"
                            alt="Cloudstore logo"
                        />
                    </Navbar.Collapse>
                    <div className="d-none d-lg-block">
                        Cloudstore
                    </div>
                </Navbar.Brand>
                <Navbar.Toggle aria-controls="responsive-navbar-nav" />
                <Navbar.Collapse id="responsive-navbar-nav">
                    <Nav >
                        <Nav.Link href="#about">About</Nav.Link>
                    </Nav>
                </Navbar.Collapse>
                <Navbar.Collapse id="responsive-navbar-nav">
                    <Nav className="me-auto">
                    </Nav>
                    <div className="d-flex justify-content-around">
                        <Nav.Link href="/login">
                            <button className="btn btn-outline-light ms-3" >Login</button>
                        </Nav.Link>
                        <Nav.Link href="/register">
                            <button className="btn btn-outline-light mx-3" >Register</button>
                        </Nav.Link>
                    </div>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default LandNavbar;
