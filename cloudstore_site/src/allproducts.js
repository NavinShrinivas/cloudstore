import axios from 'axios';
import { useState } from 'react';
function App(props) {
   const url = "http://localhost:5001"
   const [isloggedin, setIsloggedin] = useState(false)
   axios.defaults.withCredentials = true //NOTE : This is very important to be able to set cookies 
   //Do check if setting default here applies to all axios resquests from here on
   const handleLogin = () => {
      axios.get(url + "/api/account/authcheck", {}).then(function(resp) {
         if (resp.data.status) {
            setIsloggedin(true)
         } else {
            setIsloggedin(false)
         }
         console.log(resp.data.status)
      }).catch(function() {
         setIsloggedin(false)
      })
   }
   return (
      < div onLoad={handleLogin()}>
         <p hidden={!isloggedin}> To do</p>
         <p hidden={isloggedin}> Go login dumass</p>
      </div >
   );
}

export default App;
