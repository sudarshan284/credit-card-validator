package main

func luhnAlgorithm(cardNumber string) bool {
	//this function implements the luhn algortihm
	total := 0
	isSecondDigit := false
	for i := len(cardNumber) - 1;i >= 0;i-- {
		digit := int(cardNumber[i] - '0')

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		total += digit

		//Toggle the flag for the next iteration
		isSecondDigit = !isSecondDigit
	}

	return total%10 == 0
}