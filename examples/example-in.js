const request = require("snekfetch");

request.get("https://weebs.cf/random/hug").then(body => {
    console.log(body.text)
});