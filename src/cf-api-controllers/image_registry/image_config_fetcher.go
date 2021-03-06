package image_registry

import (
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/pivotal/kpack/pkg/registry"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ImageConfigFetcher
type ImageConfigFetcher interface {
	FetchImageConfig(imageReference, buildServiceAccount, buildNamespace string) (*v1.Config, error)
}

type ociImageConfigFetcher struct {
	KeychainFactory    registry.KeychainFactory
	ImageConfigFetcher ImageConfigFetcher
}

func NewImageConfigFetcher(keychainFactory registry.KeychainFactory) ociImageConfigFetcher {
	return ociImageConfigFetcher{KeychainFactory: keychainFactory}
}

func (f ociImageConfigFetcher) FetchImageConfig(imageReference, secretServiceAccount, secretNamespace string) (*v1.Config, error) {
	ref, err := name.ParseReference(imageReference)
	if err != nil {
		return nil, err
	}

	keychain, err := f.KeychainFactory.KeychainForSecretRef(registry.SecretRef{
		ServiceAccount: secretServiceAccount,
		Namespace:      secretNamespace,
	})
	if err != nil {
		return nil, err
	}

	img, err := remote.Image(ref, remote.WithAuthFromKeychain(keychain))
	if err != nil {
		return nil, err
	}

	cfgFile, err := img.ConfigFile()
	if err != nil {
		return nil, err
	}

	// TODO: address potential nil-pointer deref here
	return &cfgFile.Config, nil
}
