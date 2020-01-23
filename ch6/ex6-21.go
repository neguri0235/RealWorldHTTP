// JSON-RPC로 서버의 계산 처리를 호출하는 클라이언트

package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:18888")

	if err != nil {
		panic(err)
	}

	var result int

	args := Args{4, 5}
	err = client.Call("Calculator.Multiply", args, &result)

	if err != nil {
		panic(err)
	}

	log.Printf("4 * 5 = %d\n", result)

}
