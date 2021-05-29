import React, { useContext } from 'react'
import { currWSContext } from '../App'
import '../MyStyles.css'

const RoomButton = (props) => 
{
    const [currWS, setCurrWS] = useContext(currWSContext)

    return (
        <button className='roomButton-Style'
        onClick = { () =>
            {
                setCurrWS(props.workspace) 
                console.log(currWS)
            }
        }
        >{props.workspace.Name}</button>
    )
}

export default RoomButton;