export default () => {
    const toggle = document.getElementById('toggle')

    toggle.addEventListener('change', (event) => {
        if (event.target.checked) {
            document.body.classList.remove('black')
        } else {
            document.body.classList.add('black')
        }

    })
}