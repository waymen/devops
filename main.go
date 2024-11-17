package main

import (
	"fmt"

	"github.com/waymen/devops/base"
	"github.com/waymen/devops/base/common"
	"github.com/waymen/devops/utils"
)

func main() {

	qq, err := base.RequestQQ()
	fmt.Println(qq, err)
	baidu, err := utils.RequestBaidu()
	fmt.Println(baidu, err)
	taobao, err := common.RequestTaobao()
	fmt.Println(taobao, err)

}
