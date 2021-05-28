import React from "react"
import ChatButton from "../Components/ChatButton";
import FileButton from "../Components/FileButton";
import StatButton from "../Components/StatButton";
import TaskButton from "../Components/TaskButton";
import '../MyStyles.css'

const NavBar = () => 
{
    const navButtons = [
        <TaskButton/>,
        <ChatButton/>,
        <FileButton/>,
        <StatButton/>
    ]

    return(
        <div className='navBar-Style'>
            {navButtons}
        </div>
    );
}

export default NavBar;