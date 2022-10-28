import { product } from "./data";
import { reactive, readonly } from "vue";

// Einlesen des Produktes
const state = reactive({
  product,
});

const getters = {};

// Funktionen zum Bearbeiten
const mutations = {
  add(s) {
    state.product.push({ productName: s });
  },
};

export default {
  state: readonly(state),
  getters,
  mutations,
};
