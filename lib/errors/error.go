package errors

import (
	"database/sql"
	"net/http"
	"tinyUrlMock-go/lib/apires"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const MessageOK = "OK"

type (
	CustomMsg string
)

const (
	CODE_UNKNOWN_ERR                       = -1
	CODE_OK                                = 0
	CODE_INVALID_PARAMS                    = 10
	CODE_AUTH_ERR                          = 11
	CODE_AUTH_INVER_NOT_FOUND              = 12
	CODE_AUTH_PROMO_FAIL                   = 13
	CODE_NO_PERM                           = 14
	CODE_DB_ERR                            = 15
	CODE_NOT_EXISTS                        = 16
	CODE_DUPLICATE_KEY                     = 17
	CODE_EMAIL_ERR                         = 18
	CODE_SES_ERR                           = 19
	CODE_BLOCKED                           = 20
	CODE_REGISTER_MAX                      = 21
	CODE_TRANS_TOOL                        = 22
	CODE_GUEST_ERR                         = 23
	CODE_PAYMENT_ERR                       = 24
	CODE_CREATE_TP_ERR                     = 25
	CODE_BONUS_NOT_ENOUGH                  = 26
	CODE_GET_PAYMENT_URL_ERR               = 27
	CODE_GET_TOKEN_FAILED                  = 28
	CODE_GET_CUSTOM_FAILED                 = 29
	CODE_UPLOAD_FAILED                     = 30
	CODE_PHONE_VERIFY_FAILED               = 31
	CODE_PWD_ERR                           = 32
	CODE_BAD_VERSION_NOTICE                = 33
	CODE_BAD_VERSION_FORCE                 = 34
	CODE_REDEEM_MAX_TIMES                  = 35
	CODE_REDEEM_CODE_NOT_FOUND             = 36
	CODE_PHONE_DUPLICATE                   = 37
	CODE_NOTICE_NOT_FOUND                  = 38
	CODE_ES_ERR                            = 39
	CODE_REDEEMED_ALREADY                  = 40
	CODE_MISSION_NOT_COUNTED               = 41
	CODE_SOLD_OUT                          = 42
	CODE_NEWS_NOT_FOUND                    = 43
	CODE_ORDER_NO_TAX                      = 44
	CODE_ORDER_NOT_USED                    = 46
	CODE_TIME_NOT_FOUND                    = 47
	CODE_AGG_NOT_FOUND                     = 48
	CODE_INVOICE_ERR                       = 50
	CODE_IO_ERR                            = 51
	CODE_REDIS_ERR                         = 52
	CODE_DUPLICATE_TOKEN_ERR               = 53
	CODE_GRABPAY_ERR                       = 54
	CODE_RELEASE_ALREADY                   = 55
	CODE_IPAY88_PAYMENT_STATUS_ERR         = 56
	CODE_PAYMENT_TIMEOUT                   = 57
	CODE_REVIEW_TAG_INVALID                = 58
	CODE_REVIEW_RATING_REQUIRED            = 59
	CODE_REVIEW_COMMENT_CHOOSE_ERR         = 60
	CODE_ORDER_REVIEW_DENY                 = 61
	CODE_PAYMENT_AMOUNT_ERR                = 62
	CODE_PARSE_TOKEN_ERR                   = 63
	CODE_TOKEN_IS_EXPIRED_ERR              = 64
	CODE_INCORRECT_MANAGER_ROLE_ERR        = 65
	CODE_TOKEN_USER_NOT_MATCH_ERR          = 66
	CODE_INVALID_HEADER_ERR                = 67
	CODE_VERIFY_CODE_EXPIRED               = 68
	CODE_PWD_STRENGTH_NOT_ENOUGH           = 69
	CODE_PWD_LENGTH_NOT_ENOUGH             = 70
	CODE_NEWPWD_SAME_ERR                   = 71
	CODE_TNG_UNREGISTERED_ERR              = 72
	CODE_LOGIN_DELETING_ERR                = 73
	CODE_REGISTERED_ERR                    = 74
	CODE_WRONG_PRODUCT_SETTING_ERR         = 75
	CODE_LOGIN_ERR                         = 100
	CODE_SESSION_ERR                       = 101
	CODE_DEVICE_NUM_ERR                    = 102
	CODE_PHONE_CCODE_ERR                   = 103
	CODE_PHONE_INVALID                     = 104
	CODE_PHONE_NOT_BIND                    = 105
	CODE_SEND_SMS_ERROR                    = 106
	CODE_ACCOUNT_DELETION_NOT_ALLOW        = 107
	CODE_ACCOUNT_DELETION_APPLY_FAILED     = 108
	CODE_ACCOUNT_DELETION_TOKEN_EXPIRED    = 109
	CODE_PROMO_DUPLICATE                   = 200
	CODE_PROMO_LIMIT_MEMBER_TAG            = 201
	CODE_PROMO_LIMIT_GROUP                 = 202
	CODE_PROMO_LIMIT_CNT                   = 203
	CODE_PROMO_INVITE_INPUT_ALREADY        = 210
	CODE_PROMO_INVITE_NOT_NEW              = 211
	CODE_PROMO_INVITE_TXN_ALREADY          = 212
	CODE_BONUS_ALREADY_SEND                = 213
	CODE_PROMO_NOT_EXISTS                  = 214
	CODE_BONUS_INSUFFICIENT_PAYMENT        = 215
	CODE_BONUS_CANT_USE_NOT_BINDING        = 216
	CODE_PROMO_CREDIT_ERR                  = 230
	CODE_PROMO_ALREADY_RECEIVED            = 231
	CODE_PROMO_NOT_ELIGIBLE_TO_RECEIVE     = 232
	CODE_ORDER_CANCEL_FAILED               = 300
	CODE_ORDER_CANCEL_REDEEMED             = 301
	CODE_ORDER_CANCEL_EXPIRED              = 302
	CODE_ORDER_ALREADY_REDEEMED            = 303
	CODE_PRODUCT_NOT_EXIST                 = 304
	CODE_ORDER_UNRESCHEDULABLE_EXPIRED     = 305
	CODE_ORDER_UNRESCHEDULABLE_PRODUCT     = 306
	CODE_ORDER_UNRESCHEDULABLE_CROSSDAY    = 307
	CODE_ORDER_RESCHEDULED                 = 308
	CODE_PROMO_COMPATIBLE_FAILED           = 309
	CODE_PROMO_INSUFFICIENT                = 310
	CODE_MOBILE_PAY_FAILED                 = 311
	CODE_INVALID_CHECKOUT_CLOCK            = 312
	CODE_INVALID_ORDER_TAX                 = 313
	CODE_INVALID_ORDER_REMARK              = 314
	CODE_INVALID_BOOKINGTIME               = 315
	CODE_INVALID_ID                        = 316
	CODE_BOOKINGTIME_EXPIRED               = 317
	CODE_ORDER_NOT_EXISTS                  = 318
	CODE_INVALID_CUSTNUM                   = 319
	CODE_GROUP_CODE_EOF                    = 320
	CODE_ORDER_INVALID_STATUS              = 321
	CODE_INVALID_GROUP_CODE                = 322
	CODE_ORDERS_GROUP_NOT_EXIST            = 323
	CODE_PRODUCT_TIMEZONE_NOT_FOUND        = 324
	CODE_INVALID_AC01_USER                 = 325
	CODE_GROUP_CODE_DENY                   = 326
	CODE_ORDER_DUPLICATE                   = 327
	CODE_INVALID_REDEEM_CODE               = 328
	CODE_BRANCH_BLOCK_USER                 = 329
	CODE_ORDER_UNRESCHEDULABLE_STATUS      = 330
	CODE_INVALID_PRICE                     = 500
	CODE_INVALID_SELLLIMIT                 = 501
	CODE_INVALID_TIME_RANGE                = 502
	CODE_INVALID_TIME_FORMAT               = 503
	CODE_FLASH_TAG_NOT_EXIST               = 504
	CODE_TIME_PAIR_OVERLAP                 = 505
	CODE_INVALID_PAGINATION_RANGE          = 506
	CODE_INVALID_CLIENT_OS                 = 507
	CODE_INVALID_PLATFORM                  = 508
	CODE_INVALID_APP_VERSION               = 509
	CODE_INVALID_TIME_LIMITATION           = 510
	CODE_INVALID_EZBOOK_UID                = 511
	CODE_REMARK_CUSTOMER_NECESSARY         = 512
	CODE_INVALID_BOSSNOW_USER              = 513
	CODE_INVALID_WEEKDAY                   = 514
	CODE_PRODUCT_SET_PRICE_NOT_FOUND       = 515
	CODE_RESOURCE_GROUP_CHANGED            = 516
	CODE_INVALID_RESOURCE_NAME             = 517
	CODE_INVALID_CURRENCY                  = 518
	CODE_GROUP_CODE_DELIVERY_DENY          = 519
	CODE_INVALID_COURIER_TYPE              = 520
	CODE_INVALID_PRODUCT_STATUS            = 521
	CODE_INVALID_DATA_TYPE                 = 522
	CODE_NOT_ALLOW_FLASH_SALE              = 523
	CODE_INLINE_UNEXPECT_ERROR             = 600
	CODE_INLINE_PHONE_LIMIT                = 601
	CODE_INLINE_CAPACITY_UNAVAILABLE       = 602
	CODE_INLINE_OUT_OF_SIZE                = 603
	CODE_INLINE_BLOCK_NO_SHOW              = 604
	CODE_INLINE_BLOCK_MANUALLY             = 605
	CODE_INLINE_WEB_BOOKING_DISABLE        = 606
	CODE_RATE_LIMITER_ERROR                = 700
	CODE_RATE_LIMITER_TOO_MANY             = 701
	CODE_RATE_SMS_TOO_MANY                 = 702
	CODE_EBISOL_UNEXPECT_ERROR             = 800
	CODE_EBISOL_REMARK_FIELD_SETTING_ERROR = 801
	CODE_EBISOL_CAPACITY_UNAVAILABLE       = 802
	CODE_EBISOL_USER_INFO_INCORRECT        = 803
	CODE_EBISOL_MODIFY_DENY                = 804
	CODE_EBISOL_ORDER_NOT_EXIST            = 805
	CODE_EBISOL_RATE_LIMIT                 = 806
	CODE_APPLE_TOKEN_INVALID               = 900
	CODE_APPLE_TOKEN_EXPIRED               = 901
	CODE_APPLE_TOKEN_ISSUER_INVALID        = 902
	CODE_APPLE_TOKEN_AUDIENCE_INVALID      = 903
	CODE_APPLE_AUTH_INVALID                = 904
	CODE_DATA_NOT_READY                    = 1000
	CODE_STRIPE_EVENT_NOT_SUPPORT          = 1100
	CODE_STRIPE_AUTH_REQUIRED              = 1101
	CODE_STRIPE_PAYMENT_INTENT_INVALID     = 1102
	CODE_STRIPE_CARD_ERROR                 = 1103
	CODE_STRIPE_RELEASE_IN_WEBHOOK         = 1104
	CODE_STRIPE_RELEASE_FAILED_IN_WEBHOOK  = 1105
	CODE_STRIPE_WEBHOOK_SIGNATURE_INVALID  = 1106
	CODE_STRIPE_DIGITAL_WALLET_ERROR       = 1107
	CODE_STRIPE_CREDIT_CARD_NOT_FOUND      = 1108
	CODE_STRIPE_ENV_INVALID                = 1109
	CODE_STRIPE_GET_PAYMENT_INTENT_FAILED  = 1110
	CODE_STRIPE_NEW_PAYMENT_FAILED         = 1111
	CODE_TAPPAY_AUTH_REQUIRED              = 1200
	CODE_TAPPAY_EXCEPT_ERROR               = 1201
	CODE_TAPPAY_CARD_ERROR                 = 1202
	CODE_TAPPAY_BANK_ERROR                 = 1203
	CODE_TAPPAY_BIN_CODE_ERROR             = 1204
	CODE_TABLEAPP_UNEXPECT_ERROR           = 1300
	CODE_TABLEAPP_BUSINESS_NOT_EXIST       = 1301
	CODE_TABLEAPP_MODIFY_DENY              = 1302
	CODE_TABLEAPP_REMARK_PHONE_REQUIRED    = 1303
	CODE_FUNBOOK_NOT_ENOUGH_RESOURCE_ALERT = 1400
	CODE_INVALID_RESERVATION_STATUS        = 1401
	CODE_FUNBOOK_NOT_ENOUGH_RESOURCE       = 1402
	CODE_PRODUCT_IS_NOT_GIFTABLE           = 1500
	CODE_PRODUCT_PRICE_INVALID             = 1501
	CODE_GIFT_NOT_RECEIVABLE               = 1502
	CODE_GIFT_LINK_HASH_INVALID            = 1503
	CODE_CANT_RECEIVE_SELF_GIFT            = 1504
	CODE_GIFT_EXPIRED                      = 1505
	CODE_NO_GIFT_PERMISSION                = 1508
	CODE_GIFT_INVALID_STATUS               = 1509
	CODE_RECEIVER_CANT_DELETE_GIFT         = 1510
	CODE_INVALID_DELIVERY_STATUS           = 1600
	CODE_RWG_SLOT_ALREADY_BOOKED_BY_USER   = 1700
	CODE_RWG_INVALID_PRODUCT_SETTING       = 1701
	CODE_EZDING_UNEXPECTED_ERROR           = 1800
	CODE_EZDING_SESSION_UNAVAILABLE        = 1801
	CODE_EZDING_SESSION_FULL               = 1802
	CODE_EZDING_SESSION_NOT_EXISTS         = 1803
	CODE_EZDING_NOT_ALLOW_RELEASE_SEAT     = 1804
	CODE_EZDING_SEAT_ALREADY_RELEASED      = 1805
	CODE_EZDING_SEAT_ALREADY_OCCUPIED      = 1806
	CODE_EZDING_SESSION_TIME_EXCEEDED      = 1807
	CODE_EZDING_TICKET_INVALID             = 1808
	CODE_EZDING_PROMO_INVALID              = 1809
	CODE_EZDING_NOT_SUPPORT_DISCOUNT       = 1810
	CODE_EZDING_CINEMA_ERROR               = 1811
	CODE_EZDING_ORDER_ALREADY_PROCESS      = 1812
	CODE_EZDING_ORDER_ALREADY_COMPLETED    = 1813
	CODE_EZDING_PROPERTY_EMPTY             = 1814
	CODE_EZDING_INVALID_PROPERTY           = 1815
	CODE_EZDING_TRANS_NOT_EXIST            = 1816
	CODE_EZDING_TRANS_EXPIRED              = 1817
	CODE_EZDING_ORDER_QTY_ERROR            = 1818
	CODE_EZDING_CHECK_ORDER_FAILED         = 1819
	CODE_EZDING_ALREADY_CANCELED           = 1820
)

var ErrCodeMsgMap = map[int]string{
	CODE_PRODUCT_NOT_EXIST:                 "Product not Found",
	CODE_PROMO_COMPATIBLE_FAILED:           "Unable to compatible with promo",
	CODE_PROMO_INSUFFICIENT:                "promo code not enough",
	CODE_AUTH_ERR:                          "Auth failed",
	CODE_SESSION_ERR:                       "Session error",
	CODE_SEND_SMS_ERROR:                    "Send SMS failed",
	CODE_BONUS_CANT_USE_NOT_BINDING:        "Bonus cant use (Not binding FunNow account)",
	CODE_RELEASE_ALREADY:                   "release already",
	CODE_PAYMENT_TIMEOUT:                   "payment timeout",
	CODE_REVIEW_TAG_INVALID:                "review tag invalid",
	CODE_REDIS_ERR:                         "Redis server error",
	CODE_INVALID_CHECKOUT_CLOCK:            "Invalid checkout clock",
	CODE_INVALID_ORDER_TAX:                 "Empty company name or tax id is not allowed",
	CODE_INVALID_ORDER_REMARK:              "Empty order remark is not allowed",
	CODE_INVALID_BOOKINGTIME:               "Invalid booking time format",
	CODE_INVALID_ID:                        "Invalid ID format",
	CODE_BOOKINGTIME_EXPIRED:               "Booking time expired",
	CODE_ORDER_NOT_EXISTS:                  "Order not exist",
	CODE_ORDER_INVALID_STATUS:              "Invalid order status",
	CODE_INVALID_CUSTNUM:                   "Invalid order custnum",
	CODE_REDEEMED_ALREADY:                  "Order already redeem",
	CODE_GROUP_CODE_EOF:                    "Group code EOF",
	CODE_ORDERS_GROUP_NOT_EXIST:            "Orders group not exist",
	CODE_PRODUCT_TIMEZONE_NOT_FOUND:        "Product's timezone not found",
	CODE_INVALID_PRICE:                     "Invalid price",
	CODE_INVALID_TIME_RANGE:                "Invalid time range",
	CODE_INVALID_TIME_FORMAT:               "Invalid time format",
	CODE_FLASH_TAG_NOT_EXIST:               "Flash sale tags not found",
	CODE_INVALID_AC01_USER:                 "Invalid AC01 user",
	CODE_ORDER_DUPLICATE:                   "Order is duplicated",
	CODE_INVALID_PAGINATION_RANGE:          "Invalid pagination range",
	CODE_INVALID_GROUP_CODE:                "Invalid group code",
	CODE_INVALID_CLIENT_OS:                 "Invalid client os",
	CODE_INVALID_PLATFORM:                  "Invalid platform",
	CODE_INVALID_APP_VERSION:               "Invalid App version",
	CODE_INVALID_EZBOOK_UID:                "Invalid ezbook uid",
	CODE_REMARK_CUSTOMER_NECESSARY:         "Remark customer name is necessary",
	CODE_EBISOL_UNEXPECT_ERROR:             "Ebisol unexpect error",
	CODE_EBISOL_REMARK_FIELD_SETTING_ERROR: "Remark mobile number is required",
	CODE_EBISOL_CAPACITY_UNAVAILABLE:       "Ebisol capacity unavailable",
	CODE_EBISOL_USER_INFO_INCORRECT:        "Ebisol reservation user info is incorrect",
	CODE_EBISOL_MODIFY_DENY:                "Ebisol modify reservation deny",
	CODE_EBISOL_ORDER_NOT_EXIST:            "Ebisol order not exist",
	CODE_EBISOL_RATE_LIMIT:                 "Ebisol too many request",
	CODE_TRANS_TOOL:                        "Translation tool error",
	CODE_INVALID_BOSSNOW_USER:              "Invalid permission of operation",
	CODE_INVALID_WEEKDAY:                   "Invalid weekdays value",
	CODE_PRODUCT_SET_PRICE_NOT_FOUND:       "Product set price not found",
	CODE_ORDER_REVIEW_DENY:                 "Order is not allow to leave review",
	CODE_PARSE_TOKEN_ERR:                   "Token is not invalid",
	CODE_TOKEN_IS_EXPIRED_ERR:              "Token is expired",
	CODE_TOKEN_USER_NOT_MATCH_ERR:          "Token user is not match",
	CODE_APPLE_TOKEN_INVALID:               "Apple app token invalid",
	CODE_APPLE_TOKEN_EXPIRED:               "Apple app token expired",
	CODE_APPLE_TOKEN_ISSUER_INVALID:        "Apple app token issuer invalid",
	CODE_APPLE_TOKEN_AUDIENCE_INVALID:      "Apple app token audience invalid",
	CODE_APPLE_AUTH_INVALID:                "Apple auth invalid",
	CODE_STRIPE_EVENT_NOT_SUPPORT:          "Unsupported stripe event",
	CODE_STRIPE_AUTH_REQUIRED:              "This transaction requires authentication",
	CODE_STRIPE_PAYMENT_INTENT_INVALID:     "Stripe's request miss the requirement data or order's status is wrong",
	CODE_STRIPE_RELEASE_IN_WEBHOOK:         "Cancel the Stripe's PaymentIntent because the order has been released",
	CODE_STRIPE_RELEASE_FAILED_IN_WEBHOOK:  "Cancel the Stripe's PaymentIntent failed",
	CODE_STRIPE_WEBHOOK_SIGNATURE_INVALID:  "Invalid signature",
	CODE_STRIPE_DIGITAL_WALLET_ERROR:       "Create a PaymentIntent for digital wallet failed.",
	CODE_STRIPE_CREDIT_CARD_NOT_FOUND:      "Credit card not found",
	CODE_STRIPE_ENV_INVALID:                "Stripe environment invalid",
	CODE_STRIPE_GET_PAYMENT_INTENT_FAILED:  "Stripe get PaymentIntent failed",
	CODE_STRIPE_NEW_PAYMENT_FAILED:         "Stripe new PaymentIntent failed",
	CODE_RESOURCE_GROUP_CHANGED:            "Resource group has been changed",
	CODE_TAPPAY_AUTH_REQUIRED:              "This transaction requires authentication",
	CODE_INVALID_RESOURCE_NAME:             "Invalid resource name",
	CODE_INVALID_CURRENCY:                  "Invalid currency",
	CODE_GROUP_CODE_DELIVERY_DENY:          "Delivery type order is not allowed to be a group order",
	CODE_INVALID_COURIER_TYPE:              "Invalid courier type",
	CODE_INVALID_PRODUCT_STATUS:            "Invalid product status",
	CODE_INVALID_DATA_TYPE:                 "Invalid data type",
	CODE_NOT_ALLOW_FLASH_SALE:              "Not allow flash sale",
	CODE_TABLEAPP_UNEXPECT_ERROR:           "TableAPP unexpect error",
	CODE_TABLEAPP_BUSINESS_NOT_EXIST:       "TableAPP business not exist",
	CODE_TABLEAPP_MODIFY_DENY:              "TableAPP modify reservation deny",
	CODE_TABLEAPP_REMARK_PHONE_REQUIRED:    "Remark mobile number is required",
	CODE_INVALID_REDEEM_CODE:               "Redeem code is illegal",
	CODE_RATE_SMS_TOO_MANY:                 "Send SMS Request too many",
	CODE_INVALID_HEADER_ERR:                "Invalid header",
	CODE_VERIFY_CODE_EXPIRED:               "Verify code is expired",
	CODE_PWD_STRENGTH_NOT_ENOUGH:           "Password strength not enough",
	CODE_PWD_LENGTH_NOT_ENOUGH:             "Password length not enough",
	CODE_INVALID_RESERVATION_STATUS:        "Invalid reservation's status",
	CODE_ORDER_UNRESCHEDULABLE_STATUS:      "Rescheduling is not allowed for the orders status.",
	CODE_PRODUCT_IS_NOT_GIFTABLE:           "Product is not giftable",
	CODE_PRODUCT_PRICE_INVALID:             "Gift price is invalid",
	CODE_NO_GIFT_PERMISSION:                "No gift permission",
	CODE_GIFT_INVALID_STATUS:               "Invalid gift status",
	CODE_INVALID_DELIVERY_STATUS:           "Invalid delivery's status",
	CODE_GIFT_NOT_RECEIVABLE:               "Gift is not receivable",
	CODE_GIFT_LINK_HASH_INVALID:            "Gift link hash invalid",
	CODE_CANT_RECEIVE_SELF_GIFT:            "Cant receive self gift",
	CODE_GIFT_EXPIRED:                      "Gift is expired",
	CODE_RWG_SLOT_ALREADY_BOOKED_BY_USER:   "Slot already booked by user from RwG",
	CODE_RWG_INVALID_PRODUCT_SETTING:       "Invalid RwG product setting",
	CODE_PROMO_ALREADY_RECEIVED:            "Promo code already received",
	CODE_PROMO_NOT_ELIGIBLE_TO_RECEIVE:     "Not eligible to receive promo code",
	CODE_TNG_UNREGISTERED_ERR:              "Unregistered TNG account",
	CODE_EZDING_UNEXPECTED_ERROR:           "EZDing unexpected error",
	CODE_EZDING_SESSION_UNAVAILABLE:        "EZDing session unavailable",
	CODE_EZDING_SESSION_FULL:               "EZDing session seats are fulled",
	CODE_EZDING_SESSION_NOT_EXISTS:         "EZDing session not exists",
	CODE_EZDING_NOT_ALLOW_RELEASE_SEAT:     "The movie seat is not allow to release",
	CODE_EZDING_SEAT_ALREADY_RELEASED:      "Movie seat is already released",
	CODE_GUEST_ERR:                         "No permission for guest",
	CODE_EZDING_SEAT_ALREADY_OCCUPIED:      "Movie seat was already occupied",
	CODE_EZDING_SESSION_TIME_EXCEEDED:      "Movie session time was exceeded",
	CODE_EZDING_TICKET_INVALID:             "Movie ticket invalid",
	CODE_EZDING_PROMO_INVALID:              "Promo invalid for the movie ticket",
	CODE_EZDING_NOT_SUPPORT_DISCOUNT:       "Movie not support discount",
	CODE_EZDING_CINEMA_ERROR:               "EZDing cinema return error",
	CODE_EZDING_ORDER_ALREADY_PROCESS:      "EZDing order already process",
	CODE_EZDING_ORDER_ALREADY_COMPLETED:    "EZDing order already completed",
	CODE_EZDING_PROPERTY_EMPTY:             "EZDing property empty",
	CODE_EZDING_INVALID_PROPERTY:           "EZDing invalid property",
	CODE_EZDING_TRANS_NOT_EXIST:            "EZDing trans not exist",
	CODE_EZDING_TRANS_EXPIRED:              "EZDing trans expired",
	CODE_EZDING_ORDER_QTY_ERROR:            "EZDing order quantity error",
	CODE_EZDING_CHECK_ORDER_FAILED:         "EZDing order check failed",
	CODE_EZDING_ALREADY_CANCELED:           "EZDing order was already canceled",
}

func Throw(c *gin.Context, err error) {

	if gorm.IsRecordNotFoundError(err) || err == sql.ErrNoRows {
		DBError(c, err)
		return
	}

	switch e := err.(type) {
	case *mysql.MySQLError, gorm.Errors:
		DBError(c, e)
	// case *RedisErr:
	// 	RedisError(c, e)
	// case ICustomError:
	// 	CustomError(c, e)
	default:
		Error(c, http.StatusInternalServerError, CODE_UNKNOWN_ERR, e)
	}
}

func Error(c *gin.Context, httpCode int, code int, err interface{}) {
	var msg string
	switch err := err.(type) {
	case CustomMsg:
		// Error message and log already handled by CustomError
		msg = string(err)
	case error:
		// logs.Log(c, logs.TypeErr, err.Error())
		msg = getMsgByCode(c, code, err.Error())
	case string:
		// logs.Log(c, logs.TypeErr, err)
		msg = getMsgByCode(c, code, err)
	}

	if m, ok := err.(map[string][]interface{}); ok {
		c.JSON(httpCode, apires.Errs{
			Base: apires.Base{
				Code:    code,
				Message: msg,
			},
			Errs: m,
		})
	} else {
		c.JSON(httpCode, apires.Base{
			Code:    code,
			Message: msg,
		})
	}

	c.Abort()
}

func getMsgByCode(c *gin.Context, code int, msg string) string {

	// T := lang.GetTFunc(c)
	switch code {
	case CODE_AUTH_ERR:
		msg = "Authentication failed"
	case CODE_LOGIN_ERR:
		msg = "Login failed"
		// if T != nil {
		// 	msg = T("login_failed_err")
		// }
	case CODE_SESSION_ERR:
		msg = "Session error"
	case CODE_GUEST_ERR:
		msg = "Guest Authentication failed"
	case CODE_NOT_EXISTS:
		msg = "Data not found"
		// if T != nil {
		// 	msg = T("data_not_found")
		// }
	case CODE_PROMO_CREDIT_ERR:
		msg = "Invalid credit card to use the promo code"
		// if T != nil {
		// 	msg = T("invalid_promo_credit")
		// }
	case CODE_INVALID_PARAMS:
		if msg == "" {
			msg = "Invalid input parameter"
			// if T != nil {
			// 	msg = T("params_err")
			// }
		}
	}
	return msg
}

func DBError(c *gin.Context, err error) {

	if err == sql.ErrNoRows || gorm.IsRecordNotFoundError(err) {
		Error(c, http.StatusNotFound, CODE_NOT_EXISTS, "Data not found")
		return
	}

	myerr, ok := err.(*mysql.MySQLError)
	if !ok {
		Error(c, http.StatusInternalServerError, CODE_DB_ERR, err)
		return
	}

	if myerr.Number == 1062 {
		Error(c, http.StatusBadRequest, CODE_DUPLICATE_KEY, "Duplicate data already exists")
	} else if myerr.Number == 1452 {
		Error(c, http.StatusBadRequest, CODE_INVALID_PARAMS, "Reference ID not exists")
	} else {
		Error(c, http.StatusInternalServerError, CODE_DB_ERR, err)
	}
}
