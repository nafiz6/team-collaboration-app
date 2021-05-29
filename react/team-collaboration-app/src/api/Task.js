
let createTask = ( task, workspaceId ) => {
    const reqOpts = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            ...task,
            Subtasks : [],
            Deadline: "2022-01-01T06:00:00+06:00"
        })
    };

    console.log("adding task " , reqOpts.body);
    fetch(`http://localhost:8080/api/task/${workspaceId}`, reqOpts)
        .then(response => response.json())
        .then(data => console.log(data));

}

export { createTask };