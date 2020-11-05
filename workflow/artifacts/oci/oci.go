package oci

import (
//	"time"
	"fmt"
	"context"

	"github.com/oracle/oci-go-sdk/v27/common"
//	"github.com/oracle/oci-go-sdk/v27/example/helpers"
	"github.com/oracle/oci-go-sdk/v27/objectstorage"
	//	"github.com/oracle/oci-go-sdk/v27/objectstorage/transfer"
	//	"github.com/oracle/oci-go-sdk/v27/identity

//	log "github.com/sirupsen/logrus"
//	"k8s.io/apimachinery/pkg/util/wait"

//	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

// ArtifactDriver is a driver for Oracle OCI
type ArtifactDriver struct {
	CompartmentOCID  string
}

//func (ociDriver *ArtifactDriver) newOCIClient() (objectstorage.ObjectStorageClient) {
//func newOCIClient() (context.Context, objectstorage.ObjectStorageClient) {
func (ociDriver *ArtifactDriver) newOCIClient() (objectstorage.ObjectStorageClient, context.Context) {
	fmt.Println("creating client driver...")
	client, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	ctx := context.Background()

	return client, ctx
}

/*
func (ociDriver *OCIArtifactDriver) newOCIClient() (*oss.Client, error) {
	client, err := oss.New(ociDriver.Endpoint, ociDriver.AccessKey, ociDriver.SecretKey)
	if err != nil {
		log.Warnf("Failed to create new OSS client: %v", err)
		return nil, err
	}
	return client, err
}
*/

// Load loads an artifact
func (ociDriver *ArtifactDriver) Load(inputArtifact *wfv1.Artifact, path string) error {
/*
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
*/
return nil
}

// Save saves an artifact
func (ociDriver *ArtifactDriver) Save(path string, outputArtifact *wfv1.Artifact) error {
/*
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
*/
return nil
}
