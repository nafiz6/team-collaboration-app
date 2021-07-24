import React, { useEffect, useState } from 'react'
import '../MyStyles.css'
import SubtaskContainer from '../Containers/SubtaskContainer'
import TaskFiles from '../Containers/TaskFiles'
import axios from 'axios'
import CreateSubtaskButton from '../Components/SubtaskButton'
import { InputText } from 'primereact/inputtext';
import { Button } from 'primereact/button';
import { InputNumber } from 'primereact/inputnumber';
import { Dialog } from 'primereact/dialog';
import Deadline from '../Components/Deadline'



const TaskDetailPage = (props) => {
    const [subtasks, setSubtasks] = useState([]);
    const [displayBasic, setDisplayBasic] = useState(false);
    const [subtask, setSubtask] = useState({
        Name: '',
        Budget: 0,
        Description: ''
    });
    const [changes, setChanges] = useState(0);
    const dialogFuncMap = {
        'displayBasic': setDisplayBasic,
    }

    const onClick = (name, position) => {
        dialogFuncMap[`${name}`](true);

    }

    const onHide = (name) => {
        dialogFuncMap[`${name}`](false);
    }

    const creatingSubtask = (name) => {
        dialogFuncMap[`${name}`](false);
        createSubtask()

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
        setSubtask(prevState => ({
            ...prevState,
            [name]: value
        }));
    };

    const CreateSubtaskForm =
        <div>
            <h5>Subtask Title</h5>
            <InputText value={subtask.Name} onChange={handleChange} name="Name" />
            <h5>Description</h5>
            <InputText value={subtask.Description} onChange={handleChange} name="Description" />
            <h5>Budget</h5>
            <InputNumber value={subtask.Budget} onChange={(e) => {
                handleChange({
                    target: {
                        name: "Budget",
                        value: e.value
                    }
                })
            }} name="Budget" />
        </div>

    const getSubtasks = async () => {

        if (props.tid) {
            let res = await axios.get(`http://localhost:8080/api/subtask/${props.tid}`);
            setSubtasks(res.data);
        }
    }

    useEffect(() => {
        getSubtasks();
    }, [props.tid, changes])


    const createSubtask = async () => {
        let res = await axios.post(`http://localhost:8080/api/subtask/${props.tid}`, JSON.stringify(subtask), { withCredentials: true })

        setChanges(s => s + 1)
        window.location.reload();
    }


    let staskArr;
    if (subtasks) {
        staskArr = subtasks.map(
            subtask => <SubtaskContainer  key={subtask.id} subtask={subtask} taskId={props.tid} />)

    }

    return (
        <div class="taskPageContainer">

<Button label="Add Subtask" onClick={() => onClick('displayBasic')} />
            <div className="taskPage-Style">
                <TaskFiles taskId={props.tid} />
                {/*<CreateSubtaskButton taskId={props.tid} />*/}
                <h1>{props.taskname}</h1>
                <p className="task-page-description">{props.description}</p>
                <Deadline time={props.deadline}></Deadline>
                
                {subtasks ? staskArr : null}
            </div>
            
            <Dialog header="Create Subtask" visible={displayBasic} style={{ width: '50vw' }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateSubtaskForm}
            </Dialog>
        </div>
    )
}

export default TaskDetailPage;