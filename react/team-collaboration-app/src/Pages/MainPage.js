import React, { useState, useEffect, useCallback } from 'react'
import '../MyStyles.css'
import axios from 'axios'
import HeaderContainer from '../Containers/HeaderContainer'
import RoomsContainer from '../Containers/RoomsContainer'
import NavBar from '../Containers/NavBar'
import WorkContainer from '../Containers/WorkContainer'
import ProjectContainer from '../Containers/ProjectContainer'

const MainPage = (props) => {

    const [projects, setProjects] = useState();

    const getProjects = async () => {
        await axios.get('http://localhost:8080/api/project')
            .then((res) => {
                console.log(res.data);
                setProjects(res.data);
            })
    }

    useEffect(() => {
        getProjects();
    }, [])

    // Selectes first project for initial viewing
    let initialProject = null;
    if(projects)
    {
        initialProject = projects[0];
    }

    const [selectedProject,setSelectedProject] = useState(initialProject)
    

    // If a project is selected, that project workspaces are viewed
    useEffect(() => {
    if(Object.keys(props.match.params).length != 0)
    {
        if(projects){
        projects.forEach(element => {
            if(element.id === props.match.params.id)
            {
                setSelectedProject(element);
            }
        });
        }
    } 
    },[projects,props.match.params])

    return (
        <div className='page-Style'>
            <HeaderContainer />
            <div className='bottom-Style'>
                <ProjectContainer projects = {projects}/>
                <RoomsContainer project = {selectedProject}/> {/* This gets current selected project */}
                <div className='taskWork-Style'>
                    <NavBar />
                    <WorkContainer tab={props.tab} />
                </div>
            </div>
        </div>
    )
}

export default MainPage;