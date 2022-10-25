function newUnit() {            
    const unit = document.getElementById('unitname');
    const col = document.getElementById('unitcol')
    const val = unit.value.toLowerCase();
    if (val) {
        let txt = /*html*/`
        <div class="p-2 border-bottom">
            <button class="w-100 btn btn-primary" 
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#${val}"
                role="button"
                aria-expanded="false"
                aria-controls="${val}"
            >${val.toUpperCase()}</button>
            <div class="collapse" id="${val}">
                <form class="container border p-1">
                    <div class="mb-3">
                        <input type="text" class="form-control" placeholder="Neues Produkt eingeben" id="productname">
                    </div>
                    <button type="button" class="btn btn-success" onclick="newProduct('${val}')">Speichern</button>
                </form>
                <ul class="navbar-nav flex-column">                            
                </ul>
            </div>
        </div>
        `;
        col.insertAdjacentHTML("beforeend", txt);
    }
    console.log(unit.value);
}

export {newUnit}