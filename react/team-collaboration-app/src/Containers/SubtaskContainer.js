import axios from 'axios';
import React, { useState, useEffect } from 'react'
import AddUpdate from '../Components/AddUpdate';
import { Button } from 'primereact/button';
import { Dialog } from 'primereact/dialog';
import { MultiSelect } from 'primereact/multiselect';

const SubtaskContainer = (props) => {

    const [updates, setUpdates] = useState([]);
    const [displayBasic, setDisplayBasic] = useState(false);
    const [usersToAddToSubtask, setUsersToAddtoSubtask] = useState([]);
    const [taskUsersNotInSubtask, setTaskUsersNotInSubtask] = useState([]);
    const [changes, setChanges] = useState(0);
    const dialogFuncMap = {
        'displayBasic': setDisplayBasic,
    }

    const getUpdates = async () => {

        if (props.subtask) {
            let res = await axios.get(`http://localhost:8080/api/updates/${props.subtask.id}`)
            setUpdates(res.data);
        }
    }
    const getTaskUsers = async () => {
        let users = await axios.get(`http://localhost:8080/api/task-users/${props.taskId}`)

        console.log(users.data);

        let tUsersNotInSubtask = users.data.filter(u => !props.subtask.Assigned_users.some(a => a.id === u.id));

        setTaskUsersNotInSubtask(tUsersNotInSubtask);

        console.log(tUsersNotInSubtask);


    }
    const addUsersToSubtask = async () => {



        console.log(usersToAddToSubtask);



        usersToAddToSubtask.forEach(async user => {
            console.log(user);
            await axios.post(`http://localhost:8080/api/assign-subtask/${props.subtask.id}`, JSON.stringify({
                uid: user.id,
            }))
            setChanges(c => c + 1);
        })





    }

    useEffect(() => {
        getTaskUsers();
        getUpdates();
    }, [props.subtask,changes])


    const onClick = (name, position) => {
        dialogFuncMap[`${name}`](true);
    }
    const onHide = (name) => {
        dialogFuncMap[`${name}`](false);
    }

    const addingUsersToSubtask = (name) => {
        dialogFuncMap[`${name}`](false);
        addUsersToSubtask()

    }
    const renderFooter = (name) => {
        return (
            <div>
                <Button label="Add" icon="pi pi-check" onClick={() => addingUsersToSubtask(name)} autoFocus />
            </div>
        );
    }
    const CreateProjectFrom =
        <div>
            <h5>Add Users To Task</h5>

            <h5>Select Users to add to Project</h5>
            <MultiSelect optionLabel="name" value={usersToAddToSubtask} options={taskUsersNotInSubtask} onChange={(e) => {
                setUsersToAddtoSubtask(e.value)
                console.log(e.value);

            }} optionLabel="Name" />
        </div>


    let updateArr = [];

    if (updates) {
        updateArr = updates.map(
            update => <p key={update.id}>{update.User.Name + " : " + update.Text}
            </p>)
    }

    let assUserArr = [];

    if (props.subtask.Assigned_users.length > 0) {
        assUserArr = props.subtask.Assigned_users.map(
            user => user.Name).join(",")
        console.log(assUserArr)
    }

    //DUMMY DATA
    let userObj = {
        Name: "Marques Brownlee",
        id: "60af82dccbe1709146f71669"
    }

    return (
        <div className='subtaskPage-Style'>
            <text> Name: {props.subtask.Name}</text>
            <text> Description:  {props.subtask.Description}</text>
            <text>  Budget: {props.subtask.Budget}</text>
            <text> Designated Users: {assUserArr} </text>
            {updateArr}
            <AddUpdate user={userObj} subtaskId={props.subtask.id} taskId={props.subtask.task_id} />
            <Button label="Add User" onClick={() => onClick('displayBasic')} />
            <Dialog header="Add Users To Task" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateProjectFrom}
            </Dialog>
        </div>
    )
}

export default SubtaskContainer;
