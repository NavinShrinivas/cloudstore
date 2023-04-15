import React from 'react'
import axios from "axios";
import { useEffect } from "react";
import { useSelector, useDispatch } from 'react-redux'
import { BrowserRouter as Router, Routes, Route, Navigate, Outlet } from "react-router-dom";

import { Loading } from "./components";
import { login, logout } from './redux/features/userSlice'
import { HomePage, LandPage, LoginPage, RegisterPage, ProfilePage } from "./pages";

const ProtectedRoute = () => {
    const loggedIn = useSelector((state) => state.user.loggedIn)
    if (loggedIn == null) {
        return <Loading />
    }
    return loggedIn ? <Outlet /> : <Navigate to="/login" />;
};

function App() {
    const user = useSelector((state) => state.user.value)
    const dispatch = useDispatch()

    const checkLoginStatus = async () => {
        try {
            axios.defaults.withCredentials = true
            //NOTE : This is very important to be able to set cookies
            const response = await axios.get("/api/account/authcheck")
            if (response.data && response.data.status) {
                const user = await axios.get("/api/account/info")
                if (user.data && user.data.status) {
                    dispatch(login(user.data.record))
                } else {
                    dispatch(logout())
                }
            } else {
                dispatch(logout())
            }
        } catch (error) {
            console.log(error)
        }
    }

    useEffect(() => {
        checkLoginStatus()
        // eslint-disable-next-line
    }, [])

    return (
        <Router>
            <Routes>
                {user ? <Route path="/" element={<HomePage />} /> : <Route path="/" element={<LandPage />} />}
                <Route path="login" element={<LoginPage />} />
                <Route path="register" element={<RegisterPage />} />

                <Route element={<ProtectedRoute />}>

                    <Route exact path="/profile" element={<ProfilePage />} />
                </Route>

                <Route path="*" element={<Navigate to="/" />} />

            </Routes>
        </Router>
    );
}

export default App;
