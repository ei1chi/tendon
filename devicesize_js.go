// +build js

package tendon

import "github.com/gopherjs/gopherjs/js"

func GetDeviceSize() (float64, float64, error) {
	body := js.Global.Get("document").Get("body")
	w := body.Get("clientWidth").Float()
	h := body.Get("clientHeight").Float()
	return w, h, nil
}
