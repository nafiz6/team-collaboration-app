import React from 'react'
import '../MyStyles.css'
import { Link } from 'react-router-dom'

const TaskButton = (props) => {
    return (
        <Link to = {`/project/${props.id}/ws/${props.wsid}/tasks`}>
        <button className={props.tab === "tasks" ? "navBarButton-Style-select" : "navBarButton-Style"}> <i className="pi pi-list"></i>Tasks</button>
        </Link>    
    )
}

export default TaskButton;