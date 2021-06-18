import React from "react"
import '../MyStyles.css'
import { Link } from "react-router-dom";
import { logout } from '../api/Login.js';

const LogOutButton = () => 
{

    const logoutButton = async () =>{
        await logout();
    }


    return (
        <Link to="/">
            <button onClick={logoutButton} className='button-Style'>Log Out</button>
        </Link>
        
    );
}

export default LogOutButton;