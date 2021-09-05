<template>
  <div class="background">
    <div class="glass">
      <v-container fluid class="grow d-flex flex-column flex-nowrap">
          <v-row class="grow" align-content="center" justify="center">
          <v-img max-height="150" max-width="250" src="../assets/logo.png" />
          </v-row>
          <v-row class="grow" align-content="start" justify="center">
        <search />
        </v-row>
      </v-container>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Search from "../components/Search.vue";

@Component({
  components: {
    Search,
  },
})
export default class Banner extends Vue {}
</script>

<style scoped>
.background {
  height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  background: url("../assets/bg2.jpg") no-repeat center center / cover;
  margin: 0 auto;
}

.glass {
  height: 80%;
  width: 80%;
  max-height: 600px;
  max-width: 1200px;
  box-shadow: 0 0 1rem 0 rgba(0, 0, 0, 0.2);
  border-radius: 5px;
  background-color: rgba(255, 255, 255, 0.15);
  margin: 0 auto;
  display: flex;
  align-items: center;
  flex-direction: column;
}

.glass-container {
  height: 50%;
}
</style>
