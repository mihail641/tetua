package url_utils

import (
	"github.com/ngocphuongnb/tetua/app/server"
)

const RedirectURLConst = "redirectURL"

func GetRedirectURL(c server.Context) string {
	redirectURL := c.Query(RedirectURLConst)
	return redirectURL
}
