package main

import (
	"context"
	"fmt"

	"github.com/ekaputra07/idcloudhost-go/location"
	"github.com/ekaputra07/idcloudhost-go/objectstorage"
)

func main() {
	fmt.Println("List Locations...")
	lc := location.NewClient()
	res := lc.ListLocations(context.Background())
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Printf("%s", res.Body)

	osc := objectstorage.NewClient().ForBillingAccount("1200161630")

	fmt.Println("API URL...")
	res = osc.GetS3ApiURL(context.Background())
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Printf("%s\n", res.Body)

	fmt.Println("S3 USER INFO...")
	res = osc.GetS3UserInfo(context.Background())
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Printf("%s\n", res.Body)

	fmt.Println("S3 USER KEYS...")
	res = osc.GetS3UserKeys(context.Background())
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Printf("%s\n", res.Body)

	// fmt.Println("GENERATE USER KEY...")
	// res = osc.GenerateS3UserKey(context.Background())
	// if res.Error != nil {
	// 	panic(res.Error)
	// }
	// fmt.Printf("%s\n", res.Body)

	// fmt.Println("DELETE USER KEY...")
	// res = osc.DeleteS3UserKey(context.Background(), "P0ES5LWQSACSZY9OT7J7")
	// if res.Error != nil {
	// 	panic(res.Error)
	// }
	// fmt.Printf("%s\n", res.Body)

	fmt.Println("List Buckets...")
	res = osc.ListBuckets(context.Background())
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Printf("%s\n", res.Body)

	fmt.Println("Get Bucket...")
	res = osc.GetBucket(context.Background(), "jetform-demo-static")
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Printf("%s\n", res.Body)

	fmt.Println("Update Bucket...")
	res = osc.UpdateBucketBillingAccount(context.Background(), "snappy-demo", "1200161630")
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Printf("%s\n", res.Body)

	// fmt.Println("Create Bucket...")
	// res = osc.CreateBucket(context.Background(), "idcloudhost-go")
	// if res.Error != nil {
	// 	panic(res.Error)
	// }
	// fmt.Printf("%s\n", res.Body)

	fmt.Println("Delete Bucket...")
	res = osc.DeleteBucket(context.Background(), "idcloudhost-go")
	if res.Error != nil {
		fmt.Printf("%v", res.Error.Error())
	}
	fmt.Printf("%s\n", res.Body)
}
