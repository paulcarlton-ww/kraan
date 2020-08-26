package layers // nolint:package // unit tests should be in same package as code under test

/*

The mockgen tool generates the MockExecProvider type defined in the kubectl/mockExecProvider.go code file.

From the project root directory, you can generate mock definitions for interfaces in individual code files by calling mockgen.  Example:
	mockgen -destination=pkg/internal/kubectl/mockExecProvider.go -package=kubectl -source=pkg/internal/kubectl/execProvider.go \
	gitlab.fmr.com/common-platform/addons-manager/pkg/internal/kubectl ExecProvider

Or you can generate all the

Add a go:generate annotation above the package statement in all the code files containing interfaces that you want to mock.  Example:
//go:generate mockgen -destination=mockExecProvider.go -package=kubectl -source=execProvider.go . ExecProvider
//go:generate mockgen -destination=../mocks/logr/mockLogger.go -package=mocks github.com/go-logr/logr Logger

From the project root directory, you can then generate mocks for all the interfaces that have a go:generate annotation by running 'go generate ./...'.

*/
import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	kraanscheme "github.com/fidelity/kraan/pkg/api/v1alpha1"
	"github.com/go-logr/logr"
	testlogr "github.com/go-logr/logr/testing"
	gomock "github.com/golang/mock/gomock"
	"k8s.io/apimachinery/pkg/runtime"
	fakeK8s "k8s.io/client-go/kubernetes/fake"
	k8sscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var (
	testScheme = runtime.NewScheme()
	// testCtx    = context.Background()
)

func init() {
	_ = k8sscheme.AddToScheme(testScheme)   // nolint:errcheck // ok
	_ = kraanscheme.AddToScheme(testScheme) // nolint:errcheck // ok
}

func fakeLogger() logr.Logger {
	return testlogr.NullLogger{}
}

func TestCreateLayer(t *testing.T) {

}

func GetDataFromFile(fileName string) (interface{}, error) {
	buffer, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buffer, &x)
	if err != nil {
		return nil, err
	}
}

func TestSetRequeue(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	logger := fakeLogger()
	client := fake.NewFakeClientWithScheme(testScheme)
	fakeK8sClient := fakeK8s.NewSimpleClientSet()
	l := CreateLayer(context.Background(), client, fakeK8sClient, logger, data)

	mockLayer := NewMockLayer(mockCtl)

	// Verifies that the "findOnPath" method was called once with the exact expected string input as the parameter
	mockLayer.EXPECT().SetRequeue().Times(1)

	l.SetRequeue()
	k, ok := l.(*KraanLayer)
	if !ok {
		t.Logf("failed to cast layer interface to *KraanLayer")
	}
	if !k.requeue {

	}
	if !l.NeedsRequeue() {
		t.Logf("failed to set requeue using SetRequeue")
	}
}
