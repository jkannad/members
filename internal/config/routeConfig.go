package config

const (
	ROOT = "/"
	GET_ABOUT = "/member/about"
	GET_SEARCH = "/member/search"
	GET_REGISTER = "/member/register"
	GET_MEMBER_BY_ID = "/member/getmember/{id:[0-9]+}"
	GET_ALL_MEMBERS = "/member/getallmembers"
	GET_STATES_BY_COUNTRY = "/member/getstates/{id}"
	GET_CITIES_BY_COUNTRY_AND_STATE = "/member/getcities/{country}/{state}"
	GET_DIAL_CODE_BY_COUNTRY = "/member/getdialcode/{id}"
	GET_LOGIN = "/member/login"
	GET_LOGOUT = "/member/logout"

	POST_LOGIN = "/member/login"
	POST_UPSERT = "/member/upsert"
	POST_SEARCH_RESULT = "/member/search/result"
)