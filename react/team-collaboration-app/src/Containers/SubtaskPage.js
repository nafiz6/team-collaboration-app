import React from 'react'
import AddUpdate from '../Components/AddUpdate'
import '../MyStyles.css'

const SubtaskPage = (props) => 
{
    return (
        <div className='subtaskPage-Style'>
           <text> Name: {props.name}</text>
           <text> Description:  {props.des}</text>
          <text>  Budget: {props.budget}</text>
          <text>  Assigned to: {props.user}</text>
          <AddUpdate user="60af82dccbe1709146f71669" subtaskId=""/>
        </div>
    )
}

export default SubtaskPage;