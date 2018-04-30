package globals

import (
	"github.com/gorilla/sessions"
)

// セッションストア
// TODO main.goとusers_controller.goで使いたいので、
// グローバル変数としてみたが正しいか分からない
var Store *sessions.CookieStore
