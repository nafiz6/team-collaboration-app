import axios from "axios";


let register = async (user) => {

	let res = await axios.post(`http://localhost:8080/api/register`, 
		JSON.stringify(user), {withCredentials: true})

	return res.data;
}

let login = async (user) => {

	let res = await axios.post(`http://localhost:8080/api/login`, 
		JSON.stringify({
			Username: user.username,
			Password: user.password
		}), {withCredentials: true})

	return res.data;
}

let secretPage = async (user) => {
	let res = await axios.get(`http://localhost:8080/api/secret-page`, 
		{withCredentials: true})

	return res.data;

}

let logout = async (user) => {
	await axios.get(`http://localhost:8080/api/logout`,
		{withCredentials: true})
		

}

export { login, register, secretPage, logout };