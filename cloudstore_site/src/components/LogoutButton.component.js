import axios from "axios";
import { useDispatch } from 'react-redux'
import { logout } from '../redux/features/userSlice'

function LogoutButton() {
    const dispatch = useDispatch()

    const logoutRequest = async (e) => {
        e.preventDefault();
        await axios.post("/api/account/logout", {}).then((response) => {
            dispatch(logout())
            window.location.href = "/"
        }).catch((error) => {
            console.log(error)
        })
    }

    return (
        <>
            <div className="btn btn-outline-light mx-2" onClick={logoutRequest}>Logout</div>
        </>
    );
}

export default LogoutButton;
