package twitter

import (
	"errors"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/gologin"
	gologinOAuth1 "github.com/dghubble/gologin/oauth1"
	twitterLogin "github.com/dghubble/gologin/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/javinc/go-kit/config"
)

// User twitter
type User *twitter.User
type failHandler func(error)
type successHandler func(string, string, error)

var (
	oAuthConf *oauth1.Config
	client    *twitter.Client
)

// Init load config
func Init() {
	defaultCBURL := config.GetString("host") + config.GetString("port") + "/twitter-callback-default"
	oAuthConf = &oauth1.Config{
		ConsumerKey:    config.GetString("twitter.consumer_key"),
		ConsumerSecret: config.GetString("twitter.consumer_secret"),
		CallbackURL:    defaultCBURL,
		Endpoint:       twitterOAuth1.AuthorizeEndpoint,
	}
}

// Login handler
func Login(url string, fail failHandler) http.Handler {
	oAuthConf.CallbackURL = url
	return twitterLogin.LoginHandler(oAuthConf, onFail(fail))
}

// Callback handler
func Callback(success successHandler, fail failHandler) http.Handler {
	return twitterLogin.CallbackHandler(oAuthConf, onSuccess(success), onFail(fail))
}

// SetAccess set twitter client
func SetAccess(token, secret string) {
	t := oauth1.NewToken(token, secret)
	c := oauth1.NewConfig(oAuthConf.ConsumerKey, oAuthConf.ConsumerSecret)

	client = twitter.NewClient(c.Client(oauth1.NoContext, t))
}

// GetUser twitter user
func GetUser() (user User, err error) {
	if err = checkClient(); err != nil {
		return
	}

	user, _, err = client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	})

	return
}

// PostStatus feed
func PostStatus(link, msg string) (tweet *twitter.Tweet, err error) {
	if err = checkClient(); err != nil {
		return
	}

	status := msg + " " + link
	tweet, _, err = client.Statuses.Update(status, &twitter.StatusUpdateParams{})

	return
}

func checkClient() (err error) {
	if client == nil {
		err = errors.New("twitter client access not set")
	}

	return
}

func onSuccess(f successHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f(gologinOAuth1.AccessTokenFromContext(r.Context()))
	})
}

func onFail(f failHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f(gologin.ErrorFromContext(r.Context()))
	})
}
