const requestURL = 'https://192.168.153.128:65000/CSRF'

async function sendRequest(method, url, body = null){
    return await fetch(url).then(response =>{
        return response.json()
    })
}

// request to recieve csrf from server
sendRequest('GET', requestURL)
    .then(function(data){
        const csrf = JSON.stringify(data.token_CSRF)
        localStorage.setItem('token_CSRF', csrf);
        console.log(data.token_CSRF)
    })
    .catch(err => console.log(err))