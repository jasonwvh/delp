const express = require("express");
const cors = require("cors");
const redis = require("redis");
const app = express();
const port = 3000;

app.use(cors());

const client = redis.createClient();
client.on("error", function (err) {
    console.log("Error:", err);
});

app.get("/:country/cities", (req, res) => {
    res.setHeader("Content-Type", "application/json");

    client.lrange([req.params.country, 0, -1], function(err, cities) {
        console.log(req.params.country, cities);
        res.send(cities)
    })
})

app.get("/countries", (req, res) => {
    res.setHeader("Content-Type", "application/json");

    client.hgetall("all-countries", function(err, cities) {
        const unique = [...new Set(Object.keys(cities))]
        res.send(unique)
    })
})

app.get("/all", (req, res) => {
    res.setHeader("Content-Type", "application/json");

    client.hgetall("all-cities", function (err, cities) {
        let tmp = [];
        Object.keys(cities).forEach(function (err, city) {
            tmpObj = {
                city: city,
                country: cities[city],
            };
            tmp.push(tmpObj);
        });
        res.send(tmp);
    });
});

app.get("/hello", (req, res) => {
    res.setHeader("Content-Type", "application/json");
    const msg = { message: "Hello World!" };
    res.send(msg);
});

app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);
});
