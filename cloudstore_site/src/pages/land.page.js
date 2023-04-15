import Nav from 'react-bootstrap/Nav';

import { LandNavbar } from '../components/'
import CloudstoreLogo from '../static/cloudstore-logo.png';

function LandPage() {
    return (
        <>
            <LandNavbar />
            <div style={{ maxWidth: "100vw" }}>
                <section id="home" className="d-flex justify-content-center align-items-center" style={{ height: '100vh', backgroundColor: 'grey' }}>
                    <div className="d-flex flex-column align-items-center">
                        <div className="d-flex flex-column align-items-center" style={{ width: '80%' }}>
                            <img src={CloudstoreLogo} alt="Cloudstore" style={{ width: '100%' }} />
                            <div className="text-white text-center my-3" style={{ fontSize: '1.5rem' }}>
                                Cloudstore is a ...
                            </div>
                        </div>
                        <div className="d-flex justify-content-around">
                            <Nav.Link href="#about" className="text-white text-center" >
                                <button className="btn btn-outline-light">LEARN MORE</button>
                            </Nav.Link>
                        </div>
                    </div>
                </section>

                <section id="about" className="d-flex" style={{ minHeight: '100vh', backgroundColor: '#f5f5f5' }}>
                    <div className="ms-4 me-3 my-5">
                        <div className="mt-5" style={{ color: '#003f5c', fontSize: '2rem', fontWeight: 'bold', fontFamily: 'unset' }}>
                            <b><u>About Cloudstore</u></b>
                        </div>
                        <div className="d-flex flex-column">
                            <div className="p" style={{ fontSize: '1.1rem', textAlign: 'justify' }}>
                                Very cool website
                            </div>
                        </div>

                    </div>
                </section>

                <div className="text-white text-center bottom-0 bg-black" style={{ maxWidth: "100vw" }} >
                    Copyright © Cloudstore 2023
                </div>
            </div>
        </>
    );
}

export default LandPage;
