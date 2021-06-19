import axios from "axios";

let getUserDetails = async (uid) =>{
	let res = await axios.get(`http://localhost:8080/api/user-details/${uid}`)
	return res.data

}


export { getUserDetails };