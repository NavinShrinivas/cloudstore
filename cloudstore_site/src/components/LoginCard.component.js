import axios from "axios";
import { useState } from "react";
import { useDispatch } from 'react-redux'
import { useNavigate } from "react-router-dom";
import { login, logout } from '../redux/features/userSlice'

function LoginCard() {
    const navigate = useNavigate();
    const [errorMessage, setErrorMessage] = useState(null);
    const [passwordType, setPasswordType] = useState("password");
    const [loggedin, setloggedinUser] = useState({ username: "", password: "" });

    const dispatch = useDispatch()

    const loginRequest = async (e) => {
        e.preventDefault();

        axios.defaults.withCredentials = true //NOTE : This is very important to be able to set cookies 
        await axios.post("/api/account/login", {
            username: loggedin.username,
            password: loggedin.password
        }).then(async (res) => {
            if (res.data.status) {
                dispatch(login(res.data.user));
                setErrorMessage(null);
                setTimeout(() => {
                    navigate("/");
                }, 1000);
            } else {
                dispatch(logout());
                setErrorMessage("An error occurred. Please try again.");
            }
        }).catch(err => {
            dispatch(logout());
            console.log(err);
            if (err.response && err.response.data && err.response.data.message) {
                setErrorMessage(err.response.data.message);
            } else {
                setErrorMessage("An error occurred. Please try again.");
            }
        });
    }

    return (
        <div id="loginCard" className="text-black card shadow position-absolute top-50 start-50 translate-middle ml-5" style={{ width: "30%", minWidth: "300px", borderRadius: "30px" }}>
            <div className="card-body">
                <h2 className="card-title mb-1 fw-normal">Please Login</h2>
                <div className="card-text">
                    <form onSubmit={loginRequest}>
                        <br />
                        <div className="form-floating">
                            <input type="text" className="form-control" id="username" required placeholder="Username" value={loggedin.username} onChange={(e) => setloggedinUser({ ...loggedin, username: e.target.value })} />
                            <label htmlFor="usernameEmail">Username</label>
                        </div>
                        <div className="form-floating input-group mb-2">
                            <input type={passwordType} className="form-control" id="password" required minLength="4" value={loggedin.password} onChange={(e) => setloggedinUser({ ...loggedin, password: e.target.value })} />
                            <label htmlFor="password">Password</label>
                            <button className="btn btn-outline-secondary" type="button" style={{ border: "1px solid #ced4da" }}
                                onClick={() => { (passwordType === "password") ? setPasswordType("text") : setPasswordType("password"); }} >
                                {passwordType === "password" ? <i className="bi bi-eye" /> : <i className="bi bi-eye-slash" />}
                            </button>
                        </div>
                        {errorMessage && <div className="alert alert-danger" role="alert">{errorMessage}</div>}
                        <button className="w-100 btn btn-lg btn-dark" type="submit">Login</button>
                    </form>
                </div>
            </div>
        </div>
    );
}

export default LoginCard;
