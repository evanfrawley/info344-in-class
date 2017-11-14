#!/usr/bin/env node
"use strict";

const amqp = require('amqplib');

const qName = "testQ";
const mqAddr = process.env.MQADDR || "localhost:5672";
const mqURL = `amqp://${mqAddr}`;

(async () => {
    console.log("connecting to %s", mqURL);

    let connection = await amqp.connect(mqURL);
    let channel = await connection.createChannel();
    let qConf = await channel.assertQueue(qName, {durable: false});

    setInterval(() => {
        let msg = {
            user: "evan frawley",
            password: "memes",
        };
        channel.sendToQueue(qName, Buffer.from(JSON.stringify(msg)));
    }, 1000)
})();
