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

function closeModal() {
  // @ts-ignore
  document.querySelector(".modal-backdrop")?.click();
}
