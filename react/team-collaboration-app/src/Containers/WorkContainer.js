import React from "react"
import '../MyStyles.css'
import TaskContainer from "./TaskContainer";

const WorkContainer = () => 
{
    const tasks = [
        <TaskContainer name="Prince of Persia" time="3"/>,
        <TaskContainer name="GTA Vice City" time="4"/>,
        <TaskContainer name="Mafia" time="5"/>,
        <TaskContainer name="NFS: Most Wanted" time="6"/>
    ]


    return(
        <div className='work-Style'>
            {tasks}
        </div>
    );
}

export default WorkContainer;