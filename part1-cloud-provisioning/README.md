## S3 Bucket Creation and Deletion

This application provides functions to create or destroy an S3 bucket. The instructions below can be used to create a new bucket with your locally configured AWS credentials.

To print out the application usage you can run the application with the `-help` option.

```
kevinhartwig@kevins-mbp part1-cloud-provisioning % ./s3-bucket -help
Usage of ./s3-bucket:
  -bucket-name string
        Name of the bucket you would like to create
  -create-bucket
        Use this flag if you want to create a bucket
  -delete-bucket
        Use this flag if you want to delete a bucket
  -region string
        Region to create the bucket in (default "us-west-1")
```

Creation of a new bucket requires 3 command line arguments be passed in. The `-bucket-name`, `-region` where you would like the bucket to reside and the `-create-bucket` flag. An example and successful output can be found below. 

```
kevinhartwig@kevins-mbp part1-cloud-provisioning % ./s3-bucket -bucket-name hh-sre-test2 -region us-west-2 -create-bucket
Creating bucket "hh-sre-test2"...
Bucket: "hh-sre-test2" created successfully.
```

Deletion of an existing bucket requires 3 command line arguments be passed in. The `-bucket-name`, `-region` where you would like the bucket to be removed from and the `-delete-bucket` flag. An example of a deletion and the output can be found below.

```
kevinhartwig@kevins-mbp part1-cloud-provisioning % ./s3-bucket -bucket-name hh-sre-test2 -region us-west-2 -delete-bucket
Deleting bucket "hh-sre-test2"...
Bucket: "hh-sre-test2" deleted successfully.
```

Any errors will be printed out during run time. An example of an error can be found below. This error occured because we attempted to delete a bucket which does not exist.

```
evinhartwig@kevins-mbp part1-cloud-provisioning % ./s3-bucket -bucket-name hh-sre-test1 -region us-west-2 -delete-bucket
Error deleting bucket "hh-sre-test1", NoSuchBucket: The specified bucket does not exist
        status code: 404, request id: A581510DCBC7E565, host id: tgCeUTPVgVZPYDFf5niw29LOqODvzHNIYs0gbkgZhyu/8qirF8zzYQQ3UDZKftTXSt1r94drRcc=
```