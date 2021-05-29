import './MyStyles.css'
import React, { useEffect } from 'react'
import HeaderContainer from './Containers/HeaderContainer';
import ProjectContainer from './Containers/ProjectContainer';
import RoomsContainer from './Containers/RoomsContainer';
import NavBar from './Containers/NavBar';
import WorkContainer from './Containers/WorkContainer';
import axios from 'axios'

function App() {

  useEffect(() => {
    axios.get('http://localhost:8080/api/project')
    .then(res => {
      console.log(res.data);

    }).catch(err => {
      console.log(err);
    })
  })

  return (

   /* <div className="App">
      <header className="App-header">
        { <img src={logo} className="App-logo" alt="logo" /> }
      </header>
    </div> */
      <div className='page-Style'>
        <HeaderContainer/>
        <div className='bottom-Style'>
          <ProjectContainer/>
          <RoomsContainer/>
          <div className = 'taskWork-Style'>
            <NavBar/>
            <WorkContainer/> 
          </div>
        </div>
      </div>
    
  );
}

export default App
