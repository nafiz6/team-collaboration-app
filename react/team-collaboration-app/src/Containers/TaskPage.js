import React from 'react'
import CreateTaskButton from '../Components/CreateTask'
import '../MyStyles.css'
import SubtaskPage from './SubtaskPage'

const TaskPage = (props) => 
{
    const subtasks = [
        <SubtaskPage name="Level Make" des="Making a fun level" budget="30000" user="Somik and Nafiz"/>,
        <SubtaskPage name="Boss Make" des="Making a scary boss" budget="10000" user="Somik and Abrar"/>,
        <SubtaskPage name="Music Make" des="Making a soothing track" budget="5000" user="Abrar and Nafiz"/>,
        <SubtaskPage name="Game Make" des="Making a great game" budget="100000" user="Embros"/>,
    ]

    return (
        <div className="taskPage-Style">
            <text>Task: {props.name}</text>
            <text>Due: In {props.time} days</text>
            {subtasks}
        </div>
    )
}

export default TaskPage;