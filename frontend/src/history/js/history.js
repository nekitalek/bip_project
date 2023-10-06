window.addEventListener("DOMContentLoaded", (event) => {  // ожидание полной загрузки страницы и добавление листенера на кнопку
  document.getElementById('refresh').addEventListener('click', GetEvents); // вызов функции входа по нажатию
});

// функция отправки запроса на получение списка событий
function GetEvents(){

  // получения данных из локального хранилища
  const token = localStorage.getItem('token_CSRF') // получения токена из локального хранилища
  const auth_token = localStorage.getItem('auth_token'); // получение auth_token из локального хранилища
  
  //создаем новый запрос на регистрацию
  const xhr = new XMLHttpRequest();
  xhr.open("GET", "https://51.250.24.31:65000/api/event/",false);
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.setRequestHeader("Authorization", "Bearer " + auth_token);
  xhr.withCredentials = true;

  xhr.send(); // отправляем запрос
  var event_list = JSON.parse(xhr.responseText); // парсим список событий полученный в ответ от сервера
  
  ParseEvents(event_list) // вызываем парсинг
}

//функция обработки и вывода списка событий
function ParseEvents(event_list){

  var out = '';
  //парсим информацию полученную в json
  for (var key in event_list){
    var participants = event_list[key].participant;
    var all_part = ''
    for (var key2 in participants){
      all_part += participants[key2].username
      all_part += ', '
    }
    all_part.substring(0,all_part.length-3)
    out += '<a class="flex flex-col items-center bg-white border border-gray-200 rounded-lg shadow md:flex-row md:max-w-xl hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700">  <div class="flex flex-col justify-between p-4 leading-normal"><h4 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">Описание мероприятия ' + event_list[key].description +'</h4><p id="time_start" class="mb-3 font-normal text-gray-700 dark:text-gray-400">Время начала: '+ event_list[key].time_start +'</p><p id="time_end" class="mb-3 font-normal text-gray-700 dark:text-gray-400">Время окончания: '+ event_list[key].time_end +'</p><p id="place" class="mb-3 font-normal text-gray-700 dark:text-gray-400">Адрес: '+ event_list[key].place +'</p><p id="game" class="mb-3 font-normal text-gray-700 dark:text-gray-400">Игра: '+ event_list[key].game +'</p><p id="participant" class="mb-3 font-normal text-gray-700 dark:text-gray-400">Участники: '+ all_part +'</p><button id="'+ event_list[key].event_items_id +'_accept_event" type="button" class="focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">Записаться</button><button id="'+ event_list[key].event_items_id +'_decline_event" type="button" class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900">Покинуть мероприятие</button></div></a>';
  }
  document.getElementById('event_grid').innerHTML = out;

  var buttons = document.querySelectorAll('button'); // выбираем все кнопки с сайта
  //добавляем ивент листенер для каждой кнопки
  for (var i=0; i<buttons.length; ++i) {
    if (parseInt(buttons[i]) != NaN){
      buttons[i].addEventListener('click', clickFunc);
    }
    else{
      continue
    }
  }
}

//функция выбора присоединиться/покинуть
function clickFunc() {
  var str = this.id
  var str1 = str.substring(2,4)
  console.log(str1)
  if(str1 == 'ac'){
    JoinEvent(str); 
  }
  else if(str1 == 'de'){
    LeftEvent(str)
  }
}

//функция присоединения к событию
function JoinEvent(button_id){

    //получение айди события
    var event_id = parseInt(button_id)

    // получения данных из локального хранилища
    const token = localStorage.getItem('token_CSRF');
    const auth_token = localStorage.getItem('auth_token');
    const user_id = parseInt(localStorage.getItem('user_id'));

    //создаем новый запрос на присоединение к событию
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31:65000/api/invitation/",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.setRequestHeader("Authorization", "Bearer " + auth_token);
    xhr.withCredentials = true;

    const body = JSON.stringify({
        "event_id": event_id,
        "user_id": user_id,
        "status": "Confirmed"
    });

    // отправляем запрос
    xhr.send(body);
}

//функция подания события
function LeftEvent(button_id){

  //получение айди события
  var event_id = parseInt(button_id)

  // получения данных из локального хранилища
  const token = localStorage.getItem('token_CSRF');
  const auth_token = localStorage.getItem('auth_token');
  const user_id = parseInt(localStorage.getItem('user_id'));

  //создаем новый запрос на покидание события
  const xhr = new XMLHttpRequest();
  xhr.open("DELETE", "https://51.250.24.31:65000/api/invitation/" + event_id,false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.setRequestHeader("Authorization", "Bearer " + auth_token);
  xhr.withCredentials = true;

  const body = JSON.stringify({
      "event_id": event_id,
      "user_id": user_id,
  });

  // отправляем запрос
  xhr.send(body);
}