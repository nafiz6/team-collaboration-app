import React from "react"
import ManageButton from "../Components/ManageButton"
import NotifyButton from "../Components/NotifyButton"
import '../MyStyles.css'


const HeaderContainer = () => 
{
    return (
        <div className='header-Style'>
            <NotifyButton/>
            <ManageButton/>
        </div>
    )

}

export default HeaderContainer;