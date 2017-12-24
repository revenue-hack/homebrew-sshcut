package main

import (
	"testing"
)

func TestGetConfigFilePathOfNormal(t *testing.T) {
	filePath, isExist := getConfigFilePath()
	if isExist && filePath == "" {
		t.Error("file path is different in getConfigFilePath")
	}
	if !isExist && filePath != "" {
		t.Error("file path is not empty in getConfigFilePath")
	}
}

func TestReadFileOfNormal(t *testing.T) {
	hostList, isExist := readFile("./test.txt")
	if !isExist {
		t.Error("./test.txt nothing")
	}
	for i, host := range hostList.Hosts {
		if i == 0 && host.Alias != "test" {
			t.Error("Hosts index 0 is invalid")
		}
		if i == 1 && host.Alias != "test1" {
			t.Error("Hosts index 1 is invalid")
		}
	}
}

func TestObjectMappingOfNormal(t *testing.T) {
	host, isExist := objectMapping("Host test")
	if isExist && host.Alias != "test" {
		t.Error("return invalid ObjectMapping Host")
	}
	if !isExist {
		t.Error("return invalid ObjectMapping isExist")
	}

	_, isExist = objectMapping("Alias test")
	if isExist {
		t.Error("return invalid ObjectMapping isExist")
	}
}

