package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := NewAPI(NewRoute())

	log.Fatal(http.ListenAndServe(":8081", api.MakeHandler()))
}

func NewRoute() rest.App {
	router, err := rest.MakeRouter(
		rest.Post("/", PostCustomer),
		rest.Post("/login", Login),
	)

	if err != nil {
		log.Fatal(err)
	}
	return router
}

func NewAPI(router rest.App) (api *rest.Api) {
	api = rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders := []string{
		"Accept",
		"Authorization",
		"X-Real-IP",
		"Content-Type",
		"X-Custom-Header",
		"Language",
		"Origin",
	}
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
		AllowedMethods:                allowedMethods,
		AllowedHeaders:                allowedHeaders,
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})

	loginMiddle := &LoginMiddleware{}
	api.Use(loginMiddle)

	api.SetApp(router)
	return
}

func Login(w rest.ResponseWriter, r *rest.Request) {
	body := map[string]string{}

	err := r.DecodeJsonPayload(&body)
	if err != nil {
		w.WriteHeader(400)
		w.WriteJson(err)
	}

	response, status := checkAuthorize(body["user"], body["password"])

	w.WriteHeader(status)
	w.WriteJson(response)
}

func checkAuthorize(user, passwd string) (map[string]string, int) {
	resp := map[string]string{}

	if user == "kob@gmail.com" && passwd == "aobaob" {
		resp["token"] = CreateToken(user)
		return resp, 200
	}

	resp["error"] = "user or password wrong"
	return resp, 401
}

type LoginMiddleware struct {
}

func (login *LoginMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		fmt.Println("before execute handler")
		handler(w, r)
		fmt.Println("after execute handler")
	}
}

func PostCustomer(w rest.ResponseWriter, r *rest.Request) {
	records := []SubscriberTargetDetail{
		SubscriberTargetDetail{
			TSCID:        "tscID",
			CampaignCode: "campaignCode",
		},
	}

	bf := &bytes.Buffer{}
	c := csv.NewWriter(bf)
	c.Comma = '|'

	for _, obj := range records {
		var record []string
		record = append(record, obj.TSCID)
		record = append(record, obj.CampaignCode)
		c.Write(record)
	}
	c.Flush()

	hw := w.(http.ResponseWriter)
	hw.Header().Set("Content-type", "text/csv")
	hw.Header().Set("Content-disposition", "attachment;filename=anuchit.TXT")
	hw.Write(bf.Bytes())
}

type SubscriberTargetListDetailResponse struct {
	List []SubscriberTargetDetail `json:"list"`
}

type SubscriberTargetDetail struct {
	ARPU             float64 `bson:"arpu" json:"arpu"`
	TSCID            string  `bson:"tscID" json:"tscID"`
	CampaignCode     string  `bson:"campaignCode" json:"campaignCode"`
	CampaignName     string  `bson:"campaignName" json:"campaignName"`
	SubscriberNumber string  `bson:"subscriberNumber" json:"subscriberNumber"`
	CustomerNumber   string  `bson:"customerNumber" json:"customerNumber"`
}

func CreateToken(user string) string {
	claims := jws.Claims{}
	claims.SetIssuer("app kob")
	claims.SetAudience(user)

	tokenStruct := jws.NewJWT(claims, crypto.SigningMethodHS256)

	secret := "team"
	serialized, err := tokenStruct.Serialize([]byte(secret))
	if err != nil {
		log.Fatal("error : ", err.Error())
	}

	token := string(serialized)

	return token
}
