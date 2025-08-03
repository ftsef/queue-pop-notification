import { useEffect, useState, useRef } from 'react';
import { EventsOn, EventsOff } from '../../wailsjs/runtime/runtime';

interface EventData {
  level?: string;
  message?: string;
  timestamp?: string;
  filename?: string;
  success?: boolean;
  status?: string;
  queueType?: string;
}

interface LogEvent {
  id: number;
  type: string;
  level: string;
  message: string;
  timestamp: string;
  data: EventData;
}

export default function EventLog(): JSX.Element {
  const [events, setEvents] = useState<LogEvent[]>([]);
  const scrollRef = useRef<HTMLDivElement>(null);

  // Auto-scroll to bottom when new events are added
  useEffect(() => {
    if (scrollRef.current) {
      scrollRef.current.scrollTop = scrollRef.current.scrollHeight;
    }
  }, [events]);

  useEffect(() => {
    console.log('Setting up event listeners...');
    
    // Listen for PvP queue pop events
    const handlePvPQueuePop = (data: EventData) => {
      console.log('Received PvP Queue Pop event:', data);
      const event: LogEvent = {
        id: Date.now() + Math.random(),
        type: 'pvp-queue-pop',
        level: 'info',
        message: `PvP Queue Pop Detected -> ${data.queueType || 'Unknown Queue'}`,
        timestamp: data.timestamp || new Date().toLocaleTimeString(),
        data: data
      };
      setEvents(prev => {
        console.log('Adding event to state, current events:', prev.length);
        return [...prev, event];
      });
    };

    EventsOn('OnPvPQueuePop', handlePvPQueuePop);

    // Cleanup event listeners on unmount
    return () => {
      console.log('Cleaning up event listeners...');
      setEvents([]);
      EventsOff('OnPvPQueuePop');
    };
  }, []);

  const getEventIcon = (type: string, level: string): string => {
    switch (type) {
      case 'pvp-queue-pop':
        return '🎮';
      case 'queue-pop':
        return '🎮';
      default:
        switch (level) {
          case 'error':
            return '❌';
          case 'warn':
          case 'warning':
            return '⚠️';
          case 'success':
            return '✅';
          case 'info':
          default:
            return 'ℹ️';
        }
    }
  };

  const getEventColor = (level: string): string => {
    switch (level) {
      case 'error':
        return 'text-red-600 bg-red-50 border-red-200';
      case 'warn':
      case 'warning':
        return 'text-yellow-600 bg-yellow-50 border-yellow-200';
      case 'success':
        return 'text-green-600 bg-green-50 border-green-200';
      case 'info':
      default:
        return 'text-blue-600 bg-blue-50 border-blue-200';
    }
  };

  return (
    <div className="flex flex-col h-full">
      
      <div
        ref={scrollRef}
        className="flex-1 overflow-y-auto border border-gray-300 rounded-lg bg-white shadow-inner"
        style={{ maxHeight: 'calc(100vh - 200px)' }}
      >
        {events.length === 0 ? (
          <div className="p-4 text-center text-gray-500">
            Waiting for events...
          </div>
        ) : (
          <div className="p-2 space-y-1">
            {events.map((event) => (
              <div
                key={event.id}
                className={`p-2 rounded border-l-4 ${getEventColor(event.level)}`}
              >
                <div className="flex items-center gap-2">
                  <span className="font-mono text-xs text-gray-500 w-16 flex-shrink-0">
                    {event.timestamp}
                  </span>
                  <span className="text-xs flex-shrink-0">
                    {getEventIcon(event.type, event.level)}
                  </span>
                  <span className="text-sm font-medium flex-1 min-w-0 truncate">
                    {event.message}
                  </span>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}