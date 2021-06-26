import axios from "axios";

let createWorkspace = async (name, projectId, description) => {
    const reqOpts = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            "Name": name,
            "Users": [],
            "Description": description
        })
    };

    console.log("Name ", name);

    let res = await axios.post(`http://localhost:8080/api/workspace/${projectId}`, JSON.stringify({
        "Name": name,
        "Users": [],
        "Description": description
    }), { withCredentials: true });

    // fetch(`http://localhost:8080/api/workspace/${projectId}`, {
    //     ...reqOpts,
    //     credentials: "include"
    // })
    //     .then(response => response.json())
    //     .then(data => console.log(data));

}


export const roles = [
    {
        label: "Project Leader",
        id: 0
    },
    {
        label: "Workspace Leader",
        id: 1
    },
    {
        label: "Team Member",
        id: 2
    },
]
export { createWorkspace };