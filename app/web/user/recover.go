package webuser

import (
	"errors"
	"fmt"
	"github.com/ngocphuongnb/tetua/app/config"
	"github.com/ngocphuongnb/tetua/app/logger"
	"github.com/ngocphuongnb/tetua/app/mail"
	"github.com/ngocphuongnb/tetua/app/repositories"
	"github.com/ngocphuongnb/tetua/app/server"
	"github.com/ngocphuongnb/tetua/app/url_utils"
	"github.com/ngocphuongnb/tetua/app/utils"
	"github.com/ngocphuongnb/tetua/views"
	"net/url"
	"strings"
)

type RecoverData struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordconfirmation"`
}

func (r *RecoverData) ParsePassword(c server.Context) error {
	var err error
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if r.Password == "" || r.PasswordConfirmation == "" {
		return errors.New("Please fill all required fields")
	}
	if r.Password != r.PasswordConfirmation {
		return errors.New("Password and Password confirmation doesn't match")
	}
	if r.Password, err = utils.GenerateHash(r.Password); err != nil {
		c.Logger().Error(err)
		return err
	}
	return nil
}
func (r *RecoverData) ParseEMail(c server.Context) error {

	if err := c.BodyParser(r); err != nil {
		return err
	}
	if r.Email == "" {
		return errors.New("Please fill a field")
	}
	return nil
}
func Recover(c server.Context) (err error) {
	c.Meta().Title = "Password recovery"
	return c.Render(views.Recover(""))

}
func PostRecover(c server.Context) (err error) {
	c.Meta().Title = " Password recovery "
	redirectURL := url_utils.GetRedirectURL(c)
	messageError := "We can't send you an email to recover your password. "

	recoverData := &RecoverData{}
	if err := recoverData.ParseEMail(c); err != nil {
		c.Messages().AppendError(err.Error())
		return c.Render(views.Recover(recoverData.Email))
	}
	welcomeMessage := fmt.Sprintf("We have sent you an e-mail %s to reset your password, please follow the instructions in the letter", recoverData.Email)
	foundUsers, err := repositories.User.ByEmail(c.Context(), recoverData.Email, recoverData.Email)
	if err != nil {
		c.Logger().Error(err)
		c.Messages().AppendError("Something went wrong")
		return c.Render(views.Recover(recoverData.Email))
	}

	if foundUsers == nil {
		c.Messages().AppendError("Invalid e-mail information")
		return c.Render(views.Recover(recoverData.Email))
	}

	if !foundUsers.IsRoot() && !foundUsers.Active {
		return c.Redirect(utils.Url("/inactive"))
	}

	mailBody := []string{fmt.Sprintf("%s, you fogot your password. ", foundUsers.Username)}
	activationCode, err := GetActivationCode(foundUsers.ID)
	if err != nil {
		logger.Error(err)
		welcomeMessage = messageError
		welcomeMessage += "please contact us with your username/email and this trace id: " + c.RequestID()
	}
	link := c.BaseUrl() + "/recover/" + activationCode + "?redirectURL=" + url.QueryEscape(redirectURL)
	mailBody = append(
		mailBody,
		"Follow the link below to complete your password recovery:",
		"<a href='"+link+"'>"+link+"</a>",
	)

	mailBody = append(mailBody, fmt.Sprintf("<br><b>Cheer</b>,<br>The %s Team", config.Setting("app_name")))
	err = mail.Send(c,
		foundUsers.Username,
		foundUsers.Email,
		fmt.Sprintf("Password recovery"),
		strings.Join(mailBody, "<br>"),
	)

	if err != nil {
		logger.Get().WithContext(logger.Context{"request_id": c.RequestID()}).Error(err)
		return c.Render(views.Message("You forgot your password", messageError, redirectURL, 0))
	}

	return c.Render(views.Message("Password recovery ", welcomeMessage, redirectURL, 0))
}
func ChangePassword(c server.Context) (err error) {
	c.Meta().Title = " New password "
	return c.Render(views.ChangePassword())

}

func PostChangePassword(c server.Context) (err error) {
	c.Meta().Title = " Password changed "
	redirectURL := url_utils.GetRedirectURL(c)
	recover := &RecoverData{}
	if err = recover.ParsePassword(c); err != nil {
		c.Messages().AppendError(err.Error())
		return c.Render(views.ChangePassword())
	}
	code := c.Param("code")
	userID, err := CheckActivation(c, code)
	if err != nil {
		c.WithError("Code activate with error", err)
		return err
	}
	user, err := repositories.User.ByID(c.Context(), userID)
	if err != nil {
		c.WithError("There is no user with this ID", err)
		return c.Render(views.ChangePassword())
	}
	user.Password = recover.Password
	_, err = repositories.User.Update(c.Context(), user)
	if err != nil {
		c.WithError("Something went wrong", err)
		return c.Render(views.ChangePassword())
	}
	welcomeMessage := "Password сhanged, please do Login"
	redirectURL = fmt.Sprintf("/login?redirectURL=%s", url.QueryEscape(redirectURL))
	return c.Render(views.Message("Password changed", welcomeMessage, redirectURL, 3))
}
