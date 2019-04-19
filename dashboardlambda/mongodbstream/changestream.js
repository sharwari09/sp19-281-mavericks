const MongoClient = require('mongodb').MongoClient

const parameters = {
    mongos_query_router : "10.0.0.181:27017",
    db:"test",
    collection:"bios",
}

const pipeline = [{ $project: { documentKey: false }}];

MongoClient.connect(`mongodb://${parameters.mongos_query_router}/`).then((client)=>{
console.log("Connected correctly to server");


// specify db and collections

const db = client.db(`${parameters.db}`);
const collection = db.collection(`${parameters.db}`);
console.log(MongoClient);

const changeStream = collection.watch(pipeline);

// start listen to changes

changeStream.on("change",(change)=>{
    
    
    console.log( JSON.stringify(change)); 

})

})
