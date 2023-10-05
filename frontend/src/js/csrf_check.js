const requestURL = 'https://192.168.153.129:65000/CSRF'

async function sendRequest(method, url, body = null){
    return await fetch(url).then(response =>{
        return response.json()
    })
}

// request to recieve csrf from server
sendRequest('GET', requestURL)
    .then(function(data){
        var csrf = data.token_CSRF
        console.log(data.token_CSRF)
    })
    .catch(err => console.log(err))
