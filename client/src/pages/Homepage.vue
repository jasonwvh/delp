<template>
  <div class="container">
    <div class="banner">
      <Banner />
    </div>
    <div class="recents">
      <Recents />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Banner from "../components/Banner.vue";
import Recents from "../components/Recents.vue"

export default defineComponent({
  name: "Homepage",
  props: {
    country: {
      type: String,
      default: "Malaysia",
    },
  },
  components: {
    Banner, Recents,
  },
  created() {},
  methods: {
    getHello: async () => {
      fetch("http://localhost:3000/hello").then((response) =>
        response.json().then((data) => {
          console.log(data);
        })
      );
    },
    getCities: async (country: string) => {
      fetch(`http://localhost:3000/${country}/cities`).then((response) =>
        response.json().then((data) => {
          console.log(data);
        })
      );
    },
    getCountries: async () => {
      fetch("http://localhost:3000/countries").then((response) =>
        response.json().then((data) => {
          console.log(data);
        })
      );
    },
  },
});
</script>

<style scoped>
.banner {
  height: 100vh;
  width: 100vw;
}
</style>
