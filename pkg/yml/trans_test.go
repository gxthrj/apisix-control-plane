package yml_test

import (
	"fmt"
	"github.com/apache/apisix-control-plane/pkg/yml"
	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trans", func() {
	Describe("trans to model", func() {
		Context("trans", func() {
			It("trans to gateway no error", func() {
				b := []byte(`
kind: Gateway
name: foo-gw
servers:
 - port:
     number: 80
     name: http
     protocol: HTTP
   hosts:
   - "a.foo.com"
   - "b.foo.com"
`)
				y, err := yaml.YAMLToJSON(b)
				fmt.Println(string(y))
				ym := yml.Trans(y, b)
				Expect(err).NotTo(HaveOccurred())
				v := typeof(ym)
				fmt.Println(v)
				Expect(v).To(Equal("*yml.Gateway"))
				g, ok := ym.(*yml.Gateway)
				Expect(ok).To(Equal(true))
				Expect(len(g.Servers[0].Hosts)).To(Equal(2))
			})

			It("trans to rule no error", func() {
				b := []byte(`
kind: Rule
name: xxx-rules
hosts:
- "a.foo.com"
gateways:
- foo-gw
http:
- route:
  - destination:
     port: 28002
     host: foo-server
     subset: foo-v1
     weight: 10
  label:
    app: foo
    version: v1
  match:
  - headers:
     product_id:
       exact: v1
- route:
  - destination:
       port: 28002
       host: foo-server
       subset: v2
  label:
    app: foo
    version: v2
`)
				y, err := yaml.YAMLToJSON(b)
				fmt.Println(string(y))
				ym := yml.Trans(y, b)
				Expect(err).NotTo(HaveOccurred())
				v := typeof(ym)
				fmt.Println(v)
				Expect(v).To(Equal("*yml.Rule"))
				r, ok := ym.(*yml.Rule)
				Expect(ok).To(Equal(true))
				Expect(r.Kind).To(Equal("Rule"))
				Expect(r.Kind).To(Equal("Rule"))
			})
		})
	})
})

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
