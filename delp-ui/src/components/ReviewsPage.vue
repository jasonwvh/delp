<template>
  <v-container style="height: 100vh">
    <h1 class="pa-12 text-center">
      <span class="text-h4">What people say about: </span>
      <span class="text-h3">{{ place }}</span>
    </h1>
    <v-container class="pa-12" style="width: 65vw; min-width: 760px">
      <v-row>
        <v-col lg="4" md="6" sm="12" v-for="(comment, i) in comments" :key="i">
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
              <v-text-field v-model="newComment.author" placeholder="Sam Smith">
              </v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col style="padding-top: 30px" cols="3" align="end">
              <span> Comment </span>
            </v-col>
            <v-col cols="9">
              <v-textarea
                v-model="newComment.text"
                placeholder="The Woodman set to work at once, and so sharp was his axe that the tree was soon chopped nearly through."
              >
              </v-textarea>
            </v-col>
          </v-row>
          <v-card-actions>
            <v-spacer> </v-spacer>
            <v-btn text @click="postReview"> Submit </v-btn>
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
import getWeb3 from "../web3/getWeb3";
import ReviewsJSON from "../contracts/Reviews.json";

@Component({ components: { ReviewCard } })
export default class ReviewsPage extends Vue {
  @Prop()
  place = "Kuala Lumpur";

  private newComment: CommentInterface = {
    author: "",
    date: "",
    text: "",
    place: "",
  };

  private comments: CommentInterface[] = [];

  private web3: any = {};
  private accounts: any[] = [];
  private contract: any = {};

  async created() {
    try {
      console.log(ReviewsJSON);
      const web3 = await getWeb3();

      const accounts = await web3.eth.getAccounts();

      const networkId = await web3.eth.net.getId();
      const deployedNetwork = ReviewsJSON.networks[networkId];
      const contract = new web3.eth.Contract(
        ReviewsJSON.abi,
        deployedNetwork && deployedNetwork.address
      );

      console.log(web3);
      console.log(accounts);
      console.log(contract);

      this.web3 = web3;
      this.accounts = accounts;
      this.contract = contract;

      this.getReviews();
    } catch (error) {
      console.log(
        `Failed to load web3, accounts, or contract. Check console for details.`
      );
      console.error(error);
    }
  }

  private async getReviews() {
    const reviewCount = await this.contract.methods
      .reviewCount()
      .call(function (err, res) {
        console.log(err, res);
      });
    console.log(reviewCount);
    for (var i = 1; i <= reviewCount; i++) {
      let tmp: {
        /*id: any;*/ text: string;
        date: string;
        place: string;
        author: string;
      };
      const response = await this.contract.methods.getReviewsById(i).call();
      console.log(response);
      //if (response[2] != this.place) continue;

      tmp = {
        //id: response[0],
        date: response[0],
        text: response[1],
        place: response[2],
        author: response[2],
      };
      console.log(tmp);
      this.comments.push(tmp);
    }
  }
  private async postReview() {
    console.log(this.contract);
    await this.contract.methods
      .createReview(this.newComment.text, this.place)
      .send({ from: this.accounts[0] });
  }
  private async sendTips(author: any) {
    console.log(author);

    this.web3.eth.sendTransaction({
      from: this.accounts[0],
      to: author,
      value: this.web3.utils.toWei("5", "ether"),
    });
  }
}
</script>

<style scoped>
</style>
