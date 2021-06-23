import React from "react"
import '../MyStyles.css'
import TaskPage from "../Pages/TaskPage";
import ChatPage from "../Pages/ChatPage";
import FilePage from "../Pages/FilePage";
import StatPage from "../Pages/StatPage";
import TaskDetailPage from "../Pages/TaskDetailPage";

const WorkContainer = (props) => {

    if(props.tab === "tasks")
    {
        return (
            <TaskPage ws = {props.ws}/>
        )
    }

    if(props.tab === "taskDetail")
    {
        return (
            <TaskDetailPage tid = {props.tid} taskname={props.taskname} deadline={props.deadline}
             description={props.description}/>
        )
    }

    if(props.tab === "chats")
    {
        return (
            <ChatPage {...props} ws={props.ws} />
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