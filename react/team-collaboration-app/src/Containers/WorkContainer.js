import React from "react"
import '../MyStyles.css'
import TaskPage from "../Pages/TaskPage";
import ChatPage from "../Pages/ChatPage";
import FilePage from "../Pages/FilePage";
import StatPage from "../Pages/StatPage";

const WorkContainer = (props) => {

    if(props.tab === "tasks")
    {
        return (
            <TaskPage ws = {props.ws}/>
        )
    }

    if(props.tab === "chats")
    {
        return (
            <ChatPage />
        )
    }

    if(props.tab === "files")
    {
        return (
            <FilePage />
        )
    }

    if(props.tab === "stats")
    {
        return (
            <StatPage {...props} />
        )
    }

}

export default WorkContainer;