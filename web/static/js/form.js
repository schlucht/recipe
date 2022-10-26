const unit = document.getElementById("unitcol");
const newProduct = document.getElementById('newProduct');

const save = newProduct.querySelector('button');
const productInput = newProduct.querySelector('input');
const productDesc = newProduct.querySelector('textarea');


function Unit(data) {

    read(data);

    save.addEventListener('click', () => {
        const product = productInput.value;
        const desc = productDesc.value;
        if (!data.find(d => d.product === product) || product !== "") {
            data.push({ product, desc })
            read(data);
        }        
    });

}



function read(data) {
    unit.innerHTML = "";
    data.forEach(product => {
        let prod = product.product.toLowerCase();

        unit.insertAdjacentHTML("beforeend", readProducts(prod))
        const rws = product.unit;
        if (rws) {
            console.log(rws);
            const rwtxt = document.getElementById(`${prod}`);
            const list = document.querySelector('ul')
            rws.forEach(r => {

                list.insertAdjacentHTML('beforeend', readUnits(r.name));
            })
        }
    });
}

function readProducts(prod) {
    let del = newProduct.querySelector("#proddel");
    let txt = /*html*/`
            <div class="p-2 border-bottom">
            <button class="w-100 btn btn-primary position-relative" type="button" data-bs-toggle="collapse"
                data-bs-target="#${prod}" role="button" aria-expanded="false" aria-controls="${prod}">
                ${prod.toUpperCase()}                
                </button>
                <button type="button" class="btn-close" aria-label="Close" id="proddel" onclick="del()">              
                </button>
                <script>
                    function del() {
                        console.log('Hallo')
                    }
                </script>
            <div class="collapse" id="${prod}">
                <form class="container border p-1">
                    <div class="mb-3">
                        <input type="text" class="form-control" placeholder="Neue Unit eingeben"
                            id="productname" />
                    </div>
                    <button type="button" class="btn btn-success" onclick="newUnit('${prod}')">
                        Speichern
                    </button>
                </form>
                <ul class="navbar-nav flex-column" id="u-${prod}">

                </ul>
            </div>
        </div>
            `;

    return txt;
}


function readUnits(unit) {
    return /*html*/ `
    <li class="nav-item active ps-4">
        <a class="nav-link" href="#">${unit}</a>
    </li>
    `;
}

export { Unit }