<template>
  <div>
    <form>
      <div>
        <label for="productname">Produkt</label>
        <input
          class="form-control"
          type="text"
          placeholder="Neues Produkt..."
          name="productname"
          id="productname"
        />
        <button
          type="button"
          class="btn btn-secondary w-100"
          @click="add('Lothar')"
        >
          <i class="fa-regular fa-save"></i>Speichern
        </button>
      </div>
    </form>
    <ul v-for="prod in dat" :key="prod.productId">
      <li>{{ prod.productName }}</li>
    </ul>
  </div>
</template>

<script setup>
import Store from "../../stores/store";
const dat = Store.state.product;

function add(s) {
  Store.mutations.add(s);
  console.log(dat);
}
</script>

<style scope></style>
