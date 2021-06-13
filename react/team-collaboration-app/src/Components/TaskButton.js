import React from 'react'
import '../MyStyles.css'
import { Link } from 'react-router-dom'

const TaskButton = () => {
    return (
        <Link to = "/tasks">
        <button className="navBarButton-Style"
            onClick={() => {
                /* Set current Tab as Task */
            }
            }>Tasks</button>
        </Link>    
    )
}

export default TaskButton;