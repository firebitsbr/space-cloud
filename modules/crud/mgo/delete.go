package mgo

import (
	"context"
	"errors"

	"github.com/spaceuptech/space-cloud/model"
	"github.com/spaceuptech/space-cloud/utils"
)

// Delete removes the document(s) from the database which match the condition
func (m *Mongo) Delete(ctx context.Context, project, col string, req *model.DeleteRequest) error {
	collection := m.client.Database(project).Collection(col)

	switch req.Operation {
	case utils.One:
		_, err := collection.DeleteOne(ctx, req.Find)
		if err != nil {
			return err
		}

	case utils.All:
		_, err := collection.DeleteMany(ctx, req.Find)
		if err != nil {
			return err
		}

	default:
		return errors.New("Invalid operation")
	}

	return nil
}
