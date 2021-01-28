# bucket-review

Given a PROJECT_ID environment variable and Google application credentials, list all the GCS buckets and permissions in the project.

Used to review whether there are any GCS buckets with overly-broad permissions.

Required Environment Variables
 
* `PROJECT_ID` - GCP Project ID
* `GOOGLE_APPLICATION_CREDENTIALS` - service account json filepath

Usage

`bucket-review` will look through all buckets for IAM permissions of `allUsers` or `allAuthenticatedUsers` members and output a list of buckets that match

`bucket-review [bucket]` will look at only one bucket specified `bucket` for the same


