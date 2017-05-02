package testservice

import "fmt"

type TestService struct {
}

func (TestService) GetDate() (r string, err error) {
	return "This is the data returned from APIGateway", nil
}

func (TestService) PostData(data string) (r int32, err error) {
	fmt.Println("+++Data from Executor+++")
	fmt.Println(data)
	return 1, nil
}

func (TestService) RegExecutor() (err error) {
	fmt.Println("+++Executor+++")
	fmt.Println("exec register!")
	return nil
}
