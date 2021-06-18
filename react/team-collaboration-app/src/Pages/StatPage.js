import axios from 'axios';
import React, { useCallback, useEffect, useState } from 'react'
import '../MyStyles.css'
import { Chart } from 'primereact/chart';

const StatPage = (props) => {

    //get workspace id from url later

    const workspaceId = "60ca3b1640dfba660867877a";
    const [workspaceBudget, setWorkspaceBudget] = useState({
        Total_budget: 0,
        Total_spent: 0
    });
    const [chartData, setChartData] = useState(null)

    const dataFetch = useCallback(async () => {
        let res = await axios.get(`http://localhost:8080/api/workspace-tasks-budget-breakdown/60ca3b1640dfba660867877a`)
        setWorkspaceBudget(res.data);
        console.log(workspaceBudget);
        setChartData({
            labels: workspaceBudget.map(w => w.Task_name),
            datasets: [
                {
                    data: workspaceBudget.map(w => w.Total_spent)
                }
            ]
        })
    }, [])

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
    }, [dataFetch])




    return (

        <div>
            <div>Stat Page</div>

            <h2>Total spent: {workspaceBudget.Total_spent}</h2>
            <h2>Total workspace budget: {workspaceBudget.Total_budget}</h2>

            <Chart type="pie" data={chartData} options={lightOptions} style={{ position: 'relative', width: '40%' }} />

        </div>

    )
}

export default StatPage;