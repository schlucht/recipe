
window.addEventListener('load', function() {
    const btn = document.querySelector('button')
    const txt = document.querySelector('.txt')
    btn.addEventListener('click', () => {
        txt.innerHTML = `<p class="text-primary">Hallo Lothar</p>`
    })
})