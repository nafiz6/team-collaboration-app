import React, { useState, useEffect } from 'react'
import '../MyStyles.css'
import { InputTextarea } from 'primereact/inputtextarea';
import { InputText } from 'primereact/inputtext'; 
import {connect, sendMsg, socket} from '../api/chat.js';
import { getUserDetails} from '../api/user.js';
import { Card } from 'primereact/card';

const ChatPage = (props) => 
{

    const [messageText, setMessageText] = useState('')
    const [chatList, setChatList] = useState([])
    const [chatMessages, setChatMessages] = useState([])

    useEffect(() => {
        connect("60ca3b1640dfba660867877a");
    }, [])

    const sendMessage = e => {
        e.preventDefault();
        if (messageText === '') return;
        sendMsg(
            {
                "Type": "Text",
                "Body": messageText,
                "WorkspaceId": "60ca3b1640dfba660867877a"
            }
        );
        setMessageText('')
    }

    socket.onmessage = async (msg) => {
        let parsedMessage = JSON.parse(msg.data);
        chatList.push(parsedMessage);
        if (parsedMessage.ClientId) {
            let user = await getUserDetails(parsedMessage.ClientId)
            console.log(user);
            setChatMessages(chatList.filter(c => c.type === "Text").map(c => {
                return <div className="p-mb-2 chat-message-container">
                    <div className="p-grid" >
                        <div className="p-col-1">
                            <img className="chat-img" src={user.Dp} />
                        </div>

                        <div className="p-col-10">
                            <div className="chat-username">
                                {user.Username}
                            </div>
                            <div className="chat-message-box">
                                {c.body}
                            </div>
                        </div>
                    </div>
                </div>
            }))
        }
    };


    return (
        <div className="chat-page">
            <div className="chat-message-area">
                <div className="chat-message-groups">
                    <div className="p-d-flex p-flex-column">
                        {chatMessages}
                    </div>
                </div>

            </div>

            <form onSubmit={sendMessage}>
                <InputText onkeypress={sendMessage} className="chat-input-area" value={messageText} onChange={(e) => setMessageText(e.target.value)} />
            </form>
        </div>
    )
}

export default ChatPage;