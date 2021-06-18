import axios from "axios";


let register = async (user) => {

	let res = await axios.post(`http://localhost:8080/api/register`, 
		JSON.stringify(user), {withCredentials: true})
			.catch(err=>{
				throw err.response.data
			})

	return res.data;
}

let login = async (user) => {

		let res = await axios.post(`http://localhost:8080/api/login`,
			JSON.stringify({
				Username: user.username,
				Password: user.password
			}), { withCredentials: true })
			.catch(err=>{
				throw err.response.data
			})

		return res.data;
}

let secretPage = async () => {
	let res = await axios.get(`http://localhost:8080/api/secret-page`, 
		{withCredentials: true})
			.catch(err=>{
				throw err.response.data
			})

	return res.data;

}

let logout = async () => {
	await axios.get(`http://localhost:8080/api/logout`,
		{withCredentials: true})
			.catch(err=>{
				throw err.response.data
			})
		

}

export { login, register, secretPage, logout };