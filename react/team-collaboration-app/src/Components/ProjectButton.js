import React from 'react'
import { Link } from 'react-router-dom'
import '../MyStyles.css'


const ProjectButton = (props) => {

    let url = window.location.href.split("/")

    // console.log(url)

    // let id = url[url.length - 1];
    return (

        <Link to={`/project/${props.id}`}>
            <button className={url[4] === props.id ? 'projectButton-Style-select' : 'projectButton-Style'}>{props.name}</button>
        </Link >
    )
}

export default ProjectButton;