import React, { useCallback, useEffect, useState } from 'react'
import '../MyStyles.css'
import { Dialog } from 'primereact/dialog';
import { Button } from 'primereact/button';
import { MultiSelect } from 'primereact/multiselect';
import { InputText } from 'primereact/inputtext';
import { createProject } from '../api/Project.js';
import axios from 'axios';
import { InputTextarea } from 'primereact/inputtextarea';

const ProjectAddButton = () => {
    const [displayBasic, setDisplayBasic] = useState(false);
    const [position, setPosition] = useState('center');
    const [projectName, setProjectName] = useState('');
    const [description, setDescription] = useState('');
    const [users, setUsers] = useState([]);
    const [allUsers, setAllUsers] = useState([]);


    const dataFetch = useCallback(async () => {
        let res = await axios.get('http://localhost:8080/api/users')

        setAllUsers(res.data);

    }, [])

    useEffect(() => {
        dataFetch();
    }, [dataFetch])

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

    const creatingProject = async (name) => {
        dialogFuncMap[`${name}`](false);
        await createProject(projectName, description,  users)
        window.location.reload();

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
            <h5>Project Title</h5>
            <InputText className="form-text" value={projectName} onChange={(e) => setProjectName(e.target.value)} />
            <h5>Project Description</h5>
            <InputTextarea className="form-text" rows={5} cols={30} value={description} onChange={(e) => setDescription(e.target.value)} />

            <h5>Select Users to add to Project</h5>
            <p className="form-description">You can add members to this project later from the "Stats" Tab</p>
            <MultiSelect placeholder="Select Users" className="form-text" optionLabel="name" value={users} options={allUsers} onChange={(e) => {
                setUsers(e.value)
                console.log(e.value);

            }} optionLabel="Name" optionValue="id" />
        </div>


    return (
        <div>
            <div className = "projectAddButton-Style ">
            <Button icon="pi pi-plus" className="p-button-raised p-button-rounded p-button-lg" onClick={() => onClick('displayBasic')}></Button>
            </div>
            <Dialog header="Create Project" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateProjectFrom}
            </Dialog>
        </div>
    )
}

export default ProjectAddButton;