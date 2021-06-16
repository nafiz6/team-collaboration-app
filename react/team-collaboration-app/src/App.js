import './MyStyles.css'
import React, { useCallback, useEffect, useState } from 'react'
import axios from 'axios';
import MainPage from './Pages/MainPage';
import Login from './Pages/Login';
import SignUp from './Pages/SignUp';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'


function App() {

  const [projects, setProjects] = useState([]);
  
  const dataFetch = useCallback( async () => {
    let res = await axios.get('http://localhost:8080/api/project')
    setProjects(res.data);
    console.log(projects);
  })

  useEffect(() => {
    dataFetch();
  }, [dataFetch])


  return (

    <Router>
      <Switch>
        <Route path="/" exact render={(props) => (<Login {...props} />)} />
        <Route path="/tasks" render={(props) => (<MainPage {...props} projects = {projects} tab = "tasks"/>)} />
        <Route path="/chats" render={(props) => (<MainPage {...props} projects = {projects} tab = "chats"/>)} />
        <Route path="/files" render={(props) => (<MainPage {...props} projects = {projects} tab = "files"/>)} />
        <Route path="/stats" render={(props) => (<MainPage {...props} projects = {projects} tab = "stats"/>)} />
        <Route path="/signup" render={(props) => (<SignUp {...props} />)} />
      </Switch>
    </Router>

  );
}

export default App
