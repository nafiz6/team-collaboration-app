import axios from "axios";


let login = async (user) => {

	let res = await axios.post(`http://localhost:8080/api/login`, 
	JSON.stringify({
		Username: user.username,
		Password: user.password
	}), {withCredentials: true})

	console.log(res);
	return res;
}

export { login };