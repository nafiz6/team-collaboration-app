import React, { useState } from 'react'
import '../MyStyles.css'
import { Dialog } from 'primereact/dialog';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { InputNumber } from 'primereact/inputnumber';
import { createSubtask } from '../api/Subtask.js';

const CreateSubtaskButton = (props) => 
{
    const [displayBasic, setDisplayBasic] = useState(false);
    const [position, setPosition] = useState('center');
    const [subtask, setTask] = useState({
        Name: '',
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

    const creatingSubtask = (name) => {
        dialogFuncMap[`${name}`](false);
        createSubtask(subtask, props.taskId)

    }

    const renderFooter = (name) => {
        return (
            <div>
                <Button label="Create" icon="pi pi-check" onClick={() => creatingSubtask(name)} autoFocus />
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
                <h5>Subtask Name</h5>
                <InputText value={subtask.Name} onChange={handleChange} name="Name" />
                <h5>Description</h5>
                <InputText value={subtask.Description} onChange={handleChange} name="Description" />
                <h5>Budget</h5>
                <InputNumber value={subtask.Budget} onChange={(e)=>{
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

            <Button label="Add Subtask" onClick={() => onClick('displayBasic')} />
            <Dialog header="Create Subtask" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateTaskForm}
            </Dialog>
        </div>
    )
}

export default CreateSubtaskButton;