import React from "react"
import RoomButton from "../Components/RoomButton";
import '../MyStyles.css'
import CreateWorkspaceButton from "../Components/CreateWorkspaceButton";


const RoomsContainer = (props) => {

    if (props.project) {

        if(props.project.Workspaces){

        const createRooms = props.project.Workspaces.map(
            workspace => <RoomButton key={workspace.id} workspace={workspace} />
        )

        const rooms = [
            <div className='projName-Style'>{props.project.Name}</div>,
            ...createRooms,
            <CreateWorkspaceButton projectId={props.project.id} />
        ]

        return (
            <div className='rooms-Style'>{rooms}</div>
        );

        }
        else{
            return(
                <div>
                    <div className='projName-Style'>{props.project.Name}</div>
                    <CreateWorkspaceButton projectId={props.project.id} />
                </div>
            )
        }
    }
    else{ 
        return <div className='rooms-Style'></div>
    }

}

export default RoomsContainer;