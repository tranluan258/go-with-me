var conn;

function appendLog(item) {
  var log = document.getElementById("log");
  var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
  log.appendChild(item);
  if (doScroll) {
    log.scrollTop = log.scrollHeight - log.clientHeight;
  }
}

function handleSendMessage() {
  var msg = document.getElementById("message-input");
  if (!conn) {
    return false;
  }
  if (!msg.value) {
    return false;
  }
  conn.send(msg.value);
  msg.value = "";
  return false;
}
if (window["WebSocket"]) {
  const url = "ws://localhost:8080/ws";
  conn = new WebSocket(url);
  conn.onmessage = function (evt) {
    var messages = evt.data.split("\n");
    for (let i = 0; i < messages.length; i++) {
      let item = document.createElement("div");
      item.classList.add("message");
      item.innerText = messages[i];
      appendLog(item);
    }
  };
}

document
  .getElementById("message-input")
  .addEventListener("keypress", function (e) {
    if (e.key === "Enter") {
      document.getElementById("send-button").click();
    }
  });
