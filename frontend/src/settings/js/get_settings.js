const xhr = new XMLHttpRequest();
    const token = localStorage.getItem('token_CSRF') // получения токена из локального хранилища
    const auth_token = localStorage.getItem('auth_token'); // получение auth_token из локального хранилища

    user_id = localStorage.getItem('user_id');

    xhr.open("GET", "https://51.250.24.31:65000/api/user/" + user_id ,false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.setRequestHeader("Authorization", "Bearer " + auth_token);
    xhr.withCredentials = true;
    xhr.send();

    var jsonResponse = JSON.parse(xhr.responseText);

    document.getElementById('username').setText = jsonResponse["username"];
    document.getElementById('email_setting').setText = jsonResponse["login"];