import React, { useContext, useState } from "react"
import { currWSContext } from "../App";
import CreateTaskButton from "../Components/CreateTask";
import '../MyStyles.css'
import TaskContainer from "./TaskContainer";
import TaskPage from "./TaskPage";
import {stateContext} from "../App"



const WorkContainer = (props) => {

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
    
}

export default WorkContainer;