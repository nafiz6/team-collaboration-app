import React from 'react'
import '../MyStyles.css'

const SubtaskPage = (props) => 
{
    return (
        <div className='subtaskPage-Style'>
           <text> Name: {props.name}</text>
           <text> Description:  {props.des}</text>
          <text>  Budget: {props.budget}</text>
          <text>  Assigned to: {props.user}</text>
        </div>
    )
}

export default SubtaskPage;