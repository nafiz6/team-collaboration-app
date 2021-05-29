import React, {useContext, useState} from 'react'
import '../MyStyles.css'
import {currProjContext} from '../App'


const ProjectButton = (props) => 
{
    const [project,setProject] = useContext(currProjContext)
    console.log(project)

    return (
        <button className='projectButton-Style'
        onClick = { () =>
            {
                setProject(props.project) 
                console.log(project)
            }
        }
        >{props.project.Name}</button>
    )
}

export default ProjectButton;