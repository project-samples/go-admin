package main

import (
	"context"
	"fmt"
	"github.com/core-go/config"
	mid "github.com/core-go/log/middleware"
	ls "github.com/core-go/log/strings"
	"github.com/core-go/log/zap"
	sv "github.com/core-go/service"
	"github.com/core-go/sql/template"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"go-service/internal/app"
)

func main() {
	conf := app.Config{}
	er1 := config.Load(&conf, "configs/sql", "configs/config")
	if er1 != nil {
		panic(er1)
	}
/*
	s := `select u.userId, u.username, u.email, u.displayName
    from users u
    where
      (u.firstName like #{q} or u.lastName like #{q}
        or u.email like N"%${q}%" or u.phone like N"%{q}%" ) and u.status = #{status} and u.firstName = #{firstName} order by {sort} {sortType}
    <isNull property="status">
      u.status != "status" and
    </isNull>
    <isNotNull property="status">
      u.status = 1 and
    </isNotNull>
    <isNotNull property="firstName">
      u.firstName != #{firstName} and
    </isNotNull>
    <isEqual property="test.status" value="1">
      u.teststatus is #{test.status} and
    </isEqual>
    <isNotEqual property="status" value="1">
      u.status is not "X" and
    </isNotEqual>
    <isNotNull property="status">
      u.status = #{status} and
    </isNotNull>
    <isNotNull property="sort">
      ord er by {sort} {sortType}
    </isNotNull>
    1 = 1`

 */
	s2 := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">

<mapper namespace="com.mappers">
  <select id="user">
    select *
    from users
    where
    <if test="username != null">
      username like #{username} and
    </if>
    <if test="displayName != null">
      displayName like #{displayName} and
    </if>
    <if test="status != null">
      status in (#{status}) and
    </if>
    <if test="q != null">
      (username like #{q} or displayName like #{q} or email like #{q}) and
    </if>
    1 = 1
    <if test="sort != null">
      order by {sort}
    </if>
  </select>

  <select id="role">
    select *
    from roles
    where
    <isNotNull property="roleName" test="roleName != null">
      roleName like #{roleName} and
    </isNotNull>
    <isNotNull property="q" test="status != null">
      status in (#{status}) and
    </isNotNull>
    <isNotNull property="q">
      (roleName like #{q} or roleId like #{q} or remark like #{q}) and
    </isNotNull>
    1 = 1
  </select>
  <select id="test">
select u.userId, u.username, u.email, u.displayName
    from users u
    where
      (u.firstName like #{q} or u.lastName like #{q}
        or u.email like N"%${q}%" or u.phone like N"%{q}%" ) and u.status = #{status} and u.firstName = #{firstName} order by {sort} {sortType}
    <isNull property="status">
      u.status != "status" and
    </isNull>
    <isNotNull property="status">
      u.status = 1 and
    </isNotNull>
    <isNotNull property="firstName">
      u.firstName != #{firstName} and
    </isNotNull>
    <isEqual property="test.status" value="1">
      u.teststatus is #{test.status} and
    </isEqual>
    <isNotEqual property="status" value="1">
      u.status is not "X" and
    </isNotEqual>
    <isNotNull property="status">
      u.status = #{status} and
    </isNotNull>
    <isNotNull property="sort">
      ord er by {sort} {sortType}
    </isNotNull>
    1 = 1
  </select>
</mapper>`
	floatNumber := 325.41756
	number := big.NewFloat(floatNumber)
	n2 := RoundBigFloat(*number, 4)
	s:= fmt.Sprintf("%v", &n2)
	fmt.Printf(s)
	fmt.Printf("\n")

	floatNum := 6523147.8525641
	rat := new(big.Rat)
	rat = rat.SetFloat64(floatNum)
	fmt.Println(rat.String())
	s4 := fmt.Sprintf("%v", &rat)
	fmt.Printf(s4)
	result := RoundRat(*rat, 3)
	fmt.Println(result)
	bs, _ := rat.MarshalText()
	myString := string(bs)
	fmt.Println(myString)

	var z1 big.Int
	z1.SetUint64(123)
	z2 := new(big.Rat).SetFloat64(1.25)   // z2 := 5/4
	z3 := new(big.Float).SetFloat64(1.25)       // z3 := 123.0
	fmt.Println(z1.String())
	fmt.Println(z2.String())
	fmt.Println(z2.RatString())
	fmt.Println(z3.String())
	// t, _ := template.BuildTemplate(s)
	obj := make(map[string]interface{})
	sub := make(map[string]interface{})
	ps1 := make([]string, 0)
	obj["status"] = ps1
	sub["status"] = 1
	obj["test"] = sub
	ps2 := make([]string, 0)
	ps2 = append(ps2, "Duc")
	ps2 = append(ps2, "Tri")
	obj["firstName"] = ps2
	obj["q"] = "tien"
	obj["sort"] = "lastName"
	/*
	var t2 template.Template
	t2 = *t
	x, ps := template.Build(obj, t2, BuildDollarParam)
	fmt.Println(x)
	fmt.Println(x, ps)
	 */
	x2, _ := template.BuildTemplates(s2)
	t := x2["user"]
	x, ps := template.Build(obj, *t, BuildDollarParam)
	// fmt.Println(x2[0].Text)
	fmt.Println(x, ps)
	r := mux.NewRouter()
	log.Initialize(conf.Log)
	r.Use(func(handler http.Handler) http.Handler {
		return mid.BuildContextWithMask(handler, MaskLog)
	})
	logger := mid.NewLogger()
	if log.IsInfoEnable() {
		r.Use(mid.Logger(conf.MiddleWare, log.InfoFields, logger))
	}
	r.Use(mid.Recover(log.ErrorMsg))

	er2 := app.Route(r, context.Background(), conf)
	if er2 != nil {
		panic(er2)
	}
	/*
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	*/
	handler := cors.AllowAll().Handler(r)
	fmt.Println(sv.ServerInfo(conf.Server))
	server := sv.CreateServer(conf.Server, handler)
	if er3 := server.ListenAndServe(); er3 != nil {
		fmt.Println(er3.Error())
	}
}

func MaskLog(name, s string) string {
	return ls.Mask(s, 1, 6, "x")
}
func BuildDollarParam(i int) string {
	return "$" + strconv.Itoa(i)
}

func RoundBigFloat(num big.Float, scale int) big.Float {
	marshal, _ := num.MarshalText()
	var dot int
	for i, v := range marshal {
		if v == 46 {
			dot = i + 1
			break
		}
	}
	a := marshal[:dot]
	b := marshal[dot : dot+scale+1]
	c := b[:len(b)-1]

	if b[len(b)-1] >= 53 {
		c[len(c)-1] += 1
	}
	var r []byte
	r = append(r, a...)
	r = append(r, c...)
	num.UnmarshalText(r)
	return num
}
func RoundRat(rat big.Rat, scale int8) string {
	digits := int(math.Pow(float64(10), float64(scale)))
	floatNumString := rat.RatString()
	sl := strings.Split(floatNumString, "/")
	a := sl[0]
	b := sl[1]
	c, _ := strconv.Atoi(a)
	d, _ := strconv.Atoi(b)
	intNum := c / d
	surplus := c - d*intNum
	e := surplus * digits / d
	r := surplus * digits % d
	if r >= d/2 {
		e += 1
	}
	res := strconv.Itoa(intNum) + "." + strconv.Itoa(e)
	return res
}

