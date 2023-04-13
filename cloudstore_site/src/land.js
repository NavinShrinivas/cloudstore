import Button from 'react-bootstrap/Button';
import { useNavigate } from "react-router-dom";

function App() {
   const navigate = useNavigate();
   const toLogin = (event) => {
      navigate("/login")
   }

   const toRegister = (event) => {
      navigate("/register")
   }
   return (
      <div>
         <Button onClick={toLogin}> Login </Button>
         <Button onClick={toRegister}> Register </Button>
      </div>
   );
}

export default App;
