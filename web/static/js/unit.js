export class Unit {
    constructor(element, data) {
        this.data = data;
        this.element = element;
        console.log(this.data)
        this.printUnit();
    }

    printUnit() {
        if (this.data && this.data.length > 0) {
            this.data.forEach(unit => {
                let txt = /*html*/`
                    <li class="nav-item active p-1 d-flex align-items-center">
                        <button class="btn btn-secondary w-100 position-relative" type="button" id="${unit.name}">
                            ${unit.name}
                            <span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger" id="del-${unit.name}">
                            X
                        <span class="visually-hidden">unread messages</span>
                    </span>
                        </button>
                    </li>
                `;
                this.element.insertAdjacentHTML('beforeend', txt);
            });
        }
    }

}