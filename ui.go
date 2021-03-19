package main

import (
  "fmt"
  "os"

  "github.com/manifoldco/promptui"
  "gopkg.in/ini.v1"
)

func runUI(f *ini.File) {
  options := []string{"Set Default Key", "Delete Key", "Add a Key", "Exit"}
  prompt := promptui.Select{
    Label: "Select Day",
    Items: options,
  }

  _, result, err := prompt.Run()

  if err != nil {
    fmt.Printf("Prompt failed %v\n", err)
    return
  }
  switch result {
    case options[0]: //Set Default Key
    requestNewKeyToDefault(f)
    case options[1]: //Delete Key
    requestKeyToDelete(f)
    case options[2]: //Add a Key
    requestKeyToAdd(f)
    case options[3]: //Exit
    os.Exit(0)
  }
}

func promptKeys(f *ini.File) (string, error) {
  prompt := promptui.Select{
    Label: "Select the key to Default",
    Items: GetKeys(f),
    Size: 15,
  }

  _, result, err := prompt.Run()
  return result, err
}

func requestNewKeyToDefault(f *ini.File) {
  res, err := promptKeys(f)
  if err != nil {
    fmt.Printf("Prompt failed %v\n", err)
    return
  }
  SetDefault(f, res)

}
func requestKeyToDelete(f *ini.File) {
  res, err := promptKeys(f)
  if err != nil {
    fmt.Printf("Prompt failed %v\n", err)
    return
  }
  DeleteKey(f, res)
}

func askUserForText(title string) string {
  prompt := promptui.Prompt {
    Label: title,
  }
  result, err := prompt.Run()
  if err != nil {
    fmt.Printf("Prompt failed %v\n", err)
    return ""
  }
  return result
}
func requestKeyToAdd(f *ini.File) {

  name := askUserForText("Enter the Key Name")
  key := askUserForText("Enter the Access Key ID")
  secret := askUserForText("Enter the Secret Key ID")
  AddKey(f, name, key, secret)
}
