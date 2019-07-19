const fetchPODInfo =  async () => {
    try {
        let res = await fetch('/pod');
        let pod = await res.json()
        console.log(pod)
    } catch (err) {
        console.error(err)
    }
};

const updatePODInfo = () => {
    fetchPODInfo()
};

const loopFetch = () => {
    setInterval(()=>{
        updatePODInfo()
    }, 10 * 1000)
};

loopFetch();