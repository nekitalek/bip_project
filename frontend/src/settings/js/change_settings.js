  window.addEventListener("DOMContentLoaded", (event) => {
    // добавить if на проверку каких либо изменений данных пользователем, иначе не делать ничего по нажатию кнопки
    const email = document.getElementById('send_email');
    if (email) {
      email.addEventListener('click', ChangeMail);
    }
  });

  window.addEventListener("DOMContentLoaded", (event) => {
    // добавить if на проверку каких либо изменений данных пользователем, иначе не делать ничего по нажатию кнопки
    const password = document.getElementById('send_password');
    if (password) {
      password.addEventListener('click', ChangePassword);
    }
  });
  
  function openModal() {
    document.getElementById("2fa_modal").classList.remove("hidden");
  
  }
  
  function closeModal() {
    document.getElementById("2fa_modal").classList.add("hidden");
    }
  
  // функция проверки код второго фактора  
function SecFactorEmail(){
  var fa_code = document.getElementById("factor").value; // получение цисла из HTML
  const token = localStorage.getItem('token_CSRF') // получение CSRF токена
  const user_id = localStorage.getItem('user_id'); // получение user_id

  // формируем запрос на 2фа
  const xhr = new XMLHttpRequest();
  xhr.open("POST", "https://51.250.24.31/auth/change/email/sec_factor",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.withCredentials = true;

  //получаем переменные из html и засовываем в json
  const body = JSON.stringify({
      "user_id": parseInt(user_id), 
      "code": parseInt(fa_code),
      "device": "windows"
    });
  
  xhr.send(body); // отправляем запрос

  // парсим ответ, если ок то перенаправляем на логин
  if (xhr.status == 200){
    var jsonResponse = JSON.parse(xhr.responseText); // парсим токен полученный в ответ от сервера
    localStorage.setItem('auth_token',jsonResponse["auth_token"]); // кладем токен в локальное хранилище
    window.location.href = "https://51.250.24.31/main/"; 
  }
  else{
    alert(jsonResponse)
  }
}

  function ChangeSettings(){
  
    const token = localStorage.getItem('token_CSRF')
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31/auth/sign-up/password",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
  
    xhr.withCredentials = true;
  
    var login = document.getElementById("email").value;
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
  
    const body = JSON.stringify({
      "Login": login,
      "Username": username,
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
  }

function ChangeFistFactor(){
    const token = localStorage.getItem('token_CSRF')
    const xhr = new XMLHttpRequest();

    if(change_email = 1) //хз как определять че пользователь меняет и как обрабатывать если он меняет и пароль и почту но пока так
    {
        xhr.open("POST", "https://51.250.24.31/auth/change/email/first_factor",false);
    }
    else {
        xhr.open("POST", "https://51.250.24.31/auth/change/password/first_factor",false);
    }
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.withCredentials = true;

    var login = document.getElementById("email").value;
    var old_password = document.getElementById("old_password").value;
  
    const body = JSON.stringify({
      "Login": login,
      "Password": old_password
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

function ChangePasswordSecondFactor(){
    const token = localStorage.getItem('token_CSRF')
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31/auth/change/password/sec_factor",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.withCredentials = true;

    var new_password = document.getElementById("new_password").value;
    
    openModal() //надо как то обработать что именно меняется и изменить под это модальное окно
    var fa_code = document.getElementById("factor").value;
    const body = JSON.stringify({
        "e_conf":{
            "user_id": 5, // где брать значения юзерайди я не в курсе
            "code": parseInt(fa_code),
            "device": "windows"
        },
        "new_password": new_password
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

function ChangeEmailSecondFactor(){
    const token = localStorage.getItem('token_CSRF')
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31/auth/change/password/sec_factor",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.withCredentials = true;

    var email = document.getElementById("email").value;
    
    openModal() //надо как то обработать что именно меняется и изменить под это модальное окно
    var fa_code = document.getElementById("factor").value;
    const body = JSON.stringify({
        "e_conf":{
            "user_id": 5, // где брать значения юзерайди я не в курсе
            "code": parseInt(fa_code),
            "device": "windows"
        },
        "new_login": email
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

function ConfirmNewEmail(){
    const token = localStorage.getItem('token_CSRF')
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31/auth/change/email/verification_new_email",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.withCredentials = true;

    openModal() //надо как то обработать что именно меняется и изменить под это модальное окно
    var fa_code = document.getElementById("factor").value;
    const body = JSON.stringify({
        "user_id": 5,
        "code": 668561,
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