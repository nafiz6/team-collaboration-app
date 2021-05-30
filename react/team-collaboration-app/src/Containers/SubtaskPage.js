import React from 'react'
import AddUpdate from '../Components/AddUpdate'
import '../MyStyles.css'

const SubtaskPage = (props) => 
{
    const assignedTo = props.subtask.Assigned_users.map(
        user => user.Name ).join(",")

    const updates = props.subtask.Updates.map(
        update => <p>{update.User.Name + " : " + update.Text }
    </p>)

    let userObj = {
        Name: "Marques Brownlee",
        id: "60af82dccbe1709146f71669"
    }
        
    return (
        <div className='subtaskPage-Style'>
           <text> Name: {props.subtask.Name}</text>
           <text> Description:  {props.subtask.Description}</text>
           <text>  Budget: {props.subtask.Budget}</text>
             {updates}
          <AddUpdate user={userObj} subtaskId={props.subtask.id}/>
        </div>
    )
}

export default SubtaskPage;