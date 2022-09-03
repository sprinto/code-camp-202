package main

/*
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<rpc-reply message-id="c81f933d-8197-4dbc-94cf-2a6c3f8942b7" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
    <data>
        <toaster xmlns="http://netconfcentral.org/ns/toaster">
            <darknessFactor>4711</darknessFactor>
        </toaster>
    </data>
</rpc-reply>
*/

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/*
// array of all rpc replies
type RpcReplies struct {
	XMLName    xml.Name   `xml:"data"`
	RpcReplies []RpcReply `xml:"rpcReply"`
}

// array of all toasters
type RpcReply struct {
	MessageId    string        `xml:"message-id,attr"`
	DataElements []DataElement `xml:"data"`
}

// array of all toasters
type DataElement struct {
	Toasters []Toaster `xml:"toaster"`
}

// a toaster
type Toaster struct {
	XMLName             xml.Name `xml:"toaster"`
	DarknessFactor      int      `xml:"darknessFactor"`
	ToasterManufacturer string   `xml:"toasterManufacturer"`
	ToasterModelNumber  string   `xml:"toasterModelNumber"`
}

*/

// generated from https://www.onlinetool.io/xmltogo/
type RpcReply struct {
	XMLName   xml.Name `xml:"rpc-reply"`
	Text      string   `xml:",chardata"`
	MessageID string   `xml:"message-id,attr"`
	Xmlns     string   `xml:"xmlns,attr"`
	Data      struct {
		Text    string `xml:",chardata"`
		Toaster []struct {
			Text                string `xml:",chardata"`
			Xmlns               string `xml:"xmlns,attr"`
			DarknessFactor      int    `xml:"darknessFactor"`
			ToasterManufacturer string `xml:"toasterManufacturer"`
			ToasterModelNumber  string `xml:"toasterModelNumber"`
		} `xml:"toaster"`
	} `xml:"data"`
}

func main() {

	// Open our xmlFile
	xmlFile, err := os.Open("toast.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened toast.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err)
	}

	// we initialize our rply array
	var rpcReply RpcReply

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rpcReply)
	toasters := rpcReply.Data.Toaster

	for i := 0; i < len(toasters); i++ {
		fmt.Println("Manufacturer: " + toasters[i].ToasterManufacturer)
		fmt.Println("Model number: " + toasters[i].ToasterModelNumber)
		fmt.Printf("Darkness: %d\n", toasters[i].DarknessFactor)
	}

}
