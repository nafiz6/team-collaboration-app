import React, { useState, useEffect } from 'react'
import '../MyStyles.css'
import { Dialog } from 'primereact/dialog';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { InputNumber } from 'primereact/inputnumber';
import { addUpdate } from '../api/Subtask.js';
import { FileUpload } from 'primereact/fileupload';
import { taskFileUpload } from '../api/file.js';

const AddUpdate = (props) => 
{
    const [displayBasic, setDisplayBasic] = useState(false);
    const [position, setPosition] = useState('center');
    const [filenames, setFilenames] = useState('');
    const [files, setFiles] = useState([]);
    const [subtask, setTask] = useState({
        Text: ''
    });

    useEffect(() => {
        let fileNameList = files.map(file=> file.name)
        console.log("FILELIST")
        if (fileNameList.length > 0){
            setFilenames("Files added: " + fileNameList)
        }
        else{
            setFilenames("")
        }
    }, [files])


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
        //addUpdate(subtask.Text, props.user, props.subtaskId)
        uploadFiles();
    }

    const uploadFiles = () =>{
        files.map(file => {
            let fileDetails = {
                filename: file.name,
                taskId: props.taskId // workspace id
            }
            console.log(file)
            taskFileUpload(file, fileDetails);

        })

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

    const onFileSelect = (e) => {
        console.log(e)
        setFiles([
            ...e.files
        ])

    }


    const CreateUpdateForm =
            <div>
                <h5>Update </h5>
                <InputText value={subtask.Text} onChange={handleChange} name="Text" />
                <h5>File</h5>
                {filenames}
                <FileUpload auto mode="basic" name="files[]" multiple customUpload maxFileSize={99000000} uploadHandler={onFileSelect} />

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