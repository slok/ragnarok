package serializer

import (
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"

	"github.com/slok/ragnarok/api"
	chaosv1 "github.com/slok/ragnarok/api/chaos/v1"
	clusterv1 "github.com/slok/ragnarok/api/cluster/v1"
)

// TypeAsserter has the ability to assert type interfaces safely.
type TypeAsserter interface {
	ByteArray(data interface{}) ([]byte, error)
	Writer(data interface{}) (io.Writer, error)
	ProtoMessage(data interface{}) (proto.Message, error)
}

// safeTypeAsserter is the default type asserter implementation.
type safeTypeAsserter struct{}

// SafeTypeAsserter is a global helper to use it as a safe type assertion util in the econders/decoders.
var SafeTypeAsserter = &safeTypeAsserter{}

func (s *safeTypeAsserter) ByteArray(data interface{}) ([]byte, error) {
	bdata, ok := data.([]byte)
	if !ok {
		return bdata, fmt.Errorf("wrong type of data as argument, should be []byte")
	}
	return bdata, nil
}

func (s *safeTypeAsserter) Writer(data interface{}) (io.Writer, error) {
	w, ok := data.(io.Writer)
	if !ok {
		return nil, fmt.Errorf("wrong interface as argument, should be io.Writer interface")
	}
	return w, nil
}

func (s *safeTypeAsserter) ProtoMessage(data interface{}) (proto.Message, error) {
	pb, ok := data.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("wrong interface as argument, should be proto.Message interface")
	}
	return pb, nil
}

// Factory implements a way of obtaining objects based on the verison and kind.
type Factory interface {
	// NewPlainObject returns anew plain object based on the version and kind.
	NewPlainObject(t api.TypeMeta) (api.Object, error)
}

// objFactory will be used to get plain objects based on the version and object kind
type objFactory struct{}

// ObjFactory is a global helper to use it as a factory in the econders/decoders.
var ObjFactory = &objFactory{}

// NewPlainObject satisfies Factory interface.
func (o *objFactory) NewPlainObject(t api.TypeMeta) (api.Object, error) {
	// TODO: Make more elegant way of registering object creators.
	switch {
	case t == clusterv1.NodeTypeMeta:
		n := clusterv1.NewNode()
		return &n, nil
	case t == clusterv1.NodeListTypeMeta:
		n := clusterv1.NewNodeList([]*clusterv1.Node{}, "")
		return &n, nil
	case t == chaosv1.FailureTypeMeta:
		n := chaosv1.NewFailure()
		return &n, nil
	case t == chaosv1.FailureListTypeMeta:
		n := chaosv1.NewFailureList([]*chaosv1.Failure{}, "")
		return &n, nil
	case t == chaosv1.ExperimentTypeMeta:
		n := chaosv1.NewExperiment()
		return &n, nil
	case t == chaosv1.ExperimentListTypeMeta:
		n := chaosv1.NewExperimentList([]*chaosv1.Experiment{}, "")
		return &n, nil
	default:
		return nil, fmt.Errorf("unknown %s object type", t)
	}
}

// TypeDiscoverer discovers the type of an object from the encoded format.
type TypeDiscoverer interface {
	// Discovery will return the type of the object.
	Discover(data interface{}) (api.TypeMeta, error)
}

// Typer is the interface that knows to set the type in an object instance.
type Typer interface {
	// SetType sets the type of the object in the object.
	SetType(obj api.Object) error
}

// objTyper is the default object typer.
type objTyper struct{}

// ObjTyper is a handy instance of the default object typer.
var ObjTyper = &objTyper{}

// setTypesOnListObjects sets the typeMeta of the objects in a ObjectList.
func (o *objTyper) setTypesOnListObjects(list api.ObjectList, t api.TypeMeta) error {
	for _, item := range list.GetItems() {
		if err := o.SetType(item); err != nil {
			return err
		}
	}
	return nil
}

// SetType implements Typer interface.
func (o *objTyper) SetType(obj api.Object) error {
	// TODO: Make more elegant way of setting correct types.
	switch v := obj.(type) {
	case *clusterv1.Node:
		v.TypeMeta = clusterv1.NodeTypeMeta
	case *clusterv1.NodeList:
		v.TypeMeta = clusterv1.NodeListTypeMeta
		o.setTypesOnListObjects(v, clusterv1.NodeTypeMeta)
	case *chaosv1.Failure:
		v.TypeMeta = chaosv1.FailureTypeMeta
	case *chaosv1.FailureList:
		v.TypeMeta = chaosv1.FailureListTypeMeta
		o.setTypesOnListObjects(v, chaosv1.FailureTypeMeta)
	case *chaosv1.Experiment:
		v.TypeMeta = chaosv1.ExperimentTypeMeta
	case *chaosv1.ExperimentList:
		v.TypeMeta = chaosv1.ExperimentListTypeMeta
		o.setTypesOnListObjects(v, chaosv1.ExperimentTypeMeta)
	default:
		return fmt.Errorf("could not set the type of object because isn't a valid object type")
	}

	return nil
}
