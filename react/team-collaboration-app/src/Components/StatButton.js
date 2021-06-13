import React from 'react'
import '../MyStyles.css'
import { Link } from 'react-router-dom'

const StatButton = () => {
    return (
        <Link to="/stats">
            <button className="navBarButton-Style"
                onClick={() => {
                    /* Set current Tab as Stat */
                }
                }>Stats</button>
        </Link>
    )
}

export default StatButton;