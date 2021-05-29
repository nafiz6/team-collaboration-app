import React, { useState } from 'react'
import '../MyStyles.css'
import { Dialog } from 'primereact/dialog';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { InputNumber } from 'primereact/inputnumber';
import { addUpdate } from '../api/Subtask.js';

const AddUpdate = (props) => 
{
    const [displayBasic, setDisplayBasic] = useState(false);
    const [position, setPosition] = useState('center');
    const [subtask, setTask] = useState({
        Text: ''
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

    const creatingUpdate = (name) => {
        dialogFuncMap[`${name}`](false);
        addUpdate(subtask.Text, props.user, props.subtaskId)

    }

    const renderFooter = (name) => {
        return (
            <div>
                <Button label="Create" icon="pi pi-check" onClick={() => creatingUpdate(name)} autoFocus />
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

    const CreateUpdateForm =
            <div>
                <h5>Update </h5>
                <InputText value={subtask.Text} onChange={handleChange} name="Text" />
            </div>


    return (
        <div className="workspace-form">

            <Button label="Add Update" onClick={() => onClick('displayBasic')} />
            <Dialog header="Add Update" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateUpdateForm}
            </Dialog>
        </div>
    )
}

export default AddUpdate;