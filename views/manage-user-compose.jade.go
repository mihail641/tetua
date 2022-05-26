// Code generated by "jade.go"; DO NOT EDIT.

package views

import (
	"bufio"
	"fmt"

	"github.com/ngocphuongnb/tetua/app/asset"
	"github.com/ngocphuongnb/tetua/app/cache"
	"github.com/ngocphuongnb/tetua/app/config"
	"github.com/ngocphuongnb/tetua/app/entities"
	"github.com/ngocphuongnb/tetua/app/server"
	"github.com/ngocphuongnb/tetua/app/utils"
)

const (
	manageusercompose__21  = `<p>Auth provider<select name="provider">`
	manageusercompose__22  = `</select></p>`
	manageusercompose__23  = `<hr/><strong>To keep the old password, leave this field blank.</strong>`
	manageusercompose__26  = `</div><strong>Select roles</strong>`
	manageusercompose__27  = `<div><strong>Avatar</strong><input class="image-input" id="avatar-image" type="file" name="avatar_image"/><div class="image-upload-previewer" for="avatar-image">`
	manageusercompose__28  = `</div></div></div></div></div></form></div><div class="mobile-menu"><div class="menu-head">`
	manageusercompose__33  = `<script>listenDeleteNodeEvents('user', '/manage/users', '/manage/users')</script></body></html>`
	manageusercompose__138 = `<button class="danger delete-user" data-id="`
	manageusercompose__140 = `<div class="multi-checkbox scroll">`
	manageusercompose__142 = `<label for="`
	manageusercompose__143 = `"><input type="checkbox" name="`
	manageusercompose__145 = `" id="`
	manageusercompose__146 = `" checked="checked"/><span class="name">`
	manageusercompose__147 = `</span></label>`
	manageusercompose__152 = `"/><span class="name">`
	manageusercompose__156 = `<img/>`
)

func ManageUserCompose(ID int, user *entities.User, roles []*entities.Role, providers []server.AuthProvider) func(meta *entities.Meta, wr *bufio.Writer) {
	return func(meta *entities.Meta, wr *bufio.Writer) {
		buffer := &WriterAsBuffer{wr}

		buffer.WriteString(commentlist__0)

		var title = meta.GetTitle()
		var appName = config.Setting("app_name")
		var appLogo = config.Setting("app_logo")
		buffer.WriteString(commentlist__1)
		WriteAll(title, true, buffer)
		buffer.WriteString(commentlist__2)
		WriteAll(meta.Canonical, true, buffer)
		buffer.WriteString(commentlist__3)
		WriteAll(meta.Type, true, buffer)
		buffer.WriteString(commentlist__4)
		WriteAll(meta.Canonical, true, buffer)
		buffer.WriteString(commentlist__5)
		WriteAll(title, true, buffer)
		buffer.WriteString(commentlist__6)
		WriteAll(appName, true, buffer)
		buffer.WriteString(commentlist__7)
		WriteAll(config.Setting("twitter_site"), true, buffer)
		buffer.WriteString(commentlist__8)
		WriteAll(title, true, buffer)
		buffer.WriteString(commentlist__9)
		WriteAll(appName, true, buffer)
		buffer.WriteString(commentlist__10)
		WriteAll(appName, true, buffer)
		buffer.WriteString(commentlist__11)
		WriteAll(appName+" Feed", true, buffer)
		buffer.WriteString(commentlist__12)
		WriteAll(utils.Url("/feed"), true, buffer)
		buffer.WriteString(commentlist__13)
		if appLogo != "" {
			buffer.WriteString(commentlist__30)
			WriteAll(appLogo, true, buffer)
			buffer.WriteString(commentlist__31)
			WriteAll(appLogo, true, buffer)
			buffer.WriteString(commentlist__13)
		}
		if meta.Description != "" {
			buffer.WriteString(commentlist__33)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(commentlist__34)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(commentlist__35)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(commentlist__13)
		}
		if meta.Image != "" {
			buffer.WriteString(commentlist__37)
			WriteAll(meta.Image, true, buffer)
			buffer.WriteString(commentlist__38)
			WriteAll(meta.Image, true, buffer)
			buffer.WriteString(commentlist__13)
		}
		WriteAll(asset.CssFile("css/light.min.css"), false, buffer)
		WriteAll(asset.CssFile("css/style.css"), false, buffer)
		WriteAll(config.Setting("inject_header"), false, buffer)
		buffer.WriteString(commentlist__14)
		WriteAll(utils.Url(""), true, buffer)
		buffer.WriteString(commentlist__15)
		var logoUrl = config.Setting("app_logo")
		if logoUrl != "" {
			buffer.WriteString(commentlist__40)
			WriteAll(logoUrl, true, buffer)
			buffer.WriteString(commentlist__41)
			WriteAll(config.Setting("app_name"), true, buffer)
			buffer.WriteString(commentlist__13)
		} else {
			buffer.WriteString(commentlist__43)

		}
		buffer.WriteString(commentlist__16)
		WriteAll(meta.Query, true, buffer)
		buffer.WriteString(commentlist__17)
		WriteAll(utils.Url("/search"), true, buffer)
		buffer.WriteString(commentlist__18)

		if meta.User == nil || meta.User.ID == 0 {
			buffer.WriteString(commentlist__44)
			WriteAll(utils.Url("/login"), true, buffer)
			buffer.WriteString(commentlist__45)
			WriteAll(utils.Url("/register"), true, buffer)
			buffer.WriteString(commentlist__46)

		} else {
			buffer.WriteString(commentlist__44)
			WriteAll(utils.Url("/posts/new"), true, buffer)
			buffer.WriteString(commentlist__48)
			WriteAll(meta.User.Url(), true, buffer)
			buffer.WriteString(commentlist__49)
			WriteAll(meta.User.Username, true, buffer)
			buffer.WriteString(commentlist__50)
			if meta.User.AvatarImageUrl != "" {
				buffer.WriteString(commentlist__57)
				WriteAll(meta.User.AvatarImageUrl, true, buffer)
				buffer.WriteString(commentlist__41)
				WriteAll(meta.User.Username, true, buffer)
				buffer.WriteString(commentlist__13)
			} else {
				buffer.WriteString(commentlist__60)

			}
			buffer.WriteString(commentlist__51)

			if meta.User != nil && meta.User.IsRoot() {
				buffer.WriteString(commentlist__44)
				WriteAll(utils.Url("/manage"), true, buffer)
				buffer.WriteString(commentlist__62)

			}
			buffer.WriteString(commentlist__44)
			WriteAll(meta.User.Url(), true, buffer)
			buffer.WriteString(commentlist__53)
			WriteAll(utils.Url("/posts"), true, buffer)
			buffer.WriteString(commentlist__54)
			WriteAll(utils.Url("/settings"), true, buffer)
			buffer.WriteString(commentlist__55)
			WriteAll(utils.Url("/logout"), true, buffer)
			buffer.WriteString(commentlist__56)

		}
		buffer.WriteString(managepagecompose__19)

		{
			buffer.WriteString(commentlist__119)
			WriteAll(utils.Url("/manage"), true, buffer)
			buffer.WriteString(commentlist__120)
			WriteAll(utils.Url("/manage/topics"), true, buffer)
			buffer.WriteString(commentlist__121)
			WriteAll(utils.Url("/manage/posts"), true, buffer)
			buffer.WriteString(commentlist__122)
			WriteAll(utils.Url("/manage/pages"), true, buffer)
			buffer.WriteString(commentlist__123)
			WriteAll(utils.Url("/manage/roles"), true, buffer)
			buffer.WriteString(commentlist__124)
			WriteAll(utils.Url("/manage/users"), true, buffer)
			buffer.WriteString(commentlist__125)
			WriteAll(utils.Url("/manage/comments"), true, buffer)
			buffer.WriteString(commentlist__126)
			WriteAll(utils.Url("/manage/files"), true, buffer)
			buffer.WriteString(commentlist__127)
			WriteAll(utils.Url("/manage/settings"), true, buffer)
			buffer.WriteString(commentlist__72)

		}

		buffer.WriteString(managecommentindex__20)

		if ID > 0 {
			buffer.WriteString(managepagecompose__77)
			WriteAll("Editing user: "+user.Username, true, buffer)
			buffer.WriteString(error__20)
		} else {
			buffer.WriteString(managepagecompose__77)
			WriteEscString("Create a new user", buffer)
			buffer.WriteString(error__20)
		}
		{
			var (
				msgs = meta.Messages
			)

			if msgs.Length() > 0 {
				buffer.WriteString(commentlist__73)
				var messages = msgs.Get()
				for _, msg := range messages {
					buffer.WriteString(commentlist__75)
					WriteAll(msg.Type, true, buffer)
					buffer.WriteString(commentlist__50)
					WriteAll(msg.Message, true, buffer)
					buffer.WriteString(commentlist__77)
				}
				buffer.WriteString(commentlist__74)
			}
		}

		{
			var (
				name  = "username"
				value = user.Username
				label = "Username"
			)

			buffer.WriteString(managepagecompose__85)
			WriteEscString(label, buffer)
			buffer.WriteString(managepagecompose__86)
			WriteEscString(name, buffer)
			buffer.WriteString(managepagecompose__87)
			WriteAll(value, true, buffer)
			buffer.WriteString(managepagecompose__88)
		}

		{
			var (
				name  = "display_name"
				value = user.DisplayName
				label = "Display name"
			)

			buffer.WriteString(managepagecompose__85)
			WriteEscString(label, buffer)
			buffer.WriteString(managepagecompose__86)
			WriteEscString(name, buffer)
			buffer.WriteString(managepagecompose__87)
			WriteAll(value, true, buffer)
			buffer.WriteString(managepagecompose__88)
		}

		{
			var (
				name  = "email"
				value = user.Email
				label = "Email"
			)

			buffer.WriteString(managepagecompose__85)
			WriteEscString(label, buffer)
			buffer.WriteString(managepagecompose__86)
			WriteEscString(name, buffer)
			buffer.WriteString(managepagecompose__87)
			WriteAll(value, true, buffer)
			buffer.WriteString(managepagecompose__88)
		}

		{
			var (
				name  = "url"
				value = user.URL
				label = "Url"
			)

			buffer.WriteString(managepagecompose__85)
			WriteEscString(label, buffer)
			buffer.WriteString(managepagecompose__86)
			WriteEscString(name, buffer)
			buffer.WriteString(managepagecompose__87)
			WriteAll(value, true, buffer)
			buffer.WriteString(managepagecompose__88)
		}

		{
			var (
				name  = "bio"
				value = user.Bio
				label = "User bio"
			)

			buffer.WriteString(managepagecompose__85)
			WriteEscString(label, buffer)
			buffer.WriteString(managetopiccompose__87)
			WriteEscString(name, buffer)
			buffer.WriteString(commentlist__50)
			WriteAll(value, true, buffer)
			buffer.WriteString(managesettings__98)
		}

		buffer.WriteString(manageusercompose__21)

		for _, provider := range providers {
			{
				var (
					value    = provider.Name()
					selected = user.Provider
					label    = provider.Name()
				)

				if value == selected {
					buffer.WriteString(managepostindex__85)
					WriteAll(value, true, buffer)
					buffer.WriteString(managerolecompose__98)
					WriteAll(label, true, buffer)
					buffer.WriteString(managepostindex__87)
				} else {
					buffer.WriteString(managepostindex__85)
					WriteAll(value, true, buffer)
					buffer.WriteString(commentlist__50)
					WriteAll(label, true, buffer)
					buffer.WriteString(managepostindex__87)
				}
			}

		}
		buffer.WriteString(manageusercompose__22)

		{
			var (
				name  = "provider_id"
				value = user.ProviderID
				label = "Provider ID"
			)

			buffer.WriteString(managepagecompose__85)
			WriteEscString(label, buffer)
			buffer.WriteString(managepagecompose__86)
			WriteEscString(name, buffer)
			buffer.WriteString(managepagecompose__87)
			WriteAll(value, true, buffer)
			buffer.WriteString(managepagecompose__88)
		}

		{
			var (
				name  = "provider_username"
				value = user.ProviderUsername
				label = "Provider username"
			)

			buffer.WriteString(managepagecompose__85)
			WriteEscString(label, buffer)
			buffer.WriteString(managepagecompose__86)
			WriteEscString(name, buffer)
			buffer.WriteString(managepagecompose__87)
			WriteAll(value, true, buffer)
			buffer.WriteString(managepagecompose__88)
		}

		{
			var (
				name  = "provider_avatar"
				value = user.ProviderAvatar
				label = "Provider avatar"
			)

			buffer.WriteString(managepagecompose__85)
			WriteEscString(label, buffer)
			buffer.WriteString(managepagecompose__86)
			WriteEscString(name, buffer)
			buffer.WriteString(managepagecompose__87)
			WriteAll(value, true, buffer)
			buffer.WriteString(managepagecompose__88)
		}

		buffer.WriteString(manageusercompose__23)

		{
			var (
				name  = "password"
				value = ""
				label = "Password"
			)

			buffer.WriteString(managepagecompose__85)
			WriteEscString(label, buffer)
			buffer.WriteString(managepagecompose__86)
			WriteEscString(name, buffer)
			buffer.WriteString(managepagecompose__87)
			WriteEscString(value, buffer)
			buffer.WriteString(managepagecompose__88)
		}

		buffer.WriteString(managerolecompose__21)

		{
			var (
				label = "New User"
				link  = "/manage/users/new"
			)

			buffer.WriteString(managepagecompose__93)
			WriteEscString(link, buffer)
			buffer.WriteString(managepagecompose__94)
			WriteEscString(label, buffer)
			buffer.WriteString(commentlist__132)
		}

		{
			var (
				name      = "active"
				condition = user.Active
				label     = "Active"
			)

			buffer.WriteString(managerolecompose__118)
			WriteEscString(label, buffer)
			buffer.WriteString(managerolecompose__119)
			if condition {
				buffer.WriteString(managerolecompose__121)
				WriteEscString(name, buffer)
				buffer.WriteString(managerolecompose__122)
			} else {
				buffer.WriteString(managerolecompose__121)
				WriteEscString(name, buffer)
				buffer.WriteString(commentlist__13)
			}
			buffer.WriteString(managerolecompose__120)

		}

		buffer.WriteString(managerolecompose__22)

		if ID > 1 {
			buffer.WriteString(manageusercompose__138)
			WriteInt(int64(ID), buffer)
			buffer.WriteString(managerolecompose__126)

		}
		buffer.WriteString(manageusercompose__26)

		{
			var (
				name     = "role_ids"
				roles    = roles
				selected = user.RoleIDs
			)

			buffer.WriteString(manageusercompose__140)
			for _, role := range roles {
				var inputId = fmt.Sprintf("role-%d", role.ID)
				if utils.SliceContains(selected, role.ID) {
					buffer.WriteString(manageusercompose__142)
					WriteEscString(inputId, buffer)
					buffer.WriteString(manageusercompose__143)
					WriteEscString(name, buffer)
					buffer.WriteString(managepagecompose__87)
					WriteAll(role.ID, true, buffer)
					buffer.WriteString(manageusercompose__145)
					WriteEscString(inputId, buffer)
					buffer.WriteString(manageusercompose__146)
					WriteAll(role.Name, true, buffer)
					buffer.WriteString(manageusercompose__147)

				} else {
					buffer.WriteString(manageusercompose__142)
					WriteEscString(inputId, buffer)
					buffer.WriteString(manageusercompose__143)
					WriteEscString(name, buffer)
					buffer.WriteString(managepagecompose__87)
					WriteAll(role.ID, true, buffer)
					buffer.WriteString(manageusercompose__145)
					WriteEscString(inputId, buffer)
					buffer.WriteString(manageusercompose__152)
					WriteAll(role.Name, true, buffer)
					buffer.WriteString(manageusercompose__147)

				}
			}
			buffer.WriteString(commentlist__22)
		}

		buffer.WriteString(manageusercompose__27)

		if user.AvatarImage != nil {
			buffer.WriteString(commentlist__40)
			WriteAll(user.AvatarImage.Url(), true, buffer)
			buffer.WriteString(commentlist__13)
		} else {
			buffer.WriteString(manageusercompose__156)
		}
		buffer.WriteString(manageusercompose__28)
		WriteAll(config.Setting("app_name"), true, buffer)
		buffer.WriteString(commentlist__25)

		if meta.User == nil || meta.User.ID == 0 {
			buffer.WriteString(commentlist__106)
			WriteAll(utils.Url("/login"), true, buffer)
			buffer.WriteString(commentlist__107)
			WriteAll(utils.Url("/register"), true, buffer)
			buffer.WriteString(commentlist__108)

		} else {
			{
				buffer.WriteString(commentlist__63)
				WriteAll(meta.User.AvatarElm("32", "32", false), false, buffer)
				buffer.WriteString(commentlist__64)
				WriteAll(meta.User.Url(), true, buffer)
				buffer.WriteString(commentlist__50)
				WriteAll(meta.User.Name(), true, buffer)
				buffer.WriteString(commentlist__66)
				WriteAll("@"+meta.User.Username, true, buffer)
				buffer.WriteString(commentlist__67)
				WriteAll(utils.Url("/posts/new"), true, buffer)
				buffer.WriteString(commentlist__68)
				WriteAll(utils.Url("/posts"), true, buffer)
				buffer.WriteString(commentlist__69)
				WriteAll(utils.Url("/comments"), true, buffer)
				buffer.WriteString(commentlist__70)
				WriteAll(utils.Url("/files"), true, buffer)
				buffer.WriteString(commentlist__71)
				WriteAll(utils.Url("/settings"), true, buffer)
				buffer.WriteString(commentlist__72)

			}

			if meta.User.IsRoot() {
				{
					buffer.WriteString(commentlist__119)
					WriteAll(utils.Url("/manage"), true, buffer)
					buffer.WriteString(commentlist__120)
					WriteAll(utils.Url("/manage/topics"), true, buffer)
					buffer.WriteString(commentlist__121)
					WriteAll(utils.Url("/manage/posts"), true, buffer)
					buffer.WriteString(commentlist__122)
					WriteAll(utils.Url("/manage/pages"), true, buffer)
					buffer.WriteString(commentlist__123)
					WriteAll(utils.Url("/manage/roles"), true, buffer)
					buffer.WriteString(commentlist__124)
					WriteAll(utils.Url("/manage/users"), true, buffer)
					buffer.WriteString(commentlist__125)
					WriteAll(utils.Url("/manage/comments"), true, buffer)
					buffer.WriteString(commentlist__126)
					WriteAll(utils.Url("/manage/files"), true, buffer)
					buffer.WriteString(commentlist__127)
					WriteAll(utils.Url("/manage/settings"), true, buffer)
					buffer.WriteString(commentlist__72)

				}

			}
		}
		buffer.WriteString(commentlist__26)

		for _, topic := range cache.Topics {
			buffer.WriteString(commentlist__106)
			WriteAll(topic.Url(), true, buffer)
			buffer.WriteString(commentlist__49)
			WriteAll(topic.Name, true, buffer)
			buffer.WriteString(commentlist__50)
			WriteAll("#"+topic.Name, true, buffer)
			buffer.WriteString(commentlist__132)
		}
		buffer.WriteString(commentlist__27)
		WriteAll(config.Setting("footer_content"), false, buffer)
		buffer.WriteString(commentlist__28)
		WriteAll(config.Setting("inject_footer"), false, buffer)
		WriteAll(asset.JsFile("js/layout.js"), false, buffer)
		WriteAll(asset.JsFile("js/main.js"), false, buffer)
		buffer.WriteString(manageusercompose__33)

	}
}
