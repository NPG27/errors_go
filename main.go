package main

import (
	"errors"
	"fmt"
)

// Ejercicio 1

type SalaryError struct {
	message string
}

func (e SalaryError) Error() string {
	return e.message
}

func calculateSalary(hours int, rate float64) (float64, error) {
	if hours <= 0 {
		return 0, SalaryError{"Error: el trabajador no puede haber trabajado menos de 1 hora mensual"}
	}
	if hours < 80 {
		return 0, SalaryError{"Error: el trabajador no puede haber trabajado menos de 80 hs mensuales"}
	}
	salary := float64(hours) * rate
	if salary >= 150000 {
		salary = salary * 0.9
	}
	return salary, nil
}

func main() {
	//Ejercicio 1
	salary := 100000
	if salary < 150000 {
		err := SalaryError{"Error: el salario ingresado no alcanza el mínimo imponible"}
		panic(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	// Ejercicio 2
	salaryPoint2 := 9000
	if salaryPoint2 <= 10000 {
		err := SalaryError{"Error: el salario es menor a 10.000"}
		fmt.Println(errors.Is(err, SalaryError{}))
		if errors.Is(err, SalaryError{}) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("El salario es válido")
	}

	//Ejercicio 3
	salaryPoint3 := 9000

	if salaryPoint3 <= 10000 {
		err := errors.New("Error: el salario es menor a 10.000")
		if errors.Is(err, errors.New("Error: el salario es menor a 10.000")) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("El salario es válido")
	}

	//Ejercicio 4
	salaryPoint4 := 90000 // salario de ejemplo

	if salaryPoint4 < 150000 {
		err := fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salaryPoint4)
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	//Ejercicio 5

}
