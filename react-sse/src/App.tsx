import { useState, useEffect, useRef, FunctionComponent, MutableRefObject } from 'react'

import './App.css'

const App: FunctionComponent = () => {
  const [time, setTime] = useState<string>('')
  const eventSrcRef = useRef<EventSource | null>(null);
  
  const getTime = async () => {
    try {
      console.log('getting time')
      const res = await fetch("http://localhost:8080/time");
      console.log(res, "timmeeeeee")
      return res;
    } catch(err) {
      console.log("error fetching time: ", err)
    }
  }
  
  useEffect(() => {
    eventSrcRef.current = new EventSource("http://localhost:8080/event")
    console.log(eventSrcRef.current, "eventSrcRef.current");
    eventSrcRef.current.onmessage = function(e) {
      console.log(e.data, "e.data");
      setTime(e.data)
    }
    
    eventSrcRef.current.onerror = function (e) {
      console.log('Could not get the time', e)
    }
    
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
