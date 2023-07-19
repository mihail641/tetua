// Code generated by "jade.go"; DO NOT EDIT.

package views

import (
	"bufio"
	"fmt"

	"github.com/ngocphuongnb/tetua/app/asset"
	"github.com/ngocphuongnb/tetua/app/cache"
	"github.com/ngocphuongnb/tetua/app/config"
	"github.com/ngocphuongnb/tetua/app/entities"
	"github.com/ngocphuongnb/tetua/app/utils"
)

const (
	topicindex__14  = `<link rel="alternate" type="application/rss+xml" title="`
	topicindex__22  = `</ul><label class="menu-trigger"><svg viewBox="0 0 24 24"><path fill="currentColor" d="M3,6H21V8H3V6M3,11H21V13H3V11M3,16H21V18H3V16Z"></path></svg></label></nav></header><div class="wrapper"><div class="container"><div class="box page-desc"><h1>`
	topicindex__23  = `</h1>`
	topicindex__24  = `</div><div class="layout"><div class="left"><div class="box fixed-sidebar"><h2 class="head">Topics</h2>`
	topicindex__25  = `</div></div><main class="main">`
	topicindex__26  = `<div class="article-list">`
	topicindex__29  = `</ul></main><div class="right"><div class="box fixed-sidebar"><h2>Top posts</h2><div class="posts-list">`
	topicindex__30  = `</div></div></div></div></div><div class="mobile-menu"><div class="menu-head">`
	topicindex__35  = `</body></html>`
	topicindex__69  = `<div class="topics">`
	topicindex__80  = `<article class="box"><a class="overlay" href="`
	topicindex__84  = `<div class="box-content">`
	topicindex__85  = `<div class="info"><h3><a href="`
	topicindex__88  = `</a></h3><div class="tags">`
	topicindex__89  = `</div></div></div></article>`
	topicindex__90  = `<a class="bg" href="`
	topicindex__92  = `" style="`
	topicindex__99  = `</a><div class="stat flex"><time datetime="`
	topicindex__100 = `" class="date">`
	topicindex__101 = `</time><span class="views">`
	topicindex__102 = `</span><span class="comment">`
	topicindex__103 = `</span></div></div></div>`
	topicindex__112 = `<article><h4>`
	topicindex__116 = `</a></h4><div class="tags">`
	topicindex__117 = `</div></article>`
	topicindex__118 = `<span class="pos">`
	topicindex__119 = `</span>`
)

func TopicView(topics []*entities.Topic, topic *entities.Topic, paginate *entities.Paginate[entities.Post], topPosts []*entities.Post) func(meta *entities.Meta, wr *bufio.Writer) {
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
		buffer.WriteString(topicindex__14)
		WriteAll(topic.Name+" Feed", true, buffer)
		buffer.WriteString(commentlist__12)
		WriteAll(topic.FeedUrl(), true, buffer)
		buffer.WriteString(commentlist__13)
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
		buffer.WriteString(topicindex__22)
		WriteAll(topic.Name, true, buffer)
		buffer.WriteString(topicindex__23)
		WriteAll(topic.ContentHTML, false, buffer)
		buffer.WriteString(topicindex__24)

		{
			var (
				topics = topics
			)

			buffer.WriteString(topicindex__69)
			for _, topic := range topics {
				buffer.WriteString(commentlist__106)
				WriteAll(topic.Url(), true, buffer)
				buffer.WriteString(commentlist__49)
				WriteAll(topic.Name, true, buffer)
				buffer.WriteString(commentlist__50)
				WriteAll("# "+topic.Name, true, buffer)
				buffer.WriteString(commentlist__132)
			}
			buffer.WriteString(commentlist__22)
		}

		buffer.WriteString(topicindex__25)

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

		buffer.WriteString(topicindex__26)
		for _, post := range paginate.Data {
			{
				var (
					post = post
				)

				var postUrl = post.Url()
				var bgStyle = ""
				if post.FeaturedImage != nil {
					bgStyle = fmt.Sprintf("background-image:url('%s')", post.FeaturedImage.Url())
				}
				buffer.WriteString(topicindex__80)
				WriteAll(postUrl, true, buffer)
				buffer.WriteString(commentlist__49)
				WriteAll(post.Name, true, buffer)
				buffer.WriteString(commentlist__50)
				WriteAll(post.Name, true, buffer)
				buffer.WriteString(commentlist__132)
				if post.FeaturedImage != nil && post.FeaturedImage.ID > 0 {
					buffer.WriteString(topicindex__90)
					WriteAll(postUrl, true, buffer)
					buffer.WriteString(commentlist__49)
					WriteAll(post.Name, true, buffer)
					buffer.WriteString(topicindex__92)
					WriteEscString(bgStyle, buffer)
					buffer.WriteString(commentlist__50)
					WriteAll(post.Name, true, buffer)
					buffer.WriteString(commentlist__132)
				}
				buffer.WriteString(topicindex__84)
				{
					buffer.WriteString(commentlist__63)
					WriteAll(post.User.AvatarElm("32", "32", false), false, buffer)
					buffer.WriteString(commentlist__64)
					WriteAll(post.User.Url(), true, buffer)
					buffer.WriteString(commentlist__49)
					WriteAll(post.User.Name(), true, buffer)
					buffer.WriteString(commentlist__50)
					WriteAll(post.User.Name(), true, buffer)
					buffer.WriteString(topicindex__99)
					WriteAll(post.CreatedAt.Format("2006-01-02T15:04:05-0700"), true, buffer)
					buffer.WriteString(topicindex__100)
					WriteAll(post.CreatedAt.Format("January 2, 2006"), true, buffer)
					buffer.WriteString(topicindex__101)
					WriteEscString(fmt.Sprintf("%d views", post.ViewCount), buffer)
					buffer.WriteString(topicindex__102)
					WriteEscString(fmt.Sprintf("%d comments", post.CommentCount), buffer)
					buffer.WriteString(topicindex__103)

				}

				buffer.WriteString(topicindex__85)
				WriteAll(postUrl, true, buffer)
				buffer.WriteString(commentlist__49)
				WriteAll(post.Name, true, buffer)
				buffer.WriteString(commentlist__50)
				WriteAll(post.Name, true, buffer)
				buffer.WriteString(topicindex__88)

				for _, topic := range post.Topics {
					buffer.WriteString(commentlist__106)
					WriteAll(topic.Url(), true, buffer)
					buffer.WriteString(commentlist__49)
					WriteAll(topic.Name, true, buffer)
					buffer.WriteString(commentlist__50)
					WriteAll("#"+topic.Name, true, buffer)
					buffer.WriteString(commentlist__132)
				}
				buffer.WriteString(topicindex__89)

			}

		}
		buffer.WriteString(commentlist__22)
		var links = paginate.Links()
		buffer.WriteString(commentlist__23)
		for _, link := range links {
			buffer.WriteString(commentlist__44)
			WriteAll(link.Link, true, buffer)
			buffer.WriteString(commentlist__103)
			WriteAll(link.Class, true, buffer)
			buffer.WriteString(commentlist__50)
			WriteAll(link.Label, true, buffer)
			buffer.WriteString(commentlist__105)

		}
		buffer.WriteString(topicindex__29)

		for pos, post := range topPosts {
			{
				var (
					post = post
					pos  = pos + 1
				)

				buffer.WriteString(topicindex__112)

				if pos > 0 {
					buffer.WriteString(topicindex__118)
					WriteEscString(fmt.Sprintf("# %d", pos), buffer)
					buffer.WriteString(topicindex__119)
				}
				buffer.WriteString(commentlist__106)
				WriteAll(post.Url(), true, buffer)
				buffer.WriteString(commentlist__49)
				WriteAll(post.Name, true, buffer)
				buffer.WriteString(commentlist__50)
				WriteAll(post.Name, true, buffer)
				buffer.WriteString(topicindex__116)

				for _, topic := range post.Topics {
					buffer.WriteString(commentlist__106)
					WriteAll(topic.Url(), true, buffer)
					buffer.WriteString(commentlist__49)
					WriteAll(topic.Name, true, buffer)
					buffer.WriteString(commentlist__50)
					WriteAll("#"+topic.Name, true, buffer)
					buffer.WriteString(commentlist__132)
				}
				buffer.WriteString(topicindex__117)
			}

		}
		buffer.WriteString(topicindex__30)
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
		buffer.WriteString(topicindex__35)

	}
}
