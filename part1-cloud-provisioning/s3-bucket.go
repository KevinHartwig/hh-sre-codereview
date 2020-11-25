package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func createS3Bucket(bucket string, region string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		exitErrorf("Error creating AWS session, %v", err)
	}

	// Create S3 service client
	svc := s3.New(sess)

	_, err = svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Error creating bucket %q, %v", bucket, err)
	}

	// Wait until bucket is created before finishing
	fmt.Printf("Creating bucket %q...\n", bucket)

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Error creating bucket %q, %v", bucket, err)
	}
	fmt.Printf("Bucket: %q created successfully.\n", bucket)
}

func deleteS3Bucket(bucket string, region string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		exitErrorf("Error creating AWS session, %v", err)
	}

	// Create S3 service client
	svc := s3.New(sess)

	_, err = svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Error deleting bucket %q, %v", bucket, err)
	}

	// Wait until bucket is created before finishing
	fmt.Printf("Deleting bucket %q...\n", bucket)

	err = svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Error deleting bucket %q, %v", bucket, err)
	}
	fmt.Printf("Bucket: %q deleted successfully.\n", bucket)
}

func main() {
	bucketPtr := flag.String("bucket-name", "", "Name of the bucket you would like to create")
	regionPtr := flag.String("region", "us-west-1", "Region to create the bucket in")
	createPtr := flag.Bool("create-bucket", false, "Use this flag if you want to create a bucket")
	deletePtr := flag.Bool("delete-bucket", false, "Use this flag if you want to delete a bucket")

	flag.Parse()

	// Verify required fields exist

	if *bucketPtr == "" ||
		*regionPtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Verify bucket function is specified and unique
	if *createPtr &&
		*deletePtr {
		fmt.Println("Cannot use both create-bucket and delete-bucket simultaneously")
		os.Exit(1)
	} else if !*createPtr &&
		!*deletePtr {
		fmt.Println("Must choose either create-bucket or delete-bucket")
		os.Exit(1)
	}

	bucket := *bucketPtr
	region := *regionPtr

	if *createPtr {
		createS3Bucket(bucket, region)
	} else if *deletePtr {
		deleteS3Bucket(bucket, region)
	}

}
