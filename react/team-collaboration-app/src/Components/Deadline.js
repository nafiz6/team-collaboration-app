import React from 'react'
import '../MyStyles.css'

const Deadline = (props) => 
{
    return (
        <div className='deadline-Style'>Due in {props.time} days</div>
    )
}

export default Deadline;