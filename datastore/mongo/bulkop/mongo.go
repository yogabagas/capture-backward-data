package bulkop

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoBulkOperation interface {
	AddOperation(collectionName string, operation mongo.WriteModel)
	GetOperationList() map[string][]mongo.WriteModel
}

type CommonMongoBulkOperation struct {
	collection map[string][]mongo.WriteModel
}

func NewCommonMongoBulkOperation() MongoBulkOperation {
	bulkOp := &CommonMongoBulkOperation{
		collection: make(map[string][]mongo.WriteModel, 0),
	}

	return bulkOp
}

func (b *CommonMongoBulkOperation) AddOperation(collectionName string, operation mongo.WriteModel) {
	value, ok := b.collection[collectionName]

	if ok == true {
		value = append(value, operation)
		b.collection[collectionName] = value

		return
	}

	b.collection[collectionName] = make([]mongo.WriteModel, 0)
	value = b.collection[collectionName]
	value = append(value, operation)
	b.collection[collectionName] = value
}

func (b *CommonMongoBulkOperation) GetOperationList() map[string][]mongo.WriteModel {
	return b.collection
}
