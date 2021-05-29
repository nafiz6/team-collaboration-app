import React from "react"
import '../MyStyles.css'
import TaskContainer from "./TaskContainer";
import TaskPage from "./TaskPage";

const WorkContainer = () => 
{
    const tasks = [
        <TaskContainer name="Prince of Persia" time="3"/>,
        <TaskContainer name="GTA Vice City" time="4"/>,
        <TaskContainer name="Mafia" time="5"/>,
        <TaskContainer name="NFS: Most Wanted" time="6"/>
    ]

    const taskPage = <TaskPage name="Among Us" time="5"/>

    return(
        <div className='work-Style'>
            {tasks}
            {/*taskPage*/} 
        </div>
    );
}

export default WorkContainer;