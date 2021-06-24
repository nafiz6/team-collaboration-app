import axios from 'axios'
import React, { useEffect, useState } from 'react'
import Deadline from '../Components/Deadline'
import SubtaskButton from '../Components/SubtaskButton'
import '../MyStyles.css'
import { Link } from 'react-router-dom'

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

    const [retAddr, setRetAddr] = useState(window.location.href);
    useEffect(() => {
        let addrArr = (window.location.href).split("http://localhost:3000");
        setRetAddr(addrArr[1]);

    }, [window.location.href])



    if (subtasks) {
        const subtasksArr = subtasks.map(
            stask => <SubtaskButton key={stask.id} name={stask.Name} />
        )

        return (
            <Link to={
                {
                    pathname: `${retAddr}/taskpage/${props.task.id}`,
                    state:
                    {
                        taskname: props.task.Name,
                        deadline: props.task.Deadline,
                        description: props.task.Description
                    }
                }}  >
                <button className='taskContainer-Style'>
                    <h3 className='taskName-Style'>{props.task.Name}</h3>
                    <Deadline time={props.task.Deadline.split("T")[0]} />
                    {subtasksArr}
                </button>
            </Link>
        )

    }
    else {

        return (

            <Link to={
                {
                    pathname: `${retAddr}/taskpage/${props.task.id}`,
                    state:
                    {
                        taskname: props.task.Name,
                        deadline: props.task.Deadline,
                        description: props.task.Description
                    }
                }}  >
                <button className='taskContainer-Style'>
                    <h3 className='taskName-Style'>{props.task.Name}</h3>
                    <Deadline time={props.task.Deadline.split("T")[0]} />
                </button>
            </Link>
        )

    }

}

export default TaskContainer;