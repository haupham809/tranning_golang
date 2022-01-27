package checkvalidation

import (
	"regexp"
	"github.com/asaskevich/govalidator"
)
func Isspecialcharacters(){
	govalidator.TagMap["check2"] = govalidator.Validator(func(str string) bool {
		m, _ := regexp.MatchString("^[a-zA-Z0-9]+$", str)
		if m {
			return false
		} else {
			return true
		}

	})
}
func Nospecialcharacters(){
	govalidator.TagMap["check1"] = govalidator.Validator(func(str string) bool {
		m, _ := regexp.MatchString("^[a-zA-Z0-9]+$", str)
		if m {
			return true
		} else {
			return false
		}

	})
}
func Isletterscharacters(){
	govalidator.TagMap["check3"] = govalidator.Validator(func(str string) bool {
		for i := 0; i < len(str); i++ {
			m, _ := regexp.MatchString("^[A-Z]+$", str[i:i+1])
			if m {
				return true
			}
		}

		return false

	})
}

func SqlInjection(){
	govalidator.TagMap["SqlInjection"] = govalidator.Validator(func(str string) bool {
		for i := 0; i < len(str); i++ {
			m, _ := regexp.MatchString("[\\s]+$", str[i:i+1])
			if m {
				return false
			}
		}

		return true

	})
}