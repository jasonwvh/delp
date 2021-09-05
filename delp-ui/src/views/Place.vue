<template>
  <div>
    <header>
      <toolbar />
    </header>
    <place-details :place="city" />
    <reviews-page :place="city" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Toolbar from "../components/Toolbar.vue";
import PlaceDetails from "../components/PlaceDetails.vue";
import ReviewsPage from "../components/ReviewsPage.vue";

@Component({
  components: { Toolbar, PlaceDetails, ReviewsPage },
})
export default class Home extends Vue {
  private city = this.$route.params.city;
}
</script>

<style scoped>
</style>
