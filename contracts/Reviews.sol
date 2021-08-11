// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;
pragma experimental ABIEncoderV2;

contract Reviews {
  uint8 public reviewCount = 0;

  struct Review {
    uint8 id;
    string content;
    string city;
    address author;
  }

  mapping(uint8 => Review) public reviews;
  
  event ReviewCreated(
    uint8 id,
    string content,
    string city,
    address author
  );

  constructor() public {
    createReview("My First Dapp", "Kuala Lumpur", 0x9a06A190c898432142a4e11C02D04cd1e7D92F1e);
  }

  function createReview(string memory _content, string memory _city, address _author) public {
    reviewCount++;
    reviews[reviewCount] = Review(reviewCount, _content, _city, _author);
    emit ReviewCreated(reviewCount, _content, _city, _author);
  }


  function getReviewsById(uint8 _id) public view returns (Review memory) {
    Review storage rev = reviews[_id];
    return rev;
  }
}
