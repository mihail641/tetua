package webuser

import (
	"errors"
	"fmt"
	"github.com/ngocphuongnb/tetua/app/url_utils"
	"strings"
	"time"

	netmail "net/mail"

	"github.com/ngocphuongnb/tetua/app/auth"
	"github.com/ngocphuongnb/tetua/app/config"
	"github.com/ngocphuongnb/tetua/app/entities"
	"github.com/ngocphuongnb/tetua/app/logger"
	"github.com/ngocphuongnb/tetua/app/mail"
	"github.com/ngocphuongnb/tetua/app/repositories"
	"github.com/ngocphuongnb/tetua/app/server"
	"github.com/ngocphuongnb/tetua/app/utils"
	"github.com/ngocphuongnb/tetua/views"
)

type RegisterData struct {
	Username             string `json:"username"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordconfirmation"`
}

func (r *RegisterData) Parse(c server.Context) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if r.Username == "" || r.Email == "" || r.Password == "" || r.PasswordConfirmation == "" {
		return errors.New("Please fill all required fields")
	}

	if _, err := netmail.ParseAddress(r.Email); err != nil {
		return errors.New("Invalid email address")
	}

	if r.Password != r.PasswordConfirmation {
		return errors.New("Password and Password confirmation doesn't match")
	}

	foundUsers, err := repositories.User.ByUsernameOrEmail(c.Context(), r.Username, r.Email)

	if err != nil {
		c.Logger().Error(err)
		return errors.New("Something went wrong")
	}

	if len(foundUsers) > 0 {
		return errors.New("Username or email already exists")
	}

	if r.Password, err = utils.GenerateHash(r.Password); err != nil {
		c.Logger().Error(err)
		return errors.New("Something went wrong")
	}

	return nil
}
func Register(c server.Context) (err error) {
	redirectURL := url_utils.GetRedirectURL(c)
	if c.User() != nil && c.User().ID > 0 {
		return c.Redirect(redirectURL)
	}
	c.Meta().Title = "Register"

	return c.Render(views.Register("", "", ""))
}

func PostRegister(c server.Context) (err error) {
	redirectURL := url_utils.GetRedirectURL(c)

	register := &RegisterData{}
	autoApproveUser := config.Setting("auto_approve_user") == "yes"
	messageError := "Your account has been created. But we can't send you an email to activate your account. "
	if err = register.Parse(c); err != nil {
		c.Messages().AppendError(err.Error())
		return c.Render(views.Register(register.Username, register.Email, ""))
	}

	user, err := repositories.User.Create(c.Context(), &entities.User{
		Username: register.Username,
		Email:    register.Email,
		Password: register.Password,
		RoleIDs:  []int{auth.ROLE_USER.ID},
		Provider: "local",
		Active:   autoApproveUser,
	})

	if err != nil {
		c.WithError("Something went wrong", err)
		return c.Render(views.Register(register.Username, register.Email, redirectURL))
	}

	mailBody := []string{fmt.Sprintf("Welcome <b>%s</b>, We're happy to have you with us.", user.Username)}
	welcomeMessage := "Your account has been activated. You can now login to the site."

	if autoApproveUser {
		mailBody = append(
			mailBody,
			fmt.Sprintf("You can now login to your account at %s and start writing your first post", utils.Url("login")),
		)
	} else {
		welcomeMessage = "Your account has been created. We've sent you an email to activate your account, please check your mailbox."
		activationCode, err := GetActivationCode(user.ID)
		if err != nil {
			logger.Error(err)
			welcomeMessage = messageError
			welcomeMessage += "please contact us with your username/email and this trace id: " + c.RequestID()
		}
		link := c.BaseUrl() + "/activate?code=" + activationCode

		mailBody = append(
			mailBody,
			"Follow the link below to complete your registration:",
			"<a href='"+link+"'>"+link+"</a>",
		)
	}

	mailBody = append(mailBody, fmt.Sprintf("<br><b>Cheer</b>,<br>The %s Team", config.Setting("app_name")))
	err = mail.Send(c,
		user.Username,
		user.Email,
		fmt.Sprintf("Welcome to %s", config.Setting("app_name")),
		strings.Join(mailBody, "<br>"),
	)

	if err != nil {
		logger.Get().WithContext(logger.Context{"request_id": c.RequestID()}).Error(err)
		return c.Render(views.Message("Thank you for signing up", messageError, redirectURL, 0))
	}

	return c.Render(views.Message("Thank you for signing up", welcomeMessage, redirectURL, 0))
}
func GetActivationCode(userID int) (string, error) {
	exp := time.Now().Add(time.Hour * 24)
	activationCode, err := utils.Encrypt(fmt.Sprintf("%d_%d", userID, exp.UnixMicro()))
	if err != nil {
		return "", err
	}
	return activationCode, err
}
