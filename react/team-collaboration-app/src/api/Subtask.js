import axios from "axios";

let createSubtask = (subtask, taskId) => {
    const reqOpts = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            ...subtask,
            "Assigned_users": [],
            "Updates": [],
        })
    };

    if (subtask.Budget === 0) return;

    console.log("subtask ", subtask);
    fetch(`http://localhost:8080/api/subtask/${taskId}`, reqOpts)
        .then(response => response.json())
        .then(data => console.log(data));

}

let addUpdate = async (text, subtaskId) => {
    const reqOpts = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            "Text": text
        })
    };
    console.log(text);


    // await fetch(`http://localhost:8080/api/subtask-updates/${subtaskId}`, reqOpts)
    // .then(response => response.json())
    // .then(data => console.log(data));

    await axios.post(`http://localhost:8080/api/subtask-updates/${subtaskId}`, JSON.stringify({
        "Text": text
    }), { withCredentials: true });
    // .then(response => response.json())
    // .then(data => console.log(data));


}

export { createSubtask, addUpdate };