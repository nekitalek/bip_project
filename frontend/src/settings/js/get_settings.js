const xhr = new XMLHttpRequest();

    const token = localStorage.getItem('token_CSRF')
    const auth_token = localStorage.getItem('auth_token');
    const user_id = localStorage.getItem('user_id')

    xhr.open("GET", "https://51.250.24.31/api/user/" + user_id ,false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.setRequestHeader("Authorization", "Bearer " + auth_token);
    xhr.withCredentials = true;
    xhr.send();

    if(xhr.status!=200){
        window.location.href = "https://51.250.24.31/login/"; 
    }
    var jsonResponse = JSON.parse(xhr.responseText);

    document.getElementById("username").value = jsonResponse["username"];
    document.getElementById("email_setting").value = jsonResponse["login"];
    document.getElementById("win_name").innerHTML = jsonResponse["username"];
    document.getElementById("win_email").innerHTML = jsonResponse["login"];