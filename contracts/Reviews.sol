// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;
pragma experimental ABIEncoderV2;

contract Reviews {
  uint8 public reviewCount = 0;

  struct Review {
    uint8 id;
    string content;
    string city;
  }

  mapping(uint8 => Review) public reviews;
  
  event ReviewCreated(
    uint8 id,
    string content,
    string city
  );

  constructor() public {
    createReview("My First Dapp", "Kuala Lumpur");
  }

  function createReview(string memory _content, string memory _city) public {
    reviewCount++;
    reviews[reviewCount] = Review(reviewCount, _content, _city);
    emit ReviewCreated(reviewCount, _content, _city);
  }


  function getReviewsById(uint8 id) public view returns (Review memory) {
    Review storage rev = reviews[id];
    return rev;
  }
}
