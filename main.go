package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"testing"
	"toy/src/model"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	db["snopy"] = "snopy"
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	      Zm9vOmJhcg== is base64("foo:bar")

	   	curl -X POST \
	     	http://localhost:8080/admin \
	     	-H 'authorization: Basic Zm9vOmJhcg==' \
	     	-H 'content-type: application/json' \
	     	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

type device struct {
}

func (d *device) Write(p []byte) (n int, err error) {
	//TODO implement me
	return 0, nil
}

func (d *device) Close() error {
	//TODO implement me
	return nil
}

type ByteCounter struct {
}

func (b ByteCounter) Read(p []byte) (n int, err error) {
	//TODO implement me
	return 0, err
}

func (b ByteCounter) Write(p []byte) (n int, err error) {
	//TODO implement me
	return 0, nil
}

type Flyer interface {
	fly()
}
type Walker interface {
	walk()
}
type bird struct {
}

func (b bird) walk() {
	//TODO implement me
	fmt.Println("bird walk")
}

func (b bird) fly() {
	//TODO implement me
	fmt.Println("bird fly")
}

type pig struct {
}

func (p pig) walk() {
	//TODO implement me
	fmt.Println("pig walk")
}

type Swim interface {
}
type Dict struct {
	data map[interface{}]interface{}
}

func (d *Dict) Set(k interface{}, v interface{}) error {
	d.data[k] = v
	return nil
}
func (d *Dict) Get(k interface{}) (v interface{}, err error) {
	v = d.data[k]
	if v != nil {
		return v, nil
	}
	return nil, errors.New("not exsist")
}
func (d *Dict) Visit(callback func(k interface{}, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}
func (d *Dict) Clear() {
	d.data = make(map[interface{}]interface{})
}
func (d *Dict) Remove(k interface{}) error {
	delete(d.data, k)
	if d.data[k] == nil {
		return nil
	}
	return errors.New("删除失败")
}
func NewDict() *Dict {
	d := &Dict{}
	d.Clear()
	return d
}
func ppp(param interface{}) {
	switch param.(type) {
	case Dict:
	case error:
	case int:
	default:

	}
}

type Alipay struct {
}

/**
 *@description:TODO
 *@author:txl
 *@date:2023/2/7 17:30
 *
 */

func (a *Alipay) CanUseFaceID() {

}

type Cash struct {
}

func (c *Cash) Stolen() {

}

type CantainCanUseFaceId interface {
	CanUseFaceID()
}
type ContainStolen interface {
	Stolen()
}

func print(payMethod interface{}) {
	switch payMethod.(type) {
	case CantainCanUseFaceId:
		fmt.Printf("%T can use faceid \n", payMethod)
	case ContainStolen:
		fmt.Printf("%T may be stolen \n", payMethod)
	}
}

type dualError struct {
	Num     float64
	problem string
}

func (d dualError) Error() string {
	//TODO implement me
	return fmt.Sprintf("Wrong!because \"%f\" is a negative num ", d.Num)
}
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, dualError{Num: f}
	}
	return math.Sqrt(f), nil
}

type Expr interface {
	Eval(e Env) float64
}
type Var string
type literal float64
type unary struct {
	op rune // one of '+','-'
	x  Expr
}
type bnary struct {
	op   rune
	x, y Expr
}
type call struct {
	fn   string
	args []Expr
}
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}

	panic(fmt.Sprintf("unsupported unary operator :%q", u.op))
}
func (b bnary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)

	}
	panic(fmt.Sprintf("unsupported binary operator :%q", b.op))
}
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported call operator %q", c.fn))
}
func TestEval(t *testing.T) {
	/*tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A", 87616, "pi": math.Pi}, "167"},
		{"pow(x,3) + pow(y, 3)", Env{"x", 12, "y", 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}

	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}*/
}

func Parse(expr string) (interface{}, interface{}) {

	return nil, nil
}

func main() {
	//cmd.Execute()
	student := model.NewStudent("snopy")
	student.SetAge(18)
	student.SetSal(99999.00)
	fmt.Println(student.GetAge())
}
func index(c *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}
	//file, _ := ioutil.ReadFile("./static/index.html")
	c.JSONP(http.StatusOK, data)
}
