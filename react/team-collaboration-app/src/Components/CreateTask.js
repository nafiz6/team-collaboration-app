import React, { useState } from 'react'
import '../MyStyles.css'
import { Dialog } from 'primereact/dialog';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { InputNumber } from 'primereact/inputnumber';
import { createTask } from '../api/Task.js';
import { useHistory } from "react-router-dom";

const CreateTaskButton = (props) => 
{
    const [displayBasic, setDisplayBasic] = useState(false);

    const history = useHistory();
    const [position, setPosition] = useState('center');
    const [task, setTask] = useState({
        Name: '',
        Deadline: '',
        Budget: 0,
        Description: ''
    });

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

    const creatingTask = async (name) => {
        dialogFuncMap[`${name}`](false);

        
        await createTask(task, props.workspaceId)
        window.location.reload();


    }

    const renderFooter = (name) => {
        return (
            <div>
                <Button label="Create" icon="pi pi-check" onClick={() => creatingTask(name)} autoFocus />
            </div>
        );
    }

    const handleChange = e => {
            const { name, value } = e.target;
            setTask(prevState => ({
                ...prevState,
                [name]: value
            }));
        };

    const CreateTaskForm =
            <div>
                <h5>Task Name</h5>
                <InputText value={task.Name} onChange={handleChange} name="Name" />
                <h5>Description</h5>
                <InputText value={task.Description} onChange={handleChange} name="Description" />
                <h5>Deadline</h5>
                <InputText value={task.Deadline} onChange={handleChange} name="Deadline" />
                <h5>Budget</h5>
                <InputNumber value={task.Budget} onChange={(e)=>{
                    handleChange({
                        target: {
                            name: "Budget",
                            value: e.value
                        }
                    })
                }} name="Budget" />
            </div>


    return (
        <div className="workspace-form">

            <Button label="Add Task" onClick={() => onClick('displayBasic')} />
            <Dialog header="Create Task" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateTaskForm}
            </Dialog>
        </div>
    )
}

export default CreateTaskButton;