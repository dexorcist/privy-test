package enum

type ResponseCode int

const (
	HTTPErrorBadRequest             ResponseCode = -4000
	HTTPErrorCakeIDRequired         ResponseCode = -4001
	HTTPErrorTitleRequired          ResponseCode = -4002
	HTTPErrorDescriptionRequired    ResponseCode = -4003
	HTTPErrorRatingRequired         ResponseCode = -4004
	HTTPErrorImageRequired          ResponseCode = -4005
	HTTPErrorMaxRatingBad           ResponseCode = -4006
	HTTPErrorMinBiggerMaxRattingBad ResponseCode = -4007
	HTTPErrorCakeNotFound           ResponseCode = -4041
	HTTPErrorInternalServerError    ResponseCode = -5001
)

var messageMap = map[ResponseCode]map[string]string{
	HTTPErrorBadRequest: {
		"en": "Bad Request.",
		"id": "Permintaan Buruk.",
	},
	HTTPErrorCakeIDRequired: {
		"en": "Cake ID Required.",
		"id": "Cake ID diperlukan.",
	},
	HTTPErrorTitleRequired: {
		"en": "Title can't be empty.",
		"id": "Judul harus diisi.",
	},
	HTTPErrorDescriptionRequired: {
		"en": "Description can't be empty.",
		"id": "Deskripsi harus diisi.",
	},
	HTTPErrorImageRequired: {
		"en": "Image can't be empty.",
		"id": "Image harus diisi.",
	},
	HTTPErrorRatingRequired: {
		"en": "Rating must be greater than zero.",
		"id": "Rating Harus lebih dari 0.",
	},
	HTTPErrorMaxRatingBad: {
		"en": "max rating is 10.",
		"id": "Rating maksimal adalah 10.",
	},
	HTTPErrorMinBiggerMaxRattingBad: {
		"en": "Min rating is not bigger than max rating",
		"id": "Min rating tidak lebih besar max rating.",
	},
	HTTPErrorCakeNotFound: {
		"en": "Cake Not Found.",
		"id": "Kue tidak ditemukan.",
	},
	HTTPErrorInternalServerError: {
		"en": "Internal Server Error.",
		"id": "Kesalahan server dari dalam.",
	},
}

func (responseCode ResponseCode) StringMap(code ResponseCode) map[string]string {
	return messageMap[code]
}

func (responseCode ResponseCode) Int() int {
	return int(responseCode)
}
