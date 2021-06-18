import React from 'react'
import { Link } from 'react-router-dom'
import '../MyStyles.css'

const Login = () => {
    return (
        <div>
            <div>Name</div>
            <input type = "text"/>
            <div>Password</div>
            <input type = "password" />
            <Link to = "/project">  
            <button>Login</button>
            </Link>
            <Link to = "/signup">
            <button>Sign Up</button>
            </Link>
        </div>
    )
}

export default Login;