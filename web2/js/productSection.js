import { Product } from './model/product.js';
export class ProductSection {
    constructor (data) {
        this.container = document.getElementById('product');
        this.data = data;
        this.products = [];
    }

    save(name) {
        this.products.push(new Product(name));
    }
    print() {
        let txt;
        this.products.forEach((prod) => {
            txt += `<br><h3>${prod}</h3>`
        })
    }
}