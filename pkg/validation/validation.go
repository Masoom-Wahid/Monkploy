package validation

import (
	"fmt"
	"platform/database"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Failed string
	Tag    string
	Value  string
}

var validate = validator.New()

func Validation(data interface{}) map[string]string {
	validate.RegisterValidation("unique_field", uniqueField)
	validate.RegisterValidation("exists_value", existsValue)
	validate.RegisterValidation("exists_array", existsArray)
	validate.RegisterValidation("regex", regexValidate)
	validate.RegisterValidation("password", passwordValidate)

	validate.RegisterValidation("phone", validatePhoneNumberFormat)
	validate.RegisterValidation("format", validateDateFormat)
	validate.RegisterValidation("format24", validateTimeFormat)
	validate.RegisterValidation("today_or_after", validateTodayOrAfter)

	var errors []*ErrorResponse
	err := validate.Struct(data)

	// reflect the data to get the json tag of our struct for each field
	v := reflect.ValueOf(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			// element.Failed = strings.ToLower(err.Field()) // old version
			// element.Failed = v.Type().Field(i).Tag.Get("json")
			element.Failed = jsonName(err.Field(), v)
			element.Tag = GetMessageByTag(err.Tag(), element.Failed, err.Param())
			// element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return formatErrors(errors)
}

func uniqueField(fl validator.FieldLevel) bool {
	param := fl.Param()
	table := param[0:strings.Index(param, ":")]
	field := param[strings.Index(param, ":")+1:]

	var count int64

	db := database.GetDB()
	db.Table(table).Where(field+" = ?", fl.Field().String()).Where("deleted_at is Null").Count(&count)

	return count == 0
}

func existsValue(fl validator.FieldLevel) bool {
	param := fl.Param()
	table := param[0:strings.Index(param, ":")]
	field := param[strings.Index(param, ":")+1:]

	var count int64

	db := database.GetDB()
	if table == "users" {
		db.Table(table).Where(field+" = ?", fl.Field().String()).Where("deleted_at is Null").Count(&count)
	} else {
		if field == "id" {
			db.Table(table).Where(field+" = ?", fl.Field().Int()).Where("deleted_at is Null").Count(&count)
		} else {
			db.Table(table).Where(field+" = ?", fl.Field().String()).Where("deleted_at is Null").Count(&count)
		}

		// this commented line will use for resources that should check for its active/deactive
		// db.Table(table).Where(field+" = ?", fl.Field().String()).Where("is_active = ?", true).Where("deleted_at is Null").Count(&count)
	}

	return count >= 1
}

// only use for int types
func existsArray(fl validator.FieldLevel) bool {
	param := fl.Param()
	table := param[0:strings.Index(param, ":")]
	field := param[strings.Index(param, ":")+1:]

	// get all ints from reflect.value type
	value := getIds(fl.Field())

	db := database.GetDB()

	var count int64

	db.Unscoped().Table(table).Where(field+" IN (?)", value).Count(&count)

	return count == int64(len(value))
}

func getIds(value reflect.Value) []int64 {
	var ids []int64

	for i := 0; i < value.Len(); i++ {
		ids = append(ids, value.Index(i).Int())
	}

	return ids
}

func regexValidate(fl validator.FieldLevel) bool {
	// @TODO Complete regex validation
	param := fl.Param()

	println(param)

	return false
}

func validatePhoneNumberFormat(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	phoneNumberRegex := regexp.MustCompile(`^09\d{9}$`)

	return phoneNumberRegex.MatchString(phone)
}

func validateDateFormat(fl validator.FieldLevel) bool {
	userInput := fl.Field().String()
	layout := "2006/01/02"

	_, err := time.Parse(layout, userInput)
	if err != nil {
		fmt.Println("Invalid date format")
		return false
	}

	return true
}

func validateTimeFormat(fl validator.FieldLevel) bool {
	userInput := fl.Field().String()
	timeFormat := "15:04"

	_, err := time.Parse(timeFormat, userInput)
	if err != nil {
		fmt.Println("Invalid time format")
		return false
	}

	return true
}

func validateTodayOrAfter(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	layout := "2006/01/02"

	date, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Invalid date format")
		return false
	}

	today := time.Now().Local().Truncate(24 * time.Hour)

	if date.Before(today) {
		return false
	}

	return true
}

func passwordValidate(fl validator.FieldLevel) bool {
	secure := true
	tests := []string{"[a-z]", "[A-Z]", "[0-9]"}

	// the password length should be larger than 8
	if len(fl.Field().String()) <= 7 {
		return false
	}

	for _, test := range tests {
		t, _ := regexp.MatchString(test, fl.Field().String())
		if !t {
			secure = false
			break
		}
	}

	return secure
}

func formatErrors(errors []*ErrorResponse) map[string]string {
	var errorMap = make(map[string]string)
	for _, err := range errors {
		errorMap[err.Failed] = err.Tag + err.Value
	}

	return errorMap
}

func jsonName(field string, v reflect.Value) string {
	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Name == field {
			return v.Type().Field(i).Tag.Get("json")
		}
	}

	return "field_not_found"
}

func GetMessageByTag(tag string, filed string, param string) string {
	println(tag, filed, param)
	messages := map[string]string{
		"required":       "The %s field is required",
		"required_if":    "The %s field is required",
		"unique_field":   "The %s field is already taken",
		"exists_value":   "The selected %s is not exists",
		"exists_array":   "The selected %s is not exists",
		"min":            "The %s field length must be %s character",
		"gt":             "The %s field must be grater than %s",
		"password":       "The %s field format is invalid",
		"eqcsfield":      "The %s field should equal to %s field",
		"format":         "The %s format is invalid",
		"today_or_after": "The %s field should be today or after",
		"phone":          "The %s field format is invalid",
	}

	if tag == "min" {
		return fmt.Sprintf(messages[tag], filed, param)
	}

	if tag == "gt" {
		return fmt.Sprintf(messages[tag], filed, param)
	}

	if tag == "eqcsfield" {
		return fmt.Sprintf(messages[tag], filed, param)
	}

	return fmt.Sprintf(messages[tag], filed)
}
