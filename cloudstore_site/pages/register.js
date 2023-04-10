import React, { useState } from 'react';
import styles from '../styles/Login.module.css';

const Login = () => {
    const [query, setQuery] = useState({

    })
    return (
        <div className={styles.container}>
            <h1 style={{ paddingBottom: 10 }}>Register</h1>
            <form className={styles.formlogin}>
                <label className={styles.labellogin} htmlFor="name">Name</label>
                <input className={styles.inputlogin} type="name" id="name" name="name" required />

                <label className={styles.labellogin} htmlFor="password">Password</label>
                <input className={styles.inputlogin} type="password" id="password" name="password" required />
                <button className={styles.buttonlogin} type="submit">Login</button>
            </form>
        </div>
    );
};

export default Login;