package util

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog/log"
)

func JsonStructMapping(targetData interface{}, mappingStruct interface{}) {
	dataByte, marshalErr := json.Marshal(targetData)
	if marshalErr != nil {
		log.Error().Msgf("Failed to convert json data: %s", marshalErr.Error())
	}

	umErr := json.Unmarshal(dataByte, &mappingStruct)
	if umErr != nil {
		log.Error().Msgf("Failed to parse Target Struct: %s", umErr.Error())
	}
}

func JSONMarshalBuffer(t interface{}) (*bytes.Buffer, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(t)
	return buffer, err
}
