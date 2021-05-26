import logo from './logo.svg';
import './App.css';
import React, { useEffect, useState } from 'react'
import ReactDOM from 'react-dom'

function App() {


  const [message, setMessage] = useState()

  // console.log("AAAAA");
  // let message = ''

  useEffect(() => {
    fetch('http://localhost:8080')
      .then(res => {
        // message = res;
        // console.log(res);
        // setMessage(res);
        // console.log(message);
        // console.log(message);
        return res.text()
      }).then(res2 => setMessage(res2))
  });

  return (

    <div className="App">
      <header className="App-header">
        {/* <img src={logo} className="App-logo" alt="logo" /> */}
        {/* <p>{message}</p> */}
        <div>{message}</div>
      </header>
    </div>
  );
}

export default App;
