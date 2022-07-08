package helpers

import "go.mongodb.org/mongo-driver/mongo/options"

var after = options.After
var ReturnNewObject = options.FindOneAndUpdateOptions{
	ReturnDocument: &after,
}
