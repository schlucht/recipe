
fetch('../../web/data.json')
.then((resp) => resp.json())
.then((val) => {
    localStorage.setItem('file', JSON.stringify(val));   
})




