import axios from 'axios';
import React, { useCallback, useEffect, useState } from 'react'
import '../MyStyles.css'
import { Chart } from 'primereact/chart';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';

const StatPage = (props) => {

    //get workspace id from url later
    const workspaceId = "60ca3b1640dfba660867877a";

    const [workspaceBudget, setWorkspaceBudget] = useState([]);
    const [chartData, setChartData] = useState(null)
    const [tableData, setTableData] = useState(null)

    const dataFetch = async () => {
        let res = await axios.get(`http://localhost:8080/api/workspace-tasks-budget-breakdown/${workspaceId}`)
        setWorkspaceBudget(res.data);
        console.log(workspaceBudget);

        setTableData(workspaceBudget)
        setChartData({
            labels: workspaceBudget.map(w => w.Task_name),
            datasets: [
                {
                    data: workspaceBudget.map(w => w.Total_spent)
                }
            ]
        })
    }

    let lightOptions = {
        plugins: {
            legend: {
                labels: {
                    color: '#495057'
                }
            }
        }
    };

    useEffect(() => {
        dataFetch();
    }, [])




    return (

        <div>
            <div>Stat Page</div>

            <h2>Total spent: {workspaceBudget[0]?.Total_spent}</h2>
            <h2>Total workspace budget: {workspaceBudget[0]?.Task_budget}</h2>

            <DataTable value={workspaceBudget}>
                <Column field="Task_name" header="Task"></Column>
                <Column field="Task_budget" header="Budget"></Column>
                <Column field="Total_spent" header="Spent"></Column>
            </DataTable>

            {/* <Chart type="pie" data={chartData} options={lightOptions} style={{ position: 'relative', width: '40%' }} /> */}

        </div>

    )
}

export default StatPage;