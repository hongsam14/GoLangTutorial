package smEventXml

import "encoding/xml"

type SysmonEventXml struct {
	XMLNAME xml.Name `xml:"Event"`
	System    SystemXml `xml:"System"`
}

/***********************************************************************
*	common data
************************************************************************/

type ProviderXml struct {
	Name string `xml:"Name,attr"`
	Guid string `xml:"Guid,attr"`
}

type TimeCreatedXml struct {
	SystemTime string `xml:"SystemTime,attr"`
}

type ExecutionXml struct {
	ProcessID string `xml:"ProcessID,attr"`
	ThreadID  string `xml:"ThreadID,attr"`
}

type SecurityXml struct {
	UserID string `xml:"UserId.attr"`
}

type SystemXml struct {
	Provider      ProviderXml    `xml:"Provider"`
	EventID       int            `xml:"EventID"`
	Version       int            `xml:"Version"`
	Level         int            `xml:"Level"`
	Task          int            `xml:"Task"`
	Opcode        string         `xml:"Opcode"`
	TimeCreated   TimeCreatedXml `xml:"TimeCreated"`
	EventRecordID int            `xml:"EventRecordID"`
	Correlation   string         `xml:"Correlation"`
	Execution     ExecutionXml   `xml:"Execution"`
	Channel       string         `xml:"Channel"`
	Computer      string         `xml:"Computer"`
	Security      SecurityXml    `xml:"Security"`
}

/***********************************************************************
*	event data
************************************************************************/

type DataXml struct {
	Name string `xml:"Name,attr"`
	Data string 
}