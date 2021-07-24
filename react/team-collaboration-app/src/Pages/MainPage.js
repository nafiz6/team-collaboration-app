import React, { useState, useEffect } from 'react'
import '../MyStyles.css'
import axios from 'axios'
import HeaderContainer from '../Containers/HeaderContainer'
import RoomsContainer from '../Containers/RoomsContainer'
import NavBar from '../Containers/NavBar'
import WorkContainer from '../Containers/WorkContainer'
import ProjectContainer from '../Containers/ProjectContainer'
import { useLocation } from 'react-router-dom'

const MainPage = (props) => {


    const [projects, setProjects] = useState([]);
    const [ws, setWs] = useState([]);
    const [initialProject, setInitialProject] = useState(null);
    const [initialWS, setInitialWS] = useState(null);
    const [projId, setProjId] = useState(null);
    const [wsId, setWsId] = useState(null);



    const getWs = async () => {

        // console.log(props.match.params);



        if (projects.length !== 0) {

            let res = null;

            if (Object.keys(props.match.params).length > 0 && projId) {
                res = await axios.get(`http://localhost:8080/api/workspace/${projId}`, { withCredentials: true })
            }
            else {
                res = await axios.get(`http://localhost:8080/api/workspace/${projects[0].id}`, { withCredentials: true })
            }


            if (res.data.length > 0 ) {
                console.log("workspace RESET", projId)
                setInitialWS(res.data[0]);
                setWsId(res.data[0].id);
            }
            setWs(res.data);
        }
    }

    const getProjects = async () => {
        let res = await axios.get('http://localhost:8080/api/project', { withCredentials: true })


        if (res.data.length > 0) {
            console.log(projId + " reset1 ")
            setInitialProject(res.data[0]);
            setProjId(res.data[0].id);
            setProjects(res.data);
        }
    }

    useEffect(() => {
        getProjects();
    }, [])

    useEffect(() => {
        getWs();
    }, [projects, projId])

    const [selectedProject, setSelectedProject] = useState(initialProject)


    // If a project is selected, that project workspaces are viewed
    useEffect(() => {
        if (Object.keys(props.match.params).length > 0) {
            if (projects.length > 0) {
                projects.forEach(element => {
                    if (element.id === props.match.params.id) {
                        setSelectedProject(element);
                        console.log(projId + " reset2 ")
                        setProjId(element.id);



                        //set workspace to first one in proj
                    }
                });
            }
        }
        else {
            if (projects.length > 0) {
                setSelectedProject(projects[0]);
                console.log(projId + " reset3 ")
                setProjId(projects[0].id);
            }
        }
    }, [projects, props.match.params.id])

    const [selectedWS, setSelectedWS] = useState(initialWS)

    // If a ws is selected, that workspaces tasks are viewed
    useEffect(() => {
        if (Object.keys(props.match.params).length > 0) {
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

                console.log("workspace RESET")
                setSelectedWS(ws[0]);
                setWsId(ws[0].id);
            }
        }
    }, [ws, props.match.params.wsid])

    // console.log(selectedProject);





    const location = useLocation()
    let taskname, deadline, description
    if (location.state) {
        taskname = location.state.taskname;
        deadline = location.state.deadline;
        description = location.state.description;
    }

    return (
        <div className='page-Style'>
            <HeaderContainer {...props} />
            <div className='bottom-Style'>
                <ProjectContainer projects={projects} />
                <RoomsContainer project={selectedProject} /> {/* This gets current selected project */}
                <div className='taskWork-Style'>
                    <NavBar id={projId} wsid={wsId} tab={props.tab} />
                    <WorkContainer ws={wsId} tab={props.tab} tid={props.match.params.tid} taskname={taskname} deadline={deadline}
                        description={description} />
                </div>
            </div>
        </div>
    )
}

export default MainPage;