import React from 'react'
import '../MyStyles.css'

const RoomButton = (props) => 
{
    return (
        <button className='roomButton-Style'>{props.workspace.Name}</button>
    )
}

export default RoomButton;