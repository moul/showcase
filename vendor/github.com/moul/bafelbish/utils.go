package bafelbish

// Imported from: https://github.com/dbohdan/remarshal/blob/10e552b63d3f10231269c8151ee86ddf9f8ff317/remarshal.go#L39
// convertMapsToStringMaps recursively converts values of type
// map[interface{}]interface{} contained in item to map[string]interface{}. This
// is needed before the encoders for TOML and JSON can accept data returned by
// the YAML decoder.
func convertMapsToStringMaps(item interface{}) (res interface{}, err error) {
	switch item.(type) {
	case map[interface{}]interface{}:
		res := make(map[string]interface{})
		for k, v := range item.(map[interface{}]interface{}) {
			res[k.(string)], err = convertMapsToStringMaps(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	case []interface{}:
		res := make([]interface{}, len(item.([]interface{})))
		for i, v := range item.([]interface{}) {
			res[i], err = convertMapsToStringMaps(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	default:
		return item, nil
	}
}
