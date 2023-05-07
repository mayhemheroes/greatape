package fuzz_greatape_validators

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/reiver/greatape/app/validators"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                password, _ := fuzzConsumer.GetString()

                validators.PasswordIsValid(password)
                return 0

            case 1:
                percent, _ := fuzzConsumer.GetFloat64()

                validators.PercentIsValid(percent)
                return 0

            case 2:
                postalCode, _ := fuzzConsumer.GetString()

                validators.PostalCodeIsValid(postalCode)
                return 0

            case 3:
                testString, _ := fuzzConsumer.GetString()

                validators.RequiredStringIsValid(testString)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}