package utils

import (
	"github.com/jfrog/jfrog-cli-go/jfrog-cli/utils/config"
	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

func CreateUploadServiceManager(artDetails *config.ArtifactoryDetails, flags *UploadConfiguration, certPath string, minChecksumDeploySize int64) (*artifactory.ArtifactoryServicesManager, error) {
	artAuth, err := artDetails.CreateArtAuthConfig()
	if err != nil {
		return nil, err
	}
	servicesConfig, err := artifactory.NewConfigBuilder().
		SetArtDetails(artAuth).
		SetDryRun(flags.DryRun).
		SetCertificatesPath(certPath).
		SetMinChecksumDeploy(minChecksumDeploySize).
		SetThreads(flags.Threads).
		SetLogger(log.Logger).
		Build()

	return artifactory.New(&artAuth, servicesConfig)
}

type UploadConfiguration struct {
	Deb                   string
	Threads               int
	MinChecksumDeploySize int64
	BuildName             string
	BuildNumber           string
	DryRun                bool
	Symlink               bool
	ExplodeArchive        bool
	ArtDetails            *config.ArtifactoryDetails
	Retries               int
}
