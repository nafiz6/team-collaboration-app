import React, { useState, useEffect } from 'react';
import { Accordion, AccordionTab } from 'primereact/accordion';
import '../MyStyles.css'
import { FileUpload } from 'primereact/fileupload';
import { workspaceFileUpload, getWorkspaceFiles, getTaskFilesOfWorkspace } from '../api/file.js';

const FilePage = (props) => 
{
    const [taskView, setTaskView] = useState(<div></div>)

    const[tasks, setTasks] = useState(
        //tasks: []
         [

        ]
    )


    useEffect(() => {
        setTaskView(
            tasks.map(task=>{
                let filesView = task.files.map(file=>{
                    return <a href={file.Url} class="files-anchor">
                                <div className="files-file">
                                    <i className="pi pi-file-o" style={{'fontSize': '4em'}}></i>
                                    <br/>
                                    {file.FileName}
                                </div>
                            </a>
                })
                return <AccordionTab header={task.taskname}>
                    <div className="files-files-list">
                        {filesView}
                    </div>
                        </AccordionTab>
            })
        );

    }, [tasks])

    useEffect(async () =>{
        try {
            let workspaceFiles = await getWorkspaceFiles("60ca3b1640dfba660867877a") // workspace id
            let workspaceTaskFiles = {
                taskname: "General",
                files : workspaceFiles
            }
            console.log(workspaceTaskFiles);

            let taskFiles = await getTaskFilesOfWorkspace("60ca3b1640dfba660867877a")
            setTasks(prevTasks => [
                workspaceTaskFiles,
                ...taskFiles,
                ...prevTasks
            ])

        }
        catch(err){
            console.log(err)
        }


    }, [])


    const onUpload = (e) => {
        e.files.map(file => {
            let fileDetails = {
                filename: file.name,
                workspaceId: "60ca3b1640dfba660867877a" // workspace id
            }
            console.log(file)
            workspaceFileUpload(file, fileDetails);

        })

    }


    return (

        <div className="files-page">
            <div className="files-groups-list">
                <div className="files-groups">
                    <div className="p-d-flex p-flex-column">
                <Accordion multiple activeIndex={[0]}>
                        {taskView}
                </Accordion>
                    </div>
                </div>

            </div>

            <div className="files-upload-area">
                <FileUpload  name="files[]" url="http://localhost:8080/api/upload-file/" 
                onUpload={onUpload} multiple 
                maxFileSize={99000000}
                emptyTemplate={<p className="p-m-0">Drag and drop files to here to upload.</p>} />
            </div>

        </div>
    )
}

export default FilePage;