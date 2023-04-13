import Button from 'react-bootstrap/Button';
import axios from 'axios';
import { useNavigate } from "react-router-dom";

import { useState } from 'react';
function App(props) {
   const url = "http://localhost:5001"
   const [username, setUsername] = useState("");
   const [password, setPassword] = useState("");
   const [isVisible, setIsVisiable] = useState(false);
   const [isloggedin, setIsloggedin] = useState(false);
   const [status, setStatus] = useState("")
   const navigate = useNavigate();
   const handleUsername = (event) => {
      console.log(username)
      setUsername(event.target.value)
   }

   const handlePassword = (event) => {
      setPassword(event.target.value)
   }
   axios.defaults.withCredentials = true //NOTE : This is very important to be able to set cookies 
   //Do check if setting default here applies to all axios resquests from here on
   const handleLogin = (event) => {
      axios.post(url + "/api/account/login", {
         username: username,
         password: password
      }).then(function(resp) {
         setIsVisiable(false)
         if (resp.data.status === true) {
            console.log("Login Successfull!")
            setStatus("Logged in! Redirecting...")
            setIsVisiable(true)
            navigate("/allproducts")
         } else {
            props.loggedin = false
            setStatus("Login failed, please chek your credentials!")
            setIsVisiable(true)
         }
         console.log(resp)
      }).catch(function(resp) {
         setStatus("Login failed, please chek your credentials!")
         setIsVisiable(true)
         console.log(resp)
      })
   }
   const checkLoggedIn = () => {
      axios.get(url + "/api/account/authcheck", {}).then(function(resp) {
         console.log(resp)
         if (resp.data.status) {
            setIsloggedin(true)
         } else {
            setIsloggedin(false)
         }
         console.log("test1")
      }).catch(function() {
         console.log("test")
         setIsloggedin(false)
      })
   }
   const logout = () => {
      axios.post(url + "/api/account/logout", {}).then(function(resp) {
         if (resp.data.status) {
            setIsloggedin(false)
            window.location.reload(true)
         } else {
            setIsloggedin(true)
         }
         console.log(resp.data.status)
      }).catch(function() {
         setIsloggedin(false)
      })
   }
   checkLoggedIn()
   return (
      <div >
         <form hidden={isloggedin}>
            Username : <input type="text" name="username" placeholder="Username" onChange={handleUsername} /><br />
            Password : <input type="password" name="username" placeholder="Username" onChange={handlePassword} /> <br />
            <p hidden={!isVisible}>{status}</p>
            <Button onClick={handleLogin}>Login!</Button>
         </form>
         <div hidden={!isloggedin}>
            <p> You are already logged in, to login as a different user click <a onClick={logout} style={{"text-decoration-line":"underline"}}> here</a>.</p>
         </div>
      </div>
   );
}

export default App;
