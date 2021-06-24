import axios from 'axios';
import React, { useState, useEffect } from 'react'
import AddUpdate from '../Components/AddUpdate';

const SubtaskContainer = (props) => {

    const [updates, setUpdates] = useState([]);

    const getUpdates = async () => {

        if (props.subtask) {
            let res = await axios.get(`http://localhost:8080/api/updates/${props.subtask.id}`)
            setUpdates(res.data);
        }
    }

    useEffect(() => {
        getUpdates();
    }, [props.subtask])

    let updateArr = [];

    if (updates) {
        updateArr = updates.map(
            update => <p key={update.id}>{update.user_id + " : " + update.Text}
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
            <AddUpdate user={userObj} subtaskId={props.subtask.id} taskId={props.subtask.task_id}/>
        </div>
    )
}

export default SubtaskContainer;
