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

function prependNewRoom() {
  // @ts-ignore
  document.querySelector(".modal-backdrop")?.click();
  const listRoom = document.querySelector(".list-room");
  const room = document.querySelector(".room_id_header");
  const roomId = room?.id;
  const roomName = room?.innerHTML;
  let isExisted = false;

  listRoom?.childNodes.forEach((node) => {
    if (node instanceof HTMLElement && node.id === roomId) {
      isExisted = true;
      return;
    }
  });

  if (isExisted) return;

  const newRoom = `
<li
  id='${roomId}'
  class='flex items-center w-full h-[70px] mb-2 p-4 rounded-xl cursor-pointer hover:bg-gray-50'
  hx-get='messages?room_id={{.ID}}'
  hx-target='#chat-container'
  hx-swap='innerHTML'
>
  <div class='avatar placeholder'>
    <div class='bg-neutral text-neutral-content w-10 rounded-full'>
      <span class='text-xs'>UNE</span>
    </div>
  </div>
  <span class='ml-3 truncate'>${roomName}</span>
</li>
`;
  listRoom?.insertAdjacentHTML("beforebegin", newRoom);
}
