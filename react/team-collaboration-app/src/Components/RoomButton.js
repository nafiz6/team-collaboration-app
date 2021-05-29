import React, { useContext } from 'react'
import { currWSContext } from '../App'
import '../MyStyles.css'
import {stateContext} from "../App"

const RoomButton = (props) => 
{
    const [currWS, setCurrWS] = useContext(currWSContext)
    const [state,setState] = useContext(stateContext)

    return (
        <button className='roomButton-Style'
        onClick = { () =>
            {
                setCurrWS(props.workspace) 
                setState(0)
            }
        }
        >{props.workspace.Name}</button>
    )
}

export default RoomButton;