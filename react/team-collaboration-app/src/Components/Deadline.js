import React from 'react'
import '../MyStyles.css'

const Deadline = (props) => {
    const monthNames = ["January", "February", "March", "April", "May", "June",
        "July", "August", "September", "October", "November", "December"
    ];

    const colors = [
        {
            diff: 3 * 24 * 60 * 60 * 1000,
            color: "red"
        },
        {
            diff: 7 * 24 * 60 * 60 * 1000,
            color: "yellow"
        },
        {
            diff: -1,
            color: "green"
        }
    ]

    let time = new Date(props.time);

    let now = new Date();

    let diff = time.getTime() - now.getTime();

    let color;
    if(diff < colors[0].diff) color = colors[0].color;
    else if(diff < colors[1].diff) color = colors[1].color;
    else color = colors[2].color
    time = monthNames[time.getMonth()].slice(0, 3) + " " + time.getDate();
    return (
        <div style={
            {
                "backgroundColor": color
            }
        } className='deadline-Style'>
              <i className="pi pi-clock"></i>  
              <p>{time}</p>
              </div>
    )
}

export default Deadline;