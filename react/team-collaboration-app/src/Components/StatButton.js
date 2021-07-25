import React from 'react'
import '../MyStyles.css'
import { Link } from 'react-router-dom'

const StatButton = (props) => {
    return (
        <Link to= {`/project/${props.id}/ws/${props.wsid}/stats`}>
            <button className={props.tab === "stats" ? "navBarButton-Style-select" : "navBarButton-Style"}><i className="pi pi-cog"></i>Manage</button>
        </Link>
    )
}

export default StatButton;