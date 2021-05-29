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
        const tasks = props.ws.Tasks.map(
            tsk => <TaskContainer key={tsk.id} task={tsk} />
        )

        if(state == 0){
        return (
            <div className='work-Style'>
                <CreateTaskButton workspaceId={props.ws.id} />
                {tasks}
                
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