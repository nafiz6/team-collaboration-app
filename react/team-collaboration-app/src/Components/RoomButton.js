import React from 'react'
import '../MyStyles.css'

const RoomButton = (props) => 
{
    return (
        <button className='roomButton-Style'>{props.name}</button>
    )
}

export default RoomButton;