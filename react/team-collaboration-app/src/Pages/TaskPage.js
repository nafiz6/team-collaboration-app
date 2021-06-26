import React, { useState, useEffect } from 'react'
import CreateTaskButton from '../Components/CreateTask'
import TaskContainer from '../Containers/TaskContainer'
import '../MyStyles.css'
import axios from 'axios'

const TaskPage = (props) => {

    // console.log(props);

    const [tasks,setTasks] = useState([]);

    const getTasks = async () => {

        if(props.ws){
        let res = await axios.get(`http://localhost:8080/api/task/${props.ws}`);
        setTasks(res.data)
        }
    }

    useEffect(() => {
        getTasks();
    },[props.ws])


    if(tasks) //Tasks Present
    {
        let taskArr = tasks.map(
            tsk => <TaskContainer key={tsk.id} task={tsk} ws={props.ws} />
        )
        

        return (
            <div className="createTask">
                <CreateTaskButton key={props.ws} workspaceId={props.ws} />
                <div className='work-Style'>
                    {taskArr}
                </div>
            </div>
        )
    }
    else{  //No Tasks present

        return (
            <div className="createTask">
                <CreateTaskButton key={props.ws} workspaceId={props.ws} />
                <div className='work-Style'>
                </div>
            </div>
        )
  
    }

}

export default TaskPage;