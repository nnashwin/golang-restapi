document.addEventListener("DOMContentLoaded", () => {
    let form = document.querySelector('form')
    form.addEventListener('submit', (e) => {
        e.preventDefault()

        let request = new XMLHttpRequest()
        let inputEls = document.getElementsByClassName('formVal')
        let data = {}

        for (let i = 0; i < inputEls.length; i++) {
            data[inputEls[i].name] = inputEls[i].value
        }

        let params = Object.keys(data).map(
            (k) => {
                return encodeURIComponent(k) + '=' + encodeURIComponent(data[k])
            }
        ).join("&")

        let failure = (reqStatus) => {
            console.error("Error: " + reqStatus)
        }

        request.onreadystatechange = () => {
            const DONE = 4
            const OK = 200
            if (request.readyState === DONE) {
                if (request.status === OK) {
                    window.location = "/todos"
                } else {
                    failure(request.status)
                }
            }
        }

        request.open("POST", "/todos")

        request.setRequestHeader('X-Requested-With', 'XMLHttpRequest')
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded')

        request.send(params)
    })
})
