<template>
  <v-container style="height: 100vh">
    <h1 class="pa-12 text-center">
      <span class="text-h4">What people say about: </span>
      <span class="text-h3">{{ place }}</span>
    </h1>
    <v-container class="pa-12" style="width: 65vw; min-width: 760px">
      <v-row>
        <v-col lg="3" md="4" sm="6" v-for="(comment, i) in comments" :key="i">
          <review-card
            :author="comment.author"
            :date="comment.date"
            :text="comment.text"
          />
        </v-col>
      </v-row>
    </v-container>
    <v-container class="pa-12">
      <v-row justify="center">
        <v-card
          elevation="3"
          class="pa-12 mt-12 mb-12"
          style="min-width: 660px"
        >
          <v-row justify="center">
            <h4>Leave a comment</h4>
          </v-row>
          <v-row justify="center" align="center">
            <v-col cols="3" justify="center" align="center">
              <span> Name (optional) </span>
            </v-col>
            <v-col cols="9" justify="center" align="center">
              <v-text-field placeholder="Sam Smith"> </v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col style="padding-top: 30px" cols="3" align="end">
              <span> Comment </span>
            </v-col>
            <v-col cols="9">
              <v-textarea
                placeholder="The Woodman set to work at once, and so sharp was his axe that the tree was soon chopped nearly through."
              >
              </v-textarea>
            </v-col>
          </v-row>
          <v-card-actions>
            <v-spacer> </v-spacer>
            <v-btn text> Submit </v-btn>
          </v-card-actions>
        </v-card>
      </v-row>
    </v-container>
  </v-container>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import ReviewCard from "./ReviewCard.vue";
import CommentInterface from "../interface/CommentInterface";

@Component({ components: { ReviewCard } })
export default class Reviews extends Vue {
  @Prop()
  place = "Kuala Lumpur";

  private newComment: CommentInterface = {} as CommentInterface;

  private comments: CommentInterface[] = [];

  mounted() {
    this.comments = [
      {
        author: "Anonymous",
        date: "5/9/2021",
        text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam aliquet odio posuere velit vehicula, tincidunt euismod lectus fringilla. Suspendisse hendrerit leo non neque rutrum euismod. Integer vestibulum id mi sed blandit. Curabitur eget gravida ex. Pellentesque id orci sed mi laoreet fermentum quis id lectus. ",
      },
      {
        author: "Anonymous",
        date: "5/9/2021",
        text: "Sed porttitor rhoncus erat, sed finibus augue viverra non. Quisque consectetur nulla ut est tincidunt molestie. Maecenas vestibulum scelerisque urna. ",
      },
      {
        author: "Anonymous",
        date: "4/9/2021",
        text: "Nulla ullamcorper neque a magna convallis, vel porta sapien semper. Interdum et malesuada fames ac ante ipsum primis in faucibus. Sed viverra libero non sapien sollicitudin, nec blandit arcu feugiat.  ",
      },
      {
        author: "Anonymous",
        date: "5/8/2021",
        text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam aliquet odio posuere velit vehicula, tincidunt euismod lectus fringilla. Suspendisse hendrerit leo non neque rutrum euismod. Integer vestibulum id mi sed blandit. Curabitur eget gravida ex. Pellentesque id orci sed mi laoreet fermentum quis id lectus. ",
      },
      {
        author: "Anonymous",
        date: "5/8/2021",
        text: "Sed porttitor rhoncus erat, sed finibus augue viverra non. Quisque consectetur nulla ut est tincidunt molestie. Maecenas vestibulum scelerisque urna. ",
      },
      {
        author: "Anonymous",
        date: "4/8/2021",
        text: "Nulla ullamcorper neque a magna convallis, vel porta sapien semper. Interdum et malesuada fames ac ante ipsum primis in faucibus. Sed viverra libero non sapien sollicitudin, nec blandit arcu feugiat.  ",
      },
    ];
  }

  private postComment() {
    console.log(this.newComment);
  }
}
</script>

<style scoped>
</style>
