package main

import (
	"email/goemail/libofm"
	"fmt"
)

func main() {
	mycontent := " my dear令"

	email := libofm.NewEmail("472119740@qq.com",  "test golang email", mycontent)

	err := libofm.SendEmail(email)

	fmt.Println(err)

}
