
import { Product } from './product.js';

const unit = document.getElementById("unitcol");
const newProduct = document.getElementById('newProduct');

const save = newProduct.querySelector('button');
const productInput = newProduct.querySelector('input');
const productDesc = newProduct.querySelector('textarea');

fetch('../../web/data.json')
.then((resp) => resp.json())
.then((val) => {
    localStorage.setItem('file', JSON.stringify(val));   
})

const dataString = localStorage.getItem('file');
const data = JSON.parse(dataString);

const p = new Product(unit, data);





