package services

import (
  "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
  "github.com/ArtisanCloud/PowerWeChat/src/payment"
  "power-wechat-tutorial/config"
)

const TRANSACTION_SUCCESS = "TRANSACTION.SUCCESS"
const TRANSACTION_FAILED = "TRANSACTION.FAILED"

var PaymentApp *payment.Payment

func NewWXPaymentApp(conf *config.Configuration) (*payment.Payment, error) {

  Payment, err := payment.NewPayment(&payment.UserConfig{
    AppID:       conf.Payment.AppID,       // 小程序、公众号或者企业微信的appid
    MchID:       conf.Payment.MchID,       // 商户号 appID
    MchApiV3Key: conf.Payment.MchApiV3Key, //
    Key:         conf.Payment.Key,
    CertPath:    conf.Payment.CertPath,
    KeyPath:     conf.Payment.KeyPath,
    SerialNo:    conf.Payment.SerialNo,
    NotifyURL:   conf.Payment.NotifyURL,

    ResponseType: response.TYPE_MAP,
    Log: payment.Log{
      Level: "debug",
      File:  "./wechat.log",
    },
    Http: payment.Http{
      Timeout: 30.0,
      BaseURI: "https://api.mch.weixin.qq.com",
    },

    HttpDebug: true,
  })

  return Payment, err
}
