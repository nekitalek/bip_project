const requestURL = 'https://158.160.27.251:443/CSRF'

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
        let getting = browser.cookies.get(
            Cookie                // object
          )
        console.log(Cookie)
    })
    .catch(err => console.log(err))
