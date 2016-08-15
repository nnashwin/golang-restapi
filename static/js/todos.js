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
        console.log(deleteButton)

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
            let saveButton = todos[i].children.saveButton

        saveButton.addEventListener('click', (ev) => {
            let dueDate = todos[i].children.dueDate.value
            let todoId = todos[i].children.todoId.value
            console.log(todos[i].children)
            let todoName = encodeURI(todos[i].children.todoName.value)
            let description = encodeURI(todos[i].children.description.value)
            let completeStatus = todos[i].children.completeStatus.value
            let params = "todoName=" + todoName + "&dueDate=" + dueDate + "&completeStatus=" + completeStatus + "&todoId=" + todoId + "&description="+ description


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
            xhr.open('PUT', "/todos/edit/")
            xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
            xhr.send(params)
        })
    }
})

