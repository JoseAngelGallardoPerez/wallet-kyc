package serializers

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/Confialink/wallet-pkg-model_serializer"
	"context"
)

type adminTiers struct {
	fields    []interface{}
	fieldsOne []interface{}
}

func (t adminTiers) Serialize(ctx context.Context, model *model.Tier) map[string]interface{} {
	return model_serializer.Serialize(model, t.fieldsOne)
}

func (t adminTiers) SerializeList(ctx context.Context, models []model.Tier) []map[string]interface{} {
	serialized := make([]map[string]interface{}, len(models))

	for i, v := range models {
		serialized[i] = model_serializer.Serialize(&v, t.fields)
	}
	return serialized
}

var AdminTiers = adminTiers{
	fields: []interface{}{
		"ID",
		"CountryCode",
		"Level",
		"Name",
	},
	fieldsOne: []interface{}{
		"ID",
		"CountryCode",
		"Level",
		"Name",
		map[string][]interface{}{
			"Limitations": {
				"Value", "Name", "Index",
			},
		},
	},
}
