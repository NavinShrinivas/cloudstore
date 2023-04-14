import { Link } from 'react-router-dom';

import LoginCard from "../components/LoginCard.component";

import CloudstoreLogo from '../static/cloudstore-logo.png';

function LoginPage() {
    return (
        <>
            <div className="vh-100 d-flex justify-content-center align-items-center text-white" style={{ backgroundColor: "#070e18" }}>
                <Link to="/" className="top-0 position-absolute mt-5" >
                    <img src={CloudstoreLogo} alt="Cloudstore" style={{ minWidth: "100px", width: "10vw", maxWidth: "500px" }} />
                </Link>
                <LoginCard />
                <div className="position-absolute bottom-0 w-100 text-center mb-5">
                    <p className="mb-0">Don't have an account?</p>
                    <Link to="/register" className="text-decoration-none text-blue">Create an account!</Link>
                </div>
                <footer className="text-white text-center position-absolute bottom-0 start-50 translate-middle-x">
                    <div className="card text-white text-center start-50 translate-middle-x" style={{ borderRadius: "30px 30px 0px 0px", borderColor: "transparent", backgroundColor: "#000000", width: "100vw", fontSize: "0.7rem" }} >
                        &nbsp;This website uses cookies to store session data.&nbsp;
                        <br />
                        &nbsp;&nbsp;&nbsp;By continuing to use this website, you consent to our use of cookies.&nbsp;&nbsp;&nbsp;
                    </div>
                </footer>
            </div>
        </>
    );
}
export default LoginPage;
