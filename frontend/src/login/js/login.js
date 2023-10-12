window.addEventListener("DOMContentLoaded", (event) => {  // ожидание полной загрузки страницы и добавление листенера на кнопку
  document.getElementById('send_form').addEventListener('click', Login); // вызов функции входа по нажатию
});

//открыть модальное окно
const openModal = () => { 
  document.getElementById("2fa_modal").classList.remove("hidden");
};

// закрыть модальнео окно
const closeModal = () => { 
  document.getElementById("2fa_modal").classList.add("hidden");
};

// функция проверки код второго фактора  
function fa2POST(){
  var fa_code = document.getElementById("factor").value; // получение цисла из HTML
  const token = localStorage.getItem('token_CSRF') // получение CSRF токена
  const user_id = localStorage.getItem('user_id'); // получение user_id
  const device= window.navigator.userAgent;
  // формируем запрос на 2фа
  const xhr = new XMLHttpRequest();
  xhr.open("POST", "https://51.250.24.31/auth/sign-in/sec_factor",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.withCredentials = true;

  //получаем переменные из html и засовываем в json
  const body = JSON.stringify({
      "user_id": parseInt(user_id), 
      "code": parseInt(fa_code),
      "device": device
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

// функция получения CSRF токена и записи его в локальное хранилище
function GetCSRF(){
  var xhr = new XMLHttpRequest()
  xhr.open('GET', "https://51.250.24.31/CSRF", false);
  xhr.withCredentials = true;
  try {
  xhr.send();
  if (xhr.status != 200) { // отлов ошибки полученной при успешной отправке запроса
    console.log('Error');
  } else {
    var jsonResponse = JSON.parse(xhr.responseText); // парсим токен полученный в ответ от сервера
    localStorage.setItem('token_CSRF',jsonResponse["token_CSRF"]); // кладем токен в локальное хранилище
  }
  } catch(err) { // отлов ошибки при отправке запроса
  alert("Запрос CSRF не удался.");
}
}

// функция отправил запроса на вход
function Login(){
  
  GetCSRF() // вызываем функцию получения токена чтобы использовать его далее

  const token = localStorage.getItem('token_CSRF') // получения токена из локального хранилища
  //const user_id = localStorage.getItem('user_id'); // получение user_id из локального хранилища

  //создаем новый запрос на регистрацию
  var xhr = new XMLHttpRequest()
  xhr.open("POST", "https://51.250.24.31/auth/sign-in/password",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.withCredentials = true;

  //получаем переменные из html и засовываем в json
  var login = document.getElementById("email").value;
  var password = document.getElementById("password").value;

  const body = JSON.stringify({
    "Login": login,
    "Password": password
  });

  // отправляем запрос
  xhr.send(body);

  var jsonResponse = JSON.parse(xhr.responseText);
  if (jsonResponse["status"] = 'ok'){
  localStorage.setItem('user_id',jsonResponse["user_id"]);
  // открываем модальное окно ввода кода подтверждения
  openModal()
  }
  else {
    alert(jsonResponse);
  }
}