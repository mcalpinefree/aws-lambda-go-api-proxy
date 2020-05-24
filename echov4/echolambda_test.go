package echov4adapter_test

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/labstack/echo/v4"
	echov4adapter "github.com/mcalpinefree/aws-lambda-go-api-proxy/echov4"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EchoLambda tests", func() {
	Context("Simple ping request", func() {
		It("Proxies the event correctly", func() {
			log.Println("Starting test")
			e := echo.New()
			e.GET("/ping", func(c echo.Context) error {
				log.Println("Handler!!")
				return c.String(200, "pong")
			})

			adapter := echov4adapter.New(e)

			req := events.APIGatewayProxyRequest{
				Path:       "/ping",
				HTTPMethod: "GET",
			}

			resp, err := adapter.Proxy(req)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
		})
	})
	Context("Simple ping request v2", func() {
		It("Proxies the event correctly", func() {
			log.Println("Starting test")
			e := echo.New()
			e.GET("/v1/ping", func(c echo.Context) error {
				log.Println("Handler!!")
				return c.String(200, "pong")
			})

			adapter := echov4adapter.New(e)

			req := events.APIGatewayV2HTTPRequest{
				RequestContext: events.APIGatewayV2HTTPRequestContext{
					HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
						Path:   "/v1/ping",
						Method: "GET",
					},
				},
			}

			resp, err := adapter.ProxyV2(req)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
		})
	})
})
