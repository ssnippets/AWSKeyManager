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
func getCredentialsFile() string {
  home, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("Failed to get user home folder: %v\n", err)
		os.Exit(1)
	}
  return home + AWS_CREDENTIAL_FILENAME

}
func loadCredentials() *ini.File {
  cfg, err := ini.Load(getCredentialsFile())

	if err != nil {
		fmt.Printf("Failed to read credentials file: %v\n", err)
		os.Exit(1)
	}
  return cfg
}

func main() {
	for true {
		cfg := loadCredentials()
		runUI(cfg)
		SaveCreds(cfg)
	}
}

func GetKeys(f *ini.File) []string {
	secs := f.Sections()
	names := make([]string, len(secs))
	for i, s := range secs {
		names[i] = s.Name()
	}
	return names
}

func SaveCreds(f *ini.File) error {
	err := f.SaveTo(getCredentialsFile())
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
		fmt.Printf("Error in mapping: %v", err)
		return errors.New("Key name already exists")
	}

	return nil
}
