let wasConnected = false;
let reconnectInterval;

export const enableHotReload = () => {
    const ws = new WebSocket("ws://localhost:1337/ws");
    ws.onclose = () => {
        reconnectInterval = setInterval(() => {
            enableHotReload()
        }, 100)
    }
    ws.onopen = () => {
        if (reconnectInterval)
            clearInterval(reconnectInterval)
        if (wasConnected) {
            console.log("Hot reloading")
            window.location.reload();
        } else {
            console.log("Hot reload connected")
            wasConnected = true;
        }
    }
}