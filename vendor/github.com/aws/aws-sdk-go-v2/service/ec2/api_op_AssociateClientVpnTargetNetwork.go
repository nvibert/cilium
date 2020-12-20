// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a target network with a Client VPN endpoint. A target network is a
// subnet in a VPC. You can associate multiple subnets from the same VPC with a
// Client VPN endpoint. You can associate only one subnet in each Availability
// Zone. We recommend that you associate at least two subnets to provide
// Availability Zone redundancy. If you specified a VPC when you created the Client
// VPN endpoint or if you have previous subnet associations, the specified subnet
// must be in the same VPC. To specify a subnet that's in a different VPC, you must
// first modify the Client VPN endpoint (ModifyClientVpnEndpoint) and change the
// VPC that's associated with it.
func (c *Client) AssociateClientVpnTargetNetwork(ctx context.Context, params *AssociateClientVpnTargetNetworkInput, optFns ...func(*Options)) (*AssociateClientVpnTargetNetworkOutput, error) {
	if params == nil {
		params = &AssociateClientVpnTargetNetworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateClientVpnTargetNetwork", params, optFns, addOperationAssociateClientVpnTargetNetworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateClientVpnTargetNetworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateClientVpnTargetNetworkInput struct {

	// The ID of the Client VPN endpoint.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// The ID of the subnet to associate with the Client VPN endpoint.
	//
	// This member is required.
	SubnetId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

type AssociateClientVpnTargetNetworkOutput struct {

	// The unique ID of the target network association.
	AssociationId *string

	// The current state of the target network association.
	Status *types.AssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateClientVpnTargetNetworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssociateClientVpnTargetNetwork{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateClientVpnTargetNetwork{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opAssociateClientVpnTargetNetworkMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAssociateClientVpnTargetNetworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateClientVpnTargetNetwork(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpAssociateClientVpnTargetNetwork struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpAssociateClientVpnTargetNetwork) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpAssociateClientVpnTargetNetwork) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*AssociateClientVpnTargetNetworkInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *AssociateClientVpnTargetNetworkInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opAssociateClientVpnTargetNetworkMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpAssociateClientVpnTargetNetwork{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opAssociateClientVpnTargetNetwork(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssociateClientVpnTargetNetwork",
	}
}
