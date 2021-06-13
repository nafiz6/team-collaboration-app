import React, {useContext, useState} from 'react'
import '../MyStyles.css'


const ProjectButton = (props) => 
{
   
    return (
        <button className='projectButton-Style'
        onClick = { () =>
            {
              //Set current project
            }
        }
        >{props.project.Name}</button>
    )
}

export default ProjectButton;