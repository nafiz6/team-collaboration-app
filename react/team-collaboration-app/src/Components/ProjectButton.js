import React, {useContext, useState} from 'react'
import '../MyStyles.css'
import {currProjContext} from '../App'
import {stateContext, taskContext, currWSContext} from "../App"


const ProjectButton = (props) => 
{
    const [project,setProject] = useContext(currProjContext)
    const [state,setState] = useContext(stateContext)
    const [task,setTask] = useContext(taskContext)
    const [currWS, setCurrWS] = useContext(currWSContext)


    return (
        <button className='projectButton-Style'
        onClick = { () =>
            {
                setProject(props.project) 
                setState(0)

                if (props.project.Workspaces.length > 0){
                    setCurrWS(props.project.Workspaces[0])
                    if (props.project.Workspaces[0].Tasks.length > 0)
                        setTask(props.project.Workspaces[0].Tasks[0])
                    }
            }
        }
        >{props.project.Name}</button>
    )
}

export default ProjectButton;