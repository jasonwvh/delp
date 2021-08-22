.DEFAULT_GOAL := compile

compile:
	solc --optimize --abi ./contracts/Reviews.sol -o build --overwrite
	solc --optimize --bin ./contracts/Reviews.sol -o build --overwrite
	abigen --abi=./build/Reviews.abi --bin=./build/Reviews.bin --pkg=reviews --out=server/reviews/Reviews.go