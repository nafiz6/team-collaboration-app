import './MyStyles.css'
import React, { useCallback, useContext, useEffect, useState } from 'react'
import HeaderContainer from './Containers/HeaderContainer';
import ProjectContainer from './Containers/ProjectContainer';
import RoomsContainer from './Containers/RoomsContainer';
import NavBar from './Containers/NavBar';
import WorkContainer from './Containers/WorkContainer';
import axios from 'axios';


export const currProjContext = React.createContext();

function App() {

  const [projects, setProjects] = useState([]);
  const [currProj, setCurrProj] = useState({});
  
 /* const dataFetch = useCallback(() => {
      axios.get('http://localhost:8080/api/project')
      .then(res => {
        setProjects(res.data);
        setCurrProj(projects[0]);
        console.log(projects)
        console.log(currProj)
  
      }).catch(err => {
        console.log(err);
      })
    },[])


  useEffect(() => {dataFetch()},[dataFetch]) */

 /* const newSetProjects = () =>
  {
     
  }  */

/*  useEffect(async ()=>
  {
      let res = await axios.get('http://localhost:8080/api/project')
      
       setProjects(res.data);
       setCurrProj(projects[0]);
      

        console.log(res)
        console.log(projects)
        console.log(currProj)

     },[]) */

  
  return (
   
   // <h1>{currProj?.Name}</h1> 
    
      <div className='page-Style'>
        <HeaderContainer/>
        <div className='bottom-Style'>
          <currProjContext.Provider value = {[currProj,setCurrProj]}>
          <ProjectContainer projects = {projects}/>
          </currProjContext.Provider>
          <RoomsContainer project={currProj}/>
          <div className = 'taskWork-Style'>
            <NavBar/>
            <WorkContainer/> 
          </div>
        </div>
      </div>
     
  );
}

export default App
