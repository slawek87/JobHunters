package linkedin

// custom settings - modify them with your requirements.
//const REDIRECT_PATH = "/login/"
const REDIRECT_URI = "http://0.0.0.0:8000/api/v1/user/login"
const CLIENT_ID = "770vfbx6zalos0"
const CLIENT_SECRET = "YJK5RcXiYISsLYzz"
const STATE = "12nsd123"
const SCOPE= "r_basicprofile"

// default settings - in most cases should be the same.
const RESPONSE_TYPE = "code"
const GRANT_TYPE = "authorization_code"
const AUTHORIZATION_ENDPOINT = "https://www.linkedin.com/oauth/v2/authorization?"
const ACCESS_TOKEN_ENDPOINT = "https://www.linkedin.com/oauth/v2/accessToken"