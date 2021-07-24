import axios from 'axios';
import React, { useCallback, useEffect, useState } from 'react'
import '../MyStyles.css'
import { Chart } from 'primereact/chart';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { ProgressBar } from 'primereact/progressbar';
import { Button } from 'primereact/button';
import { Dialog } from 'primereact/dialog';
import { MultiSelect } from 'primereact/multiselect';
import { RadioButton } from 'primereact/radiobutton';
import { Dropdown } from 'primereact/dropdown';
import { roles } from '../api/Workspace';
import { Avatar } from 'primereact/avatar';

const StatPage = (props) => {

    //get workspace id from url later
    // const workspaceId = "60ca3b1640dfba660867877a";



    //create api to get users not in workspace/do it in frontend


    const [changes, setChanges] = useState(0)   //use to refetch data after changes without refreshing page
    const [workspaceBudget, setWorkspaceBudget] = useState([]);
    const [myUserDetails, setMyUserDetails] = useState({

        role: 100   //placeholder for when userDetails arent loaded
    });
    const [chartData, setChartData] = useState(null)
    const [tasksSpendingTable, setTasksSpendingTable] = useState(null)
    const [showAddUsers, setShowAddUsers] = useState(false)
    const [position, setPosition] = useState('center');
    const [usersTable, setUsersTable] = useState(null)
    const [usersToAdd, setUsersToAdd] = useState([
        // {
        //     uid: "",
        //     role: 0,
        //     Name: ""    //to help in display here
        // }
    ])
    const [allUsers, setAllUsers] = useState([
        // {
        //     id: "",
        //     Name: ""
        // }
    ])
    const [editRoles, setEditRoles] = useState(false)
    const [editRolesData, setEditRolesData] = useState([{
        uid: "",
        role: 0
    }])

    const [userIDToRemove, setUserIDToRemove] = useState(null)



    const dialogFuncMap = {
        'addUsers': setShowAddUsers,
    }

    const fetchWorkspaceTasksSpending = async () => {
        const workspaceId = props.ws;
        let res = await axios.get(`http://localhost:8080/api/task/${workspaceId}`)
        setWorkspaceBudget(res.data);
        res.data.map(w=>{
            const oneDay = 24 * 60 * 60 * 1000; // hours*minutes*seconds*milliseconds
            let createdDate = new Date(w.Date_created);
            let deadline = new Date(w.Deadline);
            let now = new Date();

            if (w.ManMonthRate == 0){
                console.log("Null")
                w.ManMonthRate = 100;
            }
            else{
                console.log("NOT NULL")
            }

            let projectDays = Math.round(Math.abs((deadline - createdDate) / oneDay)); 
            let daysFromCreation = Math.round(Math.abs((now - createdDate) / oneDay)); 
                        
            w.Budget = w.ManMonthRate * w.Assigned_users.length * projectDays;
            w.Spent = w.ManMonthRate * w.Assigned_users.length * daysFromCreation;
        })

        console.log(res.data)

        setTasksSpendingTable(res.data.map(w => ({
            ...w,
            progress: ((w.Spent / w.Budget) * 100).toFixed(2),
            spentString: w.Spent.toString() + "/" + w.Budget.toString()
        })))
        setChartData({
            labels: res.data.map(w => w.Name),
            datasets: [
                {
                    data: workspaceBudget.map(w => w.Total_spent)
                }
            ]
        })
    }

    const fetchMyDetails = async () => {

        //call this func after workspace details

        if (usersTable) {
            let res = await axios.get(`http://localhost:8080/api/my-details`, { withCredentials: true });
            // console.log(res.data);


            let workspaceUser = usersTable.find(u => u._id === res.data.id)

            setMyUserDetails({
                ...res.data,
                role: workspaceUser?.role ?? 100    //temp fix, later ill only get workspaces that this user is in
            })


        }




        //add workspace role to userDetails object


    }

    const fetchWorkspaceUserTasks = async () => {
        const workspaceId = props.ws;
        let res = await axios.get(`http://localhost:8080/api/workspace-user-tasks/${workspaceId}`)

        // let users = res.data.map(async u => {
        //     let details = await axios.get(`http://localhost:8080//api/user-details/${u._id}`)
        //     return {
        //         ...u,
        //         ...details
        //     }
        // })

        setUsersTable(res.data.map(w => ({
            ...w,
            countTasks: w.tasks.length,
        })))
    }

    const fetchAllUsers = async () => {

        let res = await axios.get('http://localhost:8080/api/users')

        if (usersTable) {
            setAllUsers(res.data.filter(r => usersTable.findIndex(u => u._id === r.id) < 0).map(r => ({
                ...r,
                uid: r.id,
                role: 2
            })));
        }

    }

    // console.log(allUsers);







    const addUsersToWorkspace = async () => {
        const workspaceId = props.ws;

        usersToAdd.forEach(async u => {
            let res = await axios.post(`http://localhost:8080/api/assign-workspace/${workspaceId}`, JSON.stringify(u));
            setChanges(s => s + 1);
        })


    }

    const setWorkspaceUserRoles = async () => {
        const workspaceId = props.ws;

        await editRolesData.forEach(async u => {
            let res = await axios.post(`http://localhost:8080/api/set-workspace-user-role/${workspaceId}`, JSON.stringify({
                uid: u.uid,
                role: u.role
            }))

            console.log("done")
            setChanges(s => s + 1)
        })





    }

    const removeUserFromWorkspace = async () => {
        const workspaceId = props.ws;


        let res = await axios.post(`http://localhost:8080/api/remove-workspace-user/${workspaceId}`, JSON.stringify({
            uid: userIDToRemove,
        }))

        console.log("done")
        setChanges(s => s + 1)






    }
    // console.log(editRolesData);
    // console.log(usersTable);

    const handleEditRoles = async () => {
        if (!editRoles) {

            setEditRolesData(usersTable.map(u => ({
                ...u,
                uid: u._id
            })))
            setEditRoles(true);
        }
        else {

            await setWorkspaceUserRoles();


            setEditRoles(false);



        }

    }

    const handleRemoveUser = (userID) => {

        setUserIDToRemove(userID)

    }

    useEffect(() => {
        if (props.ws) {
            fetchMyDetails();

            fetchWorkspaceTasksSpending();
            fetchWorkspaceUserTasks();
            // console.log("called")
        }


    }, [props.ws, changes])
    useEffect(() => {
        fetchAllUsers();    //maybe call this on add users button click only LATER

        if (usersTable) {
            fetchMyDetails();
        }
    }, [usersTable, props.ws])

    useEffect(() => {
        if (userIDToRemove) {
            console.log("removing userID", userIDToRemove)
            removeUserFromWorkspace()
        }

    }, [userIDToRemove])



    const onClick = (name, position) => {
        dialogFuncMap[`${name}`](true);

        if (position) {
            setPosition(position);
        }
    }

    const addingUsers = (name) => {
        dialogFuncMap[`${name}`](false);
        addUsersToWorkspace()

    }

    const onHide = (name) => {
        dialogFuncMap[`${name}`](false);
    }

    // console.log(usersToAdd)



    const addUsersForm =
        <div>
            <h5>Select Users to add to Workspace</h5>
            <MultiSelect value={usersToAdd} options={allUsers} onChange={(e) => {
                console.log(e.value)
                setUsersToAdd(e.value)
                // console.log(e);

            }} optionLabel="Name" />

            {usersToAdd.map((u, i) => (
                <span key={u.id}>
                    <h5>{u.Name}</h5>
                    <Dropdown placeholder="Select role" value={u.role} options={roles} optionLabel="label" optionValue="id" onChange={(e) => {

                        setUsersToAdd(users => {
                            users[i].role = parseInt(e.value)
                            return [...users]   //this is to force re render
                        })
                        // console.log(usersToAdd)

                    }} />

                </span>
            ))}




        </div>;


    const renderFooter = (name) => {
        return (
            <div>
                <Button label="Add Users" icon="pi pi-check" onClick={() => addingUsers(name)} autoFocus />
            </div>
        );
    }

    // console.log(tasksSpendingTable)







    return (

        <div>
            <br></br>
            <br></br>
            <br></br>
            {/* <div>Stat Page</div>

            <h2>Total spent: {workspaceBudget[0]?.Total_spent}</h2>
            <h2>Total workspace budget: {workspaceBudget[0]?.Task_budget}</h2> */}

            <DataTable value={tasksSpendingTable} emptyMessage="No tasks yet" header={<h2>Task Spending</h2>}>
                <Column field="Name" header="Task"></Column>
                <Column header="Spending" body={(rowData) => <ProgressBar value={isNaN(rowData?.progress) ? 0 : rowData?.progress} />}></Column>
                {/* <Column field="Task_budget" header="Budget"></Column>
                <Column field="Total_spent" header="Spent"></Column> */}
                <Column field="spentString" header=""></Column>


            </DataTable>
            {myUserDetails.role < 2 ? <>
                <Button label="Add Users" icon="pi pi-external-link" onClick={() => onClick("addUsers")} />  <Dialog header="Add users to workspace" visible={showAddUsers} style={{ width: '50vw' }} footer={renderFooter('addUsers')} onHide={() => onHide('addUsers')}>
                    {addUsersForm}
                </Dialog>
            </> : null}






            <DataTable value={usersTable} emptyMessage="No users yet" header={<h2>Workspace Users</h2>}>
                <Column body={(rowData) => <Avatar label={rowData.name[0]} size="large" image={rowData.dp} />}></Column>
                <Column field="name" header="User"></Column>
                <Column
                    header={
                        <span>
                            Role {myUserDetails.role < 2 ? <Button label={(editRoles ? "Save" : "Edit")} onClick={handleEditRoles} /> : null}
                        </span>
                    }

                    body={(rowData) => (
                        editRoles ?
                            <Dropdown value={editRolesData?.find(e => e.uid === rowData._id)?.role} options={roles} onChange={(e) => {
                                console.log(e.value);
                                setEditRolesData(r => {
                                    r.find(e => e.uid === rowData._id).role = e.value;
                                    return [...r];

                                })
                                // console.log(e);

                            }} optionLabel="label" optionValue="id" />

                            :
                            roles.find(r => r.id === rowData.role)?.label ?? "Team Member"




                    )}></Column>
                <Column field="countTasks" header="Tasks"></Column>


                {myUserDetails.role < 2 ? <Column header="" body={(rowData) => <Button label="Remove" onClick={() => handleRemoveUser(rowData._id)} />}></Column> : null}


            </DataTable>

        </div>

    )
}

export default StatPage;