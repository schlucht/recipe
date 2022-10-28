<script setup>
import productForm from "../components/product/productForm.vue";
</script>

<template>
  <main class="container">
    <div class="row min-vh-100">
      <section id="product" class="col-12 col-md-6 col-lg-4">
        <h2>Produkt <i class="fa-solid fa-star"></i></h2>
        <product-form />
      </section>
      <section id="unit" class="col-12 col-md-6 col-lg-4">
        <h2>Unit Procedure</h2>
      </section>
      <section id="parameter" class="col-12 col-md-12 col-lg-4">
        <h2>Parameters</h2>
      </section>
    </div>  
  </main>
</template>
<style scoped>

</style>
