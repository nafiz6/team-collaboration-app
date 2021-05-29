import React from 'react'
import AddUpdate from '../Components/AddUpdate'
import '../MyStyles.css'

const SubtaskPage = (props) => 
{
    const assignedTo = props.subtask.Assigned_users.map(
        user => user.Name ).join(",")

    const updates = props.subtask.Updates.map(
        update => update.User.Name + " : " + update.Text
    ).join("\n")
        
    return (
        <div className='subtaskPage-Style'>
           <text> Name: {props.subtask.Name}</text>
           <text> Description:  {props.subtask.Description}</text>
           <text>  Budget: {props.subtask.Budget}</text>
           <text>  Assigned to: {assignedTo}</text>
           <text>  Updates: {updates}</text>
          <AddUpdate user="60af82dccbe1709146f71669" subtaskId={props.subtask.id}/>
        </div>
    )
}

export default SubtaskPage;