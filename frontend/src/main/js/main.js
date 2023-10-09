window.addEventListener("DOMContentLoaded", (event) => {  // ожидание полной загрузки страницы и добавление листенера на кнопку
  document.getElementById('send_form').addEventListener('click', CreateEvent); // вызов функции проверки пароля по нажатию
});
  
  // функция отправки запроса на создание ивента
  function CreateEvent(){
    //получаем данные из локального хранилища
    const token = localStorage.getItem('token_CSRF')
    const auth_token = localStorage.getItem('auth_token');

    //создаем новый запрос на создание события
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31/api/event/",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.setRequestHeader("Authorization", "Bearer " + auth_token);
    xhr.withCredentials = true;
  
    // получаем переменные из html и засовываем в json
    var time_start = document.getElementById("start_time").value + ":00Z";
    var time_end = document.getElementById("end_time").value + ":00Z";
    var place = document.getElementById("address").value;
    var description = document.getElementById("description").value;
    var public = document.getElementById('checkbox').checked;
    var game = document.querySelector('input[name="list-radio"]:checked').value;

    const body = JSON.stringify({
        "time_start": time_start,
        "time_end": time_end,
        "place": place,
        "game": game,
        "description": description,
        "public": public
    });

    // отправляем запрос
    xhr.send(body);
  }