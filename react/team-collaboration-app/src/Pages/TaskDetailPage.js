import React, { useEffect, useState } from 'react'
import '../MyStyles.css'
import SubtaskContainer from '../Containers/SubtaskContainer'
import TaskFiles from '../Containers/TaskFiles'
import axios from 'axios'
import CreateSubtaskButton from '../Components/SubtaskButton'

const TaskDetailPage = (props) => {
    const [subtasks, setSubtasks] = useState([]);

    const getSubtasks = async () => {

        if (props.tid) {
            let res = await axios.get(`http://localhost:8080/api/subtask/${props.tid}`);
            setSubtasks(res.data);
        }
    }

    useEffect(() => {
        getSubtasks();
    }, [props.tid])

    if (subtasks) {
        const staskArr = subtasks.map(
            subtask => <SubtaskContainer key={subtask.id} subtask={subtask} />)

        return (
            <div>
                <TaskFiles taskId={props.tid} />
                <div className="taskPage-Style">
                    {/*<CreateSubtaskButton taskId={props.tid} />*/}
                    <h3>{props.taskname}</h3>
                    <h4>Deadline: {props.deadline.split("T")[0]}</h4>
                    <a>{props.description}</a>
                    {staskArr}
                </div>
            </div>
        )
    }
    else {

        return (
            <div className="taskPage-Style">
                {/*<CreateSubtaskButton taskId={props.tid} />*/}
                <h3>{props.taskname}</h3>
                <a>{props.deadline.split("T")[0]}</a>
                <h5>{props.description}</h5>
            </div>
        )

    }
}

export default TaskDetailPage;