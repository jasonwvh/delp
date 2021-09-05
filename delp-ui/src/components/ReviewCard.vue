<template>
  <v-card elevation="3" height="420" width="420" min-width="240">
    <v-card-title>
      <p>{{ author }}</p>
    </v-card-title>
    <v-card-subtitle>
      <p>{{ date }}</p>
    </v-card-subtitle>
    <v-card-text class="text-wrap">
      {{ content }}
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component({ components: {} })
export default class ReviewCard extends Vue {
  @Prop({ default: "" })
  author!: string;

  @Prop({ default: "" })
  date!: string;

  @Prop({ default: "" })
  content!: string;
}
</script>

<style scoped>
</style>
