// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a target network from the specified Client VPN endpoint. When you
// disassociate the last target network from a Client VPN, the following
// happens:
//
// * The route that was automatically added for the VPC is deleted
//
// * All
// active client connections are terminated
//
// * New client connections are
// disallowed
//
// * The Client VPN endpoint's status changes to pending-associate
func (c *Client) DisassociateClientVpnTargetNetwork(ctx context.Context, params *DisassociateClientVpnTargetNetworkInput, optFns ...func(*Options)) (*DisassociateClientVpnTargetNetworkOutput, error) {
	if params == nil {
		params = &DisassociateClientVpnTargetNetworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateClientVpnTargetNetwork", params, optFns, addOperationDisassociateClientVpnTargetNetworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateClientVpnTargetNetworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateClientVpnTargetNetworkInput struct {

	// The ID of the target network association.
	//
	// This member is required.
	AssociationId *string

	// The ID of the Client VPN endpoint from which to disassociate the target network.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

type DisassociateClientVpnTargetNetworkOutput struct {

	// The ID of the target network association.
	AssociationId *string

	// The current state of the target network association.
	Status *types.AssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateClientVpnTargetNetworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDisassociateClientVpnTargetNetwork{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDisassociateClientVpnTargetNetwork{}, middleware.After)
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
	if err = addOpDisassociateClientVpnTargetNetworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateClientVpnTargetNetwork(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateClientVpnTargetNetwork(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisassociateClientVpnTargetNetwork",
	}
}
