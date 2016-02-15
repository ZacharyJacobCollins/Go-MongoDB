package main

import "github.com/mikejs/gomongo/mongo"


func main() {
        conn, _ := mongo.Connect("127.0.0.1")
        collection := conn.GetDB("test").GetCollection("test_collection")

        doc, _ := mongo.Marshal(map[string]string{
                "_id":     "doc1",
                "title":   "A Mongo document",
                "content": "Testing, 1. 2. 3.",
        })
        collection.Insert(doc)

        query, _ := mongo.Marshal(map[string]string{"_id": "doc1"})
        got, _ := collection.FindOne(query)
        mongo.Equal(doc, got) // true!

        collection.Drop()
}
