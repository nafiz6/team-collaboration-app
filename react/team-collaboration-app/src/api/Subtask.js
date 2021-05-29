
let createSubtask = ( subtask, taskId ) => {
    const reqOpts = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            ...subtask,
            "Assigned_users" : [],
            "Updates" : [],
        })
    };

    if (subtask.Budget === 0)return;

    console.log("subtask " , subtask);
    fetch(`http://localhost:8080/api/subtask/${taskId}`, reqOpts)
        .then(response => response.json())
        .then(data => console.log(data));

}

let addUpdate = ( text, user, subtaskId) => {
    const reqOpts = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            "User": user,
            "Text": text
        })
    };
    console.log(text, user);


    fetch(`http://localhost:8080/api/subtask-updates/${subtaskId}`, reqOpts)
        .then(response => response.json())
        .then(data => console.log(data));


}

export { createSubtask, addUpdate };