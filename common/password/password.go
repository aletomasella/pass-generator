package password

import "strconv"

func TryConvertToInteger(value string) (int, error) {

	convertedValue, err := strconv.Atoi(value)

	if err != nil {
		return 0, err
	}

	return convertedValue, nil

}
