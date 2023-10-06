const xhr_get = new XMLHttpRequest();
    const token = localStorage.getItem('token_CSRF') // получения токена из локального хранилища
    const auth_token = localStorage.getItem('auth_token'); // получение auth_token из локального хранилища

    user_id = localStorage.getItem('user_id');

    xhr_get.open("GET", "https://51.250.24.31:65000/api/user/" + user_id ,false);
    xhr_get.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr_get.setRequestHeader("X-CSRF-TOKEN", token);
    xhr_get.setRequestHeader("Authorization", "Bearer " + auth_token);
    xhr_get.withCredentials = true;
    xhr_get.send();

    var jsonResponse = JSON.parse(xhr_get.responseText);

    document.getElementById("username").value = jsonResponse["username"];
    document.getElementById("email_setting").value = jsonResponse["login"];
    document.getElementById("win_name").innerHTML = jsonResponse["username"];
    document.getElementById("win_email").innerHTML = jsonResponse["login"];