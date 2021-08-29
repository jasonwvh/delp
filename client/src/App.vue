<template>
  <v-app>
    <v-main>
      <Search country="Malaysia" />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Search from "./components/Search.vue";

export default defineComponent({
  name: "App",

  components: {
    Search,
  },

  data() {
    return {
      //
    };
  },
});
</script>
