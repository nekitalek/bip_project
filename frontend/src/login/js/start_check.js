const xhr = new XMLHttpRequest();
    const token = localStorage.getItem('token_CSRF') // получения токена из локального хранилища
    const auth_token = localStorage.getItem('auth_token'); // получение auth_token из локального хранилища

    var user_id = localStorage.getItem('user_id');

    xhr.open("GET", "https://51.250.24.31/api/user/" + user_id ,false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.setRequestHeader("Authorization", "Bearer " + auth_token);
    xhr.withCredentials = true;
    xhr.send();

    if(xhr.status==200){
        window.location.href = "https://51.250.24.31/main/"; 
    }