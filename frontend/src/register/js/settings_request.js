const requestURL = 'https://jsonplaceholder.typicode.com/users'

async function sendRequest(method, url, body = null){
    return await fetch(url).then(response =>{
        return response.json()
    })
}



let k = 1

// request to recieve user settings from server
sendRequest('GET', requestURL)
    .then(function(data){
        let username = document.getElementById("username").value = data[k].username;
        let email = document.getElementById("email_setting").value = data[k].email;
        //let phone = document.getElementById("phone").value = data[k].phone;
    })
    .catch(err => console.log(err))

// requests to send changed settings to server and verify them
sendRequest('POST', requestURL)
    .then(function(){
        var login = document.getElementById("email_setting").value;
        var old_pw = document.getElementById("old_password").value;
        var xhr = new XMLHttpRequest();
        xhr.open("POST", requestURL, true);
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.send(JSON.stringify({
            value: input_val
        }));
    })
    .catch(err => console.log(err))


sendRequest('POST', requestURL)
.then(function(){
    var old_pw = document.getElementById("old_password").value;
    let new_pw = document.getElementById("new_password").value;
})
.catch(err => console.log(err))