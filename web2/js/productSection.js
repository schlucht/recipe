import { Product } from './model/product.js';
export class ProductSection {
    constructor (data) {
        this.container = document.getElementById('product');
        this.data = data;
        this.products = [];
        this.printForm();
        this.print();    
    }

    save(name) {
        this.data.push(new Product(name));
    }
    print() {
        let txt;
        const listContainer = this.container.querySelector('.listProduct');
        const ul = document.createElement('ul');
        let li = [];
        this.data.forEach((prod) => {
            let list = document.createElement('li');
            let btn = document.createElement('button');
            btn.innerText = prod.name;
            btn.classList.add('btn', 'btn-product')
            btn.setAttribute('type', 'button');
            btn.setAttribute('id', `btn-${prod.name}`);
            btn.addEventListener('click', () => {
                console.log(prod.name)
            });
            list.appendChild(btn);
            li.push(list);            
        });
        listContainer.innerHTML = "";
        li.forEach(l => listContainer.appendChild(l));        
    }

    printForm() {
        let txt = `
        <form id="newProduct">
        <fieldset>
          <input class="input-control" type="text" name="productname" id="productName" placeholder="Neues Produkt">
          <button class="btn" type="button" id="saveProduct">
            <span class="material-symbols-outlined">
              save
              </span>
          </button>
        </fieldset>
      </form>  
      <div class="listProduct">
      </div>
        `;        
        this.container.insertAdjacentHTML('beforeend', txt);
        const btn = this.container.querySelector('#saveProduct');
        const input = this.container.querySelector('#productName');
        btn.addEventListener('click', () => {
            let name = input.ariaValueMax;
            if (name.trim()) {
                this.data.push(new Product(name));
                this.print()
            }
        })
    }


}