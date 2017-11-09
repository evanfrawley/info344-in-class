"use strict";

const express = require("express");
const app = express();

const addr = process.env.ADDR || "localhost:4000";

const [host, port] = addr.split(":");

const MongoStore = require('./taskstore');
const mongodb = require('mongodb');

const mongoAddr = process.env.DBADDR || "localhost:27017";
const mongoURL = `mongodb://${mongoAddr}/tasks`;

mongodb.MongoClient.connect(mongoURL).then(db => {
    let taskStore = new MongoStore(db, "tasks");

    app.use(express.json());

    const tasksBase = "/v1/tasks";

    app.post("/v1/tasks", (req, res) => {
        let task = {
            title: req.body.title,
            tags: req.body.tags,
        };
        // res.header("Content-Type", "application/json");
        taskStore.insert(task)
            .then(newTask => res.json(newTask))
            .catch(err => {
                throw err;
            })
    });

    app.patch("/v1/tasks/:taskID", (req, res) => {
        let taskID = req.params.taskID;
        let updates = req.body;
        taskStore.update(taskID, updates)
            .then(response => {
                res.json(response);
            })
            .catch(err => {
                throw err;
            })
    });

    app.get("/v1/tasks/", (req, res) => {
        taskStore.getAll(true)
            .then(response => {
                res.json(response)
            })
            .catch(err => {
                throw err;
            })
    });

    app.delete("/v1/tasks/:taskID", (req, res) => {
        let taskID = req.params.taskID;
        taskStore.deleteByID(taskID)
            .then(response => {
                res.json(response);
            })
            .catch(err => {
                throw err;
            })
    });

    app.listen(port, host, () => {
        console.log(`server is listening at http://${addr}...`);
    })
})
.catch(err => {
    throw err;
});

