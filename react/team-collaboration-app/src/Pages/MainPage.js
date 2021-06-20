import React, { useState, useEffect } from 'react'
import '../MyStyles.css'
import axios from 'axios'
import HeaderContainer from '../Containers/HeaderContainer'
import RoomsContainer from '../Containers/RoomsContainer'
import NavBar from '../Containers/NavBar'
import WorkContainer from '../Containers/WorkContainer'
import ProjectContainer from '../Containers/ProjectContainer'

const MainPage = (props) => {

    const [projects, setProjects] = useState([]);
    const [ws, setWs] = useState([]);
    const [initialProject, setInitialProject] = useState(null);
    const [initialWS, setInitialWS] = useState(null);
    const [projId, setProjId] = useState(null);
    const [wsId, setWsId] = useState(null);

    const getWs = async () => {

        if (projects.length !== 0) {

            let res = null;

            if (Object.keys(props.match.params).length > 0) {
                 res = await axios.get(`http://localhost:8080/api/workspace/${props.match.params.id}`)
            }
            else {
                 res = await axios.get(`http://localhost:8080/api/workspace/${projects[0].id}`)
            }

            setWs(res.data);

            if (res.data.length > 0) {
                setInitialWS(res.data[0]);
                setWsId(res.data[0].id);
            }
        }
    }

    const getProjects = async () => {
        let res = await axios.get('http://localhost:8080/api/project')
        setProjects(res.data);


        if (res.data.length > 0) {
            setInitialProject(res.data[0]);
            setProjId(res.data[0].id);
        }
    }

    useEffect(() => {
        getProjects();
    }, [])

    useEffect(() => {
        getWs();
    }, [projects,props.match.params])

    const [selectedProject, setSelectedProject] = useState(initialProject)


    // If a project is selected, that project workspaces are viewed
    useEffect(() => {
        if (Object.keys(props.match.params).length != 0) {
            if (projects.length > 0) {
                projects.forEach(element => {
                    if (element.id === props.match.params.id) {
                        setSelectedProject(element);
                        setProjId(element.id);
                    }
                });
            }
        }
        else {
            if (projects.length > 0) {
                setSelectedProject(projects[0]);
                setProjId(projects[0].id);
            }
        }
    }, [projects, props.match.params.id])

    const [selectedWS, setSelectedWS] = useState(initialWS)

    // If a ws is selected, that workspaces tasks are viewed
    useEffect(() => {
        if (Object.keys(props.match.params).length != 0) {
            if (ws.length > 0) {
                ws.forEach(element => {
                    if (element.id === props.match.params.wsid) {
                        setSelectedWS(element);
                        setWsId(element.id)
                    }
                });
            }
        }
        else {
            if (ws.length > 0) {
                setSelectedWS(ws[0]);
                setWsId(ws[0].id);
            }
        }
    }, [ws, props.match.params.wsid])

    return (
        <div className='page-Style'>
            <HeaderContainer />
            <div className='bottom-Style'>
                <ProjectContainer projects={projects} />
                <RoomsContainer project={selectedProject} /> {/* This gets current selected project */}
                <div className='taskWork-Style'>
                    <NavBar id={projId} wsid={wsId} />
                    <WorkContainer ws={wsId} tab={props.tab} />
                </div>
            </div>
        </div>
    )
}

export default MainPage;