// +build !js

package tendon

import "errors"

func GetDeviceSize() (float64, float64, error) {
	return 0, 0, errors.New("not browser")
}
