document.addEventListener("DOMContentLoaded", () => {
    function sendData() {
        let xhr = new XMLHttpRequest()

        let FD = new FormData(form)

        console.log(FD)
        xhr.addEventListener('load', (e) => {
            alert(e.target.responseText)
        })

        xhr.addEventListener('error', (e) => {
            alert("there has been a mistake!!")
        })

        xhr.open("POST", "/todos")

        xhr.send(FD)
    }

    let form = document.getElementById('todoForm')
    form.addEventListener('submit', (e) => {
        e.preventDefault()

        sendData()
    })
})
