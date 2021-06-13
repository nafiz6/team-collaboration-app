import React from 'react'
import '../MyStyles.css'
import { Link } from 'react-router-dom'

const ChatButton = () => 
{
    return (
        <Link to="/chats">
        <button className="navBarButton-Style"
        onClick={() => {
            /* Set current Tab as Chat */
        }
        }>Chats</button>
        </Link>
    )
}

export default ChatButton;