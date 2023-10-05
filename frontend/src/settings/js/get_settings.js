const xhr = new XMLHttpRequest();

    user_id = localStorage.getItem('user_id');

    xhr.open("GET", "https://51.250.24.31:65000/api/user/" + user_id ,false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.withCredentials = true;

    var jsonResponse = JSON.parse(xhr.responseText);

    document.getElementById('username').innerHTML = jsonResponse["username"];
    document.getElementById('email_setting').innerHTML = jsonResponse["login"];