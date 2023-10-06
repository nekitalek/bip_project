const xhr = new XMLHttpRequest();

    xhr.open("GET", "https://51.250.24.31:65000/api/user/" + user_id ,false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.setRequestHeader("Authorization", "Bearer " + auth_token);
    xhr.withCredentials = true;
    xhr.send();
    if(xhr_get.status!=200){
        window.location.href = "https://51.250.24.31/login/"; 
    }
    var jsonResponse = JSON.parse(xhr.responseText);
    document.getElementById("win_name").innerHTML = jsonResponse["username"];
    document.getElementById("win_email").innerHTML = jsonResponse["login"];