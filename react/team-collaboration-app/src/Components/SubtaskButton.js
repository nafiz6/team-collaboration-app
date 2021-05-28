import React from 'react'
import '../MyStyles.css'

const SubtaskButton = (props) => 
{
    return (
        <div className='subtaskButton-Style'>{props.name}</div>
    )
}

export default SubtaskButton;