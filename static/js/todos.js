document.addEventListener('DOMContentLoaded', () => {
    let deleteButtons = document.getElementsByClassName('delete_button')

    for (let i = 0; i < deleteButtons.length; i++) {
        let deleteButton = deleteButtons[i]
        var todoId = deleteButton.previousElementSibling.previousElementSibling.innerHTML
        let xhr = new XMLHttpRequest()

        xhr.onreadystatechange = () => {
            const DONE = 4
            const OK = 200
            if (xhr.readyState === DONE) {
                if (xhr.status === OK) {
                    window.location.reload(true)
                } else {
                    console.log("Error: " + xhr.status)
                }
            }
        }
        xhr.open('DELETE', "/todos/delete/id=" + todoId)
        deleteButton.addEventListener('click', (ev) => {
            console.log(xhr)
            xhr.send()
        })
    }

    // bind saveButtons click

    let saveButtons = document.getElementsByClassName('save_button')

    for (let i = 0; i < saveButtons.length; i++) {
        let saveButton = saveButtons[i]
        var todoId = saveButton.previousElementSibling.previousElementSibling.innerHTML
        let xhr = new XMLHttpRequest()

        xhr.onreadystatechange = () => {
            const DONE = 4
            const OK = 200
            if (xhr.readyState === DONE) {
                if (xhr.status === OK) {
                    //window.location.reload(true)
                } else {
                    console.log("Error: " + xhr.status)
                }
            }
        }

        xhr.open('PUT', "/todos/edit/id=" + todoId)
        saveButton.addEventListener('click', (ev) => {
            console.log(ev)
            xhr.send()
        })
    }

})

