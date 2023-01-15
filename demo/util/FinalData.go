package util

var LoginExpires int64 = 3 * 24 * 60 * 60 * 1000

/**
 * 协议魔数
 */
var MagicNumber string = "20011014"

/**
 * 终止符
 */
var END string = "!end!"

/**
 * jwt token 签发方
 */
var TokenIssuer string = "TestApp-LoginServer"
