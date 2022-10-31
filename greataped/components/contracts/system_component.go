package contracts

import (
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/service"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
)

const (
	INITIALIZE = 0
	FINALIZE   = 100
)

type (
	SystemComponentType       int
	SystemAction              func() error
	SystemComponentsContainer map[string]ISystemComponent
	SystemObjectCache         map[int64]ISystemObject
	TransactionHandler        func(transaction ITransaction) error

	IConductor interface {
		Logger() ILogger
		Configuration() IConfiguration
		Atomic(handler TransactionHandler) error
		Schedule(spec string, callback func()) error
		GetSystemComponent(name string) ISystemComponent
		RequestActivityStream(method, url, keyId, privateKey string, data []byte, output interface{}) error
		LogRemoteCall(context IContext, eventType uint32, source string, input, result interface{}, err error)

		// Document
		DocumentManager() IDocumentManager
		DocumentExists(id int64) bool
		ListDocuments(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IDocumentCollection
		GetDocument(id int64, editor Identity) (IDocument, error)
		AddDocument(content string, editor Identity) (IDocument, error)
		AddDocumentAtomic(transaction ITransaction, content string, editor Identity) (IDocument, error)
		LogDocument(content string, source string, editor Identity, payload string)
		UpdateDocument(id int64, content string, editor Identity) (IDocument, error)
		UpdateDocumentAtomic(transaction ITransaction, id int64, content string, editor Identity) (IDocument, error)
		RemoveDocument(id int64, editor Identity) (IDocument, error)
		RemoveDocumentAtomic(transaction ITransaction, id int64, editor Identity) (IDocument, error)

		// SystemSchedule
		SystemScheduleManager() ISystemScheduleManager
		SystemScheduleExists(id int64) bool
		ListSystemSchedules(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ISystemScheduleCollection
		GetSystemSchedule(id int64, editor Identity) (ISystemSchedule, error)
		AddSystemSchedule(enabled bool, config string, editor Identity) (ISystemSchedule, error)
		AddSystemScheduleAtomic(transaction ITransaction, enabled bool, config string, editor Identity) (ISystemSchedule, error)
		LogSystemSchedule(enabled bool, config string, source string, editor Identity, payload string)
		UpdateSystemSchedule(id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error)
		UpdateSystemScheduleAtomic(transaction ITransaction, id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error)
		RemoveSystemSchedule(id int64, editor Identity) (ISystemSchedule, error)
		RemoveSystemScheduleAtomic(transaction ITransaction, id int64, editor Identity) (ISystemSchedule, error)

		// Identity
		IdentityManager() IIdentityManager
		IdentityExists(id int64) bool
		ListIdentities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IIdentityCollection
		GetIdentity(id int64, editor Identity) (IIdentity, error)
		AddIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		AddIdentityAtomic(transaction ITransaction, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		LogIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, source string, editor Identity, payload string)
		UpdateIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		UpdateIdentityAtomic(transaction ITransaction, id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		RemoveIdentity(id int64, editor Identity) (IIdentity, error)
		RemoveIdentityAtomic(transaction ITransaction, id int64, editor Identity) (IIdentity, error)

		// AccessControl
		AccessControlManager() IAccessControlManager
		AccessControlExists(id int64) bool
		ListAccessControls(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IAccessControlCollection
		GetAccessControl(id int64, editor Identity) (IAccessControl, error)
		AddAccessControl(key uint64, value uint64, editor Identity) (IAccessControl, error)
		AddAccessControlAtomic(transaction ITransaction, key uint64, value uint64, editor Identity) (IAccessControl, error)
		LogAccessControl(key uint64, value uint64, source string, editor Identity, payload string)
		UpdateAccessControl(id int64, key uint64, value uint64, editor Identity) (IAccessControl, error)
		UpdateAccessControlAtomic(transaction ITransaction, id int64, key uint64, value uint64, editor Identity) (IAccessControl, error)
		RemoveAccessControl(id int64, editor Identity) (IAccessControl, error)
		RemoveAccessControlAtomic(transaction ITransaction, id int64, editor Identity) (IAccessControl, error)

		// RemoteActivity
		RemoteActivityManager() IRemoteActivityManager
		RemoteActivityExists(id int64) bool
		ListRemoteActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IRemoteActivityCollection
		GetRemoteActivity(id int64, editor Identity) (IRemoteActivity, error)
		AddRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		AddRemoteActivityAtomic(transaction ITransaction, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		LogRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, source string, editor Identity, payload string)
		UpdateRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		UpdateRemoteActivityAtomic(transaction ITransaction, id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		RemoveRemoteActivity(id int64, editor Identity) (IRemoteActivity, error)
		RemoveRemoteActivityAtomic(transaction ITransaction, id int64, editor Identity) (IRemoteActivity, error)

		// CategoryType
		CategoryTypeManager() ICategoryTypeManager
		CategoryTypeExists(id int64) bool
		ListCategoryTypes(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryTypeCollection
		GetCategoryType(id int64, editor Identity) (ICategoryType, error)
		AddCategoryType(description string, editor Identity) (ICategoryType, error)
		AddCategoryTypeAtomic(transaction ITransaction, description string, editor Identity) (ICategoryType, error)
		LogCategoryType(description string, source string, editor Identity, payload string)
		UpdateCategoryType(id int64, description string, editor Identity) (ICategoryType, error)
		UpdateCategoryTypeAtomic(transaction ITransaction, id int64, description string, editor Identity) (ICategoryType, error)
		RemoveCategoryType(id int64, editor Identity) (ICategoryType, error)
		RemoveCategoryTypeAtomic(transaction ITransaction, id int64, editor Identity) (ICategoryType, error)

		// Category
		CategoryManager() ICategoryManager
		CategoryExists(id int64) bool
		ListCategories(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection
		GetCategory(id int64, editor Identity) (ICategory, error)
		AddCategory(categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		AddCategoryAtomic(transaction ITransaction, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		LogCategory(categoryTypeId int64, categoryId int64, title string, description string, source string, editor Identity, payload string)
		UpdateCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		UpdateCategoryAtomic(transaction ITransaction, id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		RemoveCategory(id int64, editor Identity) (ICategory, error)
		RemoveCategoryAtomic(transaction ITransaction, id int64, editor Identity) (ICategory, error)
		ListCategoriesByCategoryType(categoryTypeId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection
		ForEachCategoryByCategoryType(categoryTypeId int64, iterator CategoryIterator)
		ListCategoriesByCategory(categoryId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection
		ForEachCategoryByCategory(categoryId int64, iterator CategoryIterator)

		// User
		UserManager() IUserManager
		UserExists(id int64) bool
		ListUsers(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IUserCollection
		GetUser(id int64, editor Identity) (IUser, error)
		AddUser(identityId int64, github string, editor Identity) (IUser, error)
		AddUserAtomic(transaction ITransaction, identityId int64, github string, editor Identity) (IUser, error)
		LogUser(identityId int64, github string, source string, editor Identity, payload string)
		UpdateUser(id int64, github string, editor Identity) (IUser, error)
		UpdateUserAtomic(transaction ITransaction, id int64, github string, editor Identity) (IUser, error)
		RemoveUser(id int64, editor Identity) (IUser, error)
		RemoveUserAtomic(transaction ITransaction, id int64, editor Identity) (IUser, error)

		// Spi
		SpiManager() ISpiManager
		SpiExists(id int64) bool
		ListSpis(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ISpiCollection
		GetSpi(id int64, editor Identity) (ISpi, error)
		AddSpi(editor Identity) (ISpi, error)
		AddSpiAtomic(transaction ITransaction, editor Identity) (ISpi, error)
		LogSpi(source string, editor Identity, payload string)
		UpdateSpi(id int64, editor Identity) (ISpi, error)
		UpdateSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error)
		RemoveSpi(id int64, editor Identity) (ISpi, error)
		RemoveSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error)
		Echo(document IDocument, editor Identity) (IEchoResult, error)

		// CustomError
		CustomErrorManager() ICustomErrorManager
		CustomErrorExists(id int64) bool
		ListCustomErrors(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICustomErrorCollection
		GetCustomError(id int64, editor Identity) (ICustomError, error)
		AddCustomError(editor Identity) (ICustomError, error)
		AddCustomErrorAtomic(transaction ITransaction, editor Identity) (ICustomError, error)
		LogCustomError(source string, editor Identity, payload string)
		UpdateCustomError(id int64, editor Identity) (ICustomError, error)
		UpdateCustomErrorAtomic(transaction ITransaction, id int64, editor Identity) (ICustomError, error)
		RemoveCustomError(id int64, editor Identity) (ICustomError, error)
		RemoveCustomErrorAtomic(transaction ITransaction, id int64, editor Identity) (ICustomError, error)
		ResolveError(document IDocument, editor Identity) (IResolveErrorResult, error)

		NewDocument(id int64, content string) (IDocument, error)
		NewSystemSchedule(id int64, enabled bool, config string) (ISystemSchedule, error)
		NewIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) (IIdentity, error)
		NewAccessControl(id int64, key uint64, value uint64) (IAccessControl, error)
		NewRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) (IRemoteActivity, error)
		NewCategoryType(id int64, description string) (ICategoryType, error)
		NewCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string) (ICategory, error)
		NewUser(id int64, github string) (IUser, error)
		NewSpi() (ISpi, error)
		NewCustomError() (ICustomError, error)
		NewEchoResult(document IDocument, ignored interface{}) IEchoResult
		NewResolveErrorResult(ignored interface{}) IResolveErrorResult
	}

	ISystemComponent interface {
		Name() string
		ResolveDependencies(dependencies ...ISystemComponent) error
		Load() error
		Reload() error
		IsTestEnvironment() bool
		IsDevelopmentEnvironment() bool
		IsStagingEnvironment() bool
		IsProductionEnvironment() bool
		UniqueId() int64
		Logger() ILogger
		Async(task func())
		GenerateUUID() string
		GenerateSalt() string
		GenerateHash(value string, salt string) string
		GenerateJwtToken() string
		GenerateRSAKeyPair() (string, string, error)
		VerifyJwtToken(token string) error
		GenerateCode() string
		Email(destination string, format string, args ...interface{})
		SMS(destination string, format string, args ...interface{})
		Format(format string, args ...interface{}) string
		Match(pattern string, input string) (bool, error)
		Error(interface{}) error
	}

	ISystemComponentFactory interface {
		Create(SystemComponentType, IConfiguration, ILogger, ...ISystemComponent) ISystemComponent
		Components() []ISystemComponent
	}

	IAssertionResult interface {
		Or(error)
	}

	ITransaction interface {
		OnCommit(func())
	}
)

// noinspection GoSnakeCaseUsage
const (
	SYSTEM_COMPONENT_DOCUMENT_MANAGER        SystemComponentType = 0x00000001
	SYSTEM_COMPONENT_SYSTEM_SCHEDULE_MANAGER SystemComponentType = 0x00000002
	SYSTEM_COMPONENT_IDENTITY_MANAGER        SystemComponentType = 0x00000003
	SYSTEM_COMPONENT_ACCESS_CONTROL_MANAGER  SystemComponentType = 0x00000004
	SYSTEM_COMPONENT_REMOTE_ACTIVITY_MANAGER SystemComponentType = 0x00000005
	SYSTEM_COMPONENT_CATEGORY_TYPE_MANAGER   SystemComponentType = 0x00000006
	SYSTEM_COMPONENT_CATEGORY_MANAGER        SystemComponentType = 0x00000007
	SYSTEM_COMPONENT_USER_MANAGER            SystemComponentType = 0x00000008
	SYSTEM_COMPONENT_SPI_MANAGER             SystemComponentType = 0x00000009
	SYSTEM_COMPONENT_CUSTOM_ERROR_MANAGER    SystemComponentType = 0x0000000A
)
