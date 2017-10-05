package main

import "golang.org/x/sys/windows/registry"
import "log"
import "fmt"

func main() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatalf("opening registry: %v", err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue("ProxyServer")
	if err != nil {
		log.Fatalf("ProxyServer key: %v", err)
	}
	fmt.Printf("Windows ProxyServer is %q\n", s)

	i, _, err := k.GetIntegerValue("ProxyEnable")
	if err != nil {
		log.Fatalf("ProxyEnable key: %v", err)
	}
	fmt.Printf("Windows ProxyEnable is %v\n", i)

	info, err := k.Stat()
	if err != nil {
		log.Fatalf("Getting Stats: %v", err)
	}
	fmt.Printf("Last modification time: %q\n", info.ModTime())

}
