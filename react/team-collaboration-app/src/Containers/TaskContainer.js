import axios from 'axios'
import React, { useEffect, useState } from 'react'
import Deadline from '../Components/Deadline'
import SubtaskButton from '../Components/SubtaskButton'
import '../MyStyles.css'
import { Link } from 'react-router-dom'
import { Button } from 'primereact/button';
import { Dialog } from 'primereact/dialog';
import { MultiSelect } from 'primereact/multiselect';

import { Avatar } from 'primereact/avatar';
import { AvatarGroup } from 'primereact/avatargroup';
import { ProgressBar } from 'primereact/progressbar';
import { PrimeIcons } from 'primereact/api';

const TaskContainer = (props) => {

    const [subtasks, setSubtasks] = useState([])
    const [displayBasic, setDisplayBasic] = useState(false);
    const [taskIDToAddUsersTo, setTaskIDToAddUsersto] = useState(null);
    const [usersToAddToTask, setUsersToAddtoTask] = useState([]);
    const [workspaceUsersNotInTask, setWorkspaceUsersNotInTask] = useState([]);
    const [retAddr, setRetAddr] = useState(window.location.href);
    const [changes, setChanges] = useState(0);
    const [userDetails, setUserDetails] = useState([]);

    const dialogFuncMap = {
        'displayBasic': setDisplayBasic,
    }

    const getSubtasks = async () => {

        if (props.task.id) {
            let res = await axios.get(`http://localhost:8080/api/subtask/${props.task.id}`)
            setSubtasks(res.data)
        }

    }

    const getUserDetails = async () => {

        if (props.task.id) {
            props.task.Assigned_users.forEach(async u => {
                let res = await axios.get(`http://localhost:8080/api/user-details/${u.id}`);
                setUserDetails(u => [...u, res.data]);
            });
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

    const deleteTask = async () => {



        // console.log(usersToAddToTask);




        await axios.post(`http://localhost:8080/api/delete-task/${props.task.id}`);

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
            getUserDetails();
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


        let subtasksArr;
        if (subtasks) {
            subtasksArr = subtasks.map(
                stask => <SubtaskButton key={stask.id} name={stask.Name} />
            )
        }

        console.log(props.task.Assigned_users);
        console.log(userDetails);

        props.task.Assigned_users.map(u => {
            // <Avatar label={u.Name[0]} image={userDetails.find(d => d.id === u.id)?.Dp} shape="circle" size="large" />
            
            console.log( "imag", userDetails.find(d => d.id === u.id)?.Dp);
            console.log("u", u.id);

        })



        return (
            <div>
                <Link

                    className='taskContainer-Style'

                    to={
                        {
                            pathname: `${retAddr}/taskpage/${props.task.id}`,
                            state:
                            {
                                taskname: props.task.Name,
                                deadline: props.task.Deadline,
                                description: props.task.Description
                            }
                        }}
                >



                    <h3 className='taskName-Style'>{props.task.Name}</h3>
                    {/* <h5 className='taskName-Style'> Budget: {props.task.Budget}</h5> */}


                    {/* <ProgressBar value={isNaN((props.task.Spent / props.task.Budget )* 100) ? 0 : (props.task.Spent / props.task.Budget )* 100} /> */}

                    <Deadline time={props.task.Deadline.split("T")[0]} />
                    <AvatarGroup>
                        {props.task.Assigned_users.map(u => (
                            <Avatar label={userDetails.find(d => d.id === u.id)?.Dp ? null : u.Name[0]} image={userDetails.find(d => d.id === u.id)?.Dp ? userDetails.find(d => d.id === u.id)?.Dp : null} shape="circle" size="large" />

                        ))}
                    </AvatarGroup>


                    <div className="subtasks">
                        {subtasks ? subtasksArr : null}
                    </div>


                    <div className="addMemberToTaskButton">
                        <Button icon="pi pi-plus" iconPos="left" label="Assign Member" className="p-button-raised p-button-rounded" onClick={(e) => {
                            e.preventDefault();
                            onClick('displayBasic');
                        }} />
                    </div>
                    <div className="addMemberToTaskButton">
                        <Button icon="pi pi-trash" iconPos="left" label="Delete Task" className="p-button-raised p-button-rounded p-button-danger" onClick={(e) => {
                            e.preventDefault();
                            deleteTask();
                        }} />
                    </div>







                </Link>

                <Dialog header="Add Users To Task" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                    {CreateProjectFrom}
                </Dialog>
            </div>
        )

    }




    export default TaskContainer;