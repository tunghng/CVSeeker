
const wsUrl = 'ws://localhost:8080/ws';
let ws;

const connectSocket = (onMessage) => {
    ws = new WebSocket(wsUrl);

    ws.onopen = function() {
        console.log('Connected to WebSocket server at ' + wsUrl);
    };

    ws.onmessage = function(event) {
        onMessage(event.data);
    };

    ws.onclose = function() {
        console.log('Disconnected from WebSocket server');
    };

    ws.onerror = function(error) {
        console.error('WebSocket error: ' + error.message);
    };
};

const disconnect = () => {
    if (ws) {
        ws.close();
    }
};

export { connectSocket, disconnect };
