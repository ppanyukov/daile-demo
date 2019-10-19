package main

import (
	"fmt"
	"os"

	"i.love.java/no/just/kidding/pkg/assets"
	"i.love.java/no/just/kidding/pkg/version"
)

func main() {
	// Print version wired into exe at build time.
	v := version.Print("go-boot")
	fmt.Println()
	fmt.Println("VERSION:")
	fmt.Println("---------------------------------------------")
	fmt.Println(v)
	fmt.Println("=============================================")
	fmt.Println()

	// Print the list of assets embedded in the exe.
	fmt.Println()
	fmt.Println("ASSETS:")
	fmt.Println("---------------------------------------------")
	for _, a := range assets.AssetNames() {
		fmt.Println(a)
	}
	fmt.Println("=============================================")
	fmt.Println()

	// Print the file `assets/files/sample.txt` embedded into the binary at build time.
	fmt.Println()
	fmt.Println("ASSET assets/files/sample.txt")
	fmt.Println("---------------------------------------------")
	content := assets.MustAsset("assets/files/sample.txt")
	os.Stdout.Write(content)
	fmt.Println("=============================================")
}
