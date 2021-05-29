import React, { useContext, useEffect } from 'react'
import CreateSubtaskButton from '../Components/CreateSubtask'
import '../MyStyles.css'
import SubtaskPage from './SubtaskPage'
import { taskContext } from '../App'



const TaskPage = (props) => 
{
    const [task,setTask] = useContext(taskContext)

    console.log(task)

    if(task){

    const subtasks = task.Subtasks.map(
        subtask => <SubtaskPage key = {subtask.id} subtask = {subtask}/>
    )

    return (
        <div className="taskPage-Style">
            <CreateSubtaskButton taskId={task.id}/>
            <text>Task: {task.Name}</text>
            <text>{task.Deadline.split("T")[0]}</text>
            {subtasks}
        </div>
    )
    }
    else{
        return (
            <div className="taskPage-Style">
               
            </div>
        )

    }
}

export default TaskPage;