import React, { useEffect, useState } from 'react'
import '../MyStyles.css'
import SubtaskContainer from '../Containers/SubtaskContainer'
import TaskFiles from '../Containers/TaskFiles'
import axios from 'axios'
import CreateSubtaskButton from '../Components/SubtaskButton'
import { InputText } from 'primereact/inputtext';
import { InputTextarea } from 'primereact/inputtextarea';
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
        <div className="create-form">
            <h5>Subtask Title</h5>
            <InputText className="form-text" value={subtask.Name} onChange={handleChange} name="Name" />
            <h5>Description</h5>
            <InputTextarea className="form-text" rows={5} cols={30} value={subtask.Description} onChange={handleChange} name="Description" />
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
            subtask => <SubtaskContainer key={subtask.id} subtask={subtask} taskId={props.tid} />)

    }

    return (
        <div class="taskPageContainer">


            <div className="taskPage-Style">
                <TaskFiles taskId={props.tid} />
                {/*<CreateSubtaskButton taskId={props.tid} />*/}
                <h1>{props.taskname}</h1>
                <p className="task-page-description">{props.description}</p>
                <Deadline time={props.deadline}></Deadline>

                <div className="add-subtask-button">
                    <Button className="p-button-raised" label="Add Subtask" onClick={() => onClick('displayBasic')} />
                </div>


                <h2>Subtasks</h2>

                {subtasks ? staskArr : <h2 style={{
                    'height': '50px'
                }} className="centered">No Subtasks Yet</h2>}
            </div>

            <Dialog header="Create Subtask" visible={displayBasic} style={
                {
                    width: '500px',
                    // 'min-width': '300px'
                }} footer={renderFooter('displayBasic')} onHide={() => onHide('displayBasic')}>
                {CreateSubtaskForm}
            </Dialog>
        </div>
    )
}

export default TaskDetailPage;