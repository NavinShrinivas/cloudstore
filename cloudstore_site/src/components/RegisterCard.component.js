import axios from "axios";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

function RegisterCard() {
   const navigate = useNavigate();
   const [newUser, setNewUser] = useState({});
   const [errorMessage, setErrorMessage] = useState(null);
   const [successMessage, setSuccessMessage] = useState(null);
   const [passwordType, setPasswordType] = useState("password");
   const [confirmPasswordType, setConfirmPasswordType] = useState("password");

   const registerUser = async (e) => {
      e.preventDefault();

      const newUser = {
         name: e.target.name.value,
         email: e.target.email.value,
         username: e.target.username.value,
         password: e.target.password.value,
         confirmPassword: e.target.confirmPassword.value
      };

      if (newUser.username.includes(" ")) {
         setErrorMessage("Username cannot contain spaces.");
         return;
      }

      if (newUser.password !== newUser.confirmPassword) {
         setErrorMessage("Passwords do not match.");
         return;
      }

      setErrorMessage(null);
      setSuccessMessage(null);
      await axios.post('/api/account/register', newUser).then(res => {
         setErrorMessage(null);
         if (res.data.created) {
            setSuccessMessage(res.data.message + " Please login.");
            setTimeout(() => {
               navigate("/login");
            }, 1000);
         }
      }).catch(err => {
         console.log(err);
         if (err.response && err.response.data && err.response.data.message) {
            setErrorMessage(err.response.data.message);
         } else {
            setErrorMessage("An error occurred. Please try again.");
         }
      });
   }

   return (
      <div id="registerCard" className="text-black card shadow" style={{ width: "100%", minWidth: "300px", borderRadius: "30px", marginTop: "110px", marginBottom: "75px" }}>
         <div className="card-body">
            <h2 className="card-title  mb-3 fw-normal">Register</h2>
            <div className="card-text">
               <form onSubmit={registerUser}>
                  <div className="form-floating mb-1">
                     <input type="text" className="form-control" id="name" placeholder="name" required
                        value={newUser.name} onChange={(e) => setNewUser({ ...newUser, name: e.target.value })} />
                     <label htmlFor="name">Name</label>
                  </div>
                  <div className="form-floating mb-1">
                     <input type="email" className="form-control" id="email" placeholder="email" required
                        value={newUser.email} onChange={(e) => setNewUser({ ...newUser, email: e.target.value })} />
                     <label htmlFor="email">Email</label>
                  </div>
                  <div className="form-floating mb-1">
                     <input type="text" className="form-control" id="username" placeholder="username" required
                        value={newUser.username} onChange={(e) => setNewUser({ ...newUser, username: e.target.value })} />
                     <label htmlFor="name">Username</label>
                  </div>
                  <div className="form-floating input-group mb-1">
                     <input type={passwordType} className="form-control" id="password" required minLength="6" value={newUser.password} onChange={(e) => setNewUser({ ...newUser, password: e.target.value })} />
                     <label htmlFor="password">Password</label>
                     <button className="btn btn-outline-secondary" type="button" style={{ border: "1px solid #ced4da" }}
                        onClick={() => { (passwordType === "password") ? setPasswordType("text") : setPasswordType("password"); }} >
                        {passwordType === "password" ? <i className="bi bi-eye" /> : <i className="bi bi-eye-slash" />}
                     </button>
                  </div>
                  <div className="form-floating input-group mb-1">
                     <input type={confirmPasswordType} className="form-control" id="confirmPassword" required minLength="6" value={newUser.confirmPassword} onChange={(e) => setNewUser({ ...newUser, confirmPassword: e.target.value })} />
                     <label htmlFor="confirmPassword">Confirm Password</label>
                     <button className="btn btn-outline-secondary" type="button" style={{ border: "1px solid #ced4da" }}
                        onClick={() => { (confirmPasswordType === "password") ? setConfirmPasswordType("text") : setConfirmPasswordType("password"); }} >
                        {confirmPasswordType === "password" ? <i className="bi bi-eye" /> : <i className="bi bi-eye-slash" />}
                     </button>
                  </div>
                  {successMessage && <div className="alert alert-success" role="alert">{successMessage}</div>}
                  {errorMessage && <div className="alert alert-danger">{errorMessage}</div>}
                  <button className="w-100 btn btn-lg btn-dark" type="submit" >Create An Account</button>
               </form>
            </div>
         </div>


      </div>
   );
}

export default RegisterCard;
