import React from "react"
import ManageButton from "../Components/ManageButton"
import NotifyButton from "../Components/NotifyButton"
import LogOutButton from "../Components/LogOutButton"
import '../MyStyles.css'


const HeaderContainer = (props) => 
{
    return (
        <div className='header-Style'>
          {/* <NotifyButton/> */}
            <LogOutButton {...props}/>
        </div>
    )

}

export default HeaderContainer;