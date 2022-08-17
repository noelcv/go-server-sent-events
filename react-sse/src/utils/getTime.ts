export const getTime = async () => {
  try {
    const res = await fetch("http://localhost:8080/time");
    if (res.status === 204)
      return res;
  } catch(err) {
    console.log("Error at getTime Service: ", err)
  }
}