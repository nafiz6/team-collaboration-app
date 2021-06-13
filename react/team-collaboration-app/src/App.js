import './MyStyles.css'
import React from 'react'
import axios from 'axios';
import MainPage from './Pages/MainPage';
import Login from './Pages/Login';
import SignUp from './Pages/SignUp';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'


function App() {
  
  /*
   const dataFetch = useCallback(async () => {
     let res = await axios.get('http://localhost:8080/api/project')
 
     setProjects(res.data);
     setCurrProj(res.data[0]);
     setCurrWS(res.data[0].Workspaces[0]);
 
   }, [])
 
   useEffect(() => {
     dataFetch();
   }, [dataFetch])
   */


  return (

    <Router>
      <Switch>
        <Route path="/" exact render={(props) => (<Login {...props} />)} />
        <Route path="/tasks" render={(props) => (<MainPage {...props} tab = "tasks"/>)} />
        <Route path="/chats" render={(props) => (<MainPage {...props} tab = "chats"/>)} />
        <Route path="/files" render={(props) => (<MainPage {...props} tab = "files"/>)} />
        <Route path="/stats" render={(props) => (<MainPage {...props} tab = "stats"/>)} />
        <Route path="/signup" render={(props) => (<SignUp {...props} />)} />
      </Switch>
    </Router>

  );
}

export default App
