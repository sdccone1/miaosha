/**
 * @Author: David Ma
 * @Date: 2021/2/4
 */
package validatorService

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

const moblieFormat = "^1[3456789]\\d\\s?\\d{4}\\s?\\d{4}"

/**
 *对name 这个field自定义校验器
 */
func NameNotNullAndAdmin(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		if value == "" || "admin" == value {
			return false
		}
	}
	return true
}

/**
 *对mobile 这个field自定义校验器
 */

func MobileFormatIsCorrect(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(moblieFormat)
	if value, ok := fl.Field().Interface().(string); ok {
		if value == "" || !re.MatchString(value) {
			return false
		}
	}
	return true
}
