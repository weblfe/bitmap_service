package types

type (
	Code int
)

const (
	Ok                           = "Ok"
	UndefinedCode                = "Unknown"
	CodeSuccess             Code = 200
	CodeServerError         Code = 500
	CodeServiceUnKnown      Code = 404
	CodeServiceUnauthorized Code = 403
	CodeNotData             Code = 400
	CodeParamIllegality     Code = 1000
)

var (
	dataMap = map[Code]string{
		CodeSuccess:             Ok,
		CodeServerError:         "server error",
		CodeServiceUnKnown:      "service unknown",
		CodeServiceUnauthorized: "unauthorized",
		CodeNotData:             "not data",
		CodeParamIllegality:     "param Illegality",
	}
)

func (c Code) GetMsg() string {
	if v, ok := dataMap[c]; ok {
		return v
	}
	return UndefinedCode
}

func (c Code) GetCode() int {
	return int(c)
}

func (c Code) GetCode32() int32 {
	return int32(c)
}