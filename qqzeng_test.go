/*
* @Author:  老杨
* @Email:   xcapp1314@gmail.com
* @Date:    2023/8/16 19:35
* @Explain: ...
 */

package qqzeng

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	testIp := "23.106.129.59"
	fmt.Println(IpSplit(testIp))
	fmt.Println(IpInfoString(testIp))
	fmt.Println(IpInfo(testIp))
}
