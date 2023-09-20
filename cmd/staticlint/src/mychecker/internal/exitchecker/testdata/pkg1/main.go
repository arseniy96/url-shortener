package main

import "os"

func errCheckFunc() {
	// формулируем ожидания: анализатор должен находить ошибку,
	// описанную в комментарии want
	os.Exit(1) // want "call os.Exit in main package"
}
