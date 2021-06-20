import axios from 'axios';
import React, { useCallback, useEffect, useState } from 'react'
import '../MyStyles.css'
import { Chart } from 'primereact/chart';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { ProgressBar } from 'primereact/progressbar';

const StatPage = (props) => {

    //get workspace id from url later
    // const workspaceId = "60ca3b1640dfba660867877a";



    const [workspaceBudget, setWorkspaceBudget] = useState([]);
    const [chartData, setChartData] = useState(null)
    const [tasksSpendingTable, setTasksSpendingTable] = useState(null)
    const [usersTable, setUsersTable] = useState(null)

    const fetchWorkspaceTasksSpending = async () => {
        const workspaceId = props.match.params.wsid;
        let res = await axios.get(`http://localhost:8080/api/workspace-tasks-budget-breakdown/${workspaceId}`)
        setWorkspaceBudget(res.data);

        setTasksSpendingTable(res.data.map(w => ({
            ...w,
            progress: (w.Total_spent / w.Task_budget) * 100,
            spentString: w.Total_spent.toString() + "/" + w.Task_budget.toString()
        })))
        setChartData({
            labels: res.data.map(w => w.Task_name),
            datasets: [
                {
                    data: workspaceBudget.map(w => w.Total_spent)
                }
            ]
        })
    }

    const fetchWorkspaceUserTasks = async () => {
        const workspaceId = props.match.params.wsid;
        let res = await axios.get(`http://localhost:8080/api/workspace-user-tasks/${workspaceId}`)

        setUsersTable(res.data.map(w => ({
            ...w,
            countTasks: w.tasks.length,
        })))
    }

    useEffect(() => {

        fetchWorkspaceTasksSpending();
        fetchWorkspaceUserTasks();
    }, [props.match.params.wsid])




    return (

        <div>
            <br></br>
            <br></br>
            <br></br>
            {/* <div>Stat Page</div>

            <h2>Total spent: {workspaceBudget[0]?.Total_spent}</h2>
            <h2>Total workspace budget: {workspaceBudget[0]?.Task_budget}</h2> */}

            <DataTable value={tasksSpendingTable} header={<h2>Task Spending</h2>}>
                <Column field="Task_name" header="Task"></Column>
                <Column header="Spending" body={(rowData) => <ProgressBar value={rowData.progress} />}></Column>
                {/* <Column field="Task_budget" header="Budget"></Column>
                <Column field="Total_spent" header="Spent"></Column> */}
                <Column field="spentString" header=""></Column>


            </DataTable>

            <DataTable value={usersTable} header={<h2>Workspace Users</h2>}>
                <Column field="name" header="User"></Column>
                <Column header="Role" body={(rowData) => rowData.role}></Column>
                <Column field="countTasks" header="Tasks"></Column>


            </DataTable>

        </div>

    )
}

export default StatPage;