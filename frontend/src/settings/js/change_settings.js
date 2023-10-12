  window.addEventListener("DOMContentLoaded", (event) => {
    // добавить if на проверку каких либо изменений данных пользователем, иначе не делать ничего по нажатию кнопки
    const email = document.getElementById("send_email");
    if (email) {
      email.addEventListener('click', openModalpassword);
    }
  });

  window.addEventListener("DOMContentLoaded", (event) => {
    // добавить if на проверку каких либо изменений данных пользователем, иначе не делать ничего по нажатию кнопки
    const password = document.getElementById("send_password");
    if (password) {
      password.addEventListener('click', openModalpassword2);
    }
  });
  
  function openModal() {
    document.getElementById("2fa_modal").classList.remove("hidden");
  
  }
  
  function closeModal() {
    document.getElementById("2fa_modal").classList.add("hidden");
    }
  function openModal2() {
    document.getElementById("2fa_modal2").classList.remove("hidden");
  
  }
  
  function closeModal2() {
    document.getElementById("2fa_modal2").classList.add("hidden");
    }
  function openModalpassword() {
    document.getElementById("password_modal").classList.remove("hidden");
  
  }
  
  function closeModalpassword() {
    document.getElementById("password_modal").classList.add("hidden");
    }


    function openModalpassword2() {
      document.getElementById("password_modal_2").classList.remove("hidden");
    
    }
    
    function closeModalpassword2() {
      document.getElementById("password_modal_2").classList.add("hidden");
      }

function FirstFactorEmail(){
  const token = localStorage.getItem('token_CSRF') // получение CSRF токена
  const user_id = localStorage.getItem('user_id'); // получение user_id
  password = document.getElementById("password_check").value
  // формируем запрос на 2фа
  const xhr = new XMLHttpRequest();
  xhr.open("POST", "https://51.250.24.31/auth/change/email/first_factor",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.withCredentials = true;
  login = document.getElementById("win_email").innerHTML
  
  //получаем переменные из html и засовываем в json
  const body = JSON.stringify({
    "Login": login,
    "Password": password
  });
  xhr.onload = () => {
    if (xhr.readyState == 4 && xhr.status == 200) {
      console.log(JSON.parse(xhr.responseText));
    } else {
      console.log(`Error: ${xhr.status}`);
    }
  };
  xhr.send(body); // отправляем запрос
  closeModalpassword()
  openModal()
}

function SecFactorEmail(){
  const token = localStorage.getItem('token_CSRF') // получение CSRF токена
  const user_id = localStorage.getItem('user_id'); // получение user_id

  // формируем запрос на 2фа
  const xhr = new XMLHttpRequest();
  xhr.open("POST", "https://51.250.24.31/auth/change/email/sec_factor",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.withCredentials = true;
  code = document.getElementById("factor").value

  //получаем переменные из html и засовываем в json
  const body = JSON.stringify({
    "e_conf":{
        "user_id": parseInt(user_id),
        "code": parseInt(code),
        "device": "windows"
    },
    "new_login": document.getElementById("email_setting").value
  });
  xhr.onload = () => {
    if (xhr.readyState == 4 && xhr.status == 200) {
      console.log(JSON.parse(xhr.responseText));
    } else {
      console.log(`Error: ${xhr.status}`);
    }
  };
  xhr.send(body); // отправляем запрос
  closeModal()
  openModal()
  SecFactorNewEmail()
}

function SecFactorNewEmail(){
  const token = localStorage.getItem('token_CSRF') // получение CSRF токена
  const user_id = localStorage.getItem('user_id'); // получение user_id

  // формируем запрос на 2фа
  const xhr = new XMLHttpRequest();
  xhr.open("POST", "https://51.250.24.31/auth/change/email/verification_new_email",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.withCredentials = true;
  code = document.getElementById("factor").value
  closeModal()
  //получаем переменные из html и засовываем в json
  const body = JSON.stringify({
    "user_id": parseInt(user_id),
    "code": parseInt(code),
    "device": "windows"
});
  xhr.onload = () => {
  if (xhr.readyState == 4 && xhr.status == 200) {
    console.log(JSON.parse(xhr.responseText));
  } else {
    console.log(`Error: ${xhr.status}`);
  }
};
  xhr.send(body); // отправляем запрос
  if(xhr.status == 200){
    alert("Вы успешно сменили почту)"); 
  }
}


function FirstFactorPassword(){
  const token = localStorage.getItem('token_CSRF') // получение CSRF токена
  const user_id = localStorage.getItem('user_id'); // получение user_id
  password = document.getElementById("password_check2").value
  // формируем запрос на 2фа
  const xhr = new XMLHttpRequest();
  xhr.open("POST", "https://51.250.24.31/auth/change/password/first_factor",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.withCredentials = true;
  login = document.getElementById("email_setting").value
  
  //получаем переменные из html и засовываем в json
  const body = JSON.stringify({
    "Login": login,
    "Password": password
  });
  xhr.onload = () => {
    if (xhr.readyState == 4 && xhr.status == 200) {
      console.log(JSON.parse(xhr.responseText));
    } else {
      console.log(`Error: ${xhr.status}`);
    }
  };
  xhr.send(body); // отправляем запрос
  closeModal()
  openModal2()
}

function SecFactorPassword(){
  const token = localStorage.getItem('token_CSRF') // получение CSRF токена
  const user_id = localStorage.getItem('user_id'); // получение user_id

  // формируем запрос на 2фа
  const xhr = new XMLHttpRequest();
  xhr.open("POST", "https://51.250.24.31/auth/change/password/sec_factor",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.withCredentials = true;
  code = document.getElementById("factor").value
  closeModal2()
  //получаем переменные из html и засовываем в json
  const body = JSON.stringify({
    "e_conf":{
        "user_id": parseInt(user_id),
        "code": parseInt(code),
        "device": "windows"
    },
    "new_password": document.getElementById("new_password").value
  });
  xhr.onload = () => {
    if (xhr.readyState == 4 && xhr.status == 200) {
      console.log(JSON.parse(xhr.responseText));
    } else {
      console.log(`Error: ${xhr.status}`);
    }
  };
  xhr.send(body); // отправляем запрос
  if(xhr.status == 200){
    alert("Вы успешно сменили пароль)"); 
  }
}



