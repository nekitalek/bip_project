window.addEventListener("DOMContentLoaded", (event) => {
  const el = document.getElementById('refresh');
  if (el) {
    el.addEventListener('click', GetEvents);
  }
});

function GetEvents(){

  const token = localStorage.getItem('token_CSRF')
  auth_token = localStorage.getItem('auth_token');
  const xhr = new XMLHttpRequest();
  xhr.open("GET", "https://51.250.24.31:65000/api/event/",false);
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.setRequestHeader("Authorization", "Bearer " + auth_token);
  xhr.withCredentials = true;

  try {
    xhr.send();
    if (xhr.status != 200) {
      console.log('Ошибка');
    } else {
      var event_list = JSON.parse(xhr.responseText);
      console.log(event_list);
    }
    } catch(err) {
    alert("Запрос не удался");
    }
    var out = '';
    for (var key in event_list){
      var participants = event_list[key].participant;
      var all_part = ''
      for (var key2 in participants){
        all_part += participants[key2].username
        console.log(all_part)
      }

      out += '<a href="#" class="flex flex-col items-center bg-white border border-gray-200 rounded-lg shadow md:flex-row md:max-w-xl hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700">  <div class="flex flex-col justify-between p-4 leading-normal"><h4 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">Описание мероприятия ' + event_list[key].description +'</h4><p id="time_start" class="mb-3 font-normal text-gray-700 dark:text-gray-400">Время начала: '+ event_list[key].time_start +'</p><p id="time_end" class="mb-3 font-normal text-gray-700 dark:text-gray-400">Время окончания: '+ event_list[key].time_end +'</p><p id="place" class="mb-3 font-normal text-gray-700 dark:text-gray-400">Адрес: '+ event_list[key].place +'</p><p id="game" class="mb-3 font-normal text-gray-700 dark:text-gray-400">Игра: '+ event_list[key].game +'</p><p id="participant" class="mb-3 font-normal text-gray-700 dark:text-gray-400">Участники: '+ all_part +'</p><button id="'+ event_list[key].event_items_id +'_accept_event" type="button" class="focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">Записаться</button><button id="'+ event_list[key].event_items_id +'_decline_event" type="button" class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900">Покинуть мероприятие</button></div></a>';
    }
    document.getElementById('event_grid').innerHTML = out;

    window.addEventListener("DOMContentLoaded", (event) => {
      var buttons = document.querySelectorAll('button');
    
      for (var i=0; i<buttons.length; ++i) {
        buttons[i].addEventListener('click', clickFunc);
      }
    
      const comp_string = 'add'
      function clickFunc() {
        if(comp_string < this.id){
          JoinEvent(this.id); 
        }
        else{
          LeftEvent(this.id)
        }
      }
    });
}





function JoinEvent(button_id){
    var event_id = parseInt(button_id)
    console.log(event_id)
    console.log(event_id.toString())
    const token = localStorage.getItem('token_CSRF');
    const auth_token = localStorage.getItem('auth_token');
    const user_id = localStorage.getItem('user_id');
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31:65000/api/invitation/"+ event_id.toString() + '/',false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.setRequestHeader("Authorization", "Bearer " + auth_token);
    xhr.withCredentials = true;

    const body = JSON.stringify({
        "event_id": event_id,
        "user_id": user_id,
        "status": "Confirmed"
    });
    xhr.onload = () => {
      if (xhr.readyState == 4 && xhr.status == 201) {
        console.log(JSON.parse(xhr.responseText));
      } else {
        console.log(`Error: ${xhr.status}`);
      }
    };
    xhr.send(body);
    console.log(body)
}

function LeftEvent(button_id){
  var event_id = parseInt(button_id)
  console.log(event_id)
  console.log(event_id.toString())
  const token = localStorage.getItem('token_CSRF');
  const auth_token = localStorage.getItem('auth_token');
  const user_id = localStorage.getItem('user_id');
  const xhr = new XMLHttpRequest();
  xhr.open("DELETE", "https://51.250.24.31:65000/api/invitation/"+ event_id.toString() + '/',false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  xhr.setRequestHeader("Authorization", "Bearer " + auth_token);
  xhr.withCredentials = true;

  const body = JSON.stringify({
      "event_id": event_id,
      "user_id": user_id,
  });
  xhr.onload = () => {
    if (xhr.readyState == 4 && xhr.status == 201) {
      console.log(JSON.parse(xhr.responseText));
    } else {
      console.log(`Error: ${xhr.status}`);
    }
  };
  xhr.send(body);
  console.log(body)
}