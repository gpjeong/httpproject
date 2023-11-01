package rest

import (
	"bytes"
	"httpproject/util/logger"
)

func RequetApiMethod(apiUri string, method string) ([]byte, int, error) {

	resp, statusCode, err := ExecuteService(apiUri, method, &bytes.Buffer{})
	if err != nil && statusCode != 200 {
		logger.Log.Error().Msgf("Rest Execute Service Error: %s", err.Error())
		return nil, statusCode, err
	}
	if statusCode != 200 {
		return nil, statusCode, nil
	}

	return resp, statusCode, nil
}
