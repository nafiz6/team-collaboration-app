import React from 'react'
import '../MyStyles.css'

const SubtaskButton = (props) => 
{
    return (
        <div className='subtaskButton-Style'><h4>{props.name}</h4></div>
    )
}

export default SubtaskButton;