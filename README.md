# Warren Go!

Golang client library for [Warren.io](https://warren.io/) API. If your cloud hosting provider is using Warren.io (and they enabled the API) then you should be able to manage your infrastructure programatically using this library.

Progress:
- [x] Locations
- [x] Object storage
- [x] Block storage
- [x] Floating IP
- [ ] Load balancer
- [ ] Managed services
- [ ] Virtual machine
- [x] Virtual Private Cloud (VPC)

## Usage
The easiest way to getting started is to set API's base URL and API Key in environment variables:
```bash
export WARREN_API_BASE_URL=https://api.idcloudhost.com
export WARREN_API_KEY=secret123
```

> NOTE: Please consult with your hosting provider for the API base URL. In this example I'm using `https://api.idcloudhost.com` which is one of hosting providers in Indonesia that are using Warren.io

The in your Golang app:
```golang
package main

import (
	"context"
	"github.com/ekaputra07/warren-go"
)

func main() {
    ctx := context.Background()

    // Warren client
    w := warren.New()

    // list locations
    w.Location.ListLocations(ctx)

    // list S3 buckets
    w.ObjectStorage.ListBuckets(ctx)

    // list VPC networks
    w.VPC.ListNetworks(ctx, "sgp01")
}
```

### Create multiple clients
Above method works well if you're trying to connect to a single hosting provider. But what if your infrastructures are spread across multiple providers?

You can create multiple instances of Warren that points to different providers:
```golang
import (
    "github.com/ekaputra07/warren-go"
    "github.com/ekaputra07/warren-go/api"
)

// Warren client for provider A
apiA := api.New("https://api.a.com", "apiKeyFromA")
wa := warren.Init(apiA)

wa.Location.ListLocations(ctx)

// Warren client for provider B
apiB := api.New("https://api.b.com", "apiKeyFromB")
wb := warren.Init(apiB)

wb.Location.ListLocations(ctx)
```

### Create client for specific module
If you just want to create a client for specific module e.g. Object Storage, simply import and initialize your desired module.
```golang
import (
    "github.com/ekaputra07/warren-go/api"
    "github.com/ekaputra07/warren-go/objectstorage"
)

// using default API client
c := objectstorage.NewClient(api.Default)
c.ListBuckets(ctx)

// OR manually setting base URL and API key
a := api.New("https://api.idcloudhost.com", "secret")
c2 := objectstorage.NewClient(a)
c2.ListBuckets(ctx)
```

### Create client with data center location
Some resource will require us to specify data center location.
```golang
import (
    "github.com/ekaputra07/warren-go/api"
    "github.com/ekaputra07/warren-go/vpc"
)

ctx := context.Background()

// Warren client with location
w := warren.NewWithLocation("jkt01")
w.VPC.ListNetworks(ctx)

// OR manually created client
a := api.New("https://api.idcloudhost.com", "secret")
v := vpc.NewClient(a, "jkt01")
v.ListNetworks(ctx)
```