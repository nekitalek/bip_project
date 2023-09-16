const requestURL = 'https://jsonplaceholder.typicode.com/users'

async function sendRequest(method, url, body = null){
    return await fetch(url).then(response =>{
        return response.json()
    })
}



let k = 1

sendRequest('GET', requestURL)
    .then(function(data){
        let username = document.getElementById("username").value = data[k].username;
        let email = document.getElementById("email_setting").value = data[k].email;
        let phone = document.getElementById("phone").value = data[k].phone;
    })
    .catch(err => console.log(err))
