package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func main() {

	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		fmt.Println("Please define env var PROJECT_ID for GCP Project")
		os.Exit(1)
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Println("Unable to establish GCS client", err)
		os.Exit(1)
	}

	// list single bucket
	if len(os.Args) > 1 {
		printACLs(ctx, client, os.Args[1])
		os.Exit(0)
	}

	// list buckets
	var buckets []string
	it := client.Buckets(ctx, projectID)
	for {
		item, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%+v \n", item.ACL)
		buckets = append(buckets, item.Name)
	}

	// Print ACLs for bucket
	for _, bucket := range buckets {
		printACLs(ctx, client, bucket)
	}

}

func printACLs(ctx context.Context, client *storage.Client, bucket string) {
	log.Printf("gs://%s", bucket)
	/*
		// ACL and DefaultObjetACLs not used at this time
			aclRules, err := client.Bucket(bucket).DefaultObjectACL().List(ctx)
			if err != nil {
				log.Printf("Unable to get Default ACL for bucket %s: %v", bucket, err)
			}

			var acls []string
			acls = append(acls, fmt.Sprintf("gs://%s", bucket))
			for _, acl := range aclRules {
				acls = append(acls, fmt.Sprintf("%s:%s", acl.Entity, acl.Role))
			}
			fmt.Println(strings.Join(acls, ","))
			fmt.Printf("%+v\n", aclRules)

			acl, err := client.Bucket(bucket).ACL().List(ctx)
			if err != nil {
				log.Printf("Unable to get ACL for bucket %s: %v", bucket, err)
			}
			fmt.Printf("%+v\n", acl)
	*/

	// IAM Polices are what this looks at
	iamPolicy, err := client.Bucket(bucket).IAM().Policy(ctx)
	if err != nil {
		log.Printf("Unable to get policy: %v", err)
		return
	}
	//fmt.Printf("%+v\n", iamPolicy)
	for _, p := range iamPolicy.InternalProto.Bindings {
		//log.Printf("%+v", p)
		if contains(p.Members, []string{"allUsers", "allAuthenticatedUsers"}) {
			fmt.Printf("gs://%s is PUBLIC\n", bucket)
		}
	}
}

func contains(sourceList, containsList []string) bool {
	existence := false
	for _, target := range containsList {
		for _, item := range sourceList {
			if item == target {
				log.Println(target)
				existence = true
				break
			}
		}
	}
	return existence
}
