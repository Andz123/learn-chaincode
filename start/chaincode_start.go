package main

import (
	"errors"
	"fmt"
	"strings"
	//"bytes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}





// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

    err := stub.PutState("hello_world", []byte(args[0]))
    if err != nil {
        return nil, err
    }

    return nil, nil
}







// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

    // Handle different functions
    if function == "init" {
        return t.Init(stub, "init", args)
    } else if function == "write" {
        return t.write(stub, args)
    }
    fmt.Println("invoke did not find func: " + function)

    return nil, errors.New("Received unknown function invocation: " + function)
}


func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//func (t *SimpleChaincode) write(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	
	if len(args) % 2 == 0 {
		return nil, errors.New("The key and the values are not both present")
	}
	
	
	for i := 0; i < len(args); i++ { //ana 3m b22t3 3l variablet b2ra kl mara l key mn b3do l value tb3o
	
	 //arg[i] = key    --- arg[i+1] == value  (momkin ysir error fa byitsayav bl err 
	//putstate chkla putstate (string  , table byte) lhek 7awalna mn string la byte[] 3n tri2 []byte()
	err = stub.PutState(args[i], []byte(args[i+1]))

	if err != nil {
		return nil, errors.New("Error in setting value in the key") //hayde 3chen iza taj l setting byitl3
	}
	i++
	}
	return nil, nil
}



// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

    // Handle different functions
    if function == "read" {                            //read a variable
        return t.read(stub, args)
    }
    fmt.Println("query did not find func: " + function)

    return nil, errors.New("Received unknown function query: " + function)
}

/* 
func convert( b []byte ) string {
    s := make([]string,len(b))
    for i := range b {
        s[i] = strconv.Itoa(int(b[i]))
    }
    return strings.Join(s,",")
}

 */

 
 
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//func (t *SimpleChaincode) read(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var num int
 
	
	num = len(args)
	
	var result []string
	var batata1 string
	var ba []byte
	
	result = make([]string, num) 
	
	//var Buf bytes.Buffer
	
	//var key, jsonResp string
	var err error
	
	
	
	for i := 0; i < len(args); i++ { 
	batata1 = args[i]
	ba, err = stub.GetState(batata1)
	if err != nil {
		return nil, errors.New("Error bi sater la3in") 
	}
	result[i] = fmt.Sprintf("%s", ba)
	
	}
	

	var z string
	z = strings.Join(result, "-")
	//andrew = fmt.Sprintf("%s", z)
	 
    return []byte(z), nil
	
}
