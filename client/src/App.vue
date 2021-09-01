<template>
  <v-app>
    <v-main>
      <router-link to="/">Home</router-link> |
      <router-link to="/place">Place</router-link>
    </v-main>
    <router-view />
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",

  components: {

  },

  data() {
    return {
    };
  },

});
</script>