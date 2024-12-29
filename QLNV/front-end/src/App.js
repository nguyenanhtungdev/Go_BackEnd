import { useEffect, useState } from 'react';
import './App.css';

function App() {

  const [task, setTasks] = useState([])
  useEffect(()=>{
    fetch("http://localhost:8080/products")
    .then((response) => response.json())
    .then((data) => {
      console.log(data)
      setTasks(data)
    })
    .catch((error) => console.error("Error fetching tasks:", error));
  },[])
  return (
    <div className="App">
      <header className="App-header">
        <h1>Task List</h1>
        <ul>
          {
            task.map((task)=>(
              <li>{task.productName}</li>
            ))
          }
        </ul>
      </header>
    </div>
  );
}

export default App;
