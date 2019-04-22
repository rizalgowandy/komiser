// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package organizations provides the client and types for making API
// requests to AWS Organizations.
//
// AWS Organizations is a web service that enables you to consolidate your multiple
// AWS accounts into an organization and centrally manage your accounts and
// their resources.
//
// This guide provides descriptions of the Organizations API. For more information
// about using this service, see the AWS Organizations User Guide (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html).
//
// API Version
//
// This version of the Organizations API Reference documents the Organizations
// API version 2016-11-28.
//
// As an alternative to using the API directly, you can use one of the AWS SDKs,
// which consist of libraries and sample code for various programming languages
// and platforms (Java, Ruby, .NET, iOS, Android, and more). The SDKs provide
// a convenient way to create programmatic access to AWS Organizations. For
// example, the SDKs take care of cryptographically signing requests, managing
// errors, and retrying requests automatically. For more information about the
// AWS SDKs, including how to download and install them, see Tools for Amazon
// Web Services (http://aws.amazon.com/tools/).
//
// We recommend that you use the AWS SDKs to make programmatic API calls to
// Organizations. However, you also can use the Organizations Query API to make
// direct calls to the Organizations web service. To learn more about the Organizations
// Query API, see Making Query Requests (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_query-requests.html)
// in the AWS Organizations User Guide. Organizations supports GET and POST
// requests for all actions. That is, the API does not require you to use GET
// for some actions and POST for others. However, GET requests are subject to
// the limitation size of a URL. Therefore, for operations that require larger
// sizes, use a POST request.
//
// Signing Requests
//
// When you send HTTP requests to AWS, you must sign the requests so that AWS
// can identify who sent them. You sign requests with your AWS access key, which
// consists of an access key ID and a secret access key. We strongly recommend
// that you do not create an access key for your root account. Anyone who has
// the access key for your root account has unrestricted access to all the resources
// in your account. Instead, create an access key for an IAM user account that
// has administrative privileges. As another option, use AWS Security Token
// Service to generate temporary security credentials, and use those credentials
// to sign requests.
//
// To sign requests, we recommend that you use Signature Version 4 (http://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
// If you have an existing application that uses Signature Version 2, you do
// not have to update it to use Signature Version 4. However, some operations
// now require Signature Version 4. The documentation for operations that require
// version 4 indicate this requirement.
//
// When you use the AWS Command Line Interface (AWS CLI) or one of the AWS SDKs
// to make requests to AWS, these tools automatically sign the requests for
// you with the access key that you specify when you configure the tools.
//
// In this release, each organization can have only one root. In a future release,
// a single organization will support multiple roots.
//
// Support and Feedback for AWS Organizations
//
// We welcome your feedback. Send your comments to feedback-awsorganizations@amazon.com
// (mailto:feedback-awsorganizations@amazon.com) or post your feedback and questions
// in the AWS Organizations support forum (http://forums.aws.amazon.com/forum.jspa?forumID=219).
// For more information about the AWS support forums, see Forums Help (http://forums.aws.amazon.com/help.jspa).
//
// Endpoint to Call When Using the CLI or the AWS API
//
// For the current release of Organizations, you must specify the us-east-1
// region for all AWS API and CLI calls. You can do this in the CLI by using
// these parameters and commands:
//
//    * Use the following parameter with each command to specify both the endpoint
//    and its region:
//
// --endpoint-url https://organizations.us-east-1.amazonaws.com
//
//    * Use the default endpoint, but configure your default region with this
//    command:
//
// aws configure set default.region us-east-1
//
//    * Use the following parameter with each command to specify the endpoint:
//
// --region us-east-1
//
// For the various SDKs used to call the APIs, see the documentation for the
// SDK of interest to learn how to direct the requests to a specific endpoint.
// For more information, see Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#sts_region)
// in the AWS General Reference.
//
// How examples are presented
//
// The JSON returned by the AWS Organizations service as response to your requests
// is returned as a single long string without line breaks or formatting whitespace.
// Both line breaks and whitespace are included in the examples in this guide
// to improve readability. When example input parameters also would result in
// long strings that would extend beyond the screen, we insert line breaks to
// enhance readability. You should always submit the input as a single JSON
// text string.
//
// Recording API Requests
//
// AWS Organizations supports AWS CloudTrail, a service that records AWS API
// calls for your AWS account and delivers log files to an Amazon S3 bucket.
// By using information collected by AWS CloudTrail, you can determine which
// requests were successfully made to Organizations, who made the request, when
// it was made, and so on. For more about AWS Organizations and its support
// for AWS CloudTrail, see Logging AWS Organizations Events with AWS CloudTrail
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_monitoring.html#orgs_cloudtrail-integration)
// in the AWS Organizations User Guide. To learn more about CloudTrail, including
// how to turn it on and find your log files, see the AWS CloudTrail User Guide
// (http://docs.aws.amazon.com/awscloudtrail/latest/userguide/what_is_cloud_trail_top_level.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28 for more information on this service.
//
// See organizations package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/organizations/
//
// Using the Client
//
// To AWS Organizations with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Organizations client Organizations for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/organizations/#New
package organizations
