import React from "react"
import '../MyStyles.css'
import { Link } from "react-router-dom";
import { logout } from '../api/Login.js';

const LogOutButton = ({history}) => 
{

    const logoutButton = async () =>{
        history.push('/logout')
    }


    return (
        <button onClick={logoutButton} className='button-Style'>Log Out</button>
        
    );
}

export default LogOutButton;