import React, {useContext, useState} from 'react'
import '../MyStyles.css'
import {currProjContext} from '../App'
import {stateContext} from "../App"


const ProjectButton = (props) => 
{
    const [project,setProject] = useContext(currProjContext)
    const [state,setState] = useContext(stateContext)

    return (
        <button className='projectButton-Style'
        onClick = { () =>
            {
                setProject(props.project) 
                setState(0)
            }
        }
        >{props.project.Name}</button>
    )
}

export default ProjectButton;