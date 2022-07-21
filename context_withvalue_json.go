package main

import (
	"fmt"
	"reflect"
	"context"
	"encoding/json"
)

func main() {
	ctx := context.WithValue(context.Background(), "a", "b")
	//对context进行json编码会丢失value
	data, _ := json.Marshal(ctx)
	fmt.Println(string(data))
	fmt.Println(reflect.ValueOf(ctx).Type())
	fmt.Println(reflect.ValueOf(ctx).Elem())
	fmt.Println(ctx)
	fmt.Println(ctx.Value("a"))
}

/**
{"Context":0}
*context.valueCtx
{context.Background a b}
context.Background.WithValue(type string, val b)
b
*/
