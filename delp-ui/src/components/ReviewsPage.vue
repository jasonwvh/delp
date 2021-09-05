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
            :author="comment.authorName"
            :date="comment.date"
            :content="comment.content"
            @send-tip="sendTips(comment.authorAddress, '5')"
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
              <v-text-field
                v-model="commentInput.authorName"
                placeholder="Sam Smith"
              >
              </v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col style="padding-top: 30px" cols="3" align="end">
              <span> Comment </span>
            </v-col>
            <v-col cols="9">
              <v-textarea
                v-model="commentInput.content"
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
import moment from "moment";

@Component({ components: { ReviewCard } })
export default class ReviewsPage extends Vue {
  @Prop()
  place!: string;

  private commentInput = {} as CommentInterface;

  private comments: CommentInterface[] = [];

  private web3: any = {};
  private accounts: any[] = [];
  private contract: any = {};

  async created(): Promise<void> {
    try {
      const web3 = await getWeb3();
      const accounts = await web3.eth.getAccounts();
      const networkId = await web3.eth.net.getId();
      const deployedNetwork = ReviewsJSON.networks[networkId];
      const contract = new web3.eth.Contract(
        ReviewsJSON.abi,
        deployedNetwork && deployedNetwork.address
      );

      this.web3 = web3;
      this.accounts = accounts;
      this.contract = contract;

      console.log(accounts);

      this.getReviews();
    } catch (error) {
      console.error(error);
    }
  }

  private async getReviews() {
    this.comments = [];

    const reviewCount = await this.contract.methods.reviewCount().call();

    for (var i = 1; i <= reviewCount; i++) {
      let tmp: {
        authorName: string;
        authorAddress: string;
        content: string;
        date: string;
        place: string;
      };

      const response = await this.contract.methods.getReviewsById(i).call();

      if (response[5] != this.place) continue;

      tmp = {
        authorName: response[1],
        authorAddress: response[2],
        content: response[3],
        date: response[4],
        place: response[5],
      };

      this.comments.push(tmp);
    }
  }
  private async postReview() {
    let name: string;
    if (
      this.commentInput.authorName == null ||
      this.commentInput.authorName == undefined ||
      this.commentInput.authorName.length == 0
    ) {
      name = "Anonymous";
    } else {
      name = this.commentInput.authorName;
    }

    const comment: CommentInterface = {
      authorName: name,
      authorAddress: this.accounts[0],
      content: this.commentInput.content,
      date: moment(Date.now()).format("Do MMMM YYYY"),
      place: this.place,
    };

    await this.contract.methods
      .createReview(
        comment.authorName,
        comment.authorAddress,
        comment.content,
        comment.date,
        comment.place
      )
      .send({ from: this.accounts[0] });

    this.getReviews();
  }

  private async sendTips(authorAddress: string, amount: string) {
    this.web3.eth.sendTransaction({
      from: this.accounts[0],
      to: authorAddress,
      value: this.web3.utils.toWei(amount, "wei"),
    });
  }
}
</script>

<style scoped>
</style>
