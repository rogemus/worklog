package task

import "errors"

var ErrorParseInvalidId = errors.New("Parsing Error: ivalid task id. Id need to match `0001` format")
var ErrorParseInvalidCreatedDate = errors.New("Parsing Error: invalid task created date. Need to be in `YYYY-MM-DD HH:mm:ss` format")
