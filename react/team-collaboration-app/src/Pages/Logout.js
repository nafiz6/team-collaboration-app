import React, { useEffect } from 'react'
import '../MyStyles.css'
import { logout } from '../api/Login.js';
import 'primeflex/primeflex.css';

const Logout = ({history}) => {


    useEffect(async() =>{
        await logout();
        history.push('/')
    }, [])



    return (
        <div className="signup-page p-grid p-justify-center p-align-center">
            Adios!
        </div>
    )
}

export default Logout;