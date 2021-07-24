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
import TAKA from './Taka';
import { InputTextarea } from 'primereact/inputtextarea';


const CreateTaskButton = (props) => {
    const [displayBasic, setDisplayBasic] = useState(false);

    const history = useHistory();
    const [position, setPosition] = useState('center');
    const [task, setTask] = useState({
        Name: '',
        Deadline: '',
        ManMonthRate: 0,
        OverheadPercentage: 0,
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
        console.log(e)


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
    /*
    <h5>Budget</h5>
    <Dropdown options={[10, 100, 1000, 5000, 10000]} value={task.Budget} onChange={(e) => {
        handleChange({
            target: {
                name: "Budget",
                value: e.value
            }
        })
    }} name="Budget" />
    */

    const CreateTaskForm =
        <div className="create-form">
            <h5>Task Title</h5>
            <InputText className="form-text" value={task.Name} onChange={handleChange} name="Name" />
            <h5>Description</h5>
            <InputTextarea className="form-text" rows={5} cols={30} value={task.Description} onChange={handleChange} name="Description" />
            <h5>Deadline</h5>
            <Calendar className="form-text" value={task.Deadline} onChange={handleChange} name="Deadline"></Calendar>
            {/* <InputText value={task.Deadline} onChange={handleChange} name="Deadline" /> */}

            <h5>Man Month Rate ({TAKA})</h5>
            <InputNumber className="form-text" value={task.ManMonthRate} onChange={(e) => {
                handleChange({
                    target: {
                        name: "ManMonthRate",
                        value: e.value
                    }
                })
            }} />
            <h5>Overhead Percentage</h5>
            <InputNumber  className="form-text" value={task.OverheadPercentage} onChange={(e) => {
                handleChange({
                    target: {
                        name: "OverheadPercentage",
                        value: e.value
                    }
                })
            }}
            />
        </div>


    return (
        <div className="workspace-form">
            
            <Button className="p-button-raised p-button-rounded" label="Add Task" icon="pi pi-plus" onClick={() => onClick('displayBasic')} />
            <Dialog header="Create Task" visible={displayBasic} style={
                {
                    width: '500px',
                    // 'min-width': '300px'
                }
            } footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateTaskForm}
            </Dialog>
        </div>
    )
}

export default CreateTaskButton;