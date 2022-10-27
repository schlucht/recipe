import { Unit } from './unit.js';

export class Product {
    constructor(element, data) {
        this.data = data;
        this.products = data;
        
        this.element = element;
        this.txtArray = [];
        this.printProduct();
    }    
    printProduct() {
        this.element.innerHTML = "";
        this.products.forEach(product => {
            const prod = product.product;
            let txt = /*html*/`
            <div class="p-2 border-bottom">
                <button class="w-100 btn btn-primary position-relative" type="button" data-bs-toggle="collapse"
                data-bs-target="#${prod}" role="button" aria-expanded="false" aria-controls="${prod}">
                ${prod.toUpperCase()}   
                    <span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger" id="del-${prod}">
                            X
                        <span class="visually-hidden">unread messages</span>
                    </span>             
                </button>                                  
                <div class="collapse" id="${prod}">
                    <form class="container border p-1">
                        <div class="mb-3">
                            <input type="text" class="form-control" placeholder="Neue Unit eingeben"
                                id="unitname" />
                        </div>
                        <button type="button" class="btn btn-success" id="new-${prod}">
                            Speichern
                        </button>
                    </form>
                    <ul class="navbar-nav flex-column align-items-strech" id="u-${prod}">
                        <!-- Rührwerk auflistung -->
                    </ul>
                </div>
            </div>
            `;
            this.element.insertAdjacentHTML('beforeend', txt);
            let btnDel = this.element.querySelector(`#del-${prod}`);
            let btnNew = this.element.querySelector(`#new-${prod}`)
            btnDel.addEventListener('click', (e) => {
                const d = e.target.parentElement.dataset['bsTarget'].slice(1);
                const s = this.products.findIndex(f => {                    
                    return f.product === d
                })  
                let b = confirm(`Wollen sie ${d} löschen?`);
                if (b) {
                    this.products.splice(s, 1);
                    this.printProduct();
                }  
            });

            btnNew.addEventListener('click', () => {
                const ele = this.element.querySelector(`#${prod}`);
                const input = this.element.querySelector('input');
                const val = input.value;
                if (val) {
                    let b = confirm(`Wollen sie ${val} speichern?`)
                    if (b) {
                        product.unit.push(val);                    
                        this.printProduct();
                    } else {
                        input.value = "";
                    }
                    
                }
            });

            new Unit(
                this.element.querySelector(`#u-${prod}`),
                product.unit)
        });

    }
}