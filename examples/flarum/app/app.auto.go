// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package app

import (
	"flarum/model"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/platform/util"
)

// Name project
var Name = "flarum"

// FlarumPosts defined
type FlarumPosts struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewFlarumPosts defined
func NewFlarumPosts() *FlarumPosts {
	ctr := &FlarumPosts{}
	ctr.Add = FlarumPostsAdd
	ctr.Del = FlarumPostsDel
	ctr.Update = FlarumPostsUpdate
	ctr.Page = FlarumPostsPage
	ctr.Get = FlarumPostsGet
	return ctr
}

// FlarumPostsRoutes defined
func FlarumPostsRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/flarum/posts/add", Auth, FlarumPostsInstance.Add)
	group.Handle("DELETE", "/flarum/posts/del", Auth, FlarumPostsInstance.Del)
	group.Handle("PUT", "/flarum/posts/update", Auth, FlarumPostsInstance.Update)
	group.Handle("GET", "/flarum/posts/page", Auth, FlarumPostsInstance.Page)
	group.Handle("GET", "/flarum/posts/get", Auth, FlarumPostsInstance.Get)
}

// FlarumPostsInstance defined
var FlarumPostsInstance = NewFlarumPosts()

// FlarumUsers defined
type FlarumUsers struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewFlarumUsers defined
func NewFlarumUsers() *FlarumUsers {
	ctr := &FlarumUsers{}
	ctr.Add = FlarumUsersAdd
	ctr.Del = FlarumUsersDel
	ctr.Update = FlarumUsersUpdate
	ctr.Page = FlarumUsersPage
	ctr.Get = FlarumUsersGet
	return ctr
}

// FlarumUsersRoutes defined
func FlarumUsersRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/flarum/users/add", Auth, FlarumUsersInstance.Add)
	group.Handle("DELETE", "/flarum/users/del", Auth, FlarumUsersInstance.Del)
	group.Handle("PUT", "/flarum/users/update", Auth, FlarumUsersInstance.Update)
	group.Handle("GET", "/flarum/users/page", Auth, FlarumUsersInstance.Page)
	group.Handle("GET", "/flarum/users/get", Auth, FlarumUsersInstance.Get)
}

// FlarumUsersInstance defined
var FlarumUsersInstance = NewFlarumUsers()

// SyncModel defined
func SyncModel() error {
	mseti := App.Manager.MSet()
	mseti.Add(new(model.FlarumAccessTokens))
	mseti.Add(new(model.FlarumApiKeys))
	mseti.Add(new(model.FlarumDiscussionTag))
	mseti.Add(new(model.FlarumDiscussionUser))
	mseti.Add(new(model.FlarumDiscussions))
	mseti.Add(new(model.FlarumEmailTokens))
	mseti.Add(new(model.FlarumFlags))
	mseti.Add(new(model.FlarumGroupPermission))
	mseti.Add(new(model.FlarumGroupUser))
	mseti.Add(new(model.FlarumGroups))
	mseti.Add(new(model.FlarumLoginProviders))
	mseti.Add(new(model.FlarumMigrations))
	mseti.Add(new(model.FlarumNotifications))
	mseti.Add(new(model.FlarumPasswordTokens))
	mseti.Add(new(model.FlarumPostLikes))
	mseti.Add(new(model.FlarumPostMentionsPost))
	mseti.Add(new(model.FlarumPostMentionsUser))
	mseti.Add(new(model.FlarumPostUser))
	mseti.Add(new(model.FlarumPosts))
	mseti.Add(new(model.FlarumRegistrationTokens))
	mseti.Add(new(model.FlarumSettings))
	mseti.Add(new(model.FlarumTagUser))
	mseti.Add(new(model.FlarumTags))
	mseti.Add(new(model.FlarumUsers))
	return nil
}

// SyncCtr defined
func SyncCtr() error {
	FlarumPostsRoutes(App)
	FlarumUsersRoutes(App)
	return nil
}

// SyncService defined
func SyncService() error {
	return nil
}

// Executor defined
var Executor = util.NewExecutor(SyncModel, SyncCtr, SyncService)

func init() {
	if err := Executor.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
