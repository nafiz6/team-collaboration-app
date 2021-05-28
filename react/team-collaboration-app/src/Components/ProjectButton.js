import React from 'react'
import '../MyStyles.css'

const ProjectButton = (props) => 
{
    return (
        <button className='projectButton-Style'>{props.name}</button>
    )
}

export default ProjectButton;