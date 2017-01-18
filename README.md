At the moment there is only an interface for assets implemented, this can easily be extended thought.
package main

**Example:**

```go
import (
	"fmt"
	"log"

	"github/mickep76/go-workday"
)

const (
	Tenant   = "My tenant"
	Username = "My user"
	Password = "My password"
)

func main() {
	clnt := client.NewClient(client.Config{Username: Username, Password: Password, Tenant: Tenant})

	hdlr := client.NewRequestGetAssets(clnt)

        // Request references
	hdlr.Request.CompanyReference = client.NewObjectList()
	hdlr.Request.CompanyReference.Add("", []*client.RequestObjectID{
		{Type: "Company_Reference_ID", Value: "1234"},
	})

	hdlr.Request.WorktagReference = client.NewObjectList()
	hdlr.Request.WorktagReference.Add("", []*client.RequestObjectID{
		{Type: "Cost_Center_Reference_ID", Value: "1234"},
	})

	hdlr.Request.AssetStatusReference = client.NewObjectList()
	hdlr.Request.AssetStatusReference.Add("", []*client.RequestObjectID{
		{Type: "Document_Status_ID", Value: "IN_SERVICE"},
	})

        // Print XML request
	buf, err := hdlr.GetRequestXML()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Request: \n\n%s\n\n", string(buf))

        // Print response
	count := 0
	for !hdlr.Last() {
		if err := hdlr.GetNext(); err != nil {
			log.Fatal(err)
		}
		for _, a := range hdlr.Response.Assets {
			fmt.Println(count, a)
			count++
		}
	}
}
```
