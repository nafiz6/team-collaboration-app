import React, { useState } from 'react'
import '../MyStyles.css'
import { Dialog } from 'primereact/dialog';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { createWorkspace } from '../api/Workspace.js';

const CreateWorkspaceButton = (props) => 
{
    const [displayBasic, setDisplayBasic] = useState(false);
    const [position, setPosition] = useState('center');
    const [workspaceName, setWorkspaceName] = useState('');

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

    const creatingWorkspace = (name) => {
        dialogFuncMap[`${name}`](false);
        createWorkspace(workspaceName, props.projectId)

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
            </div>


    return (
        <div className="workspace-form">

            <Button label="Add Workspace" className="p-button-text" onClick={() => onClick('displayBasic')} />
            <Dialog header="Create Workspace" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateWorkspaceForm}
            </Dialog>
        </div>
    )
}

export default CreateWorkspaceButton;