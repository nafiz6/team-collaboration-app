import React, { useState, useEffect } from 'react'
import CreateTaskButton from '../Components/CreateTask'
import TaskContainer from '../Containers/TaskContainer'
import '../MyStyles.css'
import axios from 'axios'

const TaskPage = (props) => {

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

    console.log(tasks)

    if(tasks) //Tasks Present
    {
        let taskArr = tasks.map(
            tsk => <TaskContainer key={tsk.id} task={tsk} />
        )

        console.log(tasks)

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


/*const TaskPage = (props) => 
{
  //  const [task,setTask] = useContext(taskContext)

  //  console.log(task)
/*
    if(task){

    const subtasks = task.Subtasks.map(
        subtask => <SubtaskPage key = {subtask.id} subtask = {subtask}/>
    )

    return (
        <div className="taskPage-Style">
            <CreateSubtaskButton taskId={task.id}/>
            <h3>{task.Name}</h3>
            <a>{task.Deadline.split("T")[0]}</a>
            <h5>{task.Description}</h5>
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

    */
 //  return <div>Task Page</div> 
//} 

export default TaskPage;