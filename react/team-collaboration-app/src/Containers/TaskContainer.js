import React from 'react'
import Deadline from '../Components/Deadline'
import SubtaskButton from '../Components/SubtaskButton'
import '../MyStyles.css'

const TaskContainer = (props) => 
{
    const subtasks = [
        <SubtaskButton name="Level Design"/>,
        <SubtaskButton name="Character Design"/>,
        <SubtaskButton name="Fight Design"/>,
        <SubtaskButton name="Graphics Design"/>,
        <SubtaskButton name="Music Design"/>,
        <SubtaskButton name="Boss Design"/>
    ]

    return (
        <button className='taskContainer-Style'>
            <p className='taskName-Style'>{props.name}</p>
            <Deadline time={props.time}/>
            {subtasks}
        </button>
    )
}

export default TaskContainer;