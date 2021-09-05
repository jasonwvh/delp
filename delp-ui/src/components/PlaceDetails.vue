<template>
  <v-container fluid class="pa-0">
    <carousel :perPage="5" :autoplay="true" :paginationEnabled="false">
      <slide v-for="i in 10" :key="i"
        ><v-img max-height="720" :src="images + i"
      /></slide>
    </carousel>
  </v-container>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import { Carousel, Slide } from "vue-carousel";

@Component({ components: { Carousel, Slide } })
export default class PlaceDetails extends Vue {
  @Prop()
  place: string;

  @Prop()
  country: string;

  private images = `https://source.unsplash.com/720x720/?${this.place},${this.country}?sig=`;
}
</script>

<style scoped>
</style>
