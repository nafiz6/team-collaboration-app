import React, { useEffect, useState } from 'react'
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
import axios from 'axios';


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

    useEffect(() => {
        if(props.isEdit) {
            console.log(props.task)
            setTask(props.task)
        }

    }, [props.task, props.isEdit])

    // console.log(task)

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

        if (props.isEdit) {
            await axios.post(`http://localhost:8080/api/update-task/${props.taskId}`, JSON.stringify(task))

        }
        else {
            await createTask(task, props.workspaceId)
        }

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

            if (value.getTime() - new Date().getTime() < 0) {

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

            <h3>Task Details</h3>
            <h5>Task Title</h5>
            <InputText className="form-text" value={task.Name} onChange={handleChange} name="Name" />
            <h5>Description</h5>
            <InputTextarea className="form-text" rows={5} cols={30} value={task.Description} onChange={handleChange} name="Description" />
            <h5>Deadline</h5>
            <Calendar minDate={new Date()} className="form-text" value={task.Deadline} onChange={handleChange} name="Deadline"></Calendar>
            {/* <InputText value={task.Deadline} onChange={handleChange} name="Deadline" /> */}

            <h3>Budgeting</h3>

            <p className="form-description">Enter Estimated Man Month Rate and Overhead Percentage to get an estimated budget for this task</p>

            <h5 className="form-label">Man Month Rate ({TAKA})</h5>
            <p className="form-description">Average Cost spent per team member per month</p>
            <InputNumber className="form-text" value={task.ManMonthRate} onChange={(e) => {
                handleChange({
                    target: {
                        name: "ManMonthRate",
                        value: e.value
                    }
                })
            }} />
            <h5 className="form-label">Overhead Percentage (%)</h5>
            <p className="form-description">Overhead costs associated with equipment, fringe benefits etc. (This is added as a % of man month rate)</p>
            <InputNumber className="form-text" value={task.OverheadPercentage} onChange={(e) => {
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

            <Button label="Add Task" icon="pi pi-plus" className="p-button-raised p-button-rounded" onClick={(e) => {
                e.preventDefault();
                e.stopPropagation();
                onClick('displayBasic');


            }} />
            <Dialog dismissableMask={true} header="Create Task" visible={displayBasic} style={
                {
                    width: '500px',
                    // 'min-width': '300px'
                }
            } footer={renderFooter('displayBasic')} onHide={() => {
                console.log("HIDDEEE")
                onHide('displayBasic')
            }} >
                {CreateTaskForm}
            </Dialog>
        </div>
    )
}

export default CreateTaskButton;