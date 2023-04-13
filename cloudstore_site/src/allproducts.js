import axios from 'axios';
import { useState } from 'react';

function App(props) {

   // Maybe use redux to store the state of the user, may be easier to manage 

   const url = "http://localhost:5001"
   const [isloggedin, setIsloggedin] = useState(null)
   axios.defaults.withCredentials = true //NOTE : This is very important to be able to set cookies 
   //Do check if setting default here applies to all axios resquests from here on
   const handleLogin = () => {
      axios.get(url + "/api/account/authcheck", {}).then(function (resp) {
         if (resp.data.status) {
            setIsloggedin(true)
         } else {
            setIsloggedin(false)
         }
         console.log(resp.data.status)
      }).catch(function () {
         setIsloggedin(false)
      })
   }
   return (
      < div onLoad={handleLogin()}>
         {isloggedin === null ? <p> Loading </p> :
            <>
               <p hidden={!isloggedin}> To do</p>
               <p hidden={isloggedin}>
                  {/* redirect to login page */}
                  Go login dumass
               </p>
            </>
         }
      </div >
   );
}

export default App;
