import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import EventLog from "./components/event-log";

function App() {


    return (
        <div >

        <div id="App" className="h-screen flex flex-col px-2">
                <h1 className="ml-4 text-2xl font-bold text-gray-800">
                    Queue Pop Notification
                </h1>

            <main className="flex-1 overflow-scroll min-h-0 scrollbar-hide">
                <EventLog />
            </main>
        </div>
        </div>
    )
}

export default App
