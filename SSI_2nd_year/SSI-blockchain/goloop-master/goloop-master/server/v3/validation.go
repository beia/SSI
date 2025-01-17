package v3

import (
	"fmt"
	"regexp"

	"gopkg.in/go-playground/validator.v9"

	"github.com/icon-project/goloop/server/jsonrpc"
	"github.com/icon-project/goloop/service/contract"
)

var (
	hexString          = regexp.MustCompile("^0x[0-9a-f]+$")
	deployContentTypes = []string{"application/zip", "application/java"}
)

func RegisterValidationRule(v *jsonrpc.Validator) {

	v.RegisterValidation("call", isCall)
	v.RegisterValidation("deploy", isDeploy)
	v.RegisterValidation("message", isMessage)
	v.RegisterValidation("deposit", isDeposit)

	// validate : CallParam.Data, TransactionParam.Data
	v.RegisterStructValidation(DataParamValidation, CallParam{}, TransactionParam{})

}

func isCall(fl validator.FieldLevel) bool {
	return fl.Field().String() == contract.DataTypeCall
}

func isDeploy(fl validator.FieldLevel) bool {
	return fl.Field().String() == contract.DataTypeDeploy
}

func isMessage(fl validator.FieldLevel) bool {
	return fl.Field().String() == contract.DataTypeMessage
}

func isDeposit(fl validator.FieldLevel) bool {
	return fl.Field().String() == contract.DataTypeDeposit
}

func DataParamValidation(sl validator.StructLevel) {
	switch sl.Current().Interface().(type) {
	case CallParam:
		callParam := sl.Current().Interface().(CallParam)
		if data, ok := callParam.Data.(map[string]interface{}); ok {
			validateCallDataParam(sl, callParam.Data, data)
		} else {
			sl.ReportError(callParam.Data, "Data", "", "data", "")
		}
	case TransactionParam:
		txParam := sl.Current().Interface().(TransactionParam)
		if txParam.DataType != "" {
			switch txParam.DataType {
			case contract.DataTypeCall:
				if data, ok := txParam.Data.(map[string]interface{}); ok {
					validateCallDataParam(sl, txParam.Data, data)
				} else {
					sl.ReportError(txParam.Data, "Data", "", "data", "")
				}
			case contract.DataTypeDeploy:
				if data, ok := txParam.Data.(map[string]interface{}); ok {
					validateDeployDataParam(sl, txParam.Data, data)
				} else {
					sl.ReportError(txParam.Data, "Data", "", "data", "")
				}
			case contract.DataTypeMessage:
				if data, ok := txParam.Data.(string); ok {
					if !hexString.MatchString(data) {
						sl.ReportError(txParam.Data, "Data", "", "data", "")
					}
				} else {
					sl.ReportError(txParam.Data, "Data", "", "data", "")
				}
			case contract.DataTypeDeposit:
				if data, ok := txParam.Data.(map[string]interface{}); ok {
					validateDepositDataParam(sl, txParam.Data, data)
				} else {
					sl.ReportError(txParam.Data, "Data", "", "data", "")
				}
			}
		}
	}
}

func validateRPCData(sl validator.StructLevel, name string, value interface{}) {
	switch obj := value.(type) {
	case string:
		// param value : string
	case map[string]interface{}:
		for k, v := range obj {
			validateRPCData(sl, fmt.Sprintf("%s.%s", name, k), v)
		}
	case []interface{}:
		for i, v := range obj {
			validateRPCData(sl, fmt.Sprintf("%s[%d]", name, i), v)
		}
	default:
		sl.ReportError(value, name, "", "data.params", "")
	}
}

func validateCallDataParam(sl validator.StructLevel, field interface{}, data map[string]interface{}) {
	// data.method : required
	if _, ok := data["method"]; !ok {
		sl.ReportError(field, "Data", "data", "data.method", "")
	}
	// data.params : optional
	if params, ok := data["params"]; ok {
		paramsMap := params.(map[string]interface{})
		for k, pv := range paramsMap {
			validateRPCData(sl, "Data.params."+k, pv)
		}
	}
}

func validateDeployDataParam(sl validator.StructLevel, field interface{}, data map[string]interface{}) {
	// data.contentType : required
	if v, ok := data["contentType"]; ok {
		contains := func(s []string, t interface{}) bool {
			if t, ok = t.(string); !ok {
				return false
			}
			for _, v := range s {
				if v == t {
					return true
				}
			}
			return false
		}
		if !contains(deployContentTypes, v) {
			sl.ReportError(field, "Data", "Data", "data.contentType", "")
		}
	} else {
		sl.ReportError(field, "Data", "Data", "data.contentType", "")
	}
	// data.content : required
	if v, ok := data["content"]; ok {
		if content, ok := v.(string); ok {
			if !(hexString.MatchString(content) && len(content)%2 == 0) {
				sl.ReportError(field, "Data", "", "data.content", "")
			}
		} else {
			sl.ReportError(field, "Data", "", "data.content", "")
		}
	}
	// data.params : optional
	if params, ok := data["params"]; ok {
		paramsMap := params.(map[string]interface{})
		for k, pv := range paramsMap {
			validateRPCData(sl, "Data.params."+k, pv)
		}
	}
}

func validateDepositDataParam(sl validator.StructLevel, field interface{}, data map[string]interface{}) {
	action, ok := data["action"]
	if !ok {
		sl.ReportError(field, "Data", "action", "data.action", "")
		return
	}
	switch action {
	case contract.DepositActionAdd:
		if len(data) > 1 {
			sl.ReportError(field, "Data", "", "data.unknown", "")
		}
	case contract.DepositActionWithdraw:
		id, ok := data["id"]
		if !ok {
			sl.ReportError(field, "Data", "", "data.id", "")
			return
		} else {
			if s, ok := id.(string); ok {
				if !hexString.MatchString(s) {
					sl.ReportError(field, "Data", "id", "data.id", "")
				}
			} else {
				sl.ReportError(field, "Data", "id", "data.id", "")
			}
		}
		return
	default:
		sl.ReportError(field, "Data", "", "data.action", "")
	}
}
