import React, { useState, useEffect } from 'react'
import CreateTaskButton from '../Components/CreateTask'
import TaskContainer from '../Containers/TaskContainer'
import '../MyStyles.css'
import axios from 'axios'
import { Skeleton } from 'primereact/skeleton';

const TaskPage = (props) => {

    // console.log(props);

    const [tasks, setTasks] = useState([]);
    const [loading, setLoading] = useState(false);

    const getTasks = async () => {
        setLoading(true);

        if (props.ws) {
            setLoading(true);
            let res = await axios.get(`http://localhost:8080/api/task/${props.ws}`);
            setTasks(res.data)
            setLoading(false);
        }
    }

    useEffect(() => {
        getTasks();
    }, [props.ws])

    let taskArr = null;
    if (tasks) //Tasks Present
    {
        taskArr = tasks.map(
            tsk => <TaskContainer key={tsk.id} task={tsk} ws={props.ws} />
        )
    }


    return (
        <div className="createTask">
            <h1>Workspace Tasks</h1>
            <CreateTaskButton key={props.ws} workspaceId={props.ws} />
            <div className='work-Style'>
                {loading ? <Skeleton borderRadius="16px" height="500px" width="300px" /> : (tasks.length > 0 ? taskArr :
                    <div className="centered">
                       <i className="pi pi-check-circle" style={{'fontSize': '40px'}} />
                        <h2>No Tasks!</h2>
                    </div>
                )}

            </div>
        </div>
    )
}



export default TaskPage;