import { Product } from './js/model/product.js';
import './style.css'


// import { data } from './js/load.js';
import { ProductSection } from './js/productSection.js';
const prodList = [
    new Product("ASURA"),
    new Product("Foram"),
]

const prod = new ProductSection(prodList);


