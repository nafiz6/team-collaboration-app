import React from "react"
import '../MyStyles.css'
import { Link } from "react-router-dom";
import { logout } from '../api/Login.js';
import { Button } from 'primereact/button';

const LogOutButton = ({ history }) => {

    const logoutButton = async () => {
        history.push('/logout')
    }


    return (
        <div className="button-Style">
            <Button onClick={logoutButton} className='p-button-danger p-button-sm' >Log Out</Button>
        </div>


    );
}

export default LogOutButton;