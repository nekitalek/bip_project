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
      var jsonResponse = JSON.parse(xhr.responseText);
      console.log(jsonResponse);
    }
    } catch(err) {
    alert("Запрос не удался");
    }

}