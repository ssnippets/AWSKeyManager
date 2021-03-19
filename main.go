package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

/*
Features:

1. Select a key to set as default
2. Delete a key
3. Add a key
4. Create ~/.aws/credentials file if it doesn't exist
5. Save credentials file


** Have a menu!
*/

type Credential struct {
	AccessKey string `ini:"aws_access_key_id"`
	SecretKey string `init:"aws_secret_access_key"`
}

const (
	AWS_CREDENTIAL_FILENAME = "/.aws/credentials"
)

func main() {
	home, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("Failed to get user home folder: %v", err)
		os.Exit(1)
	}

	cfg, err := ini.Load(home + AWS_CREDENTIAL_FILENAME)
	if err != nil {
		fmt.Printf("Failed to read credentials file: %v", err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error mapping credentials: %v", err)
		os.Exit(1)
	}

	secs := cfg.Sections()
	for i, s := range secs {
		fmt.Println(i, s.Name())
	}

	SetDefault(cfg, "snoise")

	_ = AddKey(cfg, "Mile Two New", "key name", "secret!!")
	SaveCreds(cfg)
}

func SaveCreds(f *ini.File) error {
	err := f.SaveTo("./creds.ini")
	if err != nil {
		fmt.Printf("Error saving credentials: %v", err)
		return err
	}
	return nil
}

func SetDefault(f *ini.File, name string) error {
	//Remove the default key
	f.DeleteSection("default")
	s, err := f.GetSection(name)
	if err != nil {
		fmt.Printf("Error, Section %v doesn't exist: %v", name, err)
		return err
	}

	// Copy the values over:
	for _, k := range s.Keys() {
		f.Section("default").NewKey(k.Name(), k.Value())
	}
	return nil
}

func DeleteKey(f *ini.File, name string) {
	f.DeleteSection(name)
}

func AddKey(f *ini.File, name string, key string, secret string) error {
	_, err := f.GetSection(name)
	if err == nil {
		return errors.New("Key name already exists")
	}

	c := &Credential{key, secret}

  err = f.Section(name).ReflectFrom(c)
	if err != nil {
    fmt.Printf("Error in mapping: %v", err);
		return errors.New("Key name already exists")
	}

	return nil
}
