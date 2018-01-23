package smtp

import (
	"crypto/tls"
	"html/template"
	"io"
	"log"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

// smtp abstraction

type c struct {
	Host string
	Port int
	User string
	Pass string
}

var (
	// TmplPath tmpl path
	TmplPath = "./*.tmpl"

	conf c
	ch   = make(chan *gomail.Message)
	tmpl *template.Template
)

// Start mail service
func Start() {
	// load tmpl
	t, err := template.ParseGlob(TmplPath)
	if err != nil {
		panic(err)
	}

	tmpl = t

	conf = c{
		viper.GetString("mail.host"),
		viper.GetInt("mail.port"),
		viper.GetString("mail.user"),
		viper.GetString("mail.pass"),
	}

	// check config
	if conf.Host == "" || conf.User == "" || conf.Pass == "" {
		log.Fatalln("[mail] configuration not valid")
	}

	// default smtp port
	if conf.Port == 0 {
		conf.Port = 587
	}

	go func() {
		// delay
		time.Sleep(5 * time.Second)

		d := gomail.NewDialer(conf.Host, conf.Port, conf.User, conf.Pass)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-ch:
				if !ok {
					return
				}
				if !open {
					if s, err = d.Dial(); err != nil {
						panic(err)
					}
					open = true
				}
				if err := gomail.Send(s, m); err != nil {
					log.Print(err)
				}

			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
			case <-time.After(30 * time.Second):
				if open {
					if err := s.Close(); err != nil {
						log.Print(err)
					}
					open = false
				}
			}
		}
	}()
}

// SendHTMLTmpl mail using html template
func SendHTMLTmpl(to, subject, tmplName string, data map[string]interface{}) {
	m := new(to, subject)

	// add host
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		return tmpl.Lookup(tmplName).Execute(w, data)
	})

	// add to channel
	ch <- m
}

func baseSend(to, subject, body, contentType string) {
	m := new(to, subject)
	m.SetBody(contentType, body)

	// add to channel
	ch <- m
}

func new(to, subject string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	return m
}

// Test send
func Test() {
	SendHTMLTmpl("javinczki02@gmail.com",
		"Test Graham Email",
		"hello.tmpl",
		map[string]interface{}{
			"name": "James",
			"msg":  "like hell",
		})
}
