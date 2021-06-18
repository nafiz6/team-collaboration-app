import React, { useCallback, useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import axios from 'axios';
import '../MyStyles.css'
import {Password} from 'primereact/password';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { login } from '../api/Login.js';

const Login = () => {
    const [user, setUser] = useState({
        username: '',
        password: '',
    });
    
    const handleChange = e => {
            const { name, value } = e.target;
            setUser(prevState => ({
                ...prevState,
                [name]: value
            }));
        };

    const loginUser = async () => {
        let res = await login(user)
        console.log(res);
    }

    return (
        <div>
            <h5>Name</h5>
            <InputText name="username" value={user.username} onChange={handleChange}/>
            <div>Password</div>
            <InputText name="password" value={user.password} onChange={handleChange}/>
                <Button label="Login" onClick={() => 
                    loginUser()
                } />
            <Link to="/signup">
                <button>Sign Up</button>
            </Link>
        </div>
    )
}

export default Login;