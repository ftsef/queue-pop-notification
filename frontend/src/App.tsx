import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import EventLog from "./components/event-log";
import Config from "./components/config";

function App() {
    const [activeTab, setActiveTab] = useState<'events' | 'config'>('events');

    return (
        <div >

        <div id="App" className="h-screen flex flex-col px-2">
                <h1 className="ml-4 text-2xl font-bold text-gray-800">
                    Queue Pop Notification
                </h1>
                
                {/* Tab Navigation */}
                <div className="flex border-b border-gray-200 ml-4 mr-4 mt-4">
                    <button
                        onClick={() => setActiveTab('events')}
                        className={`px-4 py-2 text-sm font-medium border-b-2 transition-colors ${
                            activeTab === 'events'
                                ? 'border-blue-500 text-blue-600'
                                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                        }`}
                    >
                        Event Log
                    </button>
                    <button
                        onClick={() => setActiveTab('config')}
                        className={`px-4 py-2 text-sm font-medium border-b-2 transition-colors ${
                            activeTab === 'config'
                                ? 'border-blue-500 text-blue-600'
                                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                        }`}
                    >
                        Konfiguration
                    </button>
                </div>

            <main className="flex-1 overflow-scroll min-h-0 scrollbar-hide">
                {activeTab === 'events' && <EventLog />}
                {activeTab === 'config' && <Config />}
            </main>
        </div>
        </div>
    )
}

export default App
