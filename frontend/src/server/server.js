// server.js
const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 8083 });

wss.on('connection', (ws) => {
  ws.on('message', (data) => {

    wss.clients.forEach((client) => {
      if (client !== ws && client.readyState === WebSocket.OPEN) {
        client.send(data);
      }
    });
  });

  ws.on('close', () => {
    console.log('Client disconnected');
  });

  console.log('Client connected');
});

console.log('WebSocket server is running on ws://localhost:8083');
