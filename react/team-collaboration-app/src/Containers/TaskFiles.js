import axios from 'axios'
import React, { useEffect, useState } from 'react'
import Deadline from '../Components/Deadline'
import SubtaskButton from '../Components/SubtaskButton'
import '../MyStyles.css'
import { Link } from 'react-router-dom'
import { getTaskFiles } from '../api/file'
import { Accordion, AccordionTab } from 'primereact/accordion';

const TaskFiles = (props) => {

    const [taskView, setTaskView] = useState(<div></div>)
    const[task, setTask] = useState(
		{
			taskname: "Files",
			files: []
		}
    )

    useEffect(() => {
		console.log(task)
		let filesView = task.files.map(file=>{
			return <a href={file.Url} className="files-anchor">
					<div className="files-file">
						<i className="pi pi-file-o" style={{'fontSize': '4em'}}></i>
						<br/>
						{file.FileName}
					</div>
					</a>
		})

		let currentTaskView  = <AccordionTab header="Files in this task">
                    <div className="files-files-list">
                        {filesView}
                    </div>
                        </AccordionTab>
        setTaskView(
			currentTaskView
        );

    }, [task])

    useEffect(async () =>{
        try {

            let taskFiles = await getTaskFiles(props.taskId)
            setTask({
				taskname: "Files",
				files: taskFiles
			})

        }
        catch(err){
            console.log(err)
        }


    }, [])

	return (<Accordion >
		{taskView}
		</Accordion>
		)

}

export default TaskFiles;