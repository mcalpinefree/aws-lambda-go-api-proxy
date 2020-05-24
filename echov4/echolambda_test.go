package echov4adapter_test

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/labstack/echo/v4"
	"github.com/awslabs/aws-lambda-go-api-proxy/echov4"

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
})
