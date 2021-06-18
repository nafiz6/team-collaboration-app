import React from 'react'
import { Link } from 'react-router-dom'
import '../MyStyles.css'


const ProjectButton = (props) => {
    return (

        <Link to={`/project/${props.id}`}>
            <button className='projectButton-Style'>{props.name}</button>
        </Link>
    )
}

export default ProjectButton;