// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;
pragma experimental ABIEncoderV2;

contract Reviews {
    uint8 public reviewCount = 0;

    struct Review {
        uint8 id;
        string authorName;
        string authorAddress;
        string content;
        string date;
        string place;
    }

    mapping(uint8 => Review) public reviews;

    event ReviewCreated(
        uint8 id,
        string authorName,
        string authorAddress,
        string content,
        string date,
        string place
    );

    function createReview(
        string memory _authorName,
        string memory _authorAddress,
        string memory _content,
        string memory _date,
        string memory _place
    ) public {
        reviewCount++;
        reviews[reviewCount] = Review(
            reviewCount,
            _authorName,
            _authorAddress,
            _content,
            _date,
            _place
        );
        emit ReviewCreated(
            reviewCount,
            _authorName,
            _authorAddress,
            _content,
            _date,
            _place
        );
    }

    function getReviewsById(uint8 _id) public view returns (Review memory) {
        Review storage rev = reviews[_id];
        return rev;
    }
}
