package main

import (
	"context"
	"fmt"
	"graphql-to-go/generator"
	"graphql-to-go/service"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) <= 2 {
		log.Fatal("Argumentos necessários: [URL] [TOKEN] [OBJECT NAME]. **OBJECT NAME é opcional")
	}

	serviceUrl := os.Args[1]
	serviceToken := os.Args[2]

	service := service.NewService(serviceUrl, serviceToken, 30)
	// service := service.NewService(`https://countries.trevorblades.com/`, "", 30)
	ctx := context.Background()

	fmt.Println(len(os.Args))
	if len(os.Args) >= 4 {
		generateType(ctx, service, os.Args[3])
	} else {
		generateTypes(ctx, service)
	}

	// for _, v := range respData.Type.Fields {
	// 	fmt.Printf("Field %s : %s\n", v.Name, v.Type.Name)
	// }

	fmt.Printf("Tempo: %d\n", time.Since(start)/time.Millisecond)
}

func generateType(ctx context.Context, service *service.Service, typeName string) {
	responseData, err := service.GetType(ctx, typeName)

	if err != nil {
		log.Fatal(err)
	}

	result, err := generator.ParseType(*responseData)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*result)
}

func generateTypes(ctx context.Context, service *service.Service) {
	responseData, err := service.GetSchema(ctx)

	if err != nil {
		log.Fatal(err)
	}

	result, err := generator.ParseSchema(*responseData)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
