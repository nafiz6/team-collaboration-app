import React, { useEffect, useState } from "react"
import ManageButton from "../Components/ManageButton"
import NotifyButton from "../Components/NotifyButton"
import LogOutButton from "../Components/LogOutButton"
import '../MyStyles.css'
import { Avatar } from "primereact/avatar"
import axios from "axios"


const HeaderContainer = (props) => {

    const [userDetails, setUserDetails] = useState(null);

    const fetchMyDetails = async () => {


        let res = await axios.get(`http://localhost:8080/api/my-details`, { withCredentials: true });


        setUserDetails(res.data);

    }

    useEffect(() => {
        fetchMyDetails();

    }, [])




    return (
        <div className='header-Style'>
            {/* <NotifyButton/> */}
            <div class="header-user-details">
                <Avatar label={userDetails?.Dp ? null : userDetails?.Name[0]} image={userDetails?.Dp ? userDetails?.Dp : null} shape="circle" size="large" />
                <h3 className="header-user-name">{userDetails?.Name}</h3>
            </div>
            <LogOutButton {...props} />

        </div>
    )

}

export default HeaderContainer;