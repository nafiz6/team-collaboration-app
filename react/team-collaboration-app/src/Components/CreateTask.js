import React, { useState } from 'react'
import '../MyStyles.css'
import { Dialog } from 'primereact/dialog';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { InputNumber } from 'primereact/inputnumber';
import { createTask } from '../api/Task.js';
import { useHistory } from "react-router-dom";
import { Calendar } from 'primereact/calendar';
import { Dropdown } from 'primereact/dropdown';

const CreateTaskButton = (props) => {
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


        let { name, value } = e.target;
        console.log(e.target)


        if (name === "Deadline") {
            value = new Date(value);

            let month = value.getMonth() + 1;
            let date = value.getDate();

            if (month < 10) {
                month = `0${month}`
            }
            if (date < 10) {
                date = `0${date}`
            }


            value = `${value.getFullYear()}-${month}-${date}T06:00:00+06:00`
            console.log(value);
        }


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
            <Calendar value={task.Deadline} onChange={handleChange} name="Deadline"></Calendar>
            {/* <InputText value={task.Deadline} onChange={handleChange} name="Deadline" /> */}
            <h5>Budget</h5>
            
            <Dropdown options={[10, 100, 1000, 5000, 10000]} value={task.Budget} onChange={(e) => {
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

            <Button className=" newTaskButton" label="Add Task" icon="pi pi-plus"  onClick={() => onClick('displayBasic')} />
            <Dialog header="Create Task" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateTaskForm}
            </Dialog>
        </div>
    )
}

export default CreateTaskButton;