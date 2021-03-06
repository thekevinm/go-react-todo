package models

// Bring in primitive from mongo-driver
import "go.mongodb.org/mongo-driver/bson/primitive"

// Use `struct` to specify how the data will be stored in mongo
type ToDoList struct {
  // There are 3 fields, ID, Task, and Status
  ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` // objectID will be generated by mongo
  Task   string             `json:"task,omitempty"` // the task itself
  Status bool               `json:"status,omitempty"` // the status of the task can be `true` or `false`	

}
