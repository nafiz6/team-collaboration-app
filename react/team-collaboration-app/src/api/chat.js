// api/index.js
var socket;

let connect = (WorkspaceId) => {
  console.log("Attempting Connection...");

  socket = new WebSocket("ws://localhost:8080/api/chat");

  socket.onopen = () => {
    console.log("Successfully Connected");
    sendMsg({
      "Type": "Connection",
      "Body": "Connected to workspace",
      "WorkspaceId" : WorkspaceId
    })
  };
  
  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };


};

let sendMsg = msg => {
  console.log("sending msg: ", msg);
  socket.send(JSON.stringify(msg));
};

export { connect, sendMsg, socket };