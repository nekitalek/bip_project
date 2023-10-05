window.addEventListener("DOMContentLoaded", (event) => {
  const el = document.getElementById('send_form');
  if (el) {
    el.addEventListener('click', Login);
  }
});

function openModal() {
  document.getElementById("2fa_modal").classList.remove("hidden");

}

function closeModal() {
  document.getElementById("2fa_modal").classList.add("hidden");
  }

function verify2fa(){
const requestURL = 'https://51.250.24.31:65000/CSRF'

var xhr = new XMLHttpRequest()
xhr.open('GET', requestURL, false);
xhr.withCredentials = true;
try {
xhr.send();
if (xhr.status != 200) {
  console.log('Ошибка');
} else {
  var jsonResponse = JSON.parse(xhr.responseText);
  console.log(jsonResponse["token_CSRF"]);
  localStorage.setItem('token_CSRF',jsonResponse["token_CSRF"]);
  //если нужно использовать в другом скрипке, то сохранить в память браузера
  //localStorage.setItem(jsonResponse["token_CSRF"]);
  //Для получение .getItem
}
} catch(err) {
alert("Запрос не удался");
}
      var fa_code = document.getElementById("factor").value;
      if(fa_code){
          fa2POST(fa_code)
          var jsonResponse = JSON.parse(xhr.responseText);
          localStorage.setItem('auth_token',jsonResponse["auth_token"]);
            if (jsonResponse["auth_token"] != null){
              window.location.href = "https://51.250.24.31/main/main.html"; 
            }
            else{
              alert("Ошибка при регистрации")
            }
      }
      else alert("please enter the number")


}

const btnSend2FA = document.querySelector('btn-send-2fa')

function fa2POST(fa_code){
  const xhr = new XMLHttpRequest();
  const token = localStorage.getItem('token_CSRF')
  const user_id = localStorage.getItem('user_id')
  xhr.open("POST", "https://51.250.24.31:65000/auth/sign-in/sec_factor",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);

  xhr.withCredentials = true;

  const body = JSON.stringify({
      "user_id": parseInt(user_id), // я понятия не имею как брать юзер айди поэтому он захардкожен
      "code": parseInt(fa_code),
      "device": "windows"
    });
  xhr.onload = () => {
      if (xhr.readyState == 4 && xhr.status == 201) {
          console.log(JSON.parse(xhr.responseText));
      } else {
          console.log(`Error: ${xhr.status}`);
      }
  };
  xhr.send(body);
}

function Login(){

  const token = localStorage.getItem('token_CSRF')
  const xhr = new XMLHttpRequest();
  xhr.open("POST", "https://51.250.24.31:65000/auth/sign-in/password",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);

  xhr.withCredentials = true;

  var login = document.getElementById("email").value;
  var password = document.getElementById("password").value;

  const body = JSON.stringify({
    "Login": login,
    "Password": password
  });
  xhr.onload = () => {
    if (xhr.readyState == 4 && xhr.status == 201) {
      console.log(JSON.parse(xhr.responseText));
    } else {
      console.log(`Error: ${xhr.status}`);
    }
  };
  xhr.send(body);
  var jsonResponse = JSON.parse(xhr.responseText);
  const user_id = localStorage.setItem('user_id',jsonResponse["user_id"]);
  openModal()
}