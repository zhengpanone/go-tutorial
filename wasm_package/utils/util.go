package utils

import "encoding/json"

type Result struct {
	Ok   bool
	Err  string
	Data any
}

func PackageResult(ok bool, err error, data any) string {
	_err := ""
	if err != nil {
		_err = err.Error()
	}
	res := Result{
		Ok:   ok,
		Err:  _err,
		Data: data,
	}
	resByte, _ := json.Marshal(res)
	return string(resByte)

}
