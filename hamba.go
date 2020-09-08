package hamba

import (
	"errors"
	"strconv"
)

// Hamba main struct
type Hamba struct {
	key      []string
	value    []string
	readonly []bool
}

var (
	errIsNotValidType error = errors.New("that's not valid type")
	errIsNotExistKey  error = errors.New("that's not exist key")
	errItIsReadonly   error = errors.New("a readonly property of that's list is true")
)

func find(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}

func deleteElement(slice []string, i int) []string {
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkValidType(valueType string) error {
	validTypes := []string{
		"bool",
		"string",
		"int",
		"float32",
		"float64",
		"complex64",
		"complex128",
		"unit8",
		"uint16",
		"uint32",
		"uint64",
	}
	if find(validTypes, valueType) == -1 {
		return errIsNotValidType
	}
	return nil
}

// New returns the Hamba main struct
func New() *Hamba {
	return &Hamba{
		key:      []string{},
		value:    []string{},
		readonly: []bool{},
	}
}

// GetAsBool returns bool type by the key
func (h *Hamba) GetAsBool(key string) (bool, error) {
	i := find(h.key, key)
	if i == -1 {
		return false, errIsNotExistKey
	}

	value, err := strconv.ParseBool(h.value[i])
	checkErr(err)
	return value, nil
}

// GetAsString returns string type by the key
func (h *Hamba) GetAsString(key string) (string, error) {
	i := find(h.key, key)
	if i == -1 {
		return "", errIsNotExistKey
	}

	value := h.value[i]
	return value, nil
}

// GetAsInt returns int type by the key
func (h *Hamba) GetAsInt(key string) (int, error) {
	i := find(h.key, key)
	if i == -1 {
		return -1, errIsNotExistKey
	}

	value, err := strconv.Atoi(h.value[i])
	checkErr(err)
	return value, nil
}

// GetAsFloat32 returns float32 type by the key
func (h *Hamba) GetAsFloat32(key string) (float32, error) {
	i := find(h.key, key)
	if i == -1 {
		return -1, errIsNotExistKey
	}

	value, err := strconv.ParseFloat(h.value[i], 32)
	checkErr(err)
	return float32(value), nil
}

// GetAsFloat64 returns float64 type by the key
func (h *Hamba) GetAsFloat64(key string) (float64, error) {
	i := find(h.key, key)
	if i == -1 {
		return -1, errIsNotExistKey
	}

	value, err := strconv.ParseFloat(h.value[i], 64)
	checkErr(err)
	return value, nil
}

// GetAsComplex64 returns complex64 type by the key
func (h *Hamba) GetAsComplex64(key string) (complex64, error) {
	i := find(h.key, key)
	if i == -1 {
		return -1, errIsNotExistKey
	}

	value, err := strconv.ParseComplex(h.value[i], 64)
	checkErr(err)
	return complex64(value), nil
}

// GetAsComplex128 returns complex128 type by the key
func (h *Hamba) GetAsComplex128(key string) (complex128, error) {
	i := find(h.key, key)
	if i == -1 {
		return -1, errIsNotExistKey
	}

	value, err := strconv.ParseComplex(h.value[i], 128)
	checkErr(err)
	return value, nil
}

// GetAsUint8 returns uint8 type by the key
func (h *Hamba) GetAsUint8(key string, base int) (uint8, error) {
	i := find(h.key, key)
	if i == -1 {
		return 0, errIsNotExistKey
	}

	value, err := strconv.ParseUint(h.value[i], base, 8)
	checkErr(err)
	return uint8(value), nil
}

// GetAsUint16 returns uint16 type by the key
func (h *Hamba) GetAsUint16(key string, base int) (uint16, error) {
	i := find(h.key, key)
	if i == -1 {
		return 0, errIsNotExistKey
	}

	value, err := strconv.ParseUint(h.value[i], base, 16)
	checkErr(err)
	return uint16(value), nil
}

// GetAsUint32 returns uint32 type by the key
func (h *Hamba) GetAsUint32(key string, base int) (uint32, error) {
	i := find(h.key, key)
	if i == -1 {
		return 0, errIsNotExistKey
	}

	value, err := strconv.ParseUint(h.value[i], base, 32)
	checkErr(err)
	return uint32(value), nil
}

// GetAsUint64 returns uint64 type by the key
func (h *Hamba) GetAsUint64(key string, base int) (uint64, error) {
	i := find(h.key, key)
	if i == -1 {
		return 0, errIsNotExistKey
	}

	value, err := strconv.ParseUint(h.value[i], base, 64)
	checkErr(err)
	return value, nil
}

// Add adds values to hamba list
func (h *Hamba) Add(key, value string, readonly bool) {
	h.key = append(h.key, key)
	h.value = append(h.value, value)
	h.readonly = append(h.readonly, readonly)
}

// UpdateKey updates key to the newKey
func (h *Hamba) UpdateKey(key, newKey string) error {
	i := find(h.key, key)
	if i == -1 {
		return errIsNotExistKey
	}

	if !h.readonly[i] {
		h.key[i] = newKey
		return nil
	}

	return errItIsReadonly
}

// UpdateValue updates value to the newValue
func (h *Hamba) UpdateValue(key, newValue string) error {
	i := find(h.key, key)
	if i == -1 {
		return errIsNotExistKey
	}

	if !h.readonly[i] {
		h.value[i] = newValue
		return nil
	}

	return errItIsReadonly
}

// UpdateReadonly updates readonly to the newReadonly
// It doesn't care about readonly
func (h *Hamba) UpdateReadonly(key string, newReadonly bool) error {
	i := find(h.key, key)
	if i == -1 {
		return errIsNotExistKey
	}

	h.readonly[i] = newReadonly
	return nil
}

// UpdateAll updates all
func (h *Hamba) UpdateAll(key, newKey, newValue string, newReadonly bool) error {
	i := find(h.key, key)
	if i == -1 {
		return errIsNotExistKey
	}

	if !h.readonly[i] {
		h.key[i] = newKey
		h.value[i] = newValue
		h.readonly[i] = newReadonly
		return nil
	}

	return errItIsReadonly
}

// Delete deletes element
func (h *Hamba) Delete(key string) error {
	i := find(h.key, key)
	if i == -1 {
		return errIsNotExistKey
	}

	h.key = deleteElement(h.key, i)
	h.value = deleteElement(h.value, i)
	h.readonly = append(h.readonly[:i], h.readonly[i+1:]...)
	return nil
}

// Remove function is the same with Delete function
func (h *Hamba) Remove(key string) error {
	return h.Delete(key)
}
