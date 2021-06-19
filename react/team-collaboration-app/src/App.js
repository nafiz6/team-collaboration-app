import './MyStyles.css'
import React from 'react'
import MainPage from './Pages/MainPage';
import Login from './Pages/Login';
import SignUp from './Pages/SignUp';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

// import "primereact/resources/themes/bootstrap4-dark-blue/theme.css"
// import "primereact/resources/primereact.min.css"
// import "primeicons/primeicons.css"

function App() {

  return (
 
    <Router>
      <Switch>
        <Route path="/" exact render={(props) => (<Login {...props} />)} />
        <Route path="/project" exact render={(props) => (<MainPage {...props} tab = "tasks"/> ) }/>
        <Route path="/project/:id" exact render={(props) => (<MainPage {...props} tab = "tasks"/> )}/>
        <Route path="/project/:id/ws/:wsid" exact render={(props) => (<MainPage {...props} tab = "tasks"/> )}/>
        <Route path="/project/:id/ws/:wsid/tasks" exact render={(props) => (<MainPage {...props} tab = "tasks"/>)} />
        <Route path="/project/:id/ws/:wsid/chats" exact render={(props) => (<MainPage {...props} tab = "chats"/>)} />
        <Route path="/project/:id/ws/:wsid/files" exact render={(props) => (<MainPage {...props} tab = "files"/>)} />
        <Route path="/project/:id/ws/:wsid/stats" exact render={(props) => (<MainPage {...props} tab = "stats"/>)} />
        <Route path="/signup" render={(props) => (<SignUp {...props} />)} />
      </Switch>
    </Router>

  );
}

export default App
