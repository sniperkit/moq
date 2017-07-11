package example

// AUTOGENERATED BY MOQ - DO NOT EDIT
// github.com/matryer/moq

import (
	"sync"
	"context"
)

var (
	lockPersonStoreMockCreate	sync.RWMutex
	lockPersonStoreMockGet	sync.RWMutex
)

// PersonStoreMock is a mock implementation of PersonStore.
//
//     func TestSomethingThatUsesPersonStore(t *testing.T) {
//
//         // make and configure a mocked PersonStore
//         mockedPersonStore := &PersonStoreMock{ 
//             CreateFunc: func(ctx context.Context, person *Person, confirm bool) error {
// 	               panic("TODO: mock out the Create method")
//             },
//             GetFunc: func(ctx context.Context, id string) (*Person, error) {
// 	               panic("TODO: mock out the Get method")
//             },
//         }
//
//         // TODO: use mockedPersonStore in code that requires PersonStore
//         //       and then make assertions.
// 
//     }
type PersonStoreMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, person *Person, confirm bool) error

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, id string) (*Person, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Person is the person argument value.
			Person *Person
			// Confirm is the confirm argument value.
			Confirm bool
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Id is the id argument value.
			Id string
		}
	}
}

// Create calls CreateFunc.
func (mock *PersonStoreMock) Create(ctx context.Context, person *Person, confirm bool) error {
	if mock.CreateFunc == nil {
		panic("moq: PersonStoreMock.CreateFunc is nil but PersonStore.Create was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Person *Person
		Confirm bool
	}{
		Ctx: ctx,
		Person: person,
		Confirm: confirm,
	}
	lockPersonStoreMockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	lockPersonStoreMockCreate.Unlock()
	return mock.CreateFunc(ctx, person, confirm)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedPersonStore.CreateCalls())
func (mock *PersonStoreMock) CreateCalls() []struct {
		Ctx context.Context
		Person *Person
		Confirm bool
	} {
	var calls []struct {
		Ctx context.Context
		Person *Person
		Confirm bool
	}
	lockPersonStoreMockCreate.RLock()
	calls = mock.calls.Create
	lockPersonStoreMockCreate.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *PersonStoreMock) Get(ctx context.Context, id string) (*Person, error) {
	if mock.GetFunc == nil {
		panic("moq: PersonStoreMock.GetFunc is nil but PersonStore.Get was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Id string
	}{
		Ctx: ctx,
		Id: id,
	}
	lockPersonStoreMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockPersonStoreMockGet.Unlock()
	return mock.GetFunc(ctx, id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedPersonStore.GetCalls())
func (mock *PersonStoreMock) GetCalls() []struct {
		Ctx context.Context
		Id string
	} {
	var calls []struct {
		Ctx context.Context
		Id string
	}
	lockPersonStoreMockGet.RLock()
	calls = mock.calls.Get
	lockPersonStoreMockGet.RUnlock()
	return calls
}
