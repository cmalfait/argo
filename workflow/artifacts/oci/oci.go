package oci

import (
//	"time"
	"fmt"
	"context"

	"github.com/oracle/oci-go-sdk/v27/common"
	"github.com/oracle/oci-go-sdk/v27/example/helpers"
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
func newOCIClient() (objectstorage.ObjectStorageClient, context.Context) {
//func (ociDriver *ArtifactDriver) newOCIClient() (objectstorage.ObjectStorageClient, context.Context) {
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

//func getNamespace(ctx context.Context, c objectstorage.ObjectStorageClient) string {
func getNamespace(ctx context.Context, c objectstorage.ObjectStorageClient) string {
    request := objectstorage.GetNamespaceRequest{}
    r, _ := c.GetNamespace(ctx, request)
    fmt.Println()
    fmt.Print("Namespace: ")
    return *r.Value
}

func createBucket(ctx context.Context, c objectstorage.ObjectStorageClient, namespace, name string) {
	compartmentID := "ocid1.compartment.oc1..aaaaaaaaidx64b2wndiympj27i3a25riynkbpleenxu56yyjlef3joehboxa"
	var pcompartmentID *string /* pointer variable declaration */
	pcompartmentID = &compartmentID

	request := objectstorage.CreateBucketRequest{
			NamespaceName: &namespace,
	}
//      request.CompartmentId = helpers.CompartmentID()
	request.CompartmentId = pcompartmentID
	request.Name = &name
	request.Metadata = make(map[string]string)
	request.PublicAccessType = objectstorage.CreateBucketDetailsPublicAccessTypeNopublicaccess
	_, err := c.CreateBucket(ctx, request)
	helpers.FatalIfError(err)

	fmt.Println("create bucket")
}

// Load loads an artifact
func (ociDriver *ArtifactDriver) Load(inputArtifact *wfv1.Artifact, path string) error {
	client, ctx := newOCIClient()
	mynamespace := getNamespace(ctx, client)

	createBucket(ctx, client, mynamespace, "my-bucket")

	return nil
}

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
return nil
}
*/

// Save saves an artifact
func (ociDriver *ArtifactDriver) Save(path string, outputArtifact *wfv1.Artifact) error {
	client, ctx := newOCIClient()
	mynamespace := getNamespace(ctx, client)

	createBucket(ctx, client, mynamespace, "my-bucket")

	return nil
}
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
return nil
}
*/
