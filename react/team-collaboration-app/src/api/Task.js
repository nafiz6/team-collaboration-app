import axios from "axios";

let createTask = async (task, workspaceId) => {
    // const reqOpts = {
    //     method: 'POST',
    //     headers: { 'Content-Type': 'application/json' },
    //     body: JSON.stringify({
    //         ...task,
    //         Subtasks: [],
    //         Deadline: "2022-01-01T06:00:00+06:00"
    //     })
    // };

    let res = await axios.post(`http://localhost:8080/api/task/${workspaceId}`, JSON.stringify({
        ...task,
        Subtasks: [],
        // Deadline: "2022-01-01T06:00:00+06:00"
    }), { withCredentials: true })

    // console.log("adding task ", reqOpts.body);
    // fetch(`http://localhost:8080/api/task/${workspaceId}`, {
    //     ...reqOpts,
    //     credentials: "include"
    // })
    //     .then(response => response.json())
    //     .then(data => console.log(data));

}

export { createTask };