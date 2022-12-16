
/*
 * The sample smart contract for documentation topic:
 * Managing IOT Data on Hyperledger Blockchain
 */

 package main

 /* Imports
  * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
  * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
  */
 import (
	 "encoding/json"
	 "fmt"
	 "strconv"
 	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // Define the Smart Contract structure
 type SmartContract struct {
 }
  // Define the IotData structure, with 2 properties.  Structure tags are used by encoding/json library
 type IotData struct {
	 Temperature   string `json:"temperature"`
	 Humidity  string `json:"humidity"`
 }
 
 /*
  * The Init method is called when the Smart Contract "IOT" is instantiated by the blockchain network
  * Best practice is to have any Ledger initialization in separate function -- see initLedger()
  */
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }
 
 /*
  * The Invoke method is called as a result of an application request to run the Smart Contract "IOT"
  * The calling application program has also specified the particular smart contract function to be called, with arguments
  */
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 // Retrieve the requested Smart Contract function and arguments
	 function, args := APIstub.GetFunctionAndParameters()
	 // Route to the appropriate handler function to interact with the ledger appropriately
	 if function == "queryIotData" {
		 return s.queryIotData(APIstub, args)
	 } else if function == "initLedger" {
		 return s.initLedger(APIstub)
	 } else if function == "createIotData" {
		 return s.createIotData(APIstub, args)
	 } 

	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 func (s *SmartContract) queryIotData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
 
	 IotDataAsBytes, _ := APIstub.GetState(args[0])
	 return shim.Success(IotDataAsBytes)
 }
 
 func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	IotDatas := []IotData{
		IotData{Temperature: "0", Humidity: "0"},
	 }
 
	 i := 0
	 for i < len(IotDatas) {
		 fmt.Println("i is ", i)
		 IotDataAsBytes, _ := json.Marshal(IotDatas[i])
		 APIstub.PutState("IotData"+strconv.Itoa(i), IotDataAsBytes)
		 fmt.Println("Added", IotDatas[i])
		 i = i + 1
	 }
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) createIotData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 3 {
		 return shim.Error("Incorrect number of arguments. Expecting 3")
	 }
 	 var iotData = IotData{Temperature: args[1], Humidity: args[2]}
 	 iotDataAsBytes, _ := json.Marshal(iotData)
	 APIstub.PutState(args[0], iotDataAsBytes)
 	 return shim.Success(nil)
 }
 
//  The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {    
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }
 