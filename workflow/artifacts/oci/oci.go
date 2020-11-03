package oci

import (
	"time"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

// OCIArtifactDriver is a driver for Oracle OCI
type OCIArtifactDriver struct {
	Endpoint  string
	AccessKey string
	SecretKey string
}

func (ociDriver *OCIArtifactDriver) newOCIClient() (*oss.Client, error) {
	client, err := oss.New(ociDriver.Endpoint, ociDriver.AccessKey, ociDriver.SecretKey)
	if err != nil {
		log.Warnf("Failed to create new OSS client: %v", err)
		return nil, err
	}
	return client, err
}

// Load loads an artifact
func (ociDriver *OCIArtifactDriver) Load(inputArtifact *wfv1.Artifact, path string) error {
	err := wait.ExponentialBackoff(wait.Backoff{Duration: time.Second * 2, Factor: 2.0, Steps: 5, Jitter: 0.1},
		func() (bool, error) {
			log.Infof("OSS Load path: %s, key: %s", path, inputArtifact.OSS.Key)
			osscli, err := ociDriver.newOCIClient()
			if err != nil {
				return false, err
			}
			bucketName := inputArtifact.OSS.Bucket
			bucket, err := osscli.Bucket(bucketName)
			if err != nil {
				return false, err
			}
			objectName := inputArtifact.OSS.Key
			err = bucket.GetObjectToFile(objectName, path)
			if err != nil {
				return false, err
			}
			return true, nil
		})
	return err
}

// Save saves an artifact
func (ociDriver *OCIArtifactDriver) Save(path string, outputArtifact *wfv1.Artifact) error {
	err := wait.ExponentialBackoff(wait.Backoff{Duration: time.Second * 2, Factor: 2.0, Steps: 5, Jitter: 0.1},
		func() (bool, error) {
			log.Infof("OCI Save path: %s, key: %s", path, outputArtifact.OSS.Key)
			osscli, err := ociDriver.newOCIClient()
			if err != nil {
				log.Warnf("Failed to create new OSS client: %v", err)
				return false, nil
			}
			bucketName := outputArtifact.OSS.Bucket
			bucket, err := osscli.Bucket(bucketName)
			if err != nil {
				return false, err
			}
			objectName := outputArtifact.OSS.Key
			err = bucket.PutObjectFromFile(objectName, path)
			if err != nil {
				return false, err
			}
			return true, nil
		})
	return err
}
