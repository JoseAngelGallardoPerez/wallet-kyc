package serializers

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/Confialink/wallet-pkg-model_serializer"
	"context"
)

type request struct {
	fields []interface{}
}

func (t request) Serialize(ctx context.Context, model interface{}) map[string]interface{} {
	return model_serializer.Serialize(model, t.fields)
}

func (t request) SerializeList(ctx context.Context, models []model.UserRequest) []map[string]interface{} {
	serialized := make([]map[string]interface{}, len(models))

	for i, v := range models {
		serialized[i] = model_serializer.Serialize(&v, t.fields)
	}
	return serialized
}

var Request = request{
	fields: []interface{}{
		"ID",
		"UserId",
		"Status",
		"UpdatedAt",
		map[string][]interface{}{
			"Tier": {"ID", "Level", "Name"},
		},
		map[string][]interface{}{
			"User": {"FirstName", "LastName", "Email", "Role"},
		},
	},
}
