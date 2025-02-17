package routes

import (
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/controllers/official-account"
)

func InitOfficialAPIRoutes(r *gin.Engine) {

	officialRouter := r.Group("/official")
	{
		// Base
		officialRouter.GET("/base/clearQuota", official_account.ClearQuota)
		officialRouter.GET("/base/callbackIP", official_account.GetCallbackIP)

		// jssdk
		officialRouter.GET("/jssdk/config", official_account.JSSDKBuildConfig)

		// 临时素材管理
		officialRouter.POST("/media/uploadImage", official_account.APIUploadImage)
		officialRouter.POST("/media/uploadVoice", official_account.APIUploadVoice)
		officialRouter.POST("/media/uploadVideo", official_account.APIUploadVideo)
		officialRouter.POST("/media/uploadThumb", official_account.APIUploadThumb)
		officialRouter.GET("/media/get", official_account.APIGetMedia)

		// 永久素材管理
		officialRouter.POST("/material/uploadImage", official_account.APIUploadMaterialImage)
		officialRouter.POST("/material/uploadArticleImage", official_account.APIUploadArticleImage)
		officialRouter.POST("/material/uploadVoice", official_account.APIUploadMaterialVoice)
		officialRouter.POST("/material/uploadVideo", official_account.APIUploadMaterialVideo)
		officialRouter.POST("/material/uploadThumb", official_account.APIUploadMaterialThumb)
		officialRouter.GET("/material/list", official_account.APIGetMaterialList)
		officialRouter.GET("/material/get", official_account.APIGetMaterial)

		// 用户管理
		officialRouter.GET("/user/get", official_account.GetUserInfo)
		officialRouter.GET("/user/batchGet", official_account.GetBatchUserInfo)
		officialRouter.POST("/user/remark", official_account.UserRemark)
		officialRouter.GET("/user/getBlockList", official_account.GetUserBlacklist)
		officialRouter.POST("/user/block", official_account.UserBlock)
		officialRouter.POST("/user/unBlock", official_account.UserUnBlock)
		officialRouter.POST("/user/changeOpenID", official_account.UserChangeOpenID)

		// 用户标签管理
		officialRouter.GET("/userTag/list", official_account.GetUserTagList)
		officialRouter.POST("/userTag", official_account.UserTagCreate)
		officialRouter.PUT("/userTag", official_account.UserTagUpdate)
		officialRouter.DELETE("/userTag", official_account.UserTagDelete)
		officialRouter.GET("/userTag/getUserTagsByOpenID", official_account.GetUserTagsByOpenID)
		officialRouter.GET("/userTag/getUsersOfTag", official_account.GetUsersOfTag)
		officialRouter.POST("/userTag/batchTagUser", official_account.UserTagBatchTagUsers)
		officialRouter.POST("/userTag/unBatchTagUser", official_account.UserTagBatchUnTagUsers)

		// 客服消息
		officialRouter.GET("/customerService/list", official_account.GetCustomerList)
		officialRouter.GET("/customerService/online", official_account.GetCustomerOnline)
		officialRouter.POST("/customerService/create", official_account.CustomerCreate)
		officialRouter.PUT("/customerService/update", official_account.CustomerUpdate)
		officialRouter.DELETE("/customerService/delete", official_account.CustomerDelete)
		officialRouter.POST("/customerService/setAvatar", official_account.CustomerSetAvatar)
		officialRouter.POST("/customerService/messages", official_account.CustomerMessages)
		officialRouter.POST("/customerService/invite", official_account.CustomerInvite)
		officialRouter.POST("/customerService/sendText", official_account.CustomerMessageSendText)
		officialRouter.POST("/customerService/sendImage", official_account.CustomerMessageSendImage)
		officialRouter.POST("/customerService/sendVideo", official_account.CustomerMessageSendVideo)
		officialRouter.POST("/customerService/sendVoice", official_account.CustomerMessageSendVoice)
		officialRouter.POST("/customerService/sendLink", official_account.CustomerMessageSendLink)
		officialRouter.POST("/customerService/sendMusic", official_account.CustomerMessageSendMusic)
		officialRouter.POST("/customerService/sendNews", official_account.CustomerMessageSendNews)
		officialRouter.POST("/customerService/sendMsgMenu", official_account.CustomerMessageSendMsgMenu)
		officialRouter.POST("/customerService/sendRaw", official_account.CustomerMessageSendRaw)
		officialRouter.POST("/customerService/send", official_account.CustomerMessageSend)

		// 客服会话控制
		officialRouter.POST("/customerServiceSession/create", official_account.CustomerSessionCreate)
		officialRouter.DELETE("/customerServiceSession/close", official_account.CustomerSessionClose)
		officialRouter.GET("/customerServiceSession/get", official_account.GetCustomerSession)
		officialRouter.GET("/customerServiceSession/list", official_account.CustomerSessionList)
		officialRouter.GET("/customerServiceSession/waiting", official_account.CustomerSessionWaiting)

		// 草稿发布
		officialRouter.GET("/publish/draftAdd", official_account.APIDraftAdd)
		officialRouter.GET("/publish/draftGet", official_account.APIDraftGet)
		officialRouter.GET("/publish/draftDelete", official_account.APIDraftDelete)
		officialRouter.GET("/publish/draftUpdate", official_account.APIDraftUpdate)
		officialRouter.GET("/publish/draftCount", official_account.APIDraftCount)
		officialRouter.GET("/publish/draftBatchGet", official_account.APIDraftBatchGet)
		officialRouter.GET("/publish/draftSwitch", official_account.APIDraftSwitch)
		officialRouter.GET("/publish/draftCheckSwitch", official_account.APIDraftCheckSwitch)
		officialRouter.GET("/publish/publishSubmit", official_account.APIPublishSubmit)
		officialRouter.GET("/publish/publishGet", official_account.APIPublishGet)
		officialRouter.GET("/publish/publishDelete", official_account.APIPublishDelete)
		officialRouter.GET("/publish/publishGetArticle", official_account.APIPublishGetArticle)
		officialRouter.GET("/publish/publishBatchGet", official_account.APIPublishBatchGet)

		// 数据统计
		officialRouter.GET("/dateCube/getUserSummary", official_account.GetUserSummary)
		officialRouter.GET("/dateCube/getUserCumulate", official_account.GetUserCumulate)
		officialRouter.GET("/dateCube/articleSummary", official_account.ArticleSummary)
		officialRouter.GET("/dateCube/articleTotal", official_account.ArticleTotal)
		officialRouter.GET("/dateCube/userReadSummary", official_account.UserReadSummary)
		officialRouter.GET("/dateCube/userShareSummary", official_account.UserShareSummary)
		officialRouter.GET("/dateCube/userShareHourly", official_account.UserShareHourly)
		officialRouter.GET("/dateCube/upstreamMessageSummary", official_account.UpstreamMessageSummary)
		officialRouter.GET("/dateCube/upstreamMessageHourly", official_account.UpstreamMessageHourly)
		officialRouter.GET("/dateCube/upstreamMessageWeekly", official_account.UpstreamMessageWeekly)
		officialRouter.GET("/dateCube/upstreamMessageMonthly", official_account.UpstreamMessageMonthly)
		officialRouter.GET("/dateCube/upstreamMessageDistSummary", official_account.UpstreamMessageDistSummary)
		officialRouter.GET("/dateCube/upstreamMessageDistWeekly", official_account.UpstreamMessageDistWeekly)
		officialRouter.GET("/dateCube/upstreamMessageDistMonthly", official_account.UpstreamMessageDistMonthly)

		// 生成二维码
		officialRouter.GET("/qrcode/temp", official_account.GetTempQrCode)
		officialRouter.GET("/qrcode/forever", official_account.GetForeverQrCode)

		// 短key托管
		officialRouter.GET("/shorten/gen", official_account.ShortGenKey)
		officialRouter.GET("/shorten/fetch", official_account.FetchShortGen)

		// 自动回复
		officialRouter.GET("/autoReply/current", official_account.AutoReplyCurrent)

		// OAUTH2
		officialRouter.GET("/oauth/getAuthCode", official_account.GetAuthCode)
		officialRouter.GET("/oauth/userFromCode", official_account.UserFromCode)
		officialRouter.GET("/oauth/userFromToken", official_account.UserFromToken)

		// 菜单
		officialRouter.GET("/menu/list", official_account.MenuList)
		officialRouter.GET("/menu/current", official_account.MenuCurrent)
		officialRouter.POST("/menu/create", official_account.MenuCreate)
		officialRouter.POST("/menu/createConditional", official_account.MenuCreateConditional)
		officialRouter.DELETE("/menu/delete", official_account.MenuDelete)
		officialRouter.DELETE("/menu/deleteConditional", official_account.MenuDeleteConditional)
		officialRouter.DELETE("/menu/match", official_account.MenuMatch)

		// 消息群发
		officialRouter.POST("/broadcasting/text", official_account.BroadcastSendText)
		officialRouter.POST("/broadcasting/image", official_account.BroadcastSendImage)
		officialRouter.POST("/broadcasting/news", official_account.BroadcastSendNews)
		officialRouter.POST("/broadcasting/voice", official_account.BroadcastSendVoice)
		officialRouter.POST("/broadcasting/video", official_account.BroadcastSendVideo)
		officialRouter.POST("/broadcasting/card", official_account.BroadcastSendCard)
		officialRouter.POST("/broadcasting/preview", official_account.BroadcastSendPreview)
		officialRouter.DELETE("/broadcasting/delete", official_account.BroadcastDelete)
		officialRouter.GET("/broadcasting/status", official_account.BroadcastStatus)

		// 群发评论
		officialRouter.POST("/comment/open", official_account.CommentOpen)
		officialRouter.POST("/comment/close", official_account.CommentClose)
		officialRouter.DELETE("/comment/delete", official_account.CommentDelete)
		officialRouter.GET("/comment/list", official_account.CommentList)
		officialRouter.POST("/comment/markElect", official_account.CommentMarkElect)
		officialRouter.DELETE("/comment/unMarkElect", official_account.CommentUnMarkElect)
		officialRouter.POST("/comment/reply", official_account.CommentReply)
		officialRouter.DELETE("/comment/reply", official_account.CommentDeleteReply)

		// 返佣商品
		officialRouter.GET("/goods/list", official_account.GoodsList)
		officialRouter.GET("/goods/get", official_account.GoodsGet)
		officialRouter.GET("/goods/status", official_account.GoodsStatus)
		officialRouter.POST("/goods/add", official_account.GoodsAdd)
		officialRouter.PUT("/goods/update", official_account.GoodsUpdate)

		// 消息回调
		officialRouter.GET("/callback/message", official_account.CallbackVerify)
		officialRouter.POST("/callback/message", official_account.CallbackNotify)

		// 统一服务消息
		officialRouter.POST("/uniformMessage/send", official_account.UniformMessageSend)

		// 模板消息
		officialRouter.GET("/templateMessage/getIndustry", official_account.TemplateMessageGetIndustry)
		officialRouter.GET("/templateMessage/getPrivateTemplates", official_account.GetPrivateTemplates)
		officialRouter.POST("/templateMessage/send", official_account.TemplateMessageSend)
		officialRouter.POST("/templateMessage/sendSubscribe", official_account.SendSubscribe)
		officialRouter.POST("/templateMessage/setIndustry", official_account.TemplateMessageSetIndustry)
		officialRouter.POST("/templateMessage/addTemplate", official_account.TemplateMessageAddTemplate)
		officialRouter.DELETE("/templateMessage/delPrivateTemplate", official_account.DelPrivateTemplate)

	}
}
