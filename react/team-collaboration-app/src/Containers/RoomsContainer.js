import React, { useState, useEffect } from "react"
import RoomButton from "../Components/RoomButton";
import '../MyStyles.css'
import CreateWorkspaceButton from "../Components/CreateWorkspaceButton";
import axios from 'axios'


const RoomsContainer = (props) => {

    const [workspace, setWorkspace] = useState(null);

    const getWorkspace = async () => {
        let res = await axios.get(`http://localhost:8080/api/workspace/${props.project.id}`)
        setWorkspace(res.data);
    }

    useEffect(() => {
        if (props.project)
            getWorkspace();
    }, [props.project])

    if (!props.project) {
        return <div className='rooms-Style'></div>
    }


    if (workspace) {

        const createRooms = workspace.map(
            ws => <RoomButton key={ws.id} workspace={ws} projId={props.project.id} />
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
    else {
        return (
            <div className='rooms-Style'>
                <div className='projName-Style'>{props.project.Name}</div>
                <CreateWorkspaceButton projectId={props.project.id} />
            </div>
        )
    }


}

export default RoomsContainer;