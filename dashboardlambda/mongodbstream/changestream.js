const MongoClient = require('mongodb').MongoClient
var mongos_query_router="10.0.0.181"
const pipeline = [{ $project: { documentKey: false }}];

MongoClient.connect(`mongodb://${mongos_query_router}:27017/`).then((client)=>{
console.log("Connected correctly to server");

// specify db and collections

const db = client.db("test");
const collection = db.collection("bios");
console.log(MongoClient);

const changeStream = collection.watch(pipeline);

// start listen to changes

changeStream.on("change",(change)=>{
    console.log( JSON.stringify(change)); 
})

})
