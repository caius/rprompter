package main;

import (
  "fmt"
  "bytes"
  "log"
  "time"
  "os/exec"
)

func promptTime() string {
  const layout = "2006-01-02 15:04:05"
  t := time.Now()
  // Mon Jan 2 15:04:05 MST 2006
  return t.Format(layout)
}

func promptBattery() string {
  cmd := exec.Command("pmset", "-g", "batt")

  output, err := cmd.Output()
  if err != nil {
    log.Fatal(err)
  }

  output = bytes.Split(output, []byte("\n"))[1]
  output = bytes.Split(output, []byte("	"))[1]
  fields := bytes.Split(output, []byte("; "))

  return fmt.Sprintf("%s (%s)\n", fields[0], fields[2])
}

func main() {
  fmt.Printf("%s | %s", promptTime(), promptBattery())
}
