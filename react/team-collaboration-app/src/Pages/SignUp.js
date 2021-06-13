import React from 'react'
import { Link } from 'react-router-dom'
import '../MyStyles.css'

const SignUp = () => 
{
    return (
        <div>
            <div>Sign Up Page</div>
            <Link to = "/" >
            <button>Sign Up</button>
            </Link>
        </div>
    )
}

export default SignUp;