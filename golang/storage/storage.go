package storage

import (
	"errors"
	"fmt"
	"sync"
	"uuid"

	"okieoth/grpc.fun/grpc_test"

	"github.com/google/uuid"
)

var myStorage storage

type Storage interface {
	AddObject(in *grpc_test.Object) (*grpc_test.Object, error)
	DelObject(guid string) (*grpc_test.Object, error)
	UpdObject(in *grpc_test.Object) (*grpc_test.Object, error)
	ListObjects() []grpc_test.Object
}

type storage struct {
	Mutex sync.RWMutex
	S     map[string]grpc_test.Object
}

func deepCopySubType(obj *grpc_test.SubType) *grpc_test.SubType {
	if obj == nil {
		return nil
	}
	copy := &grpc_test.SubType{
		IntValue: obj.IntValue,
	}
	if obj.StringValue != nil {
		copy.StringValue = &*obj.StringValue
	}
	if obj.FloatValue != nil {
		copy.FloatValue = &*obj.FloatValue
	}
	return copy
}

func deepCopyObject(obj *grpc_test.Object) *grpc_test.Object {
	if obj == nil {
		return nil
	}
	copy := &grpc_test.Object{
		Guid:       obj.Guid,
		Date:       obj.Date,
		Timestamp:  obj.Timestamp,
		Time:       obj.Time,
		EnumValue:  obj.EnumValue,
		ArrayValue: make([]string, len(obj.ArrayValue)),
	}

	copy.ArrayValue = append(copy.ArrayValue, obj.ArrayValue...)

	copy.Complex = deepCopySubType(obj.Complex)

	return copy
}

func (s *storage) AddObject(in *grpc_test.Object) (*grpc_test.Object, error) {
	myStorage.Mutex.Lock()
	defer myStorage.Mutex.Unlock()
	newObj := deepCopyObject(in)
	if newObj == nil {
		return nil, errors.New("AddObject: input was nil")
	}
	if newObj.Guid == "" {
		newObj.Guid = uuid.NewString()
	}
	if _, exist := myStorage.S[newObj.Guid]; exist {
		msg := fmt.Sprintf("AddObject: Object with key=%s, already stored", newObj.Guid)
		return nil, errors.New(msg)
	}

	myStorage.S[newObj.Guid] = *newObj

}

func (s *storage) DelObject(guid string) (*grpc_test.Object, error) {
	myStorage.Mutex.Lock()
	defer myStorage.Mutex.Unlock()

}

func (s *storage) UpdObject(in *grpc_test.Object) (*grpc_test.Object, error) {
	myStorage.Mutex.Lock()
	defer myStorage.Mutex.Unlock()

}

func (s *storage) ListObjects() []grpc_test.Object {
	myStorage.Mutex.RLock()
	defer myStorage.Mutex.RUnlock()

}
