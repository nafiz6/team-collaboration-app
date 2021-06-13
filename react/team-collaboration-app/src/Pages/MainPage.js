import React from 'react'
import '../MyStyles.css'
import HeaderContainer from '../Containers/HeaderContainer'
import RoomsContainer from '../Containers/RoomsContainer'
import NavBar from '../Containers/NavBar'
import WorkContainer from '../Containers/WorkContainer'
import ProjectContainer from '../Containers/ProjectContainer'

const MainPage = (props) => {

    // DUMMY DATA
    var Users = [{id: 1, Name: "Somik"}, {id: 2, Name: "Sabab"}]
    var Ass_Users = [{id: 1, Name: "Somik", HasCompleted: true}, {id: 2, Name: "Sabab", HasCompleted: false}]
    var Updates = [{id: 1, User: Users[0], Text: "70% Done", Timestamp: ""},
                   {id: 2, User: Users[1], Text: "30% Done", Timestamp: ""}]
    var Subtasks = [{id: 1, Name: "Weapon Design", Description: "Building the weapon", Budget: 50000, Assigned_Users: Ass_Users, Updates: Updates},
                    {id: 2, Name: "Boss Design", Description: "Building the boss", Budget: 100000, Assigned_Users: Ass_Users, Updates: Updates}]
    var Tasks = [{id: 1, Name: "Level 1", Deadline: 10, Budget: 400000, Description: "Make Level 1", Subtask: Subtasks},
                 {id: 2, Name: "Level 2", Deadline: 4, Budget: 350000, Description: "Make Level 2", Subtask: Subtasks}]
    var Workspaces = [{id: 1, Name: "Mafia 1", User: Users, Task: Tasks},
                      {id: 2, Name: "Mafia 2", User: Users, Task: Tasks}]
    var Projects = [{id: 1, Name: "Mafia", Workspaces: Workspaces}, {id: 2, Name: "Most Wanted", Workspaces: Workspaces}]     
    
    // GLOBAL VARIABLES TO SETUP
    /* 1. Current Project
       2. Current Workspace
    */


    return (
        <div className='page-Style'>
            <HeaderContainer />
            <div className='bottom-Style'>
                <ProjectContainer projects = {Projects}/>
                <RoomsContainer project = {Projects[0]}/> {/* This gets current selected project */}
                <div className='taskWork-Style'>
                    <NavBar />
                    <WorkContainer tab = {props.tab} />
                </div>
            </div>
        </div>
    )
}

export default MainPage;