import React from "react"
import ManageButton from "../Components/ManageButton"
import NotifyButton from "../Components/NotifyButton"
import LogOutButton from "../Components/LogOutButton"
import '../MyStyles.css'


const HeaderContainer = () => 
{
    return (
        <div className='header-Style'>
            <NotifyButton/>
            <ManageButton/>
            <LogOutButton/>
        </div>
    )

}

export default HeaderContainer;