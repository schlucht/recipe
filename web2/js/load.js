fetch('../../web/data.json')
.then((resp) => resp.json())
.then((val) => {
    localStorage.setItem('file', JSON.stringify(val));   
})

const dataString = localStorage.getItem('file');
export const data = JSON.parse(dataString);


