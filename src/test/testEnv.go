package test

import (
	"env"
	"fmt"
)

func testEnv() {
	env := env.MyEnv.GetInstance()
	env.SetUp()

	fmt.Printf(" radius = %f \n", env.Radius)
}
