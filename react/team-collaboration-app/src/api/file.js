import axios from "axios";

let getTaskFilesOfWorkspace = async (workspaceId) => {
	let files = [];
	let tasks = await axios.get(`http://localhost:8080/api/task/${workspaceId}`
				)
			.catch(err=>{
				throw err.response.data
			});
	
	
	files = await Promise.all(tasks.data.map( async(task) =>{
		console.log(task)
			let file =  await getTaskFiles(task.id)
				.catch(err=>{
					throw err
				})
			return {
				taskname: task.Name,
				files: file ? file : []
			}
		}))
	console.log(files)

	files = files.filter(file => file.files != null && file.files.length > 0)

	return files;

}

let getTaskFiles = async(taskId) =>{

	let res = await axios.get(`http://localhost:8080/api/task-file/${taskId}`, 
		{withCredentials: true})
			.catch(err=>{
				console.log(err)
				throw err.response.data
			})

	return res.data;
}

let getWorkspaceFiles = async(workspaceId) =>{

	let res = await axios.get(`http://localhost:8080/api/workspace-file/${workspaceId}`, 
		{withCredentials: true})
			.catch(err=>{
				throw err.response.data
			})

	return res.data;
}

let workspaceFileUpload = async (file, details) => {
	let form = new FormData();
	form.append('file', file);
	form.append('filename', details.filename);
	form.append('workspaceId', details.workspaceId);

	let res = await axios.post(`http://localhost:8080/api/workspace-file/`, 
		form, {withCredentials: true})
			.catch(err=>{
				throw err.response.data
			})

	return res.data;
}


let fileUpload = async (file) => {
	let form = new FormData();
	form.append('file', file);

	let res = await axios.post(`http://localhost:8080/api/upload-file/`, 
		form, {withCredentials: true})
			.catch(err=>{
				throw err.response.data
			})

	return res.data;
}

export { fileUpload, workspaceFileUpload, getWorkspaceFiles, getTaskFilesOfWorkspace };