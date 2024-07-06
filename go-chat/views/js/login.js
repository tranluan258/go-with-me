/**
 * @description handle login
 */
function handleLogin() {
  const username = document.getElementById("username");
  const password = document.getElementById("password");

  if (!username.value || !password.value) {
    return;
  }

  fetch("login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username: username.value,
      password: password.value,
    }),
  })
    .then((data) => {
      if (data.status >= 400) {
        alert("Username or password invalid");
      } else {
        window.location = "/";
      }
    })
    .catch(() => {
      alert("Something wroing here!");
    });
}
