package view

import (
	"crypto/sha1"
	"encoding/binary"
	"github.com/valyala/fasthttp"
	"golang.org/x/net/xsrftoken"
	"html/template"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const CERASALT = "Crea@2019=="

func init() {
	rand.Seed(time.Now().Unix())
	sha1.New()
}

type IMethodViewer interface {
	Init()
	Before()
	Get()
	Post()
	Head()
	Options()
	Put()
	Patch()
	Delete()
	Trace()
	GetCtx() *fasthttp.RequestCtx
	SetCtx(ctx *fasthttp.RequestCtx)
	After()
	Render()
}

type IViewer interface {
	Before()
	After()
	GetCtx() *fasthttp.RequestCtx
	SetCtx(ctx *fasthttp.RequestCtx)
}

type CreaCookie struct {
	ActionId [16]byte // action id of session,generate by timestamp and random uint
	XsrfKey  [8]byte  // xsrf key
	XsrfUid  [8]byte  // xsrf uid
}

// parse cookie to struct
func ParseCookie(ck []byte) *CreaCookie {
	c := new(CreaCookie)
	copy(c.ActionId[:], ck[0:15])
	copy(c.XsrfKey[:], ck[16:23])
	copy(c.XsrfUid[:], ck[24:31])
	return c
}

// generate struct to byte slice
func (c *CreaCookie) ToByte() []byte {
	b := make([]byte, 32)
	copy(b[0:15], c.ActionId[:])
	copy(b[16:23], c.XsrfKey[:])
	copy(b[24:31], c.XsrfUid[:])
	return b
}

// generate new action id
func newActionId() [16]byte {
	buf := [16]byte{}
	r := make([]byte, 8)
	tm := make([]byte, 8)
	// covert int64 to byte
	binary.BigEndian.PutUint64(r, rand.Uint64())
	binary.BigEndian.PutUint64(tm, uint64(time.Now().Unix()))
	// add random int and timestamp to buffer
	copy(buf[0:7], r[:])
	copy(buf[8:15], tm[:])
	return buf
}

type View struct {
	JinJaTpl bool
	Data     map[string]string // stored user values
	Ctx      *fasthttp.RequestCtx
	Cookie   *fasthttp.Cookie
}

// combine this struct and rewrite those functions to reply http methods
func (r *View) Init() {
	r.Data = make(map[string]string)
}

func (r *View) Before() {}

func (r *View) Get() { r.Html404() }

func (r *View) Head() { r.Html404() }

func (r *View) Options() { r.Html404() }

func (r *View) Post() { r.Html404() }

func (r *View) Put() { r.Html404() }

func (r *View) Patch() { r.Html404() }

func (r *View) Delete() { r.Html404() }

func (r *View) Trace() { r.Html404() }

func (r *View) After() {}

func (r *View) Render() {
	t := template.Must(template.ParseGlob("./template/*.htm"))
	err := t.Execute(r.Ctx.Response.BodyWriter(), r.Data)
	if err != nil {
		log.Println(err)
		return
	}
}

// 获取参数，通过标准get url方式传值 e.g. http://xxx.com/?id=1
func (r *View) GetArgString(key string) string {
	return string(r.Ctx.Request.URI().QueryArgs().Peek(key))
}

// 获取参数，通过标准get url方式传值 e.g. http://xxx.com/?id=1
func (r *View) GetArgBytes(key string) []byte {
	return r.Ctx.Request.URI().QueryArgs().Peek(key)
}

// 获取参数，通过标准get url方式传值 e.g. http://xxx.com/?id=1
func (r *View) GetArgInt(key string) (int, error) {
	s := r.GetArgString(key)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (r *View) GetCtx() *fasthttp.RequestCtx {
	return r.Ctx
}

func (r *View) SetCtx(ctx *fasthttp.RequestCtx) {
	r.Ctx = ctx
}

func Switcher(v IMethodViewer) {
	// running before method priority
	ctx := v.GetCtx()
	v.Before()
	method := ctx.Method()
	switch string(method) {
	case fasthttp.MethodGet:
		v.Get()
		v.Render()
	case fasthttp.MethodPost:
		v.Post()
		v.Render()
	case fasthttp.MethodHead:
		v.Head()
		v.Render()
	case fasthttp.MethodOptions:
		v.Options()
		v.Render()
	case fasthttp.MethodPut:
		v.Put()
		v.Render()
	case fasthttp.MethodPatch:
		v.Patch()
		v.Render()
	case fasthttp.MethodDelete:
		v.Delete()
		v.Render()
	case fasthttp.MethodTrace:
		v.Trace()
		v.Render()
	default:
		HtmlUnknownMethod(ctx)
	}
	v.After()
}

func (r *View) GetPostArgs(key string) string {

	return string(r.Ctx.PostArgs().Peek(key))
}

func (r *View) setXsrfToken() {
	buf := make([]byte, 8)
	bufU := make([]byte, 8)
	random := rand.Uint64()
	randomU := rand.Uint64()
	binary.BigEndian.PutUint64(buf, random)
	binary.BigEndian.PutUint64(bufU, randomU)
	xsrf := xsrftoken.Generate(string(buf), strconv.FormatUint(rand.Uint64(), 16), strconv.FormatInt(time.Now().Unix(), 16))

	c := new(fasthttp.Cookie)
	c.SetKey("CreaCookie")
	// todo : set session key id
	//c.SetValue()
	r.Data["XSRF"] = xsrf
	r.Ctx.Response.Header.SetCookie(c)
}

// xsrf token check, if xsrf token is not valid, response permission denied to client
//func XsrfValidate(ctx *fasthttp.RequestCtx) {
//	// xsrf token
//	xsrf := string(ctx.Request.PostArgs("XSRF_TOKEN"))
//	// xsrf uid
//	xsu := string(r.Ctx.Request.Header.Cookie("xsu"))
//	// xsrf key
//	xsk := string(r.Ctx.Request.Header.Cookie("xsk"))
//	// xsrf action
//	xsa := string(r.Ctx.Request.Header.Cookie("xsa"))
//	if !xsrftoken.Valid(xsrf, xsk, xsu, string(xsa)) {
//		html := `
//<html>
//<head>
//<title>Permission denied!xsrf</title>
//</head>
//<body style="background:#000;text-align:center;">
//<span style="font-size:5em;color:#fff;"><b>403, sorry, xsrf token check failed! :) </b></span>
//</body>
//</html>
//`
//		r.Ctx.SetStatusCode(403)
//		if _, err := r.Ctx.Write([]byte(html)); err != nil {
//			log.Print(err)
//		}
//		r.Ctx.Done()
//	}
//}

func HtmlUnknownMethod(ctx *fasthttp.RequestCtx) {
	html := `
<html>
<head>
<title>Page not found</title>
</head>
<body style="background:#000;text-align:center;">
<span style="font-size:5em;color:#fff;"><b>500, sorry! unknown http method :) </b></span>
</body>
</html>
`
	ctx.SetStatusCode(500)
	if _, err := ctx.Write([]byte(html)); err != nil {
		log.Print(err)
	}
}

func (r *View) Html404() {
	html := `
<body style="background:#000;text-align:center;">
<span style="font-size:5em;color:#fff;"><b>404 Sorry, Page not found! :) </b></span>
</body>
`
	r.Ctx.SetStatusCode(404)
	r.Ctx.SetBodyString(html)
	if _, err := r.Ctx.Write([]byte(html)); err != nil {
		log.Print(err)
	}
}
