import React from "react"
import '../MyStyles.css'
import TaskPage from "../Pages/TaskPage";
import ChatPage from "../Pages/ChatPage";
import FilePage from "../Pages/FilePage";
import StatPage from "../Pages/StatPage";

const WorkContainer = (props) => {

    if(props.tab == "tasks")
    {
        return (
            <TaskPage />
        )
    }

    if(props.tab == "chats")
    {
        return (
            <ChatPage />
        )
    }

    if(props.tab == "files")
    {
        return (
            <FilePage />
        )
    }

    if(props.tab == "stats")
    {
        return (
            <StatPage />
        )
    }

}

/* const WorkContainer = (props) => {

    const [state,setState] = useContext(stateContext)
       
    if (props.ws) {
        let tasks = props.ws.Tasks.map(
            tsk => <TaskContainer key={tsk.id} task={tsk} />
        )

        tasks = [
                ...tasks
        ]
        if(state == 0){
        return (
            <div className="createTask">
                <CreateTaskButton key={props.ws.id} workspaceId={props.ws.id} />
                <div className='work-Style'>
                    {tasks}
                    
                </div>
            </div>
        )
        }
        else{
            const taskPage = 
            <TaskPage task = {props.ws.Tasks} />

            return(
                <div className='work-Style'>
                    {taskPage} 
                </div>
            ) 
        }
    }
    else {
        return (
            <div className='work-Style'>
            </div>
        )
    }
    
   return <div>Work Container</div>
    
} */

export default WorkContainer;