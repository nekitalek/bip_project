const xhr_get = new XMLHttpRequest();
    const token = localStorage.getItem('token_CSRF') // получения токена из локального хранилища
    const auth_token = localStorage.getItem('auth_token'); // получение auth_token из локального хранилища

    var user_id = localStorage.getItem('user_id');

    xhr_get.open("GET", "https://51.250.24.31:65000/api/user/" + user_id ,false);
    xhr_get.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr_get.setRequestHeader("X-CSRF-TOKEN", token);
    xhr_get.setRequestHeader("Authorization", "Bearer " + auth_token);
    xhr_get.withCredentials = true;
    xhr_get.send();

    if(xhr_get.status==200){
        window.location.href = "https://51.250.24.31/main/"; 
    }