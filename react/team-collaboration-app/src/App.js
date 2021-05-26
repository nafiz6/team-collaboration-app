import logo from './logo.svg';
import './App.css';

function App() {
  let message = ''
  fetch('http://localhost:8080')
    .then(res=> message = res);
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        {message}
      </header>
    </div>
  );
}

export default App;
