import React, { useState } from 'react'
import { Link } from 'react-router-dom'
import '../MyStyles.css'
import {Password} from 'primereact/password';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { login } from '../api/Login.js';
import { Card } from 'primereact/card';
import 'primeflex/primeflex.css';

const Login = ({history}) => {
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
        //history.push('/tasks')
    }

    return (
        <div className="signup-page p-grid p-justify-center p-align-center">
            <div className="p-col-4">
                <Card className="signup-card">
                    <h5>Username</h5>
                    <InputText name="username" value={user.username} onChange={handleChange}/>
                    <h5>Password</h5>
                    <Password feedback={false} name="password" value={user.password} onChange={handleChange}/>
                    <br/>
                    <Button label="Login" onClick={() => 
                        loginUser()
                    } />
                    <br/>
                    <h5>
                        Don't have an account?
                    </h5>
                    <Link to="/signup">
                        Sign Up
                    </Link>
            </Card>
            </div>
        </div>
    )
}

export default Login;