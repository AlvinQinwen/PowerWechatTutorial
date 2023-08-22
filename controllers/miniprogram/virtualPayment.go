package miniprogram

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/virtualPayment/request"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

// 小程序虚拟支付

func APIStartUploadProducts(ctx *gin.Context) {

	var items []*request.GoodItem

	items = append(items, &request.GoodItem{
		Id:      "18",
		Name:    "1000K币",
		Price:   1000,
		Remake:  "1000K币",
		ItemUrl: "https://qiniu.rongjuwh.cn/applet_kb_20230801101836.png",
	})

	params := &request.UploadProductsRequest{
		Env:        0,
		UploadItem: items,
	}

	_, err := services.MiniProgramApp.VirtualPayment.StartUploadGoods(ctx, params)

	if err != nil {

		panic(err)
	}

	//fmt.Dump(result)

}
