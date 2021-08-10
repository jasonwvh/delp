const Reviews = artifacts.require("Reviews");

/*
 * uncomment accounts to access the test accounts made available by the
 * Ethereum client
 * See docs: https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript
 */
contract("Reviews", function (/* accounts */) {
  before(async () =>{
    this.reviews = await Reviews.deployed();
  })

  it("should assert true", async function () {
    await Reviews.deployed();
    return assert.isTrue(true);
  });

  it('lists reviews', async () => {
    const reviewCount = await this.reviews.reviewCount()
    const review = await this.reviews.reviews(reviewCount)

    assert.equal(review.id.toNumber(), reviewCount.toNumber())
  })

  it('get reviews', async () => {
    const review = await this.reviews.getReviewsById()
    assert.equal(review, "Hello!%")
  })

  it('creates reviews', async () => {
    const result = await this.reviews.createReview('A new review', 'KL')
    const reviewCount = await this.reviews.reviewCount()

    assert.equal(reviewCount, 2)
  })
});
