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
      
      {/* Scrollable Content Area with hidden scrollbar and shadows */}
      <div className="flex-1 relative shadow-inner overflow-hidden">
        {/* Top shadow gradient */}
        <div className="absolute top-0 left-0 right-0 h-4 bg-gradient-to-b from-gray-200 to-transparent z-10 pointer-events-none opacity-0 transition-opacity duration-300" 
             id="top-shadow"></div>
        
        {/* Bottom shadow gradient */}
        <div className="absolute bottom-0 left-0 right-0 h-4 bg-gradient-to-t from-gray-200 to-transparent z-10 pointer-events-none opacity-100 transition-opacity duration-300" 
             id="bottom-shadow"></div>
        
        <div 
          ref={scrollRef} 
          className="h-full overflow-y-scroll scrollbar-hide"
          style={{
            scrollbarWidth: 'none', /* Firefox */
            msOverflowStyle: 'none' /* Internet Explorer 10+ */
          }}
          onScroll={(e) => {
            const element = e.target as HTMLElement;
            const topShadow = document.getElementById('top-shadow');
            const bottomShadow = document.getElementById('bottom-shadow');
            
            if (topShadow && bottomShadow) {
              // Show top shadow when scrolled down
              topShadow.style.opacity = element.scrollTop > 0 ? '1' : '0';
              
              // Show bottom shadow when not at bottom
              const isAtBottom = element.scrollTop + element.clientHeight >= element.scrollHeight - 1;
              bottomShadow.style.opacity = isAtBottom ? '0' : '1';
            }
          }}
        >
          {events.length === 0 ? (
            <div className="p-4 text-center text-gray-500">
              Waiting for events...
            </div>
          ) : (
            <div className="px-2 py-2 space-y-1">
              {events.map((event) => (
                <div
                  key={event.id}
                  className={`p-3 mx-2 rounded ${getEventColor(event.level)}`}
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
    </div>
  );
}