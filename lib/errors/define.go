package errors

import (
	"fmt"
	"net/http"
)

var (
	// ErrNoSellProduct for insert sell product to es
	ErrNoSellProduct = fmt.Errorf("no sell product insert to es")
	// ErrLoginFailed when user login with wrong data
	ErrLoginFailed = NewErr(http.StatusForbidden, CODE_LOGIN_ERR)
	// ErrLoginDeleting when login with member status in deleting
	ErrLoginDeleting = NewErr(http.StatusForbidden, CODE_LOGIN_DELETING_ERR)
	// ErrInvalidParams when getting invalid params from API request
	ErrInvalidParams = NewErr(http.StatusBadRequest, CODE_INVALID_PARAMS)
	// ErrSession when failed to deal with session data
	ErrSession = NewErr(http.StatusForbidden, CODE_SESSION_ERR)
	// ErrAuth when failed to check auth from session
	ErrAuth = NewErr(http.StatusUnauthorized, CODE_AUTH_ERR)
	// ErrGuest when guest access to member only functions
	ErrGuest = NewErr(http.StatusUnauthorized, CODE_GUEST_ERR)
	// ErrNoData when data not exists
	ErrNoData = NewErr(http.StatusNotFound, CODE_NOT_EXISTS)
	// ErrInvalidTimeFormat when getting invalid time format from params
	ErrInvalidTimeFormat = NewErr(http.StatusBadRequest, CODE_INVALID_TIME_FORMAT)
	// ErrInvalidTimeRange when getting invalid time range
	ErrInvalidTimeRange = NewErr(http.StatusBadRequest, CODE_INVALID_TIME_RANGE)
	// ErrInvalidTimeInterval when time interval is not valid
	ErrInvalidTimeLimitation = NewErr(http.StatusBadRequest, CODE_INVALID_TIME_LIMITATION)
	// ErrInvalidType when input data type is invalid
	ErrInvalidType = NewErr(http.StatusInternalServerError, CODE_INVALID_DATA_TYPE)
	// ErrIncorrectManagerRole when login as a manager user which doesn't has correct role
	ErrIncorrectManagerRole = NewErr(http.StatusForbidden, CODE_INCORRECT_MANAGER_ROLE_ERR)
	// ErrInvalidHeader when getting invalid or missing headers from API request
	ErrInvalidHeader = NewErr(http.StatusBadRequest, CODE_INVALID_HEADER_ERR)
	// ErrInvalidWeekDay when the weekday input array is out of range (0 - 6)
	ErrInvalidWeekDay = NewErr(http.StatusBadRequest, CODE_INVALID_WEEKDAY)
	// ErrInvalidPrice when the input price is not legal
	ErrInvalidPrice = NewErr(http.StatusBadRequest, CODE_INVALID_PRICE)
	// ErrInvalidPrice when the input sell limit is not legal
	ErrInvalidSellLimit = NewErr(http.StatusBadRequest, CODE_INVALID_SELLLIMIT)
	// ErrInvalidManagerPermission when user operation no permission data
	ErrInvalidManagerPermission = NewErr(http.StatusForbidden, CODE_INVALID_BOSSNOW_USER)
	// ErrTimePairOverlap when the input time range is overlap to set time range
	ErrTimePairOverlap = NewErr(http.StatusBadRequest, CODE_TIME_PAIR_OVERLAP)
	// ErrDB when getting database error
	ErrDB = NewErr(http.StatusInternalServerError, CODE_DB_ERR)
	// ErrSoldOut when ES sellproduct not found
	ErrSoldOut = NewErr(http.StatusOK, CODE_SOLD_OUT)
	// ErrPasswordStrength when password strength not enough
	ErrPasswordStrength = NewErr(http.StatusBadRequest, CODE_PWD_STRENGTH_NOT_ENOUGH)
	// ErrPasswordStrength when password length not enough
	ErrPasswordLength = NewErr(http.StatusBadRequest, CODE_PWD_LENGTH_NOT_ENOUGH)
	// ErrMemberDeletionNotAllow when FunNow member is not allow to delete account
	ErrMemberDeletionNotAllow = NewErr(http.StatusBadRequest, CODE_ACCOUNT_DELETION_NOT_ALLOW)
	// ErrMemberDeletionApplyFailed when FunNow member apply deletion failed
	ErrMemberDeletionApplyFailed = NewErr(http.StatusBadRequest, CODE_ACCOUNT_DELETION_APPLY_FAILED)
	// ErrMemberDeletionTokenExpired when FunNow member deletion token expired
	ErrMemberDeletionTokenExpired = NewErr(http.StatusBadRequest, CODE_ACCOUNT_DELETION_TOKEN_EXPIRED)
	// ErrBonusNoEnough
	ErrBonusNoEnough = NewErr(http.StatusBadRequest, CODE_BONUS_NOT_ENOUGH)
	// ErrRegistered when FunNow member register with already registered AuthID
	// ErrRegistered = NewErr(http.StatusBadRequest, CODE_DUPLICATE_KEY).SetMsgCode(CODE_REGISTERED_ERR)
	// ErrRedeemCodeNotFound when input redeem code not found
	ErrRedeemCodeNotFound = NewErr(http.StatusBadRequest, CODE_REDEEM_CODE_NOT_FOUND)
	// ErrReleaseAlready
	ErrReleaseAlready = NewErr(http.StatusConflict, CODE_RELEASE_ALREADY)
	// ErrPaymentTimeout
	ErrPaymentTimeout = NewErr(http.StatusConflict, CODE_PAYMENT_TIMEOUT)
	// ErrGrabPay
	ErrGrabPay = NewErr(http.StatusBadGateway, CODE_GRABPAY_ERR)
	// ErrGetToken
	ErrGetToken = NewErr(http.StatusInternalServerError, CODE_GET_TOKEN_FAILED)
	// ErrPayment
	ErrPayment = NewErr(http.StatusInternalServerError, CODE_PAYMENT_ERR)
	// ErrPaymentCodeOK get some error but need to return HTTP status OK
	ErrPaymentCodeOK = NewErr(http.StatusOK, CODE_PAYMENT_ERR)
	// ErrInvalidBinCode
	ErrInvalidBinCode = NewErr(http.StatusOK, CODE_TAPPAY_BIN_CODE_ERROR)
	// ErrMobilePay
	// ErrMobilePay = NewErr(http.StatusOK, CODE_PAYMENT_ERR).SetMsgCode(CODE_MOBILE_PAY_FAILED)
	// ErrMovieTicketInvalid
	ErrMovieTicketInvalid = NewErr(http.StatusBadRequest, CODE_EZDING_TICKET_INVALID)
	// ErrMoviePromoInvalid
	ErrMoviePromoInvalid = NewErr(http.StatusBadRequest, CODE_EZDING_PROMO_INVALID)
	// ErrMovieNotSupportDiscount
	ErrMovieNotSupportDiscount = NewErr(http.StatusBadRequest, CODE_EZDING_NOT_SUPPORT_DISCOUNT)
	// ErrEZDingSeatOccupied
	ErrEZDingSeatOccupied = NewErr(http.StatusBadRequest, CODE_EZDING_SEAT_ALREADY_OCCUPIED)
	// ErrEZDingSessionTimeExceeded
	ErrEZDingSessionTimeExceeded = NewErr(http.StatusBadRequest, CODE_EZDING_SESSION_TIME_EXCEEDED)
	// ErrInvalidEmail
	ErrInvalidEmail = NewErr(http.StatusBadRequest, CODE_EMAIL_ERR)
	// ErrInvalidPhone
	ErrInvalidPhone = NewErr(http.StatusBadRequest, CODE_PHONE_INVALID)
)
