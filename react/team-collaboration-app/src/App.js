import './MyStyles.css'
import React, { useCallback, useContext, useEffect, useState } from 'react'
import HeaderContainer from './Containers/HeaderContainer';
import ProjectContainer from './Containers/ProjectContainer';
import RoomsContainer from './Containers/RoomsContainer';
import NavBar from './Containers/NavBar';
import WorkContainer from './Containers/WorkContainer';
import axios from 'axios';


export const currProjContext = React.createContext();
export const currWSContext = React.createContext();
export const stateContext = React.createContext();
export const taskContext = React.createContext();

function App() {

  const [projects, setProjects] = useState([]);
  const [currProj, setCurrProj] = useState(null);
  const [currWS, setCurrWS] = useState(null);
  const [state, setState] = useState(0);
  const [task, setTask] = useState(null);

  const dataFetch = useCallback(async () => {
    let res = await axios.get('http://localhost:8080/api/project')

    setProjects(res.data);
    setCurrProj(res.data[0]);

  }, [])

  useEffect(() => {
    dataFetch();
  }, [dataFetch])


  return (

    <div className='page-Style'>
      <HeaderContainer />
      <div className='bottom-Style'>
        <currProjContext.Provider value={[currProj, setCurrProj]}>
          <stateContext.Provider value={[state, setState]}>
            <ProjectContainer projects={projects} />
          </stateContext.Provider>
        </currProjContext.Provider>
        <currWSContext.Provider value={[currWS, setCurrWS]}>
          <stateContext.Provider value={[state, setState]}>
            <RoomsContainer project={currProj} />
          </stateContext.Provider>
        </currWSContext.Provider>
        <div className='taskWork-Style'>
          <NavBar />
          <stateContext.Provider value={[state, setState]}>
            <taskContext.Provider value={[task,setTask]}>
            <WorkContainer ws={currWS} />
            </taskContext.Provider>
          </stateContext.Provider>
        </div>
      </div>
    </div>

  );
}

export default App
