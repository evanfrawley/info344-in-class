'use strict';

const mongodb = require('mongodb');

class MongoStore {
    constructor(db, colName) {
        this.collection = db.collection(colName);
    }

    insert(task) {
        task._id = new mongodb.ObjectID();
        return this.collection.insertOne(task).then(() => task);
    }

    update(id, updates) {
        let updateDoc = {"$set": updates};
        return this.collection.findOneAndUpdate(
            {_id: id},
            updateDoc,
            {returnOriginal: false}
        ).then((result) => result.value);
    }

    getByID(id) {
        return this.collection.findOne({_id: id});
    }

    deleteByID(id) {
        return this.collection.deleteOne({_id: id});
    }

    getAll(completed) {
        return this.collection.find({completed: completed})
            .limit(1000)
            .toArray();
    }
}

module.exports = MongoStore;
