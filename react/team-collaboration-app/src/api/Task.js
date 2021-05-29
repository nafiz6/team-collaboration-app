
let createTask = ( task, workspaceId ) => {
    const reqOpts = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            ...task,
            "Subtasks" : []
        })
    };

    console.log("task " , task);
    fetch(`http://localhost:8080/api/task/${workspaceId}`, reqOpts)
        .then(response => response.json())
        .then(data => console.log(data));

}

export { createTask };