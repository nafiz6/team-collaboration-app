import React, { useState } from 'react'
import '../MyStyles.css'
import { Dialog } from 'primereact/dialog';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { createProject } from '../api/Project.js';

const ProjectAddButton = () => 
{
    const [displayBasic, setDisplayBasic] = useState(false);
    const [position, setPosition] = useState('center');
    const [projectName, setProjectName] = useState('');

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

    const creatingProject = (name) => {
        dialogFuncMap[`${name}`](false);
        createProject(projectName)

    }

    const renderFooter = (name) => {
        return (
            <div>
                <Button label="Create" icon="pi pi-check" onClick={() => creatingProject(name)} autoFocus />
            </div>
        );
    }

    const CreateProjectFrom =
            <div>
                <h5>Project Name</h5>
                <InputText value={projectName} onChange={(e) => setProjectName(e.target.value)} />
            </div>


    return (
        <div>
            <button className='projectAddButton-Style' position={position} onClick={() => onClick('displayBasic')}>+</button>
            <Dialog header="Create Project" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateProjectFrom}
            </Dialog>
        </div>
    )
}

export default ProjectAddButton;