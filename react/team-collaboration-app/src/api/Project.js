import axios from "axios";


let createProject = async (name, description, users) => {
    const reqOpts = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ "Name": name, "Description": description })

    };



    console.log("Name ", name);
    fetch('http://localhost:8080/api/project', {
        ...reqOpts,
        credentials: "include"
    })
        .then(response =>
            // response.json();
            response.json()
        )
        .then(data => {
            console.log(data)
            users.forEach(async uid => {


                let res = await axios.post(`http://localhost:8080/api/assign-projects/${data}`, JSON.stringify({
                    uid: uid,
                    role: 2
                }), { withCredentials: true })

                console.log(res)



            })



        });



}

export { createProject };