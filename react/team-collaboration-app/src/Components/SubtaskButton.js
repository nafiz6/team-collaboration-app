import React from 'react'
import '../MyStyles.css'

const SubtaskButton = (props) => 
{
    return (
        <div className='subtaskButton-Style'><p>{props.name}</p></div>
    )
}

export default SubtaskButton;