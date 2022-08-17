<script>
  import { onMount } from "svelte";
  let time = "";
  
  onMount(() => {
    const eventSrc = new EventSource("http://localhost:8080/event");
    eventSrc.onmessage = function(event) {
      console.log(event);
      time = event.data;
    }
    
    eventSrc.onerror = function(event) {
      console.log(event);
    }
  })
  
  async function getTime() {
    console.log("clicking")
    const res = await fetch("http://localhost:8080/time");
    console.log(res)
    if (res.status !== 200) {
      console.log("could not connect to server")
    } else {
      console.log("OK");
    }
  } 
  
</script>

<main>
  <h1>Server-Sent Events</h1>
  <button on:click="{ getTime }">Get Time</button>
  <p>Time : { time }</p>
</main>

