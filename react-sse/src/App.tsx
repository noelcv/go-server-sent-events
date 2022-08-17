import { useState, useEffect, useRef, FunctionComponent, MutableRefObject } from 'react'
import { getTime } from './utils/getTime'
import './App.css'

const App: FunctionComponent = () => {
  const [time, setTime] = useState<string>('')
  const eventSrcRef = useRef<EventSource | null>(null);
  
  useEffect(() => {
    eventSrcRef.current = new EventSource("http://localhost:8080/event")
    eventSrcRef.current.onmessage = (e) => {
      setTime(e.data)
    }
    
    eventSrcRef.current.onerror = (e) => {
      console.log('Could not get the time:', e)
    }
    
    //clean up function
    return () => eventSrcRef.current?.close()
  }, [])
  
  return (
    <div className="App">
      <h1>Server-Sent Events</h1>
      <button onClick={getTime}>Get Time</button>
      <h2>{time}</h2>
      
    </div>
  )
}

export default App
