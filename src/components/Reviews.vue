<template>
  <div class="reviews">
    <span> Showing recommendations for <span v-if="city"> {{ city }} </span> </span>
    <div class="cards">
      <article class="card" v-for="review in reviews" :key="review.id">
        <div class="content">
          {{review.content}}
        </div>
        <div class="tip">
          <button @click="sendTips(review.author)">Send Tips</button>
        </div>
      </article>
    </div>

    <div class="content-area">
      <textarea v-model="content" placeholder="Review..." />
    </div>
    <div class="post-btn">
      <button @click="createReview(this.content)">Post Review</button>
    </div>
  </div>
</template>


<script>
import Reviews from "../../build/contracts/Reviews.json";
import getWeb3 from "../getWeb3";

export default {
	name: 'Reviews',
	data() {
		return {
      accounts: null,
      contract: null,
      web3: null,
      reviews: [],
      content: null,
      city: null,
		}
	},
  async mounted() {
    await fetch('http://www.geoplugin.net/json.gp')
      .then(response => response.json())
      .then(data => (this.city = data.geoplugin_city));
    },
  async created() {
    try {
      const web3 = await getWeb3();

      const accounts = await web3.eth.getAccounts();

      const networkId = await web3.eth.net.getId();
      const deployedNetwork = Reviews.networks[networkId];
      const contract = new web3.eth.Contract(
        Reviews.abi,
        deployedNetwork && deployedNetwork.address,
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
        `Failed to load web3, accounts, or contract. Check console for details.`,
      );
      console.error(error);
    }
  },
  methods: {
    async getReviews() {
      const reviewCount = await this.contract.methods.reviewCount().call()

      for (var i=1; i <= reviewCount; i++) {
        let tmp;
        const response = await this.contract.methods.getReviewsById(i).call();

        // filter by city
        if (response[2] != this.city) continue;

        tmp = {
          id: response[0],
          content: response[1],
          city: response[2],
          author: response[3],
        }
        console.log(tmp)
        this.reviews.push(tmp);
      }
    },
    async createReview(content) {
      await this.contract.methods.createReview(content, this.city, this.accounts[0]).send({from: this.accounts[0] });
    },
    async sendTips(author) {
      console.log(author)

      this.web3.eth.sendTransaction({from: this.accounts[0], to:author, value:this.web3.utils.toWei('5', "ether")});
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.cards {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(230px, 1fr));
    grid-gap: 20px;
}

.card {
    display: grid;
    grid-template-rows: max-content 200px 1fr;
}
</style>
