import React from "react"
import ChatButton from "../Components/ChatButton";
import FileButton from "../Components/FileButton";
import StatButton from "../Components/StatButton";
import TaskButton from "../Components/TaskButton";
import '../MyStyles.css'

const NavBar = (props) => 
{
    const navButtons = [
        <TaskButton id={props.id} wsid={props.wsid} {...props}/>,
        <ChatButton id={props.id} wsid={props.wsid} {...props}/>,
        <FileButton id={props.id} wsid={props.wsid} {...props}/>,
        <StatButton id={props.id} wsid={props.wsid} {...props}/>
    ]

    return(
        <div className='navBar-Style'>
            {navButtons}
        </div>
    );
}

export default NavBar;