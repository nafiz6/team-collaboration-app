import React from 'react'
import '../MyStyles.css'
import { Link } from 'react-router-dom'

const ChatButton = (props) => 
{
    return (
        <Link to = {`/project/${props.id}/ws/${props.wsid}/chats`}>
        <button className="navBarButton-Style">Chats</button>
        </Link>
    )
}

export default ChatButton;