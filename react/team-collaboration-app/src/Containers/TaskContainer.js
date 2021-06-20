import axios from 'axios'
import React, { useEffect, useState } from 'react'
import Deadline from '../Components/Deadline'
import SubtaskButton from '../Components/SubtaskButton'
import '../MyStyles.css'

const TaskContainer = (props) => {

    const [subtasks, setSubtasks] = useState([])

    const getSubtasks = async () => {

        if (props.task.id) {
            let res = await axios.get(`http://localhost:8080/api/subtask/${props.task.id}`)
            setSubtasks(res.data)
        }

    }

    useEffect(() => {
        getSubtasks();
    }, [props.task.id])

    if (subtasks) {
        const subtasksArr = subtasks.map(
            stask => <SubtaskButton key={stask.id} name={stask.Name} />
        )

        return (
            <button className='taskContainer-Style'
                onClick=
                {
                    () => {
                    }
                }
            >
                <h3 className='taskName-Style'>{props.task.Name}</h3>
                <Deadline time={props.task.Deadline.split("T")[0]} />
                {subtasksArr}
            </button>
        )

    }
    else {

        return (
            <button className='taskContainer-Style'
            onClick=
            {
                () => {
                }
            }
        >
            <h3 className='taskName-Style'>{props.task.Name}</h3>
            <Deadline time={props.task.Deadline.split("T")[0]} />
        </button>  
        )

    }

}

export default TaskContainer;