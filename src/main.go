package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/bxcodec/faker/v3"

	User "builder-pattern/src/user"
	Address "builder-pattern/src/user/address"
	"builder-pattern/src/user/gender"
)

func main() {
	userBuilder := User.NewBuilder()

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		var wg2 sync.WaitGroup
		wg2.Add(1)

		go func() {
			defer wg2.Done()
			userBuilder.
				SetName(faker.Name())
		}()

		wg2.Add(1)
		go func() {
			defer wg2.Done()
			userBuilder.
				SetAge(30)
		}()

		wg2.Add(1)
		go func() {
			defer wg2.Done()

			genderChoices := []gender.Gender{gender.Male, gender.Female}
			idx := rand.Intn(len(genderChoices))

			userBuilder.
				SetGender(genderChoices[idx])
		}()

		wg2.Add(1)
		go func() {
			defer wg2.Done()

			addressBuilder := Address.NewBuilder()
			addressBuilder.
				// SetLine1("No. 111").
				// SetLine2("Jalan Taman").
				SetCity("Kuala Lumpur").
				SetState("Wilayah Persekutuan").
				SetCountry("Malaysia").
				SetPincode("123456")

			address := addressBuilder.Build()
			if err := addressBuilder.Error(); err != nil {
				fmt.Println(err.String(""))
				return
			}

			userBuilder.SetAddress(address)
		}()

		wg2.Wait()

		user := userBuilder.Build()
		if err := userBuilder.Error(); err != nil {
			fmt.Println(err.String(""))
			return
		}

		wg2.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(user.String())
		}(&wg2)
		wg2.Wait()
	}()

	wg.Wait()
	fmt.Println("All done")
}
