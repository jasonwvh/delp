<template>
  <div class="search">
    
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "HelloWorld",
  props: {
    msg: String,
  },
  created() {
    this.getCountries();
  },
  methods: {
    getHello: async () => {
      fetch("http://localhost:3000/hello").then((response) =>
        response.json().then((data) => {
          console.log(data);
        })
      );
    },
    getCities: async () => {
      fetch("http://localhost:3000/cities").then((response) =>
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

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
