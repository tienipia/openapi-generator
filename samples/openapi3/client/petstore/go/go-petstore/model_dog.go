/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the Dog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dog{}

// Dog struct for Dog
type Dog struct {
	Animal
	Breed *string `json:"breed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Dog Dog

// NewDog instantiates a new Dog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDog(className string) *Dog {
	this := Dog{}
	this.ClassName = className
	var color string = "red"
	this.Color = &color
	return &this
}

// NewDogWithDefaults instantiates a new Dog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDogWithDefaults() *Dog {
	this := Dog{}
	return &this
}

// GetBreed returns the Breed field value if set, zero value otherwise.
func (o *Dog) GetBreed() string {
	if o == nil || IsNil(o.Breed) {
		var ret string
		return ret
	}
	return *o.Breed
}

// GetBreedOk returns a tuple with the Breed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dog) GetBreedOk() (*string, bool) {
	if o == nil || IsNil(o.Breed) {
		return nil, false
	}
	return o.Breed, true
}

// HasBreed returns a boolean if a field has been set.
func (o *Dog) HasBreed() bool {
	if o != nil && !IsNil(o.Breed) {
		return true
	}

	return false
}

// SetBreed gets a reference to the given string and assigns it to the Breed field.
func (o *Dog) SetBreed(v string) {
	o.Breed = &v
}

func (o Dog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAnimal, errAnimal := json.Marshal(o.Animal)
	if errAnimal != nil {
		return map[string]interface{}{}, errAnimal
	}
	errAnimal = json.Unmarshal([]byte(serializedAnimal), &toSerialize)
	if errAnimal != nil {
		return map[string]interface{}{}, errAnimal
	}
	if !IsNil(o.Breed) {
		toSerialize["breed"] = o.Breed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Dog) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"className",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type DogWithoutEmbeddedStruct struct {
		Breed *string `json:"breed,omitempty"`
	}

	varDogWithoutEmbeddedStruct := DogWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varDogWithoutEmbeddedStruct)
	if err == nil {
		varDog := _Dog{}
		varDog.Breed = varDogWithoutEmbeddedStruct.Breed
		*o = Dog(varDog)
	} else {
		return err
	}

	varDog := _Dog{}

	err = json.Unmarshal(bytes, &varDog)
	if err == nil {
		o.Animal = varDog.Animal
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "breed")

		// remove fields from embedded structs
		reflectAnimal := reflect.ValueOf(o.Animal)
		for i := 0; i < reflectAnimal.Type().NumField(); i++ {
			t := reflectAnimal.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDog struct {
	value *Dog
	isSet bool
}

func (v NullableDog) Get() *Dog {
	return v.value
}

func (v *NullableDog) Set(val *Dog) {
	v.value = val
	v.isSet = true
}

func (v NullableDog) IsSet() bool {
	return v.isSet
}

func (v *NullableDog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDog(val *Dog) *NullableDog {
	return &NullableDog{value: val, isSet: true}
}

func (v NullableDog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


