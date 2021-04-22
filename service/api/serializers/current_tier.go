package serializers

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/Confialink/wallet-pkg-model_serializer"
	"context"
)

type currentTier struct {
	fields []interface{}
}

func (t currentTier) Serialize(ctx context.Context, model interface{}) map[string]interface{} {
	return model_serializer.Serialize(model, t.fields)
}

func (t currentTier) SerializeList(ctx context.Context, models []model.Tier) []map[string]interface{} {
	serialized := make([]map[string]interface{}, len(models))

	for i, v := range models {
		serialized[i] = model_serializer.Serialize(&v, t.fields)
	}
	return serialized
}

var CurrentTier = currentTier{
	fields: []interface{}{
		"ID",
		"Name",
		"Level",
		map[string][]interface{}{
			"Limitations": {
				"Value", "Name",
			},
		},
	},
}
