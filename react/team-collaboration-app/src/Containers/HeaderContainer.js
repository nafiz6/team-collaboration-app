import React, { useState } from "react"
import ManageButton from "../Components/ManageButton"
import NotifyButton from "../Components/NotifyButton"
import LogOutButton from "../Components/LogOutButton"
import '../MyStyles.css'
import { Avatar } from "primereact/avatar"
import axios from "axios"


const HeaderContainer = (props) => 
{

    const [me, setMe] = useState(null);

    const fetchMyDetails = async () => {

      
        let res = await axios.get(`http://localhost:8080/api/my-details`, { withCredentials: true });


        setMe(res.data);

    }


    return (
        <div className='header-Style'>
          {/* <NotifyButton/> */}
            <LogOutButton {...props}/>
            <div>
                <Avatar />
            </div>
        </div>
    )

}

export default HeaderContainer;