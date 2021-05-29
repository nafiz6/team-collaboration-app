import React from "react"
import RoomButton from "../Components/RoomButton";
import '../MyStyles.css'
import CreateWorkspaceButton from "../Components/CreateWorkspace";


const RoomsContainer = () => 
{
    const rooms = [
        <div className='projName-Style'>My Project</div>,
        <RoomButton name="ROOM 1"/>,
        <RoomButton name="ROOM 2"/>,
        <RoomButton name="ROOM 3"/>,
        <RoomButton name="ROOM 4"/>,
    ]   
    return(
        <div className='rooms-Style'>{rooms}
            <CreateWorkspaceButton projectId='60b21eb1da6a7a91ae769b3d'/>
        </div>
    );
}

export default RoomsContainer;