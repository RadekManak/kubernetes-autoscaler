// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package neptunedataiface provides an interface to enable mocking the Amazon NeptuneData service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package neptunedataiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/neptunedata"
)

// NeptunedataAPI provides an interface to enable mocking the
// neptunedata.Neptunedata service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon NeptuneData.
//	func myFunc(svc neptunedataiface.NeptunedataAPI) bool {
//	    // Make svc.CancelGremlinQuery request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := neptunedata.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockNeptunedataClient struct {
//	    neptunedataiface.NeptunedataAPI
//	}
//	func (m *mockNeptunedataClient) CancelGremlinQuery(input *neptunedata.CancelGremlinQueryInput) (*neptunedata.CancelGremlinQueryOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockNeptunedataClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type NeptunedataAPI interface {
	CancelGremlinQuery(*neptunedata.CancelGremlinQueryInput) (*neptunedata.CancelGremlinQueryOutput, error)
	CancelGremlinQueryWithContext(aws.Context, *neptunedata.CancelGremlinQueryInput, ...request.Option) (*neptunedata.CancelGremlinQueryOutput, error)
	CancelGremlinQueryRequest(*neptunedata.CancelGremlinQueryInput) (*request.Request, *neptunedata.CancelGremlinQueryOutput)

	CancelLoaderJob(*neptunedata.CancelLoaderJobInput) (*neptunedata.CancelLoaderJobOutput, error)
	CancelLoaderJobWithContext(aws.Context, *neptunedata.CancelLoaderJobInput, ...request.Option) (*neptunedata.CancelLoaderJobOutput, error)
	CancelLoaderJobRequest(*neptunedata.CancelLoaderJobInput) (*request.Request, *neptunedata.CancelLoaderJobOutput)

	CancelMLDataProcessingJob(*neptunedata.CancelMLDataProcessingJobInput) (*neptunedata.CancelMLDataProcessingJobOutput, error)
	CancelMLDataProcessingJobWithContext(aws.Context, *neptunedata.CancelMLDataProcessingJobInput, ...request.Option) (*neptunedata.CancelMLDataProcessingJobOutput, error)
	CancelMLDataProcessingJobRequest(*neptunedata.CancelMLDataProcessingJobInput) (*request.Request, *neptunedata.CancelMLDataProcessingJobOutput)

	CancelMLModelTrainingJob(*neptunedata.CancelMLModelTrainingJobInput) (*neptunedata.CancelMLModelTrainingJobOutput, error)
	CancelMLModelTrainingJobWithContext(aws.Context, *neptunedata.CancelMLModelTrainingJobInput, ...request.Option) (*neptunedata.CancelMLModelTrainingJobOutput, error)
	CancelMLModelTrainingJobRequest(*neptunedata.CancelMLModelTrainingJobInput) (*request.Request, *neptunedata.CancelMLModelTrainingJobOutput)

	CancelMLModelTransformJob(*neptunedata.CancelMLModelTransformJobInput) (*neptunedata.CancelMLModelTransformJobOutput, error)
	CancelMLModelTransformJobWithContext(aws.Context, *neptunedata.CancelMLModelTransformJobInput, ...request.Option) (*neptunedata.CancelMLModelTransformJobOutput, error)
	CancelMLModelTransformJobRequest(*neptunedata.CancelMLModelTransformJobInput) (*request.Request, *neptunedata.CancelMLModelTransformJobOutput)

	CancelOpenCypherQuery(*neptunedata.CancelOpenCypherQueryInput) (*neptunedata.CancelOpenCypherQueryOutput, error)
	CancelOpenCypherQueryWithContext(aws.Context, *neptunedata.CancelOpenCypherQueryInput, ...request.Option) (*neptunedata.CancelOpenCypherQueryOutput, error)
	CancelOpenCypherQueryRequest(*neptunedata.CancelOpenCypherQueryInput) (*request.Request, *neptunedata.CancelOpenCypherQueryOutput)

	CreateMLEndpoint(*neptunedata.CreateMLEndpointInput) (*neptunedata.CreateMLEndpointOutput, error)
	CreateMLEndpointWithContext(aws.Context, *neptunedata.CreateMLEndpointInput, ...request.Option) (*neptunedata.CreateMLEndpointOutput, error)
	CreateMLEndpointRequest(*neptunedata.CreateMLEndpointInput) (*request.Request, *neptunedata.CreateMLEndpointOutput)

	DeleteMLEndpoint(*neptunedata.DeleteMLEndpointInput) (*neptunedata.DeleteMLEndpointOutput, error)
	DeleteMLEndpointWithContext(aws.Context, *neptunedata.DeleteMLEndpointInput, ...request.Option) (*neptunedata.DeleteMLEndpointOutput, error)
	DeleteMLEndpointRequest(*neptunedata.DeleteMLEndpointInput) (*request.Request, *neptunedata.DeleteMLEndpointOutput)

	DeletePropertygraphStatistics(*neptunedata.DeletePropertygraphStatisticsInput) (*neptunedata.DeletePropertygraphStatisticsOutput, error)
	DeletePropertygraphStatisticsWithContext(aws.Context, *neptunedata.DeletePropertygraphStatisticsInput, ...request.Option) (*neptunedata.DeletePropertygraphStatisticsOutput, error)
	DeletePropertygraphStatisticsRequest(*neptunedata.DeletePropertygraphStatisticsInput) (*request.Request, *neptunedata.DeletePropertygraphStatisticsOutput)

	DeleteSparqlStatistics(*neptunedata.DeleteSparqlStatisticsInput) (*neptunedata.DeleteSparqlStatisticsOutput, error)
	DeleteSparqlStatisticsWithContext(aws.Context, *neptunedata.DeleteSparqlStatisticsInput, ...request.Option) (*neptunedata.DeleteSparqlStatisticsOutput, error)
	DeleteSparqlStatisticsRequest(*neptunedata.DeleteSparqlStatisticsInput) (*request.Request, *neptunedata.DeleteSparqlStatisticsOutput)

	ExecuteFastReset(*neptunedata.ExecuteFastResetInput) (*neptunedata.ExecuteFastResetOutput, error)
	ExecuteFastResetWithContext(aws.Context, *neptunedata.ExecuteFastResetInput, ...request.Option) (*neptunedata.ExecuteFastResetOutput, error)
	ExecuteFastResetRequest(*neptunedata.ExecuteFastResetInput) (*request.Request, *neptunedata.ExecuteFastResetOutput)

	ExecuteGremlinExplainQuery(*neptunedata.ExecuteGremlinExplainQueryInput) (*neptunedata.ExecuteGremlinExplainQueryOutput, error)
	ExecuteGremlinExplainQueryWithContext(aws.Context, *neptunedata.ExecuteGremlinExplainQueryInput, ...request.Option) (*neptunedata.ExecuteGremlinExplainQueryOutput, error)
	ExecuteGremlinExplainQueryRequest(*neptunedata.ExecuteGremlinExplainQueryInput) (*request.Request, *neptunedata.ExecuteGremlinExplainQueryOutput)

	ExecuteGremlinProfileQuery(*neptunedata.ExecuteGremlinProfileQueryInput) (*neptunedata.ExecuteGremlinProfileQueryOutput, error)
	ExecuteGremlinProfileQueryWithContext(aws.Context, *neptunedata.ExecuteGremlinProfileQueryInput, ...request.Option) (*neptunedata.ExecuteGremlinProfileQueryOutput, error)
	ExecuteGremlinProfileQueryRequest(*neptunedata.ExecuteGremlinProfileQueryInput) (*request.Request, *neptunedata.ExecuteGremlinProfileQueryOutput)

	ExecuteGremlinQuery(*neptunedata.ExecuteGremlinQueryInput) (*neptunedata.ExecuteGremlinQueryOutput, error)
	ExecuteGremlinQueryWithContext(aws.Context, *neptunedata.ExecuteGremlinQueryInput, ...request.Option) (*neptunedata.ExecuteGremlinQueryOutput, error)
	ExecuteGremlinQueryRequest(*neptunedata.ExecuteGremlinQueryInput) (*request.Request, *neptunedata.ExecuteGremlinQueryOutput)

	ExecuteOpenCypherExplainQuery(*neptunedata.ExecuteOpenCypherExplainQueryInput) (*neptunedata.ExecuteOpenCypherExplainQueryOutput, error)
	ExecuteOpenCypherExplainQueryWithContext(aws.Context, *neptunedata.ExecuteOpenCypherExplainQueryInput, ...request.Option) (*neptunedata.ExecuteOpenCypherExplainQueryOutput, error)
	ExecuteOpenCypherExplainQueryRequest(*neptunedata.ExecuteOpenCypherExplainQueryInput) (*request.Request, *neptunedata.ExecuteOpenCypherExplainQueryOutput)

	GetEngineStatus(*neptunedata.GetEngineStatusInput) (*neptunedata.GetEngineStatusOutput, error)
	GetEngineStatusWithContext(aws.Context, *neptunedata.GetEngineStatusInput, ...request.Option) (*neptunedata.GetEngineStatusOutput, error)
	GetEngineStatusRequest(*neptunedata.GetEngineStatusInput) (*request.Request, *neptunedata.GetEngineStatusOutput)

	GetGremlinQueryStatus(*neptunedata.GetGremlinQueryStatusInput) (*neptunedata.GetGremlinQueryStatusOutput, error)
	GetGremlinQueryStatusWithContext(aws.Context, *neptunedata.GetGremlinQueryStatusInput, ...request.Option) (*neptunedata.GetGremlinQueryStatusOutput, error)
	GetGremlinQueryStatusRequest(*neptunedata.GetGremlinQueryStatusInput) (*request.Request, *neptunedata.GetGremlinQueryStatusOutput)

	GetMLDataProcessingJob(*neptunedata.GetMLDataProcessingJobInput) (*neptunedata.GetMLDataProcessingJobOutput, error)
	GetMLDataProcessingJobWithContext(aws.Context, *neptunedata.GetMLDataProcessingJobInput, ...request.Option) (*neptunedata.GetMLDataProcessingJobOutput, error)
	GetMLDataProcessingJobRequest(*neptunedata.GetMLDataProcessingJobInput) (*request.Request, *neptunedata.GetMLDataProcessingJobOutput)

	GetMLEndpoint(*neptunedata.GetMLEndpointInput) (*neptunedata.GetMLEndpointOutput, error)
	GetMLEndpointWithContext(aws.Context, *neptunedata.GetMLEndpointInput, ...request.Option) (*neptunedata.GetMLEndpointOutput, error)
	GetMLEndpointRequest(*neptunedata.GetMLEndpointInput) (*request.Request, *neptunedata.GetMLEndpointOutput)

	GetMLModelTrainingJob(*neptunedata.GetMLModelTrainingJobInput) (*neptunedata.GetMLModelTrainingJobOutput, error)
	GetMLModelTrainingJobWithContext(aws.Context, *neptunedata.GetMLModelTrainingJobInput, ...request.Option) (*neptunedata.GetMLModelTrainingJobOutput, error)
	GetMLModelTrainingJobRequest(*neptunedata.GetMLModelTrainingJobInput) (*request.Request, *neptunedata.GetMLModelTrainingJobOutput)

	GetMLModelTransformJob(*neptunedata.GetMLModelTransformJobInput) (*neptunedata.GetMLModelTransformJobOutput, error)
	GetMLModelTransformJobWithContext(aws.Context, *neptunedata.GetMLModelTransformJobInput, ...request.Option) (*neptunedata.GetMLModelTransformJobOutput, error)
	GetMLModelTransformJobRequest(*neptunedata.GetMLModelTransformJobInput) (*request.Request, *neptunedata.GetMLModelTransformJobOutput)

	GetOpenCypherQueryStatus(*neptunedata.GetOpenCypherQueryStatusInput) (*neptunedata.GetOpenCypherQueryStatusOutput, error)
	GetOpenCypherQueryStatusWithContext(aws.Context, *neptunedata.GetOpenCypherQueryStatusInput, ...request.Option) (*neptunedata.GetOpenCypherQueryStatusOutput, error)
	GetOpenCypherQueryStatusRequest(*neptunedata.GetOpenCypherQueryStatusInput) (*request.Request, *neptunedata.GetOpenCypherQueryStatusOutput)

	GetPropertygraphStatistics(*neptunedata.GetPropertygraphStatisticsInput) (*neptunedata.GetPropertygraphStatisticsOutput, error)
	GetPropertygraphStatisticsWithContext(aws.Context, *neptunedata.GetPropertygraphStatisticsInput, ...request.Option) (*neptunedata.GetPropertygraphStatisticsOutput, error)
	GetPropertygraphStatisticsRequest(*neptunedata.GetPropertygraphStatisticsInput) (*request.Request, *neptunedata.GetPropertygraphStatisticsOutput)

	GetPropertygraphSummary(*neptunedata.GetPropertygraphSummaryInput) (*neptunedata.GetPropertygraphSummaryOutput, error)
	GetPropertygraphSummaryWithContext(aws.Context, *neptunedata.GetPropertygraphSummaryInput, ...request.Option) (*neptunedata.GetPropertygraphSummaryOutput, error)
	GetPropertygraphSummaryRequest(*neptunedata.GetPropertygraphSummaryInput) (*request.Request, *neptunedata.GetPropertygraphSummaryOutput)

	GetRDFGraphSummary(*neptunedata.GetRDFGraphSummaryInput) (*neptunedata.GetRDFGraphSummaryOutput, error)
	GetRDFGraphSummaryWithContext(aws.Context, *neptunedata.GetRDFGraphSummaryInput, ...request.Option) (*neptunedata.GetRDFGraphSummaryOutput, error)
	GetRDFGraphSummaryRequest(*neptunedata.GetRDFGraphSummaryInput) (*request.Request, *neptunedata.GetRDFGraphSummaryOutput)

	GetSparqlStatistics(*neptunedata.GetSparqlStatisticsInput) (*neptunedata.GetSparqlStatisticsOutput, error)
	GetSparqlStatisticsWithContext(aws.Context, *neptunedata.GetSparqlStatisticsInput, ...request.Option) (*neptunedata.GetSparqlStatisticsOutput, error)
	GetSparqlStatisticsRequest(*neptunedata.GetSparqlStatisticsInput) (*request.Request, *neptunedata.GetSparqlStatisticsOutput)

	GetSparqlStream(*neptunedata.GetSparqlStreamInput) (*neptunedata.GetSparqlStreamOutput, error)
	GetSparqlStreamWithContext(aws.Context, *neptunedata.GetSparqlStreamInput, ...request.Option) (*neptunedata.GetSparqlStreamOutput, error)
	GetSparqlStreamRequest(*neptunedata.GetSparqlStreamInput) (*request.Request, *neptunedata.GetSparqlStreamOutput)

	ListGremlinQueries(*neptunedata.ListGremlinQueriesInput) (*neptunedata.ListGremlinQueriesOutput, error)
	ListGremlinQueriesWithContext(aws.Context, *neptunedata.ListGremlinQueriesInput, ...request.Option) (*neptunedata.ListGremlinQueriesOutput, error)
	ListGremlinQueriesRequest(*neptunedata.ListGremlinQueriesInput) (*request.Request, *neptunedata.ListGremlinQueriesOutput)

	ListLoaderJobs(*neptunedata.ListLoaderJobsInput) (*neptunedata.ListLoaderJobsOutput, error)
	ListLoaderJobsWithContext(aws.Context, *neptunedata.ListLoaderJobsInput, ...request.Option) (*neptunedata.ListLoaderJobsOutput, error)
	ListLoaderJobsRequest(*neptunedata.ListLoaderJobsInput) (*request.Request, *neptunedata.ListLoaderJobsOutput)

	ListMLDataProcessingJobs(*neptunedata.ListMLDataProcessingJobsInput) (*neptunedata.ListMLDataProcessingJobsOutput, error)
	ListMLDataProcessingJobsWithContext(aws.Context, *neptunedata.ListMLDataProcessingJobsInput, ...request.Option) (*neptunedata.ListMLDataProcessingJobsOutput, error)
	ListMLDataProcessingJobsRequest(*neptunedata.ListMLDataProcessingJobsInput) (*request.Request, *neptunedata.ListMLDataProcessingJobsOutput)

	ListMLEndpoints(*neptunedata.ListMLEndpointsInput) (*neptunedata.ListMLEndpointsOutput, error)
	ListMLEndpointsWithContext(aws.Context, *neptunedata.ListMLEndpointsInput, ...request.Option) (*neptunedata.ListMLEndpointsOutput, error)
	ListMLEndpointsRequest(*neptunedata.ListMLEndpointsInput) (*request.Request, *neptunedata.ListMLEndpointsOutput)

	ListMLModelTrainingJobs(*neptunedata.ListMLModelTrainingJobsInput) (*neptunedata.ListMLModelTrainingJobsOutput, error)
	ListMLModelTrainingJobsWithContext(aws.Context, *neptunedata.ListMLModelTrainingJobsInput, ...request.Option) (*neptunedata.ListMLModelTrainingJobsOutput, error)
	ListMLModelTrainingJobsRequest(*neptunedata.ListMLModelTrainingJobsInput) (*request.Request, *neptunedata.ListMLModelTrainingJobsOutput)

	ListMLModelTransformJobs(*neptunedata.ListMLModelTransformJobsInput) (*neptunedata.ListMLModelTransformJobsOutput, error)
	ListMLModelTransformJobsWithContext(aws.Context, *neptunedata.ListMLModelTransformJobsInput, ...request.Option) (*neptunedata.ListMLModelTransformJobsOutput, error)
	ListMLModelTransformJobsRequest(*neptunedata.ListMLModelTransformJobsInput) (*request.Request, *neptunedata.ListMLModelTransformJobsOutput)

	ListOpenCypherQueries(*neptunedata.ListOpenCypherQueriesInput) (*neptunedata.ListOpenCypherQueriesOutput, error)
	ListOpenCypherQueriesWithContext(aws.Context, *neptunedata.ListOpenCypherQueriesInput, ...request.Option) (*neptunedata.ListOpenCypherQueriesOutput, error)
	ListOpenCypherQueriesRequest(*neptunedata.ListOpenCypherQueriesInput) (*request.Request, *neptunedata.ListOpenCypherQueriesOutput)

	ManagePropertygraphStatistics(*neptunedata.ManagePropertygraphStatisticsInput) (*neptunedata.ManagePropertygraphStatisticsOutput, error)
	ManagePropertygraphStatisticsWithContext(aws.Context, *neptunedata.ManagePropertygraphStatisticsInput, ...request.Option) (*neptunedata.ManagePropertygraphStatisticsOutput, error)
	ManagePropertygraphStatisticsRequest(*neptunedata.ManagePropertygraphStatisticsInput) (*request.Request, *neptunedata.ManagePropertygraphStatisticsOutput)

	ManageSparqlStatistics(*neptunedata.ManageSparqlStatisticsInput) (*neptunedata.ManageSparqlStatisticsOutput, error)
	ManageSparqlStatisticsWithContext(aws.Context, *neptunedata.ManageSparqlStatisticsInput, ...request.Option) (*neptunedata.ManageSparqlStatisticsOutput, error)
	ManageSparqlStatisticsRequest(*neptunedata.ManageSparqlStatisticsInput) (*request.Request, *neptunedata.ManageSparqlStatisticsOutput)

	StartLoaderJob(*neptunedata.StartLoaderJobInput) (*neptunedata.StartLoaderJobOutput, error)
	StartLoaderJobWithContext(aws.Context, *neptunedata.StartLoaderJobInput, ...request.Option) (*neptunedata.StartLoaderJobOutput, error)
	StartLoaderJobRequest(*neptunedata.StartLoaderJobInput) (*request.Request, *neptunedata.StartLoaderJobOutput)

	StartMLDataProcessingJob(*neptunedata.StartMLDataProcessingJobInput) (*neptunedata.StartMLDataProcessingJobOutput, error)
	StartMLDataProcessingJobWithContext(aws.Context, *neptunedata.StartMLDataProcessingJobInput, ...request.Option) (*neptunedata.StartMLDataProcessingJobOutput, error)
	StartMLDataProcessingJobRequest(*neptunedata.StartMLDataProcessingJobInput) (*request.Request, *neptunedata.StartMLDataProcessingJobOutput)

	StartMLModelTrainingJob(*neptunedata.StartMLModelTrainingJobInput) (*neptunedata.StartMLModelTrainingJobOutput, error)
	StartMLModelTrainingJobWithContext(aws.Context, *neptunedata.StartMLModelTrainingJobInput, ...request.Option) (*neptunedata.StartMLModelTrainingJobOutput, error)
	StartMLModelTrainingJobRequest(*neptunedata.StartMLModelTrainingJobInput) (*request.Request, *neptunedata.StartMLModelTrainingJobOutput)

	StartMLModelTransformJob(*neptunedata.StartMLModelTransformJobInput) (*neptunedata.StartMLModelTransformJobOutput, error)
	StartMLModelTransformJobWithContext(aws.Context, *neptunedata.StartMLModelTransformJobInput, ...request.Option) (*neptunedata.StartMLModelTransformJobOutput, error)
	StartMLModelTransformJobRequest(*neptunedata.StartMLModelTransformJobInput) (*request.Request, *neptunedata.StartMLModelTransformJobOutput)
}

var _ NeptunedataAPI = (*neptunedata.Neptunedata)(nil)