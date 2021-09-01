<template>
  <div class="container" :style="{backgroundImage: wallpaper}">
    <div class="background">
      <div class="wrapper">
        <input placeholder="Search a place" />
        <input placeholder="Near... " />
        <button>Search</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Banner",
  data() {
    return {
      wallpaper: `url(https://source.unsplash.com/${window.innerWidth}x${window.innerHeight}/?travel)`,
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
  align-items: center;
}
.background {
  height: 80%;
  width: 80%;
  background-color: aliceblue;
  margin: 0 auto;
  display: flex;
  align-items: center;
}
.wrapper {
  margin: 0 auto;
  display: flex;
  align-items: center;
}
</style>
