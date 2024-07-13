/**
 * @type {WebSocket}
 */
var conn;

/**
 * @type {Array<string>}
 */
var users = [];

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
  conn.send(
    JSON.stringify({
      msg: msg.value,
    }),
  );

  msg.value = "";
}
// TODO: handle connect when select chat room
// if (window["WebSocket"]) {
//   const url = "ws://localhost:8080/ws/1";
//   conn = new WebSocket(url);
//
//   /**
//    * @param {MessageEvent<{data: string}>} evt
//    */
//   conn.onmessage = function (evt) {
//     /**
//      * @type {{full_name: string; msg: string,type?:string}}
//      */
//     const data = JSON.parse(evt.data);
//     createMessageElement(data.msg, true, data.full_name);
//   };
// }
//
document
  .getElementById("message-input")
  .addEventListener("keypress", function (e) {
    if (e.key === "Enter") {
      document.getElementById("send-button").click();
    }
  });

function toggleEmojiPicker() {
  const existPicket = document.querySelector(".emoji-picker");
  if (existPicket) {
    existPicket.remove();
  } else {
    const pickerOptions = { onEmojiSelect: selectEmoji };
    const picker = new EmojiMart.Picker(pickerOptions);
    picker.className = "emoji-picker";

    const chatApp = document.querySelector(".chat-input");
    chatApp.appendChild(picker);
  }
}

/**
 * @param {{native: string}} data
 */
function selectEmoji(data) {
  const messageInput = document.getElementById("message-input");
  messageInput.value += data.native;
}

function toogleDropdown() {
  const dropdownMenu = document.getElementById("dropdownMenu");
  dropdownMenu.style.display =
    dropdownMenu.style.display === "block" ? "none" : "block";
}
