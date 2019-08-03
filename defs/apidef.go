/*
* @Author: sottxiong
* @Date:   2019-08-03 01:28:55
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-08-03 01:31:00
 */
package defs

//requests

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}
