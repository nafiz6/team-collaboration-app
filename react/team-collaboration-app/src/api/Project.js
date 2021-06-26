import axios from "axios";


let createProject = async (name, users) => {
    const reqOpts = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ "Name": name })
        
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
        // .then(data => {  //ALREADY DOING THIS IN BACKEND
        //     console.log(data)
        //     users.forEach(async uid => {


        //         let res = await axios.post(`http://localhost:8080/api/assign-projects/${data}`, JSON.stringify({
        //             Uid: uid,
        //             Role: 100
        //         }))

        //         console.log(res)



        //     })



        // });



}

export { createProject };