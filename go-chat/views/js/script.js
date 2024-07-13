/**
 * @type {WebSocket}
 */
let conn;

/**
 * @param {HTMLElement} item
 * @description Append messageElement to chatContainer
 */
function appendMessageToContainer(item) {
  /**
   * @type {HTMLElement | null}
   */
  const chatContainer = document.getElementById("chat-messages");
  if (!chatContainer) return;

  const doScroll =
    chatContainer.scrollTop >
    chatContainer.scrollHeight - chatContainer.clientHeight - 1;

  chatContainer.appendChild(item);
  if (doScroll) {
    chatContainer.scrollTop =
      chatContainer.scrollHeight - chatContainer.clientHeight;
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
  /**
   * @type {HTMLInputElement | null}
   * */
  // @ts-ignore
  const msg = document.getElementById("message-input");
  if (!conn) return;
  if (!msg) return;

  if (!msg.value) return;

  createMessageElement(msg.value, false);
  conn.send(
    JSON.stringify({
      msg: msg.value,
    }),
  );

  msg.value = "";
}

/**
 * @param {Event} e
 */
function connectWs(e) {
  if (window["WebSocket"]) {
    // @ts-ignore
    const roomId = e.target?.id;
    const url = "ws://localhost:8080/ws/" + roomId;
    conn = new WebSocket(url);

    /**
     * @param {MessageEvent<{data: string}>} evt
     */
    conn.onmessage = function (evt) {
      /**
       * @type {{full_name: string; msg: string,type?:string}}
       */
      // @ts-ignore
      const data = JSON.parse(evt.data);
      createMessageElement(data.msg, true, data.full_name);
    };
  }
}

document
  .getElementById("message-input")
  ?.addEventListener("keypress", function (e) {
    if (e.key === "Enter") {
      document.getElementById("send-button")?.click();
    }
  });

function toggleEmojiPicker() {
  const existPicket = document.querySelector(".emoji-picker");
  if (existPicket) {
    existPicket.remove();
  } else {
    const pickerOptions = { onEmojiSelect: selectEmoji };
    // @ts-ignore
    const picker = new EmojiMart.Picker(pickerOptions);
    picker.className = "emoji-picker";

    const chatApp = document.querySelector(".chat-input");
    chatApp?.appendChild(picker);
  }
}

/**
 * @param {{native: string}} data
 */
function selectEmoji(data) {
  const messageInput = document.getElementById("message-input");
  if (!messageInput) return;
  // @ts-ignore
  messageInput.value += data.native;
}

function toogleDropdown() {
  const dropdownMenu = document.getElementById("dropdownMenu");
  if (!dropdownMenu) return;
  dropdownMenu.style.display =
    dropdownMenu.style.display === "block" ? "none" : "block";
}
