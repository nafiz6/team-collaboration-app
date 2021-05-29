import React, { useContext, useState } from 'react'
import Deadline from '../Components/Deadline'
import SubtaskButton from '../Components/SubtaskButton'
import '../MyStyles.css'
import { stateContext, taskContext } from "../App"


const TaskContainer = (props) => {
    const [state, setState] = useContext(stateContext);
    const [task,setTask] = useContext(taskContext)

    if (props.task) {
        const subtasks = props.task.Subtasks.map(
            subtask => <SubtaskButton key={subtask.id} name={subtask.Name} />
        )

        return (
            <button className='taskContainer-Style'
                onClick=
                {
                    () => {
                        setState(1);
                        setTask(props.task);
                    }
                }
            >
                <p className='taskName-Style'>{props.task.Name}</p>
                <Deadline time={props.task.Deadline.split("T")[0]} />
                {subtasks}
            </button>
        )
    }
    else {
        <button className='taskContainer-Style'>
        </button>
    }

}

export default TaskContainer;