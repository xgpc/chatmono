/**
 * @Author: smono
 * @Description:
 * @File:  render
 * @Version: 1.0.0
 * @Date: 2022/10/9 22:42
 */

package render

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

type ResList struct {
	List  []interface{}
	Total int
}
