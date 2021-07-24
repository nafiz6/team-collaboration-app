import React from 'react'
import '../MyStyles.css'

const Time = (props) => {
    const monthNames = ["January", "February", "March", "April", "May", "June",
        "July", "August", "September", "October", "November", "December"
    ];


    let time = new Date(props.time);

    let now = new Date();

    let diff = time.getTime() - now.getTime();

    let color;

    time = time.getDate() + " " +   monthNames[time.getMonth()].slice(0, 3) + " " + time.getHours() + ":" + time.getMinutes();
    return (
        <div
         className={props.className}>
              <p>{time}</p>
              </div>
    )
}

export default Time;