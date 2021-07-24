import axios from 'axios';
import React, { useState, useEffect } from 'react'
import AddUpdate from '../Components/AddUpdate';
import { Button } from 'primereact/button';
import { Dialog } from 'primereact/dialog';
import { MultiSelect } from 'primereact/multiselect';
import { Avatar } from 'primereact/avatar';
import Time from '../Components/Time';


const SubtaskContainer = (props) => {

    const [updates, setUpdates] = useState([]);
    const [displayBasic, setDisplayBasic] = useState(false);
    const [usersToAddToSubtask, setUsersToAddtoSubtask] = useState([]);
    const [taskUsersNotInSubtask, setTaskUsersNotInSubtask] = useState([]);
    const [changes, setChanges] = useState(0);
    const dialogFuncMap = {
        'displayBasic': setDisplayBasic,
    }

    const [taskUsers, setTaskUsers] = useState([]);

    const getUpdates = async () => {

        if (props.subtask) {
            let res = await axios.get(`http://localhost:8080/api/updates/${props.subtask.id}`)
            setUpdates(res.data);
        }
    }
    const getTaskUsers = async () => {
        let users = await axios.get(`http://localhost:8080/api/task-users/${props.taskId}`)

        console.log(users.data);
        users.data.forEach(async u => {
            let deets = await axios.get(`http://localhost:8080/api/user-details/${u.id}`);
            setTaskUsers(users => [...users, deets.data]);
        });

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
    }, [props.subtask, changes])


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

            <h3>Select Team members to add to Subtask</h3>
            <MultiSelect placeholder="select team members" className="form-text" optionLabel="name" value={usersToAddToSubtask} options={taskUsersNotInSubtask} onChange={(e) => {
                setUsersToAddtoSubtask(e.value)
                console.log(e.value);

            }} optionLabel="Name" />
        </div>


    let updateArr = [];

    if (updates) {
        updateArr = updates.map(
            update =>
                <div className="subtask-update">
                    <div className="subtask-update-content">
                        <Avatar style={{
                            'margin': '0px 20px'
                        }} label={taskUsers.find(u => u.id === update.User.id)?.Dp ? null : update.User.Name[0]} image={taskUsers.find(u => u.id === update.User.id)?.Dp} />
                        {/* <p>{update.User.Name}</p> */}
                        <p className="subtask-update-text" key={update.id}>{update.Text}</p>
                    </div>


                    <Time className="subtask-update-time" time={update.Timestamp} />
                </div>

        )
    }
    else {
        updateArr = 
        <div className="centered" style={{
            height: '100px'
        }}>
            <h3>No Updates yet</h3>
        </div> 
    }

    let assUserArr = [];

    console.log(taskUsers);
    console.log(props.subtask.Assigned_users);

    if (props.subtask.Assigned_users.length > 0) {
        assUserArr = props.subtask.Assigned_users.map(
            user =>
                <div class="subtask-user">
                    <Avatar style={{
                        'margin-right': '20px'
                    }} label={taskUsers.find(u => u.id === user.id)?.Dp ? null : user.Name[0]} image={taskUsers.find(u => u.id === user.id)?.Dp} />
                    <p>{user.Name}</p>

                </div>
        )
    }

    //DUMMY DATA
    let userObj = {
        Name: "Marques Brownlee",
        id: "60af82dccbe1709146f71669"
    }

    return (
        <div className='subtaskPage-Style'>
            <h1 className="subtask-name">{props.subtask.Name}</h1>
            <p className="task-page-description">{props.subtask.Description}</p>
            {/* <text> Budget: {props.subtask.Budget}</text> */}
            <div class="subtask-users-container">
                {/* <h2>Assigned Users</h2> */}
                <div className="subtask-users">
                    {assUserArr}
                </div>
            </div>


            <h1>Updates</h1>

            {updateArr}
            <AddUpdate user={userObj} subtaskId={props.subtask.id} taskId={props.subtask.task_id} />
            <Button className="addUserToTaskButton" label="Assign User to Subtask" onClick={() => onClick('displayBasic')} />
            <Dialog header="Assign Team Members to Subtask" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateProjectFrom}
            </Dialog>
        </div>
    )
}

export default SubtaskContainer;
