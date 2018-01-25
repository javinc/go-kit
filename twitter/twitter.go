package twitter

import (
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/gologin"
	twitterLogin "github.com/dghubble/gologin/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/javinc/go-kit/config"
)

var (
	oAuthConf = &oauth1.Config{
		ConsumerKey:    config.GetString("twitter.consumer_key"),
		ConsumerSecret: config.GetString("twitter.consumer_secret"),
		CallbackURL:    "/twitter-callback-default",
		Endpoint:       twitterOAuth1.AuthorizeEndpoint,
	}
)

// User twitter
type User *twitter.User
type failHandler func(error)
type successHandler func(User, error)

// Login handler
func Login(url string, fail failHandler) http.Handler {
	oAuthConf.CallbackURL = url
	oAuthConf.CallbackURL = "http://mashdrop.com/cb"
	return twitterLogin.LoginHandler(oAuthConf, onFail(fail))
}

// Callback handler
func Callback(success successHandler, fail failHandler) http.Handler {
	return twitterLogin.CallbackHandler(oAuthConf, onSuccess(success), onFail(fail))
}

// PostStatus feed
func PostStatus(link, msg string) (id string, err error) {
	return
}

func newClient(token, secret string) *twitter.Client {
	t := oauth1.NewToken(token, secret)
	c := oauth1.NewConfig(oAuthConf.ConsumerKey, oAuthConf.ConsumerSecret)

	return twitter.NewClient(c.Client(oauth1.NoContext, t))
}

func onSuccess(f successHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get user
		f(twitterLogin.UserFromContext(r.Context()))
	})
}

func onFail(f failHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f(gologin.ErrorFromContext(r.Context()))
	})
}
