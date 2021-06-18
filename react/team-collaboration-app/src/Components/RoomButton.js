import React from 'react'
import { Link } from 'react-router-dom'
import '../MyStyles.css'

const RoomButton = (props) => 
{
    return (
        <Link to = {`/project/${props.projId}/ws/${props.workspace.id}`}>
        <button className='roomButton-Style'>{props.workspace.Name}</button>
        </Link>
    )
}

export default RoomButton;