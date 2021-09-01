<template>
  <div class="container">
      <h1> {{ city }} </h1>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Reviews",
  data() {
    return {
        city: "Kuala Lumpur"
    };
  },
  created() {

  },
  methods: {

  },
});
</script>

<style scoped>
.container {
  height: 100vh;
  width: 100vw;
  display: flex;
  justify-content: center;
}

</style>
