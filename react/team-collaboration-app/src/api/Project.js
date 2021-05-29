
let createProject = ( name ) => {
    const reqOpts = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ "Name": name })
    };

    console.log("Name " , name);
    fetch('http://localhost:8080/api/project', reqOpts)
        .then(response => response.json())
        .then(data => console.log(data));

}

export { createProject };