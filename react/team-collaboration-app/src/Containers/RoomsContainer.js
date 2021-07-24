import React, { useState, useEffect, useRef  } from "react"
import RoomButton from "../Components/RoomButton";
import '../MyStyles.css'
import CreateWorkspaceButton from "../Components/CreateWorkspaceButton";
import axios from 'axios'
import { Button } from 'primereact/button';
import { Toast } from 'primereact/toast';
import { ConfirmDialog, confirmDialog } from 'primereact/confirmdialog';


const RoomsContainer = (props) => {
    const [visible, setVisible] = useState(false);
    const toast = useRef(null);
    const [workspace, setWorkspace] = useState(null);

    const getWorkspace = async () => {
        let res = await axios.get(`http://localhost:8080/api/workspace/${props.project.id}`, { withCredentials: true })
        setWorkspace(res.data);
    }

    useEffect(() => {
        if (props.project)
            getWorkspace();
    }, [props.project])

    useEffect(() => {
        if (workspace) fetchMyDetails();
    }, [workspace])

    const [myUserDetails, setMyUserDetails] = useState({ role: 100 });

    if (!props.project) {
        return <div className='rooms-Style'></div>
    }



    const fetchMyDetails = async () => {

        //call this func after workspace details

        let users = await axios.get(`http://localhost:8080/api/workspace-user-tasks/${workspace[0].id}`)

        // if (usersTable) {
        let res = await axios.get(`http://localhost:8080/api/my-details`, { withCredentials: true });
        // console.log(res.data);


        let workspaceUser = users.data.find(u => u._id === res.data.id)

        setMyUserDetails(workspaceUser);

    }




    const deleteProject = async () => {

        // console.log(usersToAddToTask);

        await axios.post(`http://localhost:8080/api/delete-project/${props.project.id}`);

        window.location.reload();

    }
    const accept = () => {
        console.log("ACCEPTED")
        deleteProject();
    }

    const reject = () => {
        console.log("REJECTED")
    }

    const confirm2 = () => {
        confirmDialog({
            message: 'Do you want to delete this Project?',
            header: 'Delete Confirmation',
            icon: 'pi pi-info-circle',
            acceptClassName: 'p-button-danger',
            accept,
            reject
        });
    };

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

            <div className='rooms-Style'>{rooms}
            {myUserDetails.role < 1 ? 
                <div className="addMemberToTaskButton">
                            <Button icon="pi pi-trash" className="p-button-danger p-button-raised p-button-rounded" onClick={(e) => {
                                e.preventDefault();
                                confirm2();
                              //deleteProject();
                                
                            }} />
                </div>
                : null}
            </div>    
        );

    }
    else {
        return (
            <div className='rooms-Style'>
                <div className='projName-Style'>{props.project.Name}</div>
                <CreateWorkspaceButton projectId={props.project.id} />
                {myUserDetails.role < 1 ? 
                <div className="addMemberToTaskButton">
                            <Button icon="pi pi-trash" className="p-button-danger p-button-raised p-button-rounded p-button-lg" onClick={(e) => {
                                e.preventDefault();
                                confirm2();
                               //deleteProject();
                            }} />
                </div>
                : null}
            </div>
        )
    }


}

export default RoomsContainer;