<script>
  import { onMount } from "svelte";
  let time = "";
  
  onMount (() => {
    const eventSrc = new EventSource("http://localhost:8080/event");
    eventSrc.onmessage = function (event) {
      time = event.data;
    }
    
    eventSrc.onerror = function (event) {
      console.log(event);
    }
  })
  
  async function getTime() {
    const res = await fetch("http://localhost:8080/time");
    if (res.status !== 200) {
      console.log("could not connect to server")
    }
    return res;
  } 
  
</script>

<main>
  
  <h1>Server-Sent Events</h1>
  <p>Time : { time }</p>
  <button on:click="{getTime}">Get Time</button>
</main>

<style>
  
</style>
