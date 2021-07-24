import React, { useState } from 'react'
import '../MyStyles.css'
import { Dialog } from 'primereact/dialog';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { createWorkspace } from '../api/Workspace.js';

const CreateWorkspaceButton = (props) => {
    const [displayBasic, setDisplayBasic] = useState(false);
    const [position, setPosition] = useState('center');
    const [workspaceName, setWorkspaceName] = useState('');
    const [description, setDescription] = useState('');

    const dialogFuncMap = {
        'displayBasic': setDisplayBasic,
    }

    const onClick = (name, position) => {
        dialogFuncMap[`${name}`](true);

        if (position) {
            setPosition(position);
        }
    }

    const onHide = (name) => {
        dialogFuncMap[`${name}`](false);
    }

    const creatingWorkspace = async (name) => {
        dialogFuncMap[`${name}`](false);
        await createWorkspace(workspaceName, props.projectId, description)
        window.location.reload();
        
    }

    const renderFooter = (name) => {
        return (
            <div>
                <Button label="Create" icon="pi pi-check" onClick={() => creatingWorkspace(name)} autoFocus />
            </div>
        );
    }

    const CreateWorkspaceForm =
        <div>
            <h5>Workspace Name</h5>
            <InputText value={workspaceName} onChange={(e) => setWorkspaceName(e.target.value)} />
            <h5>Workspace Description</h5>
            <InputText value={description} onChange={(e) => setDescription(e.target.value)} />
        </div>


    return (
        <div className="workspace-form">

            
            <Button label="Workspace" icon="pi pi-plus" className="p-button-raised p-button-rounded p-button-sm" onClick={() => onClick('displayBasic')} />
            
            <Dialog header="Create Workspace" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateWorkspaceForm}
            </Dialog>
        </div>
    )
}

export default CreateWorkspaceButton;