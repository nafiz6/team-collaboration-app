import React from "react"
import ProjectAddButton from "../Components/ProjectAddButton";
import ProjectButton from "../Components/ProjectButton";
import '../MyStyles.css'


const ProjectContainer = (props) => 
{
    const projButtons = props.projects.map(
        project => <ProjectButton key={project.id} project={project} />
    )

    return(
        <div className='project-Style'>
            {projButtons}
            <ProjectAddButton/>
        </div>
    );

}

export default ProjectContainer;