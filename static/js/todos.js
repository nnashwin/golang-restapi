window.onload = () => {
    let deleteButtons = document.getElementsByClassName('delete_button')

    for (let i = 0; i < deleteButtons.length; i++) {
        let deleteButton = deleteButtons[i]
        console.log(deleteButton)
        console.log(i)
        deleteButton.addEventListener('click', (ev) => {
            console.log(ev)
        })
    }
}

