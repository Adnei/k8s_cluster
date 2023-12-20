package k8_cluster

import "fmt"

func Hello(name string) string {
  message := fmt.Sprintf("Hello %v. Welcome!", name)
  return message
}
