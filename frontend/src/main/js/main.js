window.addEventListener("DOMContentLoaded", (event) => {
    const el = document.getElementById('send_form');
    if (el) {
      el.addEventListener('click', CreateEvent);
    }
  });
  
  function CreateEvent(){
  
    const token = localStorage.getItem('token_CSRF')
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31:65000/api/event",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
  
    xhr.withCredentials = true;
  
    var time_start = document.getElementById("start_time").value;
    var time_end = document.getElementById("end_time").value;
    var place = document.getElementById("address").value;
    var description = document.getElementById("description").value;
    var public = document.getElementsById('checkbox').checked;
    const game = document.querySelector('input[name="list-radio"]:checked').value;

    const body = JSON.stringify({
        "time_start": time_start,
        "time_end": time_end,
        "place": place,
        "game": game,
        "description": description,
        "public": public
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