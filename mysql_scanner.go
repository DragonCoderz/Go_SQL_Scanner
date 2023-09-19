package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	host := "127.0.0.1" // Replace with the target host's IP address or "localhost"
	port := "3306"      // Replace with the target port (default is 3306)

	target := fmt.Sprintf("%s:%s", host, port)

	// Send a client handshake packet to initiate communication
	clientHandshake := []byte{0x0a, 0x00, 0x00, 0x00, 0x0a}

	// Create a connection to the target
	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Printf("MySQL does not appear to be running on host %s at port %s\n", host, port)
		return
	}
	defer conn.Close()

	// Write the handshake packet to the connection
	_, err = conn.Write(clientHandshake)
	if err != nil {
		fmt.Printf("Error sending handshake: %v\n", err)
		return
	}

	// Read the server's response
	serverResponse := make([]byte, 1024)
	_, err = conn.Read(serverResponse)
	if err != nil {
		fmt.Printf("Error reading server response: %v\n", err)
		return
	}

	// Extract MySQL version and protocol information from the response
	versionStart := bytes.Index(serverResponse, []byte{0x0a, 0x00, 0x00, 0x00, 0x0a}) + 5
	versionEnd := bytes.Index(serverResponse[versionStart:], []byte{0x00}) + versionStart
	mysqlVersion := string(serverResponse[versionStart:versionEnd])

	// Extract server capabilities (4 bytes, little-endian)
	serverCapabilities := binary.LittleEndian.Uint32(serverResponse[5:9])

	// Hard code the translation of server capabilities to actual output
	capabilities := map[uint32]string{
		0x00000001: "CLIENT_LONG_PASSWORD",
		0x00000002: "CLIENT_FOUND_ROWS",
		0x00000004: "CLIENT_LONG_FLAG_COUNT",
		0x00000008: "CLIENT_CONNECT_WITH_DB",
		0x00000010: "CLIENT_NO_SCHEMA",
		0x00000020: "CLIENT_COMPRESS",
		0x00000040: "CLIENT_MULTI_STATEMENTS",
		0x00000080: "CLIENT_MULTI_RESULTS",
		0x00000100: "CLIENT_PS_MULTI_RESULTS",
		0x00000200: "CLIENT_PLUGIN_AUTH",
		0x00000400: "CLIENT_SESSION_TRACK",
		0x00000800: "CLIENT_SSL",
	}

	// Extract server status (2 bytes, little-endian)
	serverStatus := binary.LittleEndian.Uint16(serverResponse[9:11])

	// Hard code the translation of server status to actual output
	serverStatusFlags := map[uint16]string{
		0x0001: "SERVER_STATUS_IN_TRANS",
		0x0002: "SERVER_STATUS_AUTOCOMMIT",
		0x0004: "SERVER_STATUS_MORE_RESULTS_EXISTS",
		0x0008: "SERVER_STATUS_NO_GOOD_INDEX_USED",
		0x0010: "SERVER_STATUS_NO_GOOD_INDEX_FOUND",
		0x0020: "SERVER_STATUS_NOT_QUERIED",
		0x0040: "SERVER_STATUS_CURSOR_EXISTS",
		0x0080: "SERVER_STATUS_LAST_ROW_SENT",
		0x0100: "SERVER_STATUS_DB_DROPPED",
		0x0200: "SERVER_STATUS_NO_BACKSLASH_ESCAPES",
		0x0400: "SERVER_STATUS_METADATA_CHANGED",
		0x0800: "SERVER_STATUS_QUERY_WAS_SLOW",
		0x1000: "SERVER_STATUS_PS_OUT_PARAMS",
		0x2000: "SERVER_STATUS_IN_TRANS_READONLY",
		0x4000: "SERVER_STATUS_SESSION_STATE_CHANGED",
	}

	// Extract server plugin name (null-terminated string)
	pluginNameStart := bytes.Index(serverResponse[10:], []byte{0x00}) + 10
	pluginNameEnd := bytes.Index(serverResponse[pluginNameStart:], []byte{0x00}) + pluginNameStart
	pluginName := string(serverResponse[pluginNameStart:pluginNameEnd])
	if pluginName == "" {
		pluginName = "None"
	}

	// Print extracted information
	fmt.Printf("MySQL is running on host %s at port %s\n", host, port)
	fmt.Printf("MySQL version: %s\n", mysqlVersion)
	fmt.Printf("Server capabilities:\n")
	for capability, name := range capabilities {
		if serverCapabilities&capability != 0 {
			fmt.Printf("    %s\n", name)
		}
	}
	// Print the translated server status
	fmt.Printf("Server status:\n")
	for flag, name := range serverStatusFlags {
		if serverStatus&flag != 0 {
			fmt.Printf("    %s\n", name)
		}
	}

	fmt.Printf("Server plugin name: %s\n", pluginName)

	if err != nil {
		fmt.Printf("Error writing packet to file: %v\n", err)
	}
}
