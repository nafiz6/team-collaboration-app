import axios from "axios";


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

export { fileUpload };