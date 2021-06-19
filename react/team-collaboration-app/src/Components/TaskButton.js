import React from 'react'
import '../MyStyles.css'
import { Link } from 'react-router-dom'

const TaskButton = (props) => {
    return (
        <Link to = {`/project/${props.id}/ws/${props.wsid}/tasks`}>
        <button className="navBarButton-Style">Tasks</button>
        </Link>    
    )
}

export default TaskButton;