import React from 'react'
import { Link } from 'react-router-dom'
import '../MyStyles.css'

const RoomButton = (props) => {

    let url = window.location.href.split("/")

    let id = url[6];
    return (
        <Link to={`/project/${props.projId}/ws/${props.workspace.id}`}>
            <button className={id === props.workspace.id ? 'roomButton-Style-select' : 'roomButton-Style'}>{props.workspace.Name}</button>
        </Link>
    )
}

export default RoomButton;