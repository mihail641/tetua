// Code generated by "jade.go"; DO NOT EDIT.

package views

import (
	"bufio"

	"github.com/ngocphuongnb/tetua/app/asset"
	"github.com/ngocphuongnb/tetua/app/cache"
	"github.com/ngocphuongnb/tetua/app/config"
	"github.com/ngocphuongnb/tetua/app/entities"
	"github.com/ngocphuongnb/tetua/app/utils"
)

const (
	managetopiccompose__0   = `<!DOCTYPE html><html lang="en">`
	managetopiccompose__1   = `<head><meta charset="utf-8"/><meta name="viewport" content="width=device-width, initial-scale=1.0, viewport-fit=cover"/><title>`
	managetopiccompose__2   = `</title><meta name="keywords" content="software development, devloper community"/><link rel="canonical" href="`
	managetopiccompose__3   = `"/><meta property="og:type" content="`
	managetopiccompose__4   = `"/><meta property="og:url" content="`
	managetopiccompose__5   = `"/><meta property="og:title" content="`
	managetopiccompose__6   = `"/><meta property="og:site_name" content="`
	managetopiccompose__7   = `"/><meta name="twitter:site" content="`
	managetopiccompose__8   = `"/><meta name="twitter:title" content="`
	managetopiccompose__9   = `"/><meta name="twitter:card" content="summary_large_image"/><meta name="apple-mobile-web-app-title" content="`
	managetopiccompose__10  = `"/><meta name="application-name" content="`
	managetopiccompose__11  = `"/><link rel="alternate" type="application/rss+xml" title="`
	managetopiccompose__12  = `" href="`
	managetopiccompose__13  = `"/>`
	managetopiccompose__14  = `</head><body><header><nav class="main container"><a class="logo" href="`
	managetopiccompose__15  = `" title="Home">`
	managetopiccompose__16  = `</a><form class="search-form" method="get" action="/search" accept-charset="UTF-8"><input class="search-input" type="text" name="q" placeholder="Search..." autocomplete="off" value="`
	managetopiccompose__17  = `"/><button class="search-btn" type="submit" aria-label="Search"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M9.5,3A6.5,6.5 0 0,1 16,9.5C16,11.11 15.41,12.59 14.44,13.73L14.71,14H15.5L20.5,19L19,20.5L14,15.5V14.71L13.73,14.44C12.59,15.41 11.11,16 9.5,16A6.5,6.5 0 0,1 3,9.5A6.5,6.5 0 0,1 9.5,3M9.5,5C7,5 5,7 5,9.5C5,12 7,14 9.5,14C12,14 14,12 14,9.5C14,7 12,5 9.5,5Z"></path></svg></button></form><ul><li class="search-mobile"><a href="`
	managetopiccompose__18  = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M9.5,3A6.5,6.5 0 0,1 16,9.5C16,11.11 15.41,12.59 14.44,13.73L14.71,14H15.5L20.5,19L19,20.5L14,15.5V14.71L13.73,14.44C12.59,15.41 11.11,16 9.5,16A6.5,6.5 0 0,1 3,9.5A6.5,6.5 0 0,1 9.5,3M9.5,5C7,5 5,7 5,9.5C5,12 7,14 9.5,14C12,14 14,12 14,9.5C14,7 12,5 9.5,5Z"></path></svg></a></li>`
	managetopiccompose__19  = `</ul><label class="menu-trigger"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M3,6H21V8H3V6M3,11H21V13H3V11M3,16H21V18H3V16Z"></path></svg></label></nav></header><div class="wrapper"><div class="container"><form method="POST"><div class="layout"><div class="left"><div class="box fixed-sidebar">`
	managetopiccompose__20  = `</div></div><div class="main"><div class="box">`
	managetopiccompose__21  = `</div></div><div class="right"><div class="box fixed-sidebar"><div class="flex">`
	managetopiccompose__22  = `</div><div><label>Parent topic</label>`
	managetopiccompose__23  = `</div><div class="save-actions"><button>Save</button>`
	managetopiccompose__24  = `</div>`
	managetopiccompose__25  = `</div></div></div></form></div><div class="mobile-menu"><div class="menu-head">`
	managetopiccompose__26  = `<label class="menu-trigger menu-close"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M19,6.41L17.59,5L12,10.59L6.41,5L5,6.41L10.59,12L5,17.59L6.41,19L12,13.41L17.59,19L19,17.59L13.41,12L19,6.41Z"></path></svg></label></div>`
	managetopiccompose__27  = `<strong>Topics</strong><div class="menu-topics">`
	managetopiccompose__28  = `</div></div></div><div class="overlay menu-trigger"></div><footer><div class="container"><div>`
	managetopiccompose__29  = `</div><p>Proudly powered by <a href="https://tetua.net" title="Tetua - CMS Platform for Blogging">Tetua</a></p></div></footer>`
	managetopiccompose__30  = `<script src="/static/js/manage.js"></script><script>listenDeleteNodeEvents('topic','/manage/topics', '/manage/topics')</script></body></html>`
	managetopiccompose__31  = `<link rel="icon" type="image/png" href="`
	managetopiccompose__32  = `"/><link rel="apple-touch-icon" href="`
	managetopiccompose__34  = `<meta name="description" content="`
	managetopiccompose__35  = `"/><meta property="og:description" content="`
	managetopiccompose__36  = `"/><meta name="twitter:description" content="`
	managetopiccompose__38  = `<meta property="og:image" content="`
	managetopiccompose__39  = `"/><meta name="twitter:image:src" content="`
	managetopiccompose__41  = `<img src="`
	managetopiccompose__42  = `" alt="`
	managetopiccompose__44  = `<svg viewBox="0 0 24 24"><path fill="#164e63" d="M11,6.5V9.33L8.33,12L11,14.67V17.5L5.5,12M13,6.43L18.57,12L13,17.57V14.74L15.74,12L13,9.26M5,3C3.89,3 3,3.9 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3H5Z"></path></svg>`
	managetopiccompose__45  = `<li><a href="`
	managetopiccompose__46  = `">Login</a></li><li><a href="`
	managetopiccompose__47  = `">Register</a></li>`
	managetopiccompose__49  = `">New</a></li><li><div class="user-menu"><a href="`
	managetopiccompose__50  = `" title="`
	managetopiccompose__51  = `">`
	managetopiccompose__52  = `</a><svg viewBox="0 0 24 24"><path fill="currentColor" d="M7.41,8.58L12,13.17L16.59,8.58L18,10L12,16L6,10L7.41,8.58Z"></path></svg><ul class="sub">`
	managetopiccompose__54  = `">Profile</a></li><li><a href="`
	managetopiccompose__55  = `">Posts</a></li><li><a href="`
	managetopiccompose__56  = `">Setting</a></li><li><a href="`
	managetopiccompose__57  = `">Logout</a></li></ul></div></li>`
	managetopiccompose__58  = `<img class="avatar" src="`
	managetopiccompose__61  = `<span class="avatar none"></span>`
	managetopiccompose__63  = `">Manage</a></li>`
	managetopiccompose__64  = `<h2 class="header"><a href="`
	managetopiccompose__65  = `">Manage</a></h2><ul class="manage-features"><li><a href="`
	managetopiccompose__66  = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M9,1H19A2,2 0 0,1 21,3V19L19,18.13V3H7A2,2 0 0,1 9,1M15,20V7H5V20L10,17.82L15,20M15,5C16.11,5 17,5.9 17,7V23L10,20L3,23V7A2,2 0 0,1 5,5H15Z"></path></svg>Topics</a></li><li><a href="`
	managetopiccompose__67  = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M20 5L20 19L4 19L4 5H20M20 3H4C2.89 3 2 3.89 2 5V19C2 20.11 2.89 21 4 21H20C21.11 21 22 20.11 22 19V5C22 3.89 21.11 3 20 3M18 15H6V17H18V15M10 7H6V13H10V7M12 9H18V7H12V9M18 11H12V13H18V11Z"></path></svg>Posts</a></li><li><a href="`
	managetopiccompose__68  = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z"></path></svg>Pages</a></li><li><a href="`
	managetopiccompose__69  = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M17 14.4C17.6 14.4 18.1 14.9 18.1 15.5S17.6 16.6 17 16.6 15.9 16.1 15.9 15.5 16.4 14.4 17 14.4M17 17.5C16.3 17.5 14.8 17.9 14.8 18.6C15.3 19.3 16.1 19.8 17 19.8S18.7 19.3 19.2 18.6C19.2 17.9 17.7 17.5 17 17.5M18 11.1V6.3L10.5 3L3 6.3V11.2C3 15.7 6.2 20 10.5 21C11.1 20.9 11.6 20.7 12.1 20.5C13.2 22 15 23 17 23C20.3 23 23 20.3 23 17C23 14 20.8 11.6 18 11.1M11 17C11 17.6 11.1 18.1 11.2 18.6C11 18.7 10.7 18.8 10.5 18.9C7.3 17.9 5 14.7 5 11.2V7.6L10.5 5.2L16 7.6V11.1C13.2 11.6 11 14 11 17M17 21C14.8 21 13 19.2 13 17S14.8 13 17 13 21 14.8 21 17 19.2 21 17 21Z"></path></svg>Roles</a></li><li><a href="`
	managetopiccompose__70  = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M13.07 10.41A5 5 0 0 0 13.07 4.59A3.39 3.39 0 0 1 15 4A3.5 3.5 0 0 1 15 11A3.39 3.39 0 0 1 13.07 10.41M5.5 7.5A3.5 3.5 0 1 1 9 11A3.5 3.5 0 0 1 5.5 7.5M7.5 7.5A1.5 1.5 0 1 0 9 6A1.5 1.5 0 0 0 7.5 7.5M16 17V19H2V17S2 13 9 13 16 17 16 17M14 17C13.86 16.22 12.67 15 9 15S4.07 16.31 4 17M15.95 13A5.32 5.32 0 0 1 18 17V19H22V17S22 13.37 15.94 13Z"></path></svg>Users</a></li><li><a href="`
	managetopiccompose__71  = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M20 2H4C2.9 2 2 2.9 2 4V22L6 18H20C21.1 18 22 17.1 22 16V4C22 2.9 21.1 2 20 2M20 16H5.2L4 17.2V4H20V16Z"></path></svg>Comments</a></li><li><a href="`
	managetopiccompose__72  = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M21,17H7V3H21M21,1H7A2,2 0 0,0 5,3V17A2,2 0 0,0 7,19H21A2,2 0 0,0 23,17V3A2,2 0 0,0 21,1M3,5H1V21A2,2 0 0,0 3,23H19V21H3M15.96,10.29L13.21,13.83L11.25,11.47L8.5,15H19.5L15.96,10.29Z"></path></svg>Files</a></li><li><a href="`
	managetopiccompose__73  = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M12,8A4,4 0 0,1 16,12A4,4 0 0,1 12,16A4,4 0 0,1 8,12A4,4 0 0,1 12,8M12,10A2,2 0 0,0 10,12A2,2 0 0,0 12,14A2,2 0 0,0 14,12A2,2 0 0,0 12,10M10,22C9.75,22 9.54,21.82 9.5,21.58L9.13,18.93C8.5,18.68 7.96,18.34 7.44,17.94L4.95,18.95C4.73,19.03 4.46,18.95 4.34,18.73L2.34,15.27C2.21,15.05 2.27,14.78 2.46,14.63L4.57,12.97L4.5,12L4.57,11L2.46,9.37C2.27,9.22 2.21,8.95 2.34,8.73L4.34,5.27C4.46,5.05 4.73,4.96 4.95,5.05L7.44,6.05C7.96,5.66 8.5,5.32 9.13,5.07L9.5,2.42C9.54,2.18 9.75,2 10,2H14C14.25,2 14.46,2.18 14.5,2.42L14.87,5.07C15.5,5.32 16.04,5.66 16.56,6.05L19.05,5.05C19.27,4.96 19.54,5.05 19.66,5.27L21.66,8.73C21.79,8.95 21.73,9.22 21.54,9.37L19.43,11L19.5,12L19.43,13L21.54,14.63C21.73,14.78 21.79,15.05 21.66,15.27L19.66,18.73C19.54,18.95 19.27,19.04 19.05,18.95L16.56,17.95C16.04,18.34 15.5,18.68 14.87,18.93L14.5,21.58C14.46,21.82 14.25,22 14,22H10M11.25,4L10.88,6.61C9.68,6.86 8.62,7.5 7.85,8.39L5.44,7.35L4.69,8.65L6.8,10.2C6.4,11.37 6.4,12.64 6.8,13.8L4.68,15.36L5.43,16.66L7.86,15.62C8.63,16.5 9.68,17.14 10.87,17.38L11.24,20H12.76L13.13,17.39C14.32,17.14 15.37,16.5 16.14,15.62L18.57,16.66L19.32,15.36L17.2,13.81C17.6,12.64 17.6,11.37 17.2,10.2L19.31,8.65L18.56,7.35L16.15,8.39C15.38,7.5 14.32,6.86 13.12,6.62L12.75,4H11.25Z"></path></svg>Settings</a></li></ul>`
	managetopiccompose__74  = `<h1>`
	managetopiccompose__75  = `</h1>`
	managetopiccompose__76  = `<h1>Create new topic</h1>`
	managetopiccompose__77  = `<ul class="messages">`
	managetopiccompose__78  = `</ul>`
	managetopiccompose__79  = `<li class="`
	managetopiccompose__81  = `</li>`
	managetopiccompose__82  = `<p><label>`
	managetopiccompose__83  = `</label><input name="`
	managetopiccompose__84  = `" value="`
	managetopiccompose__85  = `"/></p>`
	managetopiccompose__87  = `</label><textarea name="`
	managetopiccompose__89  = `</textarea></p>`
	managetopiccompose__90  = `<a class="link-icon" href="`
	managetopiccompose__91  = `"><svg style="width:24px;height:24px" viewBox="0 0 24 24"><path fill="currentColor" d="M17,13H13V17H11V13H7V11H11V7H13V11H17M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z"></path></svg>`
	managetopiccompose__92  = `</a>`
	managetopiccompose__93  = `<select name="`
	managetopiccompose__94  = `"><option value="">--</option>`
	managetopiccompose__95  = `</select>`
	managetopiccompose__96  = `<option value="`
	managetopiccompose__97  = `" selected="selected">`
	managetopiccompose__98  = `</option>`
	managetopiccompose__102 = `<button class="danger delete-topic" data-id="`
	managetopiccompose__103 = `">Delete</button>`
	managetopiccompose__104 = `<strong style="margin-bottom:5px;display:block">Writing a Great Post</strong><ul><li>Using markdown shortcut to compose your post.</li><li>Use the Markdown switch button to toggle between rich text and plain markdown editor mode.</li><li>Select the topics that represent for the post content.</li></ul>`
	managetopiccompose__105 = `<a href="`
	managetopiccompose__106 = `">Login</a><a href="`
	managetopiccompose__107 = `">Register</a>`
	managetopiccompose__108 = `<div class="meta flex">`
	managetopiccompose__109 = `<div><a class="author" href="`
	managetopiccompose__111 = `</a><div class="stat flex"><span>`
	managetopiccompose__112 = `</span></div></div></div><ul class="manage-features"><li><a href="`
	managetopiccompose__113 = `"><svg style="width:24px;height:24px" viewBox="0 0 24 24"><path fill="currentColor" d="M12,20C7.59,20 4,16.41 4,12C4,7.59 7.59,4 12,4C16.41,4 20,7.59 20,12C20,16.41 16.41,20 12,20M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M13,7H11V11H7V13H11V17H13V13H17V11H13V7Z"></path></svg>New post</a></li><li><a href="`
	managetopiccompose__114 = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M20 5L20 19L4 19L4 5H20M20 3H4C2.89 3 2 3.89 2 5V19C2 20.11 2.89 21 4 21H20C21.11 21 22 20.11 22 19V5C22 3.89 21.11 3 20 3M18 15H6V17H18V15M10 7H6V13H10V7M12 9H18V7H12V9M18 11H12V13H18V11Z"></path></svg>My Posts</a></li><li><a href="`
	managetopiccompose__115 = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M20 2H4C2.9 2 2 2.9 2 4V22L6 18H20C21.1 18 22 17.1 22 16V4C22 2.9 21.1 2 20 2M20 16H5.2L4 17.2V4H20V16Z"></path></svg>My Comments</a></li><li><a href="`
	managetopiccompose__116 = `"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M21,17H7V3H21M21,1H7A2,2 0 0,0 5,3V17A2,2 0 0,0 7,19H21A2,2 0 0,0 23,17V3A2,2 0 0,0 21,1M3,5H1V21A2,2 0 0,0 3,23H19V21H3M15.96,10.29L13.21,13.83L11.25,11.47L8.5,15H19.5L15.96,10.29Z"></path></svg>My Files</a></li><li><a href="`
)

func ManageTopicCompose(topics []*entities.Topic, topic *entities.TopicMutation) func(meta *entities.Meta, wr *bufio.Writer) {
	return func(meta *entities.Meta, wr *bufio.Writer) {
		buffer := &WriterAsBuffer{wr}

		buffer.WriteString(managetopiccompose__0)

		var title = meta.GetTitle()
		var appName = config.Setting("app_name")
		var appLogo = config.Setting("app_logo")
		buffer.WriteString(managetopiccompose__1)
		WriteAll(title, true, buffer)
		buffer.WriteString(managetopiccompose__2)
		WriteAll(meta.Canonical, true, buffer)
		buffer.WriteString(managetopiccompose__3)
		WriteAll(meta.Type, true, buffer)
		buffer.WriteString(managetopiccompose__4)
		WriteAll(meta.Canonical, true, buffer)
		buffer.WriteString(managetopiccompose__5)
		WriteAll(title, true, buffer)
		buffer.WriteString(managetopiccompose__6)
		WriteAll(appName, true, buffer)
		buffer.WriteString(managetopiccompose__7)
		WriteAll(config.Setting("twitter_site"), true, buffer)
		buffer.WriteString(managetopiccompose__8)
		WriteAll(title, true, buffer)
		buffer.WriteString(managetopiccompose__9)
		WriteAll(appName, true, buffer)
		buffer.WriteString(managetopiccompose__10)
		WriteAll(appName, true, buffer)
		buffer.WriteString(managetopiccompose__11)
		WriteAll(appName+" Feed", true, buffer)
		buffer.WriteString(managetopiccompose__12)
		WriteAll(utils.Url("/feed"), true, buffer)
		buffer.WriteString(managetopiccompose__13)
		if appLogo != "" {
			buffer.WriteString(managetopiccompose__31)
			WriteAll(appLogo, true, buffer)
			buffer.WriteString(managetopiccompose__32)
			WriteAll(appLogo, true, buffer)
			buffer.WriteString(managetopiccompose__13)
		}
		if meta.Description != "" {
			buffer.WriteString(managetopiccompose__34)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(managetopiccompose__35)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(managetopiccompose__36)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(managetopiccompose__13)
		}
		if meta.Image != "" {
			buffer.WriteString(managetopiccompose__38)
			WriteAll(meta.Image, true, buffer)
			buffer.WriteString(managetopiccompose__39)
			WriteAll(meta.Image, true, buffer)
			buffer.WriteString(managetopiccompose__13)
		}
		WriteAll(asset.CssFile("css/light.min.css"), false, buffer)
		WriteAll(asset.CssFile("css/style.css"), false, buffer)
		WriteAll(config.Setting("inject_header"), false, buffer)
		buffer.WriteString(managetopiccompose__14)
		WriteAll(utils.Url(""), true, buffer)
		buffer.WriteString(managetopiccompose__15)
		var logoUrl = config.Setting("app_logo")
		if logoUrl != "" {
			buffer.WriteString(managetopiccompose__41)
			WriteAll(logoUrl, true, buffer)
			buffer.WriteString(managetopiccompose__42)
			WriteAll(config.Setting("app_name"), true, buffer)
			buffer.WriteString(managetopiccompose__13)
		} else {
			buffer.WriteString(managetopiccompose__44)

		}
		buffer.WriteString(managetopiccompose__16)
		WriteAll(meta.Query, true, buffer)
		buffer.WriteString(managetopiccompose__17)
		WriteAll(utils.Url("/search"), true, buffer)
		buffer.WriteString(managetopiccompose__18)

		if meta.User == nil || meta.User.ID == 0 {
			buffer.WriteString(managetopiccompose__45)
			WriteAll(utils.Url("/login"), true, buffer)
			buffer.WriteString(managetopiccompose__46)
			WriteAll(utils.Url("/register"), true, buffer)
			buffer.WriteString(managetopiccompose__47)

		} else {
			buffer.WriteString(managetopiccompose__45)
			WriteAll(utils.Url("/posts/new"), true, buffer)
			buffer.WriteString(managetopiccompose__49)
			WriteAll(meta.User.Url(), true, buffer)
			buffer.WriteString(managetopiccompose__50)
			WriteAll(meta.User.Username, true, buffer)
			buffer.WriteString(managetopiccompose__51)
			if meta.User.AvatarImageUrl != "" {
				buffer.WriteString(managetopiccompose__58)
				WriteAll(meta.User.AvatarImageUrl, true, buffer)
				buffer.WriteString(managetopiccompose__42)
				WriteAll(meta.User.Username, true, buffer)
				buffer.WriteString(managetopiccompose__13)
			} else {
				buffer.WriteString(managetopiccompose__61)

			}
			buffer.WriteString(managetopiccompose__52)

			if meta.User != nil && meta.User.IsRoot() {
				buffer.WriteString(managetopiccompose__45)
				WriteAll(utils.Url("/manage"), true, buffer)
				buffer.WriteString(managetopiccompose__63)

			}
			buffer.WriteString(managetopiccompose__45)
			WriteAll(meta.User.Url(), true, buffer)
			buffer.WriteString(managetopiccompose__54)
			WriteAll(utils.Url("/posts"), true, buffer)
			buffer.WriteString(managetopiccompose__55)
			WriteAll(utils.Url("/settings"), true, buffer)
			buffer.WriteString(managetopiccompose__56)
			WriteAll(utils.Url("/logout"), true, buffer)
			buffer.WriteString(managetopiccompose__57)

		}
		buffer.WriteString(managetopiccompose__19)

		{
			buffer.WriteString(managetopiccompose__64)
			WriteAll(utils.Url("/manage"), true, buffer)
			buffer.WriteString(managetopiccompose__65)
			WriteAll(utils.Url("/manage/topics"), true, buffer)
			buffer.WriteString(managetopiccompose__66)
			WriteAll(utils.Url("/manage/posts"), true, buffer)
			buffer.WriteString(managetopiccompose__67)
			WriteAll(utils.Url("/manage/pages"), true, buffer)
			buffer.WriteString(managetopiccompose__68)
			WriteAll(utils.Url("/manage/roles"), true, buffer)
			buffer.WriteString(managetopiccompose__69)
			WriteAll(utils.Url("/manage/users"), true, buffer)
			buffer.WriteString(managetopiccompose__70)
			WriteAll(utils.Url("/manage/comments"), true, buffer)
			buffer.WriteString(managetopiccompose__71)
			WriteAll(utils.Url("/manage/files"), true, buffer)
			buffer.WriteString(managetopiccompose__72)
			WriteAll(utils.Url("/manage/settings"), true, buffer)
			buffer.WriteString(managetopiccompose__73)

		}

		buffer.WriteString(managetopiccompose__20)

		if topic.ID > 0 {
			buffer.WriteString(managetopiccompose__74)
			WriteAll("Editing topic: "+topic.Name, true, buffer)
			buffer.WriteString(managetopiccompose__75)
		} else {
			buffer.WriteString(managetopiccompose__76)

		}
		{
			var (
				msgs = meta.Messages
			)

			if msgs.Length() > 0 {
				buffer.WriteString(managetopiccompose__77)
				var messages = msgs.Get()
				for _, msg := range messages {
					buffer.WriteString(managetopiccompose__79)
					WriteAll(msg.Type, true, buffer)
					buffer.WriteString(managetopiccompose__51)
					WriteAll(msg.Message, true, buffer)
					buffer.WriteString(managetopiccompose__81)
				}
				buffer.WriteString(managetopiccompose__78)
			}
		}

		{
			var (
				name  = "name"
				value = topic.Name
				label = "Name"
			)

			buffer.WriteString(managetopiccompose__82)
			WriteEscString(label, buffer)
			buffer.WriteString(managetopiccompose__83)
			WriteEscString(name, buffer)
			buffer.WriteString(managetopiccompose__84)
			WriteAll(value, true, buffer)
			buffer.WriteString(managetopiccompose__85)
		}

		{
			var (
				name  = "content"
				value = topic.Content
				label = "Description"
			)

			buffer.WriteString(managetopiccompose__82)
			WriteEscString(label, buffer)
			buffer.WriteString(managetopiccompose__87)
			WriteEscString(name, buffer)
			buffer.WriteString(managetopiccompose__51)
			WriteAll(value, true, buffer)
			buffer.WriteString(managetopiccompose__89)
		}

		buffer.WriteString(managetopiccompose__21)

		{
			var (
				label = "New Topic"
				link  = "/manage/topics/new"
			)

			buffer.WriteString(managetopiccompose__90)
			WriteEscString(link, buffer)
			buffer.WriteString(managetopiccompose__91)
			WriteEscString(label, buffer)
			buffer.WriteString(managetopiccompose__92)
		}

		buffer.WriteString(managetopiccompose__22)

		{
			var (
				name    = "parent_id"
				topics  = topics
				current = topic
			)

			buffer.WriteString(managetopiccompose__93)
			WriteEscString(name, buffer)
			buffer.WriteString(managetopiccompose__94)

			for _, t := range topics {
				if t.ID == current.ParentID {
					buffer.WriteString(managetopiccompose__96)
					WriteAll(t.ID, true, buffer)
					buffer.WriteString(managetopiccompose__97)
					WriteAll(t.Name, true, buffer)
					buffer.WriteString(managetopiccompose__98)
				} else {
					buffer.WriteString(managetopiccompose__96)
					WriteAll(t.ID, true, buffer)
					buffer.WriteString(managetopiccompose__51)
					WriteAll(t.Name, true, buffer)
					buffer.WriteString(managetopiccompose__98)
				}
			}
			buffer.WriteString(managetopiccompose__95)
		}

		buffer.WriteString(managetopiccompose__23)

		if topic.ID > 0 {
			buffer.WriteString(managetopiccompose__102)
			WriteAll(topic.ID, true, buffer)
			buffer.WriteString(managetopiccompose__103)

		}
		buffer.WriteString(managetopiccompose__24)
		{
			buffer.WriteString(managetopiccompose__104)

		}

		buffer.WriteString(managetopiccompose__25)
		WriteAll(config.Setting("app_name"), true, buffer)
		buffer.WriteString(managetopiccompose__26)

		if meta.User == nil || meta.User.ID == 0 {
			buffer.WriteString(managetopiccompose__105)
			WriteAll(utils.Url("/login"), true, buffer)
			buffer.WriteString(managetopiccompose__106)
			WriteAll(utils.Url("/register"), true, buffer)
			buffer.WriteString(managetopiccompose__107)

		} else {
			{
				buffer.WriteString(managetopiccompose__108)
				WriteAll(meta.User.AvatarElm("32", "32", false), false, buffer)
				buffer.WriteString(managetopiccompose__109)
				WriteAll(meta.User.Url(), true, buffer)
				buffer.WriteString(managetopiccompose__51)
				WriteAll(meta.User.Name(), true, buffer)
				buffer.WriteString(managetopiccompose__111)
				WriteAll("@"+meta.User.Username, true, buffer)
				buffer.WriteString(managetopiccompose__112)
				WriteAll(utils.Url("/posts/new"), true, buffer)
				buffer.WriteString(managetopiccompose__113)
				WriteAll(utils.Url("/posts"), true, buffer)
				buffer.WriteString(managetopiccompose__114)
				WriteAll(utils.Url("/comments"), true, buffer)
				buffer.WriteString(managetopiccompose__115)
				WriteAll(utils.Url("/files"), true, buffer)
				buffer.WriteString(managetopiccompose__116)
				WriteAll(utils.Url("/settings"), true, buffer)
				buffer.WriteString(managetopiccompose__73)

			}

			if meta.User.IsRoot() {
				{
					buffer.WriteString(managetopiccompose__64)
					WriteAll(utils.Url("/manage"), true, buffer)
					buffer.WriteString(managetopiccompose__65)
					WriteAll(utils.Url("/manage/topics"), true, buffer)
					buffer.WriteString(managetopiccompose__66)
					WriteAll(utils.Url("/manage/posts"), true, buffer)
					buffer.WriteString(managetopiccompose__67)
					WriteAll(utils.Url("/manage/pages"), true, buffer)
					buffer.WriteString(managetopiccompose__68)
					WriteAll(utils.Url("/manage/roles"), true, buffer)
					buffer.WriteString(managetopiccompose__69)
					WriteAll(utils.Url("/manage/users"), true, buffer)
					buffer.WriteString(managetopiccompose__70)
					WriteAll(utils.Url("/manage/comments"), true, buffer)
					buffer.WriteString(managetopiccompose__71)
					WriteAll(utils.Url("/manage/files"), true, buffer)
					buffer.WriteString(managetopiccompose__72)
					WriteAll(utils.Url("/manage/settings"), true, buffer)
					buffer.WriteString(managetopiccompose__73)

				}

			}
		}
		buffer.WriteString(managetopiccompose__27)

		for _, topic := range cache.Topics {
			buffer.WriteString(managetopiccompose__105)
			WriteAll(topic.Url(), true, buffer)
			buffer.WriteString(managetopiccompose__50)
			WriteAll(topic.Name, true, buffer)
			buffer.WriteString(managetopiccompose__51)
			WriteAll("#"+topic.Name, true, buffer)
			buffer.WriteString(managetopiccompose__92)
		}
		buffer.WriteString(managetopiccompose__28)
		WriteAll(config.Setting("footer_content"), false, buffer)
		buffer.WriteString(managetopiccompose__29)
		WriteAll(config.Setting("inject_footer"), false, buffer)
		WriteAll(asset.JsFile("js/layout.js"), false, buffer)
		WriteAll(asset.JsFile("js/main.js"), false, buffer)
		buffer.WriteString(managetopiccompose__30)

	}
}
