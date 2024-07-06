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
if (window["WebSocket"]) {
  const url = "ws://localhost:8080/ws/1";
  conn = new WebSocket(url);

  /**
   * @param {MessageEvent<{data: string}>} evt
   */
  conn.onmessage = function (evt) {
    /**
     * @type {{full_name: string; msg: string,type?:string}}
     */
    const data = JSON.parse(evt.data);

    switch (data?.type) {
      case "joined":
        addUserToList(data.full_name);
        return;
      case "left":
        removeUserFromList(data.full_name);
        return;
      case "user-list":
        addUserToList(data.full_name, false);
        return;
      default:
        createMessageElement(data.msg, true, data.full_name);
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

/**
 * @param {string} message
 * @param {'left'| 'joined'} event
 */
function showToast(message, event) {
  const toastContainer = document.getElementById("toastContainer");
  const toast = document.createElement("div");
  toast.className = event === "left" ? "toast-left" : "toast-joined";
  toast.textContent = message;

  toastContainer.appendChild(toast);

  setTimeout(() => {
    toast.remove();
  }, 3000); // Toast duration
}

/**
 * @param {string} username
 * @param {boolean} [isShow]
 */
function addUserToList(username, isShow) {
  const userListContainer = document.getElementById("userListContainer");
  const userItem = document.createElement("li");
  userItem.id = `user-${username}`;
  userItem.textContent = username;
  userListContainer.appendChild(userItem);

  if (isShow) {
    showToast(`${username} has joined the chat!`, "joined");
  }
}

/**
 * @param {string} username
 */
function removeUserFromList(username) {
  const userItem = document.getElementById(`user-${username}`);
  if (userItem) {
    userItem.remove();
    showToast(`${username} has left the chat!`, "left");
  }
}

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

function selectEmoji(data) {
  console.log(data);
  const messageInput = document.getElementById("message-input");
  messageInput.value += data.native;
}

function toogleDropdown() {
  const dropdownMenu = document.getElementById("dropdownMenu");
  const logoutOption = document.getElementById("logoutOption");

  dropdownMenu.style.display =
    dropdownMenu.style.display === "block" ? "none" : "block";
}
