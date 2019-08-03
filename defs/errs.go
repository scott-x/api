/*
* @Author: sottxiong
* @Date:   2019-08-03 01:29:15
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-08-03 01:37:37
 */
package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed", ErrorCode: "002"}}
)
