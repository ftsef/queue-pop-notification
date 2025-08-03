import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import EventLog from "./components/event-log";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    return (
        <div id="App" className="h-screen flex flex-col">
                <h1 className="ml-4 text-2xl font-bold text-gray-800">
                    Queue Pop Notification
                </h1>
            
            <main className="flex-1 p-4 overflow-hidden">
                <EventLog />
            </main>
        </div>
    )
}

export default App
