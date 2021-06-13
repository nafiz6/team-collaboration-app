import React from 'react'
import '../MyStyles.css'
import { Link } from 'react-router-dom'

const FileButton = () => {
    return (
        <Link to="/files">
            <button className="navBarButton-Style"
                onClick={() => {
                    /* Set current Tab as File */
                }
                }>Files</button>
        </Link>
    )
}

export default FileButton;