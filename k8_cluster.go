package k8_scluster

import "fmt"

func Hello(name string) string {
  message := fmt.Sprintf("Hello %v. Welcome!", name)
  return message
}
