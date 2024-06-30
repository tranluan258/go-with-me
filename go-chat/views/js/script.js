/**
 * @type {WebSocket}
 */
var conn;

/**
 * @param {HTMLElement} item
 * @description Append messageElement to chatContainer
 */
function appendMessageToContainer(item) {
  var log = document.getElementById("container");
  var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
  log.appendChild(item);
  if (doScroll) {
    log.scrollTop = log.scrollHeight - log.clientHeight;
  }
}

/**
 * @param {string} msg
 * @param {boolean} isRec
 * @param {string} [username]
 * @description Create messageElement for sender and receiver
 */
function createMessageElement(msg, isRec, username) {
  let messageElement = document.createElement("div");
  messageElement.classList.add("message");

  isRec
    ? messageElement.classList.add("received")
    : messageElement.classList.add("sent");

  const metadataElement = document.createElement("div");
  metadataElement.classList.add("metadata");
  metadataElement.innerHTML = `<span class="time">${isRec ? username : "Me"}</span>`;

  const textElement = document.createElement("span");
  textElement.textContent = msg;

  messageElement.appendChild(metadataElement);
  messageElement.appendChild(textElement);
  appendMessageToContainer(messageElement);
}

/**
 * @description Handle send message websocket
 */
function handleSendMessage() {
  var msg = document.getElementById("message-input");
  if (!conn) {
    return;
  }
  if (!msg.value) {
    return;
  }
  createMessageElement(msg.value, false);
  conn.send(msg.value);

  msg.value = "";
}
if (window["WebSocket"]) {
  const url = "ws://localhost:8080/ws";
  conn = new WebSocket(url);

  /**
   * @param {MessageEvent<{data: string}>} evt
   */
  conn.onmessage = function (evt) {
    /**
     * @type {{username: string; msg: string}}
     */
    const data = JSON.parse(evt.data);
    createMessageElement(data.msg, true, data.username);
  };
}

document
  .getElementById("message-input")
  .addEventListener("keypress", function (e) {
    if (e.key === "Enter") {
      document.getElementById("send-button").click();
    }
  });
