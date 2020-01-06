package main

import (
	// For printing messages on console
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// peer.Response is in the peer package
    "github.com/hyperledger/fabric/protos/peer"

	// Conversion functions
	"strconv"

	// JSON Encoding
	"encoding/json"
)


// EnergyOriginChaincode example simple Chaincode implementation
type EnergyOriginChaincode struct {
}

type Details struct {
	latitude    uint32
	longitude   uint32
	trafoid     uint32
	ckt         uint8  //In case of Prosumers' ckt rep. --> 1st ckt:0; 2nd ckt:1; In case of Trafo 4branches --> 4 ckt
	types       uint8  //0:PV, 1:Wind, 2:CCP, 3:CHP, 4:Coal, 5:battery+, 6: battery-, 7:Consumer, 8:DNO
	pload       uint16 // if _type=8 then index:0->trafoload
}

// Init Implements the Init method
func (token *EnergyOriginChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed in v1")

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *EnergyOriginChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the function name and parameters
	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Invoke executed : ", function, " args=", args)

	switch {

	// Query function
	case    function == "addEnergyOriginCertificate":
			return addEnergyOriginCertificate(stub, args)
	case    function == "getAgentDetail":
			return getAgentDetail(stub, args)
	}

	return errorResponse("Invalid function",1)
}

func addEnergyOriginCertificate(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 7   {
		return errorResponse("Needs to have at least seven arguments", 700)
	}
	agent_id := string(args[0])
	lat, _ := strconv.ParseUint(args[1], 10, 32)
	nlat := uint32(lat)
	long, _ := strconv.ParseUint(args[2], 10, 32)
	nlong := uint32(long)
	trafo, _ := strconv.ParseUint(args[3], 10, 32)
	ntrafo := uint32(trafo)
	ckt, _ := strconv.ParseUint(args[4], 10, 8)
	nckt := uint8(ckt)
	typex, _ := strconv.ParseUint(args[5], 10, 8)
	ntypex := uint8(typex)
	pload, _ := strconv.ParseUint(args[6], 10, 16)
	npload := uint16(pload)


    // Create an instance of the Details struct
	var detail = Details{latitude: nlat, longitude: nlong, trafoid: ntrafo, ckt: nckt, types: ntypex, pload: npload}

	// Convert to JSON and store detail in the state
	jsondetail, _ := json.Marshal(detail)
	stub.PutState(agent_id, []byte(jsondetail))

	return shim.Success([]byte(jsondetail))

}

func getAgentDetail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1   {
		return errorResponse("Needs to have at least one arguments", 700)
	}
	agent_id := string(args[0])
    bytes, err := stub.GetState(agent_id)
    if err != nil {
		return errorResponse(err.Error(), 7)
	}

	 return shim.Success([]byte(bytes))
}


func errorResponse(err string, code  uint ) peer.Response {
	codeStr := strconv.FormatUint(uint64(code), 10)
	errorString := "{\"error\":" + err +", \"code\":"+codeStr+" \" }"
	return shim.Error(errorString)
}

func successResponse(dat string) peer.Response {
	success := "{\"response\": " + dat +", \"code\": 0 }"
	return shim.Success([]byte(success))
}

func main() {
	err := shim.Start(new(EnergyOriginChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
