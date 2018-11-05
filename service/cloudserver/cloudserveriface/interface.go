// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudserveriface provides an interface to enable mocking the Cloudserver Client service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudserveriface

import (
	"github.com/arekisannda/zenkocli/service/cloudserver"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// CloudserverAPI provides an interface to enable mocking the
// cloudserver.Cloudserver service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Cloudserver Client.
//    func myFunc(svc cloudserveriface.CloudserverAPI) bool {
//        // Make svc.AbortMultipartUpload request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudserver.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudserverClient struct {
//        cloudserveriface.CloudserverAPI
//    }
//    func (m *mockCloudserverClient) AbortMultipartUpload(input *cloudserver.AbortMultipartUploadInput) (*cloudserver.AbortMultipartUploadOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudserverClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CloudserverAPI interface {
	AbortMultipartUpload(*cloudserver.AbortMultipartUploadInput) (*cloudserver.AbortMultipartUploadOutput, error)
	AbortMultipartUploadWithContext(aws.Context, *cloudserver.AbortMultipartUploadInput, ...request.Option) (*cloudserver.AbortMultipartUploadOutput, error)
	AbortMultipartUploadRequest(*cloudserver.AbortMultipartUploadInput) (*request.Request, *cloudserver.AbortMultipartUploadOutput)

	CompleteMultipartUpload(*cloudserver.CompleteMultipartUploadInput) (*cloudserver.CompleteMultipartUploadOutput, error)
	CompleteMultipartUploadWithContext(aws.Context, *cloudserver.CompleteMultipartUploadInput, ...request.Option) (*cloudserver.CompleteMultipartUploadOutput, error)
	CompleteMultipartUploadRequest(*cloudserver.CompleteMultipartUploadInput) (*request.Request, *cloudserver.CompleteMultipartUploadOutput)

	CopyObject(*cloudserver.CopyObjectInput) (*cloudserver.CopyObjectOutput, error)
	CopyObjectWithContext(aws.Context, *cloudserver.CopyObjectInput, ...request.Option) (*cloudserver.CopyObjectOutput, error)
	CopyObjectRequest(*cloudserver.CopyObjectInput) (*request.Request, *cloudserver.CopyObjectOutput)

	CreateBucket(*cloudserver.CreateBucketInput) (*cloudserver.CreateBucketOutput, error)
	CreateBucketWithContext(aws.Context, *cloudserver.CreateBucketInput, ...request.Option) (*cloudserver.CreateBucketOutput, error)
	CreateBucketRequest(*cloudserver.CreateBucketInput) (*request.Request, *cloudserver.CreateBucketOutput)

	CreateMultipartUpload(*cloudserver.CreateMultipartUploadInput) (*cloudserver.CreateMultipartUploadOutput, error)
	CreateMultipartUploadWithContext(aws.Context, *cloudserver.CreateMultipartUploadInput, ...request.Option) (*cloudserver.CreateMultipartUploadOutput, error)
	CreateMultipartUploadRequest(*cloudserver.CreateMultipartUploadInput) (*request.Request, *cloudserver.CreateMultipartUploadOutput)

	DeleteBucket(*cloudserver.DeleteBucketInput) (*cloudserver.DeleteBucketOutput, error)
	DeleteBucketWithContext(aws.Context, *cloudserver.DeleteBucketInput, ...request.Option) (*cloudserver.DeleteBucketOutput, error)
	DeleteBucketRequest(*cloudserver.DeleteBucketInput) (*request.Request, *cloudserver.DeleteBucketOutput)

	DeleteBucketAnalyticsConfiguration(*cloudserver.DeleteBucketAnalyticsConfigurationInput) (*cloudserver.DeleteBucketAnalyticsConfigurationOutput, error)
	DeleteBucketAnalyticsConfigurationWithContext(aws.Context, *cloudserver.DeleteBucketAnalyticsConfigurationInput, ...request.Option) (*cloudserver.DeleteBucketAnalyticsConfigurationOutput, error)
	DeleteBucketAnalyticsConfigurationRequest(*cloudserver.DeleteBucketAnalyticsConfigurationInput) (*request.Request, *cloudserver.DeleteBucketAnalyticsConfigurationOutput)

	DeleteBucketCors(*cloudserver.DeleteBucketCorsInput) (*cloudserver.DeleteBucketCorsOutput, error)
	DeleteBucketCorsWithContext(aws.Context, *cloudserver.DeleteBucketCorsInput, ...request.Option) (*cloudserver.DeleteBucketCorsOutput, error)
	DeleteBucketCorsRequest(*cloudserver.DeleteBucketCorsInput) (*request.Request, *cloudserver.DeleteBucketCorsOutput)

	DeleteBucketEncryption(*cloudserver.DeleteBucketEncryptionInput) (*cloudserver.DeleteBucketEncryptionOutput, error)
	DeleteBucketEncryptionWithContext(aws.Context, *cloudserver.DeleteBucketEncryptionInput, ...request.Option) (*cloudserver.DeleteBucketEncryptionOutput, error)
	DeleteBucketEncryptionRequest(*cloudserver.DeleteBucketEncryptionInput) (*request.Request, *cloudserver.DeleteBucketEncryptionOutput)

	DeleteBucketInventoryConfiguration(*cloudserver.DeleteBucketInventoryConfigurationInput) (*cloudserver.DeleteBucketInventoryConfigurationOutput, error)
	DeleteBucketInventoryConfigurationWithContext(aws.Context, *cloudserver.DeleteBucketInventoryConfigurationInput, ...request.Option) (*cloudserver.DeleteBucketInventoryConfigurationOutput, error)
	DeleteBucketInventoryConfigurationRequest(*cloudserver.DeleteBucketInventoryConfigurationInput) (*request.Request, *cloudserver.DeleteBucketInventoryConfigurationOutput)

	DeleteBucketLifecycle(*cloudserver.DeleteBucketLifecycleInput) (*cloudserver.DeleteBucketLifecycleOutput, error)
	DeleteBucketLifecycleWithContext(aws.Context, *cloudserver.DeleteBucketLifecycleInput, ...request.Option) (*cloudserver.DeleteBucketLifecycleOutput, error)
	DeleteBucketLifecycleRequest(*cloudserver.DeleteBucketLifecycleInput) (*request.Request, *cloudserver.DeleteBucketLifecycleOutput)

	DeleteBucketMetricsConfiguration(*cloudserver.DeleteBucketMetricsConfigurationInput) (*cloudserver.DeleteBucketMetricsConfigurationOutput, error)
	DeleteBucketMetricsConfigurationWithContext(aws.Context, *cloudserver.DeleteBucketMetricsConfigurationInput, ...request.Option) (*cloudserver.DeleteBucketMetricsConfigurationOutput, error)
	DeleteBucketMetricsConfigurationRequest(*cloudserver.DeleteBucketMetricsConfigurationInput) (*request.Request, *cloudserver.DeleteBucketMetricsConfigurationOutput)

	DeleteBucketPolicy(*cloudserver.DeleteBucketPolicyInput) (*cloudserver.DeleteBucketPolicyOutput, error)
	DeleteBucketPolicyWithContext(aws.Context, *cloudserver.DeleteBucketPolicyInput, ...request.Option) (*cloudserver.DeleteBucketPolicyOutput, error)
	DeleteBucketPolicyRequest(*cloudserver.DeleteBucketPolicyInput) (*request.Request, *cloudserver.DeleteBucketPolicyOutput)

	DeleteBucketReplication(*cloudserver.DeleteBucketReplicationInput) (*cloudserver.DeleteBucketReplicationOutput, error)
	DeleteBucketReplicationWithContext(aws.Context, *cloudserver.DeleteBucketReplicationInput, ...request.Option) (*cloudserver.DeleteBucketReplicationOutput, error)
	DeleteBucketReplicationRequest(*cloudserver.DeleteBucketReplicationInput) (*request.Request, *cloudserver.DeleteBucketReplicationOutput)

	DeleteBucketTagging(*cloudserver.DeleteBucketTaggingInput) (*cloudserver.DeleteBucketTaggingOutput, error)
	DeleteBucketTaggingWithContext(aws.Context, *cloudserver.DeleteBucketTaggingInput, ...request.Option) (*cloudserver.DeleteBucketTaggingOutput, error)
	DeleteBucketTaggingRequest(*cloudserver.DeleteBucketTaggingInput) (*request.Request, *cloudserver.DeleteBucketTaggingOutput)

	DeleteBucketWebsite(*cloudserver.DeleteBucketWebsiteInput) (*cloudserver.DeleteBucketWebsiteOutput, error)
	DeleteBucketWebsiteWithContext(aws.Context, *cloudserver.DeleteBucketWebsiteInput, ...request.Option) (*cloudserver.DeleteBucketWebsiteOutput, error)
	DeleteBucketWebsiteRequest(*cloudserver.DeleteBucketWebsiteInput) (*request.Request, *cloudserver.DeleteBucketWebsiteOutput)

	DeleteObject(*cloudserver.DeleteObjectInput) (*cloudserver.DeleteObjectOutput, error)
	DeleteObjectWithContext(aws.Context, *cloudserver.DeleteObjectInput, ...request.Option) (*cloudserver.DeleteObjectOutput, error)
	DeleteObjectRequest(*cloudserver.DeleteObjectInput) (*request.Request, *cloudserver.DeleteObjectOutput)

	DeleteObjectTagging(*cloudserver.DeleteObjectTaggingInput) (*cloudserver.DeleteObjectTaggingOutput, error)
	DeleteObjectTaggingWithContext(aws.Context, *cloudserver.DeleteObjectTaggingInput, ...request.Option) (*cloudserver.DeleteObjectTaggingOutput, error)
	DeleteObjectTaggingRequest(*cloudserver.DeleteObjectTaggingInput) (*request.Request, *cloudserver.DeleteObjectTaggingOutput)

	DeleteObjects(*cloudserver.DeleteObjectsInput) (*cloudserver.DeleteObjectsOutput, error)
	DeleteObjectsWithContext(aws.Context, *cloudserver.DeleteObjectsInput, ...request.Option) (*cloudserver.DeleteObjectsOutput, error)
	DeleteObjectsRequest(*cloudserver.DeleteObjectsInput) (*request.Request, *cloudserver.DeleteObjectsOutput)

	GetBucketAccelerateConfiguration(*cloudserver.GetBucketAccelerateConfigurationInput) (*cloudserver.GetBucketAccelerateConfigurationOutput, error)
	GetBucketAccelerateConfigurationWithContext(aws.Context, *cloudserver.GetBucketAccelerateConfigurationInput, ...request.Option) (*cloudserver.GetBucketAccelerateConfigurationOutput, error)
	GetBucketAccelerateConfigurationRequest(*cloudserver.GetBucketAccelerateConfigurationInput) (*request.Request, *cloudserver.GetBucketAccelerateConfigurationOutput)

	GetBucketAcl(*cloudserver.GetBucketAclInput) (*cloudserver.GetBucketAclOutput, error)
	GetBucketAclWithContext(aws.Context, *cloudserver.GetBucketAclInput, ...request.Option) (*cloudserver.GetBucketAclOutput, error)
	GetBucketAclRequest(*cloudserver.GetBucketAclInput) (*request.Request, *cloudserver.GetBucketAclOutput)

	GetBucketAnalyticsConfiguration(*cloudserver.GetBucketAnalyticsConfigurationInput) (*cloudserver.GetBucketAnalyticsConfigurationOutput, error)
	GetBucketAnalyticsConfigurationWithContext(aws.Context, *cloudserver.GetBucketAnalyticsConfigurationInput, ...request.Option) (*cloudserver.GetBucketAnalyticsConfigurationOutput, error)
	GetBucketAnalyticsConfigurationRequest(*cloudserver.GetBucketAnalyticsConfigurationInput) (*request.Request, *cloudserver.GetBucketAnalyticsConfigurationOutput)

	GetBucketCors(*cloudserver.GetBucketCorsInput) (*cloudserver.GetBucketCorsOutput, error)
	GetBucketCorsWithContext(aws.Context, *cloudserver.GetBucketCorsInput, ...request.Option) (*cloudserver.GetBucketCorsOutput, error)
	GetBucketCorsRequest(*cloudserver.GetBucketCorsInput) (*request.Request, *cloudserver.GetBucketCorsOutput)

	GetBucketEncryption(*cloudserver.GetBucketEncryptionInput) (*cloudserver.GetBucketEncryptionOutput, error)
	GetBucketEncryptionWithContext(aws.Context, *cloudserver.GetBucketEncryptionInput, ...request.Option) (*cloudserver.GetBucketEncryptionOutput, error)
	GetBucketEncryptionRequest(*cloudserver.GetBucketEncryptionInput) (*request.Request, *cloudserver.GetBucketEncryptionOutput)

	GetBucketInventoryConfiguration(*cloudserver.GetBucketInventoryConfigurationInput) (*cloudserver.GetBucketInventoryConfigurationOutput, error)
	GetBucketInventoryConfigurationWithContext(aws.Context, *cloudserver.GetBucketInventoryConfigurationInput, ...request.Option) (*cloudserver.GetBucketInventoryConfigurationOutput, error)
	GetBucketInventoryConfigurationRequest(*cloudserver.GetBucketInventoryConfigurationInput) (*request.Request, *cloudserver.GetBucketInventoryConfigurationOutput)

	GetBucketLifecycle(*cloudserver.GetBucketLifecycleInput) (*cloudserver.GetBucketLifecycleOutput, error)
	GetBucketLifecycleWithContext(aws.Context, *cloudserver.GetBucketLifecycleInput, ...request.Option) (*cloudserver.GetBucketLifecycleOutput, error)
	GetBucketLifecycleRequest(*cloudserver.GetBucketLifecycleInput) (*request.Request, *cloudserver.GetBucketLifecycleOutput)

	GetBucketLifecycleConfiguration(*cloudserver.GetBucketLifecycleConfigurationInput) (*cloudserver.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationWithContext(aws.Context, *cloudserver.GetBucketLifecycleConfigurationInput, ...request.Option) (*cloudserver.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationRequest(*cloudserver.GetBucketLifecycleConfigurationInput) (*request.Request, *cloudserver.GetBucketLifecycleConfigurationOutput)

	GetBucketLocation(*cloudserver.GetBucketLocationInput) (*cloudserver.GetBucketLocationOutput, error)
	GetBucketLocationWithContext(aws.Context, *cloudserver.GetBucketLocationInput, ...request.Option) (*cloudserver.GetBucketLocationOutput, error)
	GetBucketLocationRequest(*cloudserver.GetBucketLocationInput) (*request.Request, *cloudserver.GetBucketLocationOutput)

	GetBucketLogging(*cloudserver.GetBucketLoggingInput) (*cloudserver.GetBucketLoggingOutput, error)
	GetBucketLoggingWithContext(aws.Context, *cloudserver.GetBucketLoggingInput, ...request.Option) (*cloudserver.GetBucketLoggingOutput, error)
	GetBucketLoggingRequest(*cloudserver.GetBucketLoggingInput) (*request.Request, *cloudserver.GetBucketLoggingOutput)

	GetBucketMetricsConfiguration(*cloudserver.GetBucketMetricsConfigurationInput) (*cloudserver.GetBucketMetricsConfigurationOutput, error)
	GetBucketMetricsConfigurationWithContext(aws.Context, *cloudserver.GetBucketMetricsConfigurationInput, ...request.Option) (*cloudserver.GetBucketMetricsConfigurationOutput, error)
	GetBucketMetricsConfigurationRequest(*cloudserver.GetBucketMetricsConfigurationInput) (*request.Request, *cloudserver.GetBucketMetricsConfigurationOutput)

	GetBucketNotification(*cloudserver.GetBucketNotificationInput) (*cloudserver.GetBucketNotificationOutput, error)
	GetBucketNotificationWithContext(aws.Context, *cloudserver.GetBucketNotificationInput, ...request.Option) (*cloudserver.GetBucketNotificationOutput, error)
	GetBucketNotificationRequest(*cloudserver.GetBucketNotificationInput) (*request.Request, *cloudserver.GetBucketNotificationOutput)

	GetBucketNotificationConfiguration(*cloudserver.GetBucketNotificationConfigurationInput) (*cloudserver.GetBucketNotificationConfigurationOutput, error)
	GetBucketNotificationConfigurationWithContext(aws.Context, *cloudserver.GetBucketNotificationConfigurationInput, ...request.Option) (*cloudserver.GetBucketNotificationConfigurationOutput, error)
	GetBucketNotificationConfigurationRequest(*cloudserver.GetBucketNotificationConfigurationInput) (*request.Request, *cloudserver.GetBucketNotificationConfigurationOutput)

	GetBucketPolicy(*cloudserver.GetBucketPolicyInput) (*cloudserver.GetBucketPolicyOutput, error)
	GetBucketPolicyWithContext(aws.Context, *cloudserver.GetBucketPolicyInput, ...request.Option) (*cloudserver.GetBucketPolicyOutput, error)
	GetBucketPolicyRequest(*cloudserver.GetBucketPolicyInput) (*request.Request, *cloudserver.GetBucketPolicyOutput)

	GetBucketReplication(*cloudserver.GetBucketReplicationInput) (*cloudserver.GetBucketReplicationOutput, error)
	GetBucketReplicationWithContext(aws.Context, *cloudserver.GetBucketReplicationInput, ...request.Option) (*cloudserver.GetBucketReplicationOutput, error)
	GetBucketReplicationRequest(*cloudserver.GetBucketReplicationInput) (*request.Request, *cloudserver.GetBucketReplicationOutput)

	GetBucketRequestPayment(*cloudserver.GetBucketRequestPaymentInput) (*cloudserver.GetBucketRequestPaymentOutput, error)
	GetBucketRequestPaymentWithContext(aws.Context, *cloudserver.GetBucketRequestPaymentInput, ...request.Option) (*cloudserver.GetBucketRequestPaymentOutput, error)
	GetBucketRequestPaymentRequest(*cloudserver.GetBucketRequestPaymentInput) (*request.Request, *cloudserver.GetBucketRequestPaymentOutput)

	GetBucketTagging(*cloudserver.GetBucketTaggingInput) (*cloudserver.GetBucketTaggingOutput, error)
	GetBucketTaggingWithContext(aws.Context, *cloudserver.GetBucketTaggingInput, ...request.Option) (*cloudserver.GetBucketTaggingOutput, error)
	GetBucketTaggingRequest(*cloudserver.GetBucketTaggingInput) (*request.Request, *cloudserver.GetBucketTaggingOutput)

	GetBucketVersioning(*cloudserver.GetBucketVersioningInput) (*cloudserver.GetBucketVersioningOutput, error)
	GetBucketVersioningWithContext(aws.Context, *cloudserver.GetBucketVersioningInput, ...request.Option) (*cloudserver.GetBucketVersioningOutput, error)
	GetBucketVersioningRequest(*cloudserver.GetBucketVersioningInput) (*request.Request, *cloudserver.GetBucketVersioningOutput)

	GetBucketWebsite(*cloudserver.GetBucketWebsiteInput) (*cloudserver.GetBucketWebsiteOutput, error)
	GetBucketWebsiteWithContext(aws.Context, *cloudserver.GetBucketWebsiteInput, ...request.Option) (*cloudserver.GetBucketWebsiteOutput, error)
	GetBucketWebsiteRequest(*cloudserver.GetBucketWebsiteInput) (*request.Request, *cloudserver.GetBucketWebsiteOutput)

	GetObject(*cloudserver.GetObjectInput) (*cloudserver.GetObjectOutput, error)
	GetObjectWithContext(aws.Context, *cloudserver.GetObjectInput, ...request.Option) (*cloudserver.GetObjectOutput, error)
	GetObjectRequest(*cloudserver.GetObjectInput) (*request.Request, *cloudserver.GetObjectOutput)

	GetObjectAcl(*cloudserver.GetObjectAclInput) (*cloudserver.GetObjectAclOutput, error)
	GetObjectAclWithContext(aws.Context, *cloudserver.GetObjectAclInput, ...request.Option) (*cloudserver.GetObjectAclOutput, error)
	GetObjectAclRequest(*cloudserver.GetObjectAclInput) (*request.Request, *cloudserver.GetObjectAclOutput)

	GetObjectTagging(*cloudserver.GetObjectTaggingInput) (*cloudserver.GetObjectTaggingOutput, error)
	GetObjectTaggingWithContext(aws.Context, *cloudserver.GetObjectTaggingInput, ...request.Option) (*cloudserver.GetObjectTaggingOutput, error)
	GetObjectTaggingRequest(*cloudserver.GetObjectTaggingInput) (*request.Request, *cloudserver.GetObjectTaggingOutput)

	GetObjectTorrent(*cloudserver.GetObjectTorrentInput) (*cloudserver.GetObjectTorrentOutput, error)
	GetObjectTorrentWithContext(aws.Context, *cloudserver.GetObjectTorrentInput, ...request.Option) (*cloudserver.GetObjectTorrentOutput, error)
	GetObjectTorrentRequest(*cloudserver.GetObjectTorrentInput) (*request.Request, *cloudserver.GetObjectTorrentOutput)

	HeadBucket(*cloudserver.HeadBucketInput) (*cloudserver.HeadBucketOutput, error)
	HeadBucketWithContext(aws.Context, *cloudserver.HeadBucketInput, ...request.Option) (*cloudserver.HeadBucketOutput, error)
	HeadBucketRequest(*cloudserver.HeadBucketInput) (*request.Request, *cloudserver.HeadBucketOutput)

	HeadObject(*cloudserver.HeadObjectInput) (*cloudserver.HeadObjectOutput, error)
	HeadObjectWithContext(aws.Context, *cloudserver.HeadObjectInput, ...request.Option) (*cloudserver.HeadObjectOutput, error)
	HeadObjectRequest(*cloudserver.HeadObjectInput) (*request.Request, *cloudserver.HeadObjectOutput)

	ListBucketAnalyticsConfigurations(*cloudserver.ListBucketAnalyticsConfigurationsInput) (*cloudserver.ListBucketAnalyticsConfigurationsOutput, error)
	ListBucketAnalyticsConfigurationsWithContext(aws.Context, *cloudserver.ListBucketAnalyticsConfigurationsInput, ...request.Option) (*cloudserver.ListBucketAnalyticsConfigurationsOutput, error)
	ListBucketAnalyticsConfigurationsRequest(*cloudserver.ListBucketAnalyticsConfigurationsInput) (*request.Request, *cloudserver.ListBucketAnalyticsConfigurationsOutput)

	ListBucketInventoryConfigurations(*cloudserver.ListBucketInventoryConfigurationsInput) (*cloudserver.ListBucketInventoryConfigurationsOutput, error)
	ListBucketInventoryConfigurationsWithContext(aws.Context, *cloudserver.ListBucketInventoryConfigurationsInput, ...request.Option) (*cloudserver.ListBucketInventoryConfigurationsOutput, error)
	ListBucketInventoryConfigurationsRequest(*cloudserver.ListBucketInventoryConfigurationsInput) (*request.Request, *cloudserver.ListBucketInventoryConfigurationsOutput)

	ListBucketMetricsConfigurations(*cloudserver.ListBucketMetricsConfigurationsInput) (*cloudserver.ListBucketMetricsConfigurationsOutput, error)
	ListBucketMetricsConfigurationsWithContext(aws.Context, *cloudserver.ListBucketMetricsConfigurationsInput, ...request.Option) (*cloudserver.ListBucketMetricsConfigurationsOutput, error)
	ListBucketMetricsConfigurationsRequest(*cloudserver.ListBucketMetricsConfigurationsInput) (*request.Request, *cloudserver.ListBucketMetricsConfigurationsOutput)

	ListBuckets(*cloudserver.ListBucketsInput) (*cloudserver.ListBucketsOutput, error)
	ListBucketsWithContext(aws.Context, *cloudserver.ListBucketsInput, ...request.Option) (*cloudserver.ListBucketsOutput, error)
	ListBucketsRequest(*cloudserver.ListBucketsInput) (*request.Request, *cloudserver.ListBucketsOutput)

	ListMultipartUploads(*cloudserver.ListMultipartUploadsInput) (*cloudserver.ListMultipartUploadsOutput, error)
	ListMultipartUploadsWithContext(aws.Context, *cloudserver.ListMultipartUploadsInput, ...request.Option) (*cloudserver.ListMultipartUploadsOutput, error)
	ListMultipartUploadsRequest(*cloudserver.ListMultipartUploadsInput) (*request.Request, *cloudserver.ListMultipartUploadsOutput)

	ListMultipartUploadsPages(*cloudserver.ListMultipartUploadsInput, func(*cloudserver.ListMultipartUploadsOutput, bool) bool) error
	ListMultipartUploadsPagesWithContext(aws.Context, *cloudserver.ListMultipartUploadsInput, func(*cloudserver.ListMultipartUploadsOutput, bool) bool, ...request.Option) error

	ListObjectVersions(*cloudserver.ListObjectVersionsInput) (*cloudserver.ListObjectVersionsOutput, error)
	ListObjectVersionsWithContext(aws.Context, *cloudserver.ListObjectVersionsInput, ...request.Option) (*cloudserver.ListObjectVersionsOutput, error)
	ListObjectVersionsRequest(*cloudserver.ListObjectVersionsInput) (*request.Request, *cloudserver.ListObjectVersionsOutput)

	ListObjectVersionsPages(*cloudserver.ListObjectVersionsInput, func(*cloudserver.ListObjectVersionsOutput, bool) bool) error
	ListObjectVersionsPagesWithContext(aws.Context, *cloudserver.ListObjectVersionsInput, func(*cloudserver.ListObjectVersionsOutput, bool) bool, ...request.Option) error

	ListObjects(*cloudserver.ListObjectsInput) (*cloudserver.ListObjectsOutput, error)
	ListObjectsWithContext(aws.Context, *cloudserver.ListObjectsInput, ...request.Option) (*cloudserver.ListObjectsOutput, error)
	ListObjectsRequest(*cloudserver.ListObjectsInput) (*request.Request, *cloudserver.ListObjectsOutput)

	ListObjectsPages(*cloudserver.ListObjectsInput, func(*cloudserver.ListObjectsOutput, bool) bool) error
	ListObjectsPagesWithContext(aws.Context, *cloudserver.ListObjectsInput, func(*cloudserver.ListObjectsOutput, bool) bool, ...request.Option) error

	ListObjectsV2(*cloudserver.ListObjectsV2Input) (*cloudserver.ListObjectsV2Output, error)
	ListObjectsV2WithContext(aws.Context, *cloudserver.ListObjectsV2Input, ...request.Option) (*cloudserver.ListObjectsV2Output, error)
	ListObjectsV2Request(*cloudserver.ListObjectsV2Input) (*request.Request, *cloudserver.ListObjectsV2Output)

	ListObjectsV2Pages(*cloudserver.ListObjectsV2Input, func(*cloudserver.ListObjectsV2Output, bool) bool) error
	ListObjectsV2PagesWithContext(aws.Context, *cloudserver.ListObjectsV2Input, func(*cloudserver.ListObjectsV2Output, bool) bool, ...request.Option) error

	ListParts(*cloudserver.ListPartsInput) (*cloudserver.ListPartsOutput, error)
	ListPartsWithContext(aws.Context, *cloudserver.ListPartsInput, ...request.Option) (*cloudserver.ListPartsOutput, error)
	ListPartsRequest(*cloudserver.ListPartsInput) (*request.Request, *cloudserver.ListPartsOutput)

	ListPartsPages(*cloudserver.ListPartsInput, func(*cloudserver.ListPartsOutput, bool) bool) error
	ListPartsPagesWithContext(aws.Context, *cloudserver.ListPartsInput, func(*cloudserver.ListPartsOutput, bool) bool, ...request.Option) error

	PutBucketAccelerateConfiguration(*cloudserver.PutBucketAccelerateConfigurationInput) (*cloudserver.PutBucketAccelerateConfigurationOutput, error)
	PutBucketAccelerateConfigurationWithContext(aws.Context, *cloudserver.PutBucketAccelerateConfigurationInput, ...request.Option) (*cloudserver.PutBucketAccelerateConfigurationOutput, error)
	PutBucketAccelerateConfigurationRequest(*cloudserver.PutBucketAccelerateConfigurationInput) (*request.Request, *cloudserver.PutBucketAccelerateConfigurationOutput)

	PutBucketAcl(*cloudserver.PutBucketAclInput) (*cloudserver.PutBucketAclOutput, error)
	PutBucketAclWithContext(aws.Context, *cloudserver.PutBucketAclInput, ...request.Option) (*cloudserver.PutBucketAclOutput, error)
	PutBucketAclRequest(*cloudserver.PutBucketAclInput) (*request.Request, *cloudserver.PutBucketAclOutput)

	PutBucketAnalyticsConfiguration(*cloudserver.PutBucketAnalyticsConfigurationInput) (*cloudserver.PutBucketAnalyticsConfigurationOutput, error)
	PutBucketAnalyticsConfigurationWithContext(aws.Context, *cloudserver.PutBucketAnalyticsConfigurationInput, ...request.Option) (*cloudserver.PutBucketAnalyticsConfigurationOutput, error)
	PutBucketAnalyticsConfigurationRequest(*cloudserver.PutBucketAnalyticsConfigurationInput) (*request.Request, *cloudserver.PutBucketAnalyticsConfigurationOutput)

	PutBucketCors(*cloudserver.PutBucketCorsInput) (*cloudserver.PutBucketCorsOutput, error)
	PutBucketCorsWithContext(aws.Context, *cloudserver.PutBucketCorsInput, ...request.Option) (*cloudserver.PutBucketCorsOutput, error)
	PutBucketCorsRequest(*cloudserver.PutBucketCorsInput) (*request.Request, *cloudserver.PutBucketCorsOutput)

	PutBucketEncryption(*cloudserver.PutBucketEncryptionInput) (*cloudserver.PutBucketEncryptionOutput, error)
	PutBucketEncryptionWithContext(aws.Context, *cloudserver.PutBucketEncryptionInput, ...request.Option) (*cloudserver.PutBucketEncryptionOutput, error)
	PutBucketEncryptionRequest(*cloudserver.PutBucketEncryptionInput) (*request.Request, *cloudserver.PutBucketEncryptionOutput)

	PutBucketInventoryConfiguration(*cloudserver.PutBucketInventoryConfigurationInput) (*cloudserver.PutBucketInventoryConfigurationOutput, error)
	PutBucketInventoryConfigurationWithContext(aws.Context, *cloudserver.PutBucketInventoryConfigurationInput, ...request.Option) (*cloudserver.PutBucketInventoryConfigurationOutput, error)
	PutBucketInventoryConfigurationRequest(*cloudserver.PutBucketInventoryConfigurationInput) (*request.Request, *cloudserver.PutBucketInventoryConfigurationOutput)

	PutBucketLifecycle(*cloudserver.PutBucketLifecycleInput) (*cloudserver.PutBucketLifecycleOutput, error)
	PutBucketLifecycleWithContext(aws.Context, *cloudserver.PutBucketLifecycleInput, ...request.Option) (*cloudserver.PutBucketLifecycleOutput, error)
	PutBucketLifecycleRequest(*cloudserver.PutBucketLifecycleInput) (*request.Request, *cloudserver.PutBucketLifecycleOutput)

	PutBucketLifecycleConfiguration(*cloudserver.PutBucketLifecycleConfigurationInput) (*cloudserver.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationWithContext(aws.Context, *cloudserver.PutBucketLifecycleConfigurationInput, ...request.Option) (*cloudserver.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationRequest(*cloudserver.PutBucketLifecycleConfigurationInput) (*request.Request, *cloudserver.PutBucketLifecycleConfigurationOutput)

	PutBucketLogging(*cloudserver.PutBucketLoggingInput) (*cloudserver.PutBucketLoggingOutput, error)
	PutBucketLoggingWithContext(aws.Context, *cloudserver.PutBucketLoggingInput, ...request.Option) (*cloudserver.PutBucketLoggingOutput, error)
	PutBucketLoggingRequest(*cloudserver.PutBucketLoggingInput) (*request.Request, *cloudserver.PutBucketLoggingOutput)

	PutBucketMetricsConfiguration(*cloudserver.PutBucketMetricsConfigurationInput) (*cloudserver.PutBucketMetricsConfigurationOutput, error)
	PutBucketMetricsConfigurationWithContext(aws.Context, *cloudserver.PutBucketMetricsConfigurationInput, ...request.Option) (*cloudserver.PutBucketMetricsConfigurationOutput, error)
	PutBucketMetricsConfigurationRequest(*cloudserver.PutBucketMetricsConfigurationInput) (*request.Request, *cloudserver.PutBucketMetricsConfigurationOutput)

	PutBucketNotification(*cloudserver.PutBucketNotificationInput) (*cloudserver.PutBucketNotificationOutput, error)
	PutBucketNotificationWithContext(aws.Context, *cloudserver.PutBucketNotificationInput, ...request.Option) (*cloudserver.PutBucketNotificationOutput, error)
	PutBucketNotificationRequest(*cloudserver.PutBucketNotificationInput) (*request.Request, *cloudserver.PutBucketNotificationOutput)

	PutBucketNotificationConfiguration(*cloudserver.PutBucketNotificationConfigurationInput) (*cloudserver.PutBucketNotificationConfigurationOutput, error)
	PutBucketNotificationConfigurationWithContext(aws.Context, *cloudserver.PutBucketNotificationConfigurationInput, ...request.Option) (*cloudserver.PutBucketNotificationConfigurationOutput, error)
	PutBucketNotificationConfigurationRequest(*cloudserver.PutBucketNotificationConfigurationInput) (*request.Request, *cloudserver.PutBucketNotificationConfigurationOutput)

	PutBucketPolicy(*cloudserver.PutBucketPolicyInput) (*cloudserver.PutBucketPolicyOutput, error)
	PutBucketPolicyWithContext(aws.Context, *cloudserver.PutBucketPolicyInput, ...request.Option) (*cloudserver.PutBucketPolicyOutput, error)
	PutBucketPolicyRequest(*cloudserver.PutBucketPolicyInput) (*request.Request, *cloudserver.PutBucketPolicyOutput)

	PutBucketReplication(*cloudserver.PutBucketReplicationInput) (*cloudserver.PutBucketReplicationOutput, error)
	PutBucketReplicationWithContext(aws.Context, *cloudserver.PutBucketReplicationInput, ...request.Option) (*cloudserver.PutBucketReplicationOutput, error)
	PutBucketReplicationRequest(*cloudserver.PutBucketReplicationInput) (*request.Request, *cloudserver.PutBucketReplicationOutput)

	PutBucketRequestPayment(*cloudserver.PutBucketRequestPaymentInput) (*cloudserver.PutBucketRequestPaymentOutput, error)
	PutBucketRequestPaymentWithContext(aws.Context, *cloudserver.PutBucketRequestPaymentInput, ...request.Option) (*cloudserver.PutBucketRequestPaymentOutput, error)
	PutBucketRequestPaymentRequest(*cloudserver.PutBucketRequestPaymentInput) (*request.Request, *cloudserver.PutBucketRequestPaymentOutput)

	PutBucketTagging(*cloudserver.PutBucketTaggingInput) (*cloudserver.PutBucketTaggingOutput, error)
	PutBucketTaggingWithContext(aws.Context, *cloudserver.PutBucketTaggingInput, ...request.Option) (*cloudserver.PutBucketTaggingOutput, error)
	PutBucketTaggingRequest(*cloudserver.PutBucketTaggingInput) (*request.Request, *cloudserver.PutBucketTaggingOutput)

	PutBucketVersioning(*cloudserver.PutBucketVersioningInput) (*cloudserver.PutBucketVersioningOutput, error)
	PutBucketVersioningWithContext(aws.Context, *cloudserver.PutBucketVersioningInput, ...request.Option) (*cloudserver.PutBucketVersioningOutput, error)
	PutBucketVersioningRequest(*cloudserver.PutBucketVersioningInput) (*request.Request, *cloudserver.PutBucketVersioningOutput)

	PutBucketWebsite(*cloudserver.PutBucketWebsiteInput) (*cloudserver.PutBucketWebsiteOutput, error)
	PutBucketWebsiteWithContext(aws.Context, *cloudserver.PutBucketWebsiteInput, ...request.Option) (*cloudserver.PutBucketWebsiteOutput, error)
	PutBucketWebsiteRequest(*cloudserver.PutBucketWebsiteInput) (*request.Request, *cloudserver.PutBucketWebsiteOutput)

	PutObject(*cloudserver.PutObjectInput) (*cloudserver.PutObjectOutput, error)
	PutObjectWithContext(aws.Context, *cloudserver.PutObjectInput, ...request.Option) (*cloudserver.PutObjectOutput, error)
	PutObjectRequest(*cloudserver.PutObjectInput) (*request.Request, *cloudserver.PutObjectOutput)

	PutObjectAcl(*cloudserver.PutObjectAclInput) (*cloudserver.PutObjectAclOutput, error)
	PutObjectAclWithContext(aws.Context, *cloudserver.PutObjectAclInput, ...request.Option) (*cloudserver.PutObjectAclOutput, error)
	PutObjectAclRequest(*cloudserver.PutObjectAclInput) (*request.Request, *cloudserver.PutObjectAclOutput)

	PutObjectTagging(*cloudserver.PutObjectTaggingInput) (*cloudserver.PutObjectTaggingOutput, error)
	PutObjectTaggingWithContext(aws.Context, *cloudserver.PutObjectTaggingInput, ...request.Option) (*cloudserver.PutObjectTaggingOutput, error)
	PutObjectTaggingRequest(*cloudserver.PutObjectTaggingInput) (*request.Request, *cloudserver.PutObjectTaggingOutput)

	RestoreObject(*cloudserver.RestoreObjectInput) (*cloudserver.RestoreObjectOutput, error)
	RestoreObjectWithContext(aws.Context, *cloudserver.RestoreObjectInput, ...request.Option) (*cloudserver.RestoreObjectOutput, error)
	RestoreObjectRequest(*cloudserver.RestoreObjectInput) (*request.Request, *cloudserver.RestoreObjectOutput)

	SelectObjectContent(*cloudserver.SelectObjectContentInput) (*cloudserver.SelectObjectContentOutput, error)
	SelectObjectContentWithContext(aws.Context, *cloudserver.SelectObjectContentInput, ...request.Option) (*cloudserver.SelectObjectContentOutput, error)
	SelectObjectContentRequest(*cloudserver.SelectObjectContentInput) (*request.Request, *cloudserver.SelectObjectContentOutput)

	UploadPart(*cloudserver.UploadPartInput) (*cloudserver.UploadPartOutput, error)
	UploadPartWithContext(aws.Context, *cloudserver.UploadPartInput, ...request.Option) (*cloudserver.UploadPartOutput, error)
	UploadPartRequest(*cloudserver.UploadPartInput) (*request.Request, *cloudserver.UploadPartOutput)

	UploadPartCopy(*cloudserver.UploadPartCopyInput) (*cloudserver.UploadPartCopyOutput, error)
	UploadPartCopyWithContext(aws.Context, *cloudserver.UploadPartCopyInput, ...request.Option) (*cloudserver.UploadPartCopyOutput, error)
	UploadPartCopyRequest(*cloudserver.UploadPartCopyInput) (*request.Request, *cloudserver.UploadPartCopyOutput)

	WaitUntilBucketExists(*cloudserver.HeadBucketInput) error
	WaitUntilBucketExistsWithContext(aws.Context, *cloudserver.HeadBucketInput, ...request.WaiterOption) error

	WaitUntilBucketNotExists(*cloudserver.HeadBucketInput) error
	WaitUntilBucketNotExistsWithContext(aws.Context, *cloudserver.HeadBucketInput, ...request.WaiterOption) error

	WaitUntilObjectExists(*cloudserver.HeadObjectInput) error
	WaitUntilObjectExistsWithContext(aws.Context, *cloudserver.HeadObjectInput, ...request.WaiterOption) error

	WaitUntilObjectNotExists(*cloudserver.HeadObjectInput) error
	WaitUntilObjectNotExistsWithContext(aws.Context, *cloudserver.HeadObjectInput, ...request.WaiterOption) error
}

var _ CloudserverAPI = (*cloudserver.Cloudserver)(nil)