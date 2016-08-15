document.addEventListener('DOMContentLoaded', () => {
    var todos = document.getElementsByClassName('todo')

    for (let i = 0; i < todos.length; i++) {
        let completeStatus = todos[i].children.completeStatus
        completeStatus.addEventListener('click', (ev) => {
            if (completeStatus.value === 'false') {
                completeStatus.value = 'true'
            } else {
                completeStatus.value = 'false'
            }
        })
    }

    for (let i = 0; i < todos.length; i++) {

        let todoId = todos[i].children.todoId.value
        let deleteButton = todos[i].children.deleteButton

        deleteButton.addEventListener('click', (ev) => {
            let xhr = new XMLHttpRequest()

            xhr.onreadystatechange = () => {
                const DONE = 4
                const OK = 200
                if (xhr.readyState === DONE) {
                    if (xhr.status === OK) {
                        console.log('success')
                        //window.location.reload(true)
                    } else {
                        console.log("Error: " + xhr.status)
                    }
                }
            }
            xhr.open('DELETE', "/todos/delete/id=" + todoId)
            console.log(xhr)
            xhr.send()
        })
    }

    // bind saveButtons click

    let saveButtons = document.getElementsByClassName('save_button')

    for (let i = 0; i < todos.length; i++) {
        let dueDate = todos[i].children.dueDate.value
        let completeStatus = todos[i].children.completeStatus
        
        let todoId = todos[i].children.todoId.value
        let description = encodeURI(todos[i].children.description.value)

        let saveButton = todos[i].children.saveButton

        saveButton.addEventListener('click', (ev) => {
            let xhr = new XMLHttpRequest()
            let params = "dueDate=" + dueDate + "&completeStatus=" + completeStatus + "&todoId=" + todoId + "&description="+ description
            console.log(params)

            xhr.onreadystatechange = () => {
                const DONE = 4
                const OK = 200
                if (xhr.readyState === DONE) {
                    if (xhr.status === OK) {
                        console.log('success')
                        //window.location.reload(true)
                    } else {
                        console.log("Error: " + xhr.status)
                    }
                }
            }
            xhr.open('PUT', "/todos/edit/")
            xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
            xhr.send(params)
        })
    }
})

