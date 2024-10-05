// Code generated by Kitex v0.11.3. DO NOT EDIT.

package flashkill

import (
	flashkill "client/rpc/kitex_gen/FlashKill"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, b *flashkill.Buyer, seller *flashkill.Seller, callOptions ...callopt.Option) (err error)
	Login(ctx context.Context, b *flashkill.Buyer, seller *flashkill.Seller, callOptions ...callopt.Option) (err error)
	GenToken(ctx context.Context, b *flashkill.Buyer, seller *flashkill.Seller, callOptions ...callopt.Option) (r string, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kFlashKillClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFlashKillClient struct {
	*kClient
}

func (p *kFlashKillClient) Register(ctx context.Context, b *flashkill.Buyer, seller *flashkill.Seller, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, b, seller)
}

func (p *kFlashKillClient) Login(ctx context.Context, b *flashkill.Buyer, seller *flashkill.Seller, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, b, seller)
}

func (p *kFlashKillClient) GenToken(ctx context.Context, b *flashkill.Buyer, seller *flashkill.Seller, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GenToken(ctx, b, seller)
}
