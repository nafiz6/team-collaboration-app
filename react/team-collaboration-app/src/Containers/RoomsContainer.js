import React from "react"
import RoomButton from "../Components/RoomButton";
import '../MyStyles.css'
import CreateWorkspaceButton from "../Components/CreateWorkspace";


const RoomsContainer = (props) => {

    if (props.project) {

        const createRooms = props.project.Workspaces.map(
            workspace => <RoomButton key={workspace.id} workspace={workspace} />
        )


        const rooms = [
            <div className='projName-Style'>{props.project.Name}</div>,
            ...createRooms
        ]
        return (
            <div className='rooms-Style'>{rooms}</div>
        );
    }
    else return <div className='rooms-Style'></div>

}

export default RoomsContainer;