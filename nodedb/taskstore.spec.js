'use strict';

const MongoStore = require('./taskstore');
const mongodb = require('mongodb');

const mongoAddr = process.env.DBADDR || "localhost:27017";
const mongoURL = `mongodb://${mongoAddr}/tasks`;

describe("Mongo Task Store", () => {
    test("CRUD cycle", () => {
        return mongodb.MongoClient.connect(mongoURL)
            .then((db) => {
                let store = new MongoStore(db, "tasks");
                let task = {
                    title: "learn nodejs to mongodb",
                    tags: [
                        "node",
                        "mongo",
                        "memes"
                    ]
                };

                return store.insert(task)
                    .then(task => {
                        expect(task._id).toBeDefined();
                        return task._id;
                    })
                    .then(taskId => {
                        return store.getByID(taskId)
                    })
                    .then(fetchedTask => {
                        expect(fetchedTask).toEqual(task);
                        return store.update(task._id, {completed: true})
                    })
                    .then(updatedTask => {
                        expect(updatedTask.completed).toBe(true);
                        return store.deleteByID(task._id)
                    })
                    .then(() => {
                        return store.getByID(task._id)
                    })
                    .then(fetchedTask => {
                        expect(fetchedTask).toBeFalsy();
                    })
                    .then(() => {
                        db.close();
                    })
                    .catch(err => {
                        db.close();
                        throw err;
                    })
            });
    });
});
