import React from "react"
import ProjectAddButton from "../Components/ProjectAddButton";
import ProjectButton from "../Components/ProjectButton";
import '../MyStyles.css'


const ProjectContainer = (props) => 
{
    if(props.projects){
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
    else{
        return(
            <div className='project-Style'>
                <ProjectAddButton/>
            </div>
        )
    }

}

export default ProjectContainer;