import axios from 'axios'
import React, { useEffect, useState } from 'react'
import Deadline from '../Components/Deadline'
import SubtaskButton from '../Components/SubtaskButton'
import '../MyStyles.css'
import { Link } from 'react-router-dom'
import { Button } from 'primereact/button';
import { Dialog } from 'primereact/dialog';
import { MultiSelect } from 'primereact/multiselect';

const TaskContainer = (props) => {

    const [subtasks, setSubtasks] = useState([])
    const [displayBasic, setDisplayBasic] = useState(false);
    const [taskIDToAddUsersTo, setTaskIDToAddUsersto] = useState(null);
    const [usersToAddToTask, setUsersToAddtoTask] = useState([]);
    const [workspaceUsersNotInTask, setWorkspaceUsersNotInTask] = useState([]);
    const [retAddr, setRetAddr] = useState(window.location.href);
    const [changes, setChanges] = useState(0);

    const dialogFuncMap = {
        'displayBasic': setDisplayBasic,
    }

    const getSubtasks = async () => {

        if (props.task.id) {
            let res = await axios.get(`http://localhost:8080/api/subtask/${props.task.id}`)
            setSubtasks(res.data)
        }

    }

    const addUsersToTask = async () => {



        // console.log(usersToAddToTask);



        usersToAddToTask.forEach(async user => {
            console.log(user);
            await axios.post(`http://localhost:8080/api/assign-task/${props.task.id}`, JSON.stringify({
                uid: user._id,
                role: 2 //no roles in task for now
            }))
            setChanges(c => c + 1);
        })

        window.location.reload();





    }

    const getWorkspaceUsers = async () => {
        let users = await axios.get(`http://localhost:8080/api/workspace-user-tasks/${props.ws}`)

        // console.log(users.data);
// 
        let wsUsersNotInTask = users.data.filter(u => !props.task.Assigned_users.some(a => a.id === u._id));

        setWorkspaceUsersNotInTask(wsUsersNotInTask);

        // console.log(wsUsersNotInTask);


    }

    useEffect(() => {
        getWorkspaceUsers();
        getSubtasks();
    }, [props.task.id, changes])


    useEffect(() => {
        let addrArr = (window.location.href).split("http://localhost:3000");
        setRetAddr(addrArr[1]);

    }, [window.location.href])

    const onClick = (name, position) => {
        dialogFuncMap[`${name}`](true);
    }
    const onHide = (name) => {
        dialogFuncMap[`${name}`](false);
    }

    const addingUsersToTask = (name) => {
        dialogFuncMap[`${name}`](false);
        addUsersToTask()

    }
    const renderFooter = (name) => {
        return (
            <div>
                <Button label="Add" icon="pi pi-check" onClick={() => addingUsersToTask(name)} autoFocus />
            </div>
        );
    }
    const CreateProjectFrom =
        <div>
            <h5>Add Users To Task</h5>

            <h5>Select Users to add to Project</h5>
            <MultiSelect optionLabel="name" value={usersToAddToTask} options={workspaceUsersNotInTask} onChange={(e) => {
                setUsersToAddtoTask(e.value)
                console.log(e.value);

            }} optionLabel="name" />
        </div>



    if (subtasks) {
        const subtasksArr = subtasks.map(
            stask => <SubtaskButton key={stask.id} name={stask.Name} />
        )

        return (
            <div>
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
                        {
                            props.task.Assigned_users.map(u => (
                                <div>
                                    <p>{u.Name}</p>
                                </div>
                            ))
                        }



                        {subtasksArr}



                    </button>

                </Link>
                <Button label="Add User" onClick={() => onClick('displayBasic')} />
                <Dialog header="Add Users To Task" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                    {CreateProjectFrom}
                </Dialog>
            </div>
        )

    }
    else {

        return (

            <div>
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
                        {
                            props.task.Assigned_users.map(u => (
                                <div>
                                    <p>{u.Name}</p>
                                </div>
                            ))
                        }
                    </button>
                </Link>
                <Button label="Add User" onClick={() => onClick('displayBasic')} />
                <Dialog header="Add Users To Task" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                    {CreateProjectFrom}
                </Dialog>


            </div>



        )


    }

}

export default TaskContainer;