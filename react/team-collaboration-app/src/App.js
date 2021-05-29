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

function App() {

  const [projects, setProjects] = useState([]);
  const [currProj, setCurrProj] = useState(null);
  const [currWS, setCurrWS] = useState(null);

  const dataFetch = useCallback(async () => {
    let res = await axios.get('http://localhost:8080/api/project')

    setProjects(res.data);
    setCurrProj(res.data[0]);

  }, [])

  useEffect(() => {
    dataFetch();
  }, [dataFetch])


  return (

    // <h1>{currProj?.Name}</h1>


    <div className='page-Style'>
      <HeaderContainer />
      <div className='bottom-Style'>
        <currProjContext.Provider value={[currProj, setCurrProj]}>
          <ProjectContainer projects={projects} />
        </currProjContext.Provider>
        <currWSContext.Provider value={[currWS, setCurrWS]}>
          <RoomsContainer project={currProj} />
        </currWSContext.Provider>
        <div className='taskWork-Style'>
          <NavBar />
          <WorkContainer ws={currWS} />
        </div>
      </div>
    </div>


  );
}

export default App
