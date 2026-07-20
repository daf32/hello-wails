import "./style.css";
import "./app.css";

import logo from "./assets/images/logo-universal.png";
import { GetUser, Divide, Greet, ListUsers } from "../wailsjs/go/main/App";

document.querySelector("#app").innerHTML = `
    <img id="logo" class="logo">
      <div class="result" id="result">Please enter your name below 👇</div>
      <div class="input-box" id="input">
        <input class="input" id="name" type="text" autocomplete="off" />
        <button class="btn" onclick="greet()">Greet</button>
      </div>
    </div>
`;
document.getElementById("logo").src = logo;

let nameElement = document.getElementById("name");
nameElement.focus();
let resultElement = document.getElementById("result");

// Setup the greet function
window.greet = function () {
  // Get name
  let name = nameElement.value;

  // Check if the input is empty
  if (name === "") return;

  // Call App.Greet(name)
  try {
    Greet(name)
      .then((result) => {
        // Update result with data back from App.Greet()
        resultElement.innerText = result;
      })
      .catch((err) => {
        console.error(err);
      });
  } catch (err) {
    console.error(err);
  }
};

async function testDivide() {
  try {
    const ok = await Divide(10, 2);
    console.log("10 / 2 =", ok);
  } catch (err) {
    console.error("Не должно сюда попасть:", err);
  }

  try {
    const bad = await Divide(10, 0);
    console.log("Не должно сюда попасть:", bad);
  } catch (err) {
    console.error("Поймал ошибку из Go:", err);
  }
}

testDivide();

try {
  const user = await GetUser(1);
  console.log(user.name, user.email);
} catch (err) {
  console.error(err);
}
try {
  const users = await ListUsers(10);
  console.log(users);
} catch (err) {
  console.error(err);
}
