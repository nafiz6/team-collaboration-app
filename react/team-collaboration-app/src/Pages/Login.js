import React, { useState, useEffect} from 'react'
import { Link } from 'react-router-dom'
import '../MyStyles.css'
import {Password} from 'primereact/password';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { login, secretPage } from '../api/Login.js';
import { Card } from 'primereact/card';
import 'primeflex/primeflex.css';

const Login = ({history}) => {
    const [user, setUser] = useState({
        username: '',
        password: '',
    });
    
    const [err, setErr] = useState('');
    
    const handleChange = e => {
            const { name, value } = e.target;
            setUser(prevState => ({
                ...prevState,
                [name]: value
            }));
        };

    const loginUser = async () => {
        try{
            let res = await login(user)
            history.push("/project")
        }
        catch(err){
            setErr(err);
        }
        //history.push('/tasks')
    }

    useEffect(async () => {
        try{
            let res = await secretPage(user)
            if (res.toLowerCase().includes('hello')){
                history.push("/project")
            }
        }
        catch(err){
            //setErr(err);
            console.log("login verify error " + err)
        }
    }, [])

    return (
        <div className="signup-page p-grid p-justify-center p-align-center">
            <div className="p-col-4">
                <Card className="signup-card">
                        <h2> Login </h2>
                    <h5>Username</h5>
                    <InputText name="username" value={user.username} onChange={handleChange}/>
                    <h5>Password</h5>
                    <Password feedback={false} name="password" value={user.password} onChange={handleChange}/>
                    <br/>
                    <a className="err">{err} </a>
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