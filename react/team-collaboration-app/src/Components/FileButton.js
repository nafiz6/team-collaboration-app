import React from 'react'
import '../MyStyles.css'
import { Link } from 'react-router-dom'

const FileButton = (props) => {
    return (
        <Link to = {`/project/${props.id}/ws/${props.wsid}/files`}>
            <button className={props.tab === "files" ? "navBarButton-Style-select" : "navBarButton-Style"}><i className="pi pi-file"></i>Files</button>
        </Link>
    )
}

export default FileButton;