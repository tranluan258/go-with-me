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
    .then(() => {
      window.location = "/";
    })
    .catch((err) => {
      alert(err);
    });
}
