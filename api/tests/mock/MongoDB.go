package mock

import (
	"context"
	"github.com/elissonalvesilva/interview-meli/api/infra/db/mongodb"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

// MockMongoDB is a mock of MongoDB interface.
type MockMongoDB struct {
	ctrl     *gomock.Controller
	recorder *MockMongoDBMockRecorder
}

// MockMongoDBMockRecorder is the mock recorder for MockMongoDB.
type MockMongoDBMockRecorder struct {
	mock *MockMongoDB
}

// NewMockMongoDB creates a new mock instance.
func NewMockMongoDB(ctrl *gomock.Controller) *MockMongoDB {
	mock := &MockMongoDB{ctrl: ctrl}
	mock.recorder = &MockMongoDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMongoDB) EXPECT() *MockMongoDBMockRecorder {
	return m.recorder
}

// Client mocks base method.
func (m *MockMongoDB) Client() mongodb.ClientHelper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(mongodb.ClientHelper)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockMongoDBMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockMongoDB)(nil).Client))
}

// Collection mocks base method.
func (m *MockMongoDB) Collection(name string) mongodb.CollectionHelper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collection", name)
	ret0, _ := ret[0].(mongodb.CollectionHelper)
	return ret0
}

// Collection indicates an expected call of Collection.
func (mr *MockMongoDBMockRecorder) Collection(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collection", reflect.TypeOf((*MockMongoDB)(nil).Collection), name)
}

// MockCollectionHelper is a mock of CollectionHelper interface.
type MockCollectionHelper struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionHelperMockRecorder
}

// MockCollectionHelperMockRecorder is the mock recorder for MockCollectionHelper.
type MockCollectionHelperMockRecorder struct {
	mock *MockCollectionHelper
}

// NewMockCollectionHelper creates a new mock instance.
func NewMockCollectionHelper(ctrl *gomock.Controller) *MockCollectionHelper {
	mock := &MockCollectionHelper{ctrl: ctrl}
	mock.recorder = &MockCollectionHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCollectionHelper) EXPECT() *MockCollectionHelperMockRecorder {
	return m.recorder
}

// Aggregate mocks base method.
func (m *MockCollectionHelper) Aggregate(ctx context.Context, pipeline interface{}) (*mongo.Cursor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Aggregate", ctx, pipeline)
	ret0, _ := ret[0].(*mongo.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Aggregate indicates an expected call of Aggregate.
func (mr *MockCollectionHelperMockRecorder) Aggregate(ctx, pipeline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aggregate", reflect.TypeOf((*MockCollectionHelper)(nil).Aggregate), ctx, pipeline)
}

// CountDocuments mocks base method.
func (m *MockCollectionHelper) CountDocuments(ctx context.Context, filter interface{}) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountDocuments", ctx, filter)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountDocuments indicates an expected call of CountDocuments.
func (mr *MockCollectionHelperMockRecorder) CountDocuments(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountDocuments", reflect.TypeOf((*MockCollectionHelper)(nil).CountDocuments), ctx, filter)
}

// Find mocks base method.
func (m *MockCollectionHelper) Find(arg0 context.Context, arg1 interface{}) (*mongo.Cursor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*mongo.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCollectionHelperMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCollectionHelper)(nil).Find), arg0, arg1)
}

// FindOne mocks base method.
func (m *MockCollectionHelper) FindOne(arg0 context.Context, arg1 interface{}) mongodb.SingleResultHelper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(mongodb.SingleResultHelper)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockCollectionHelperMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockCollectionHelper)(nil).FindOne), arg0, arg1)
}

// InsertOne mocks base method.
func (m *MockCollectionHelper) InsertOne(arg0 context.Context, arg1 interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockCollectionHelperMockRecorder) InsertOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockCollectionHelper)(nil).InsertOne), arg0, arg1)
}

// MockSingleResultHelper is a mock of SingleResultHelper interface.
type MockSingleResultHelper struct {
	ctrl     *gomock.Controller
	recorder *MockSingleResultHelperMockRecorder
}

// MockSingleResultHelperMockRecorder is the mock recorder for MockSingleResultHelper.
type MockSingleResultHelperMockRecorder struct {
	mock *MockSingleResultHelper
}

// NewMockSingleResultHelper creates a new mock instance.
func NewMockSingleResultHelper(ctrl *gomock.Controller) *MockSingleResultHelper {
	mock := &MockSingleResultHelper{ctrl: ctrl}
	mock.recorder = &MockSingleResultHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSingleResultHelper) EXPECT() *MockSingleResultHelperMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockSingleResultHelper) Decode(v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockSingleResultHelperMockRecorder) Decode(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockSingleResultHelper)(nil).Decode), v)
}

// MockClientHelper is a mock of ClientHelper interface.
type MockClientHelper struct {
	ctrl     *gomock.Controller
	recorder *MockClientHelperMockRecorder
}

// MockClientHelperMockRecorder is the mock recorder for MockClientHelper.
type MockClientHelperMockRecorder struct {
	mock *MockClientHelper
}

// NewMockClientHelper creates a new mock instance.
func NewMockClientHelper(ctrl *gomock.Controller) *MockClientHelper {
	mock := &MockClientHelper{ctrl: ctrl}
	mock.recorder = &MockClientHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientHelper) EXPECT() *MockClientHelperMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockClientHelper) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockClientHelperMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockClientHelper)(nil).Connect))
}

// Database mocks base method.
func (m *MockClientHelper) Database(arg0 string) mongodb.MongoDB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Database", arg0)
	ret0, _ := ret[0].(mongodb.MongoDB)
	return ret0
}

// Database indicates an expected call of Database.
func (mr *MockClientHelperMockRecorder) Database(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Database", reflect.TypeOf((*MockClientHelper)(nil).Database), arg0)
}

// StartSession mocks base method.
func (m *MockClientHelper) StartSession() (mongo.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartSession")
	ret0, _ := ret[0].(mongo.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartSession indicates an expected call of StartSession.
func (mr *MockClientHelperMockRecorder) StartSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSession", reflect.TypeOf((*MockClientHelper)(nil).StartSession))
}
