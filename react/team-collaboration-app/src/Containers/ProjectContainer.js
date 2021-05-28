import React from "react"
import ProjectAddButton from "../Components/ProjectAddButton";
import ProjectButton from "../Components/ProjectButton";
import '../MyStyles.css'


const ProjectContainer = () => 
{
    const projButtons = [
        <ProjectButton name="1"/>, <ProjectButton name = "2"/>, <ProjectButton name ="3"/>, <ProjectAddButton/>
    ]

    return(
        <div className='project-Style'>
            {projButtons}
        </div>
    );

}

export default ProjectContainer;