import { Link } from 'react-router-dom';

import RegisterCard from "../components/RegisterCard.component";

import CloudstoreLogo from '../static/cloudstore-logo.png';

function RegisterPage() {
    return (
        <>
            <div className="d-flex justify-content-center align-items-center text-white" style={{ backgroundColor: "#070e18", minHeight: "100vh", maxHeight: "100vh" }}>
                <Link to="/" className="top-0 position-absolute mt-3 start-50 translate-middle-x" >
                    <img src={CloudstoreLogo} alt="Cloudstore" style={{ minWidth: "100px", width: "10vw", maxWidth: "500px" }} />
                </Link>
                <div className="bottom-0" style={{ position: "absolute", width: "30vw" }}>
                    <RegisterCard />
                    <footer className="text-white text-center position-absolute bottom-0 start-50 translate-middle-x">
                        <div className="card text-white text-center start-50 translate-middle-x" style={{ borderRadius: "30px 30px 0px 0px", borderColor: "transparent", backgroundColor: "#000000", width: "100vw", fontSize: "0.7rem" }} >
                            &nbsp;This website uses cookies to store session data.&nbsp;
                            <br />
                            &nbsp;&nbsp;&nbsp;By continuing to use this website, you consent to our use of cookies.&nbsp;&nbsp;&nbsp;
                        </div>
                    </footer>
                </div>
            </div>
        </>
    );
}
export default RegisterPage;
