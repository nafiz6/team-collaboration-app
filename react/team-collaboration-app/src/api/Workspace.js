
let createWorkspace = (name, projectId) => {
    const reqOpts = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            "Name": name,
            "Users": [],
            "Tasks": []
        })
    };

    console.log("Name ", name);
    fetch(`http://localhost:8080/api/workspace/${projectId}`, reqOpts)
        .then(response => response.json())
        .then(data => console.log(data));

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