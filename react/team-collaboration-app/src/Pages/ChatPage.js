import React, { useState, useEffect } from 'react'
import '../MyStyles.css'
import { InputTextarea } from 'primereact/inputtextarea';
import { InputText } from 'primereact/inputtext';
import { connect, sendMsg, socket } from '../api/chat.js';
import { getUserDetails } from '../api/user.js';
import { Card } from 'primereact/card';
import axios from 'axios';
import { DataScroller } from 'primereact/datascroller';

const ChatPage = (props) => {

    const [messageText, setMessageText] = useState('')
    const [chatList, setChatList] = useState([])
    const [chatMessages, setChatMessages] = useState([])
    const [chatUsers, setChatUsers] = useState([])
    const [myUserDetails, setMyUserDetails] = useState(null);

    const setChats = async () => {

        console.log("calling")
        console.log(chatList);


        let parsedMessage = chatList[chatList.length - 1];

        if (!chatUsers.some(u => u.id === parsedMessage.ClientId)) {

            if (parsedMessage.ClientId) {
                let user = await getUserDetails(parsedMessage.ClientId);
                setChatUsers(users => [...users, user])
            }
        }


        setChatMessages(chatList.filter(c => c.Type === "Text").map(c => {
            console.log(chatUsers);
            let user = chatUsers.find(u => u.id === c.ClientId);



            return <div className="p-mb-2 chat-message-container">

                {
                    user?.id === myUserDetails?.id ?
                        <div className="p-grid ">
                            <div className="p-col-1">
                                <img className="chat-img" src={user?.Dp} />
                            </div>


                            <div className="p-col-10"
                            >
                                <div className="chat-username">
                                    {user?.Username}
                                </div>
                                <div className="chat-message-box">
                                    {c.Body}
                                </div>
                            </div>

                        </div>

                        :
                        <div className="p-grid" >
                            <div className="p-col-1">
                                <img className="chat-img" src={user?.Dp} />
                            </div>

                            <div className="p-col-10">
                                <div className="chat-username">
                                    {user?.Username}
                                </div>
                                <div className="chat-message-box">
                                    {c.Body}
                                </div>
                            </div>
                        </div>

                }

            </div>
        }))
        // }



    }


    useEffect(() => {



        if (chatList.length > 0) {
            setChats();
        }


    }, [chatList, chatUsers, props.ws])


    useEffect(() => {
        connect(props.ws); // workspace id


        socket.onmessage = async (msg) => {
            let parsedMessage = JSON.parse(msg.data);

            console.log(parsedMessage)
            console.log(chatUsers)

            setChatList(chats => [...chats, parsedMessage]);
            // chatList.push(parsedMessage);
            //might have bug since using chatUsers here, debug later





        };

        fetchMyDetails();



        fetchWorkspaceChats();






        return function cleanup() {
            socket.close();
        }
    }, [props.ws])

    const fetchMyDetails = async () => {

        //call this func after workspace details


        let res = await axios.get(`http://localhost:8080/api/my-details`, { withCredentials: true });
        console.log(res.data);

        setMyUserDetails(res.data)







        //add workspace role to userDetails object


    }





    const fetchWorkspaceChats = async () => {
        const workspaceId = props.ws;
        let res = await axios.get(`http://localhost:8080/api/workspace-chats/${workspaceId}`);

        let users = []

        console.log(res.data);  

        if(!res.data) {
            return;
        }


        //populate users array
        res.data.forEach(async chat => {
            if (!users.some(u => u.id === chat.ClientId)) {
                let user = await getUserDetails(chat.ClientId);
                console.log(user);
                users.push(user);
            }
        })
        setChatUsers(users);


        console.log(res.data)


        setChatList(res.data);

    }

    const sendMessage = e => {
        e.preventDefault();
        if (messageText === '') return;
        sendMsg(
            {
                "Type": "Text",
                "Body": messageText,
                "WorkspaceId": props.ws
            }
        );
        setMessageText('')
    }


    return (
        <div className="chat-page">
            <div className="chat-message-area">
                <div className="chat-message-groups">
                    <div className="chat-container">
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