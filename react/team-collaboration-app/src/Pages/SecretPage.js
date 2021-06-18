import React, { useEffect, useState } from 'react'
import '../MyStyles.css'
import { secretPage } from '../api/Login.js';
import 'primeflex/primeflex.css';

const Secret = ({history}) => {

    const [id, setId] = useState("")

    useEffect(async() =>{
        let res = await secretPage();
        console.log(res);
        setId(res)
    }, [])



    return (
        <div className="signup-page p-grid p-justify-center p-align-center">
            {id}
        </div>
    )
}

export default Secret;