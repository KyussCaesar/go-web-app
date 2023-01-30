# IUpdateUserDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | New value for the user&#39;s username.  | [optional] 
**Email** | Pointer to **string** | New value for the user&#39;s email.  | [optional] 

## Methods

### NewIUpdateUserDTO

`func NewIUpdateUserDTO() *IUpdateUserDTO`

NewIUpdateUserDTO instantiates a new IUpdateUserDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIUpdateUserDTOWithDefaults

`func NewIUpdateUserDTOWithDefaults() *IUpdateUserDTO`

NewIUpdateUserDTOWithDefaults instantiates a new IUpdateUserDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *IUpdateUserDTO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IUpdateUserDTO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IUpdateUserDTO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IUpdateUserDTO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *IUpdateUserDTO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IUpdateUserDTO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IUpdateUserDTO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IUpdateUserDTO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


