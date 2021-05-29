import React from "react"
import RoomButton from "../Components/RoomButton";
import '../MyStyles.css'


const RoomsContainer = (props) => 
{
    const rooms = [
        <div className='projName-Style'>{props.project?.Name}</div>,
        <RoomButton name="ROOM 1"/>,
        <RoomButton name="ROOM 2"/>,
        <RoomButton name="ROOM 3"/>,
        <RoomButton name="ROOM 4"/>,
    ]   
    return(
        <div className='rooms-Style'>{rooms}</div>
    );
}

export default RoomsContainer;