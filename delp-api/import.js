const redis = require("redis");
const client = redis.createClient();

client.on("error", function (err) {
    console.log("Error:", err);
});

const csv = require("csv-parser");
const fs = require("fs");

fs.createReadStream("data/worldcities_clean.csv")
    .pipe(csv())
    .on("data", (data) => {
        client.hmset([data.city, "country", data.country], function (err, res) {
            if (err) {
                console.error("error");
            } else {
                console.log("city:", data.city, "res:", res);
            }
        });
        client.lpush([data.country, data.city], function (err, res) {
            if (err) {
                console.error("error");
            } else {
                console.log("city:", data.country, "res:", res);
            }
        });
        client.hmset(
            ["all-cities", data.city, data.country],
            function (err, res) {
                if (err) {
                    console.error("error");
                } else {
                    console.log("city:", data.city, "res:", res);
                }
            }
        );
    })
    .on("end", () => {
        console.log("uploaded all data");
    });
/*
client.lrange(["France", 0, -1], function(err, value) {
    if (err) {
        console.error("error");
    } else {
        console.log(value);
    }
});

client.hmget("Paris", "country", function(err, value) {
    if (err) {
        console.error("error");
    } else {
        console.log(value[0]);
    }
});

client.keys("*", function(err, value) {
    if (err) {
        console.error("error");
    } else {
        console.log(value);
    }
});*/
