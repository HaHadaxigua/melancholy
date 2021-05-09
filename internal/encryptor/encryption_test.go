package encryptor

import (
	"fmt"
	"testing"
)

func TestCalcLargeFileHashInSHA1ByName(t *testing.T) {
	pngFilename := "D:\\data\\pictures\\wallpaper\\pixiv\\82847627_p0.png"
	hash, err := CalcLargeFileHashInSHA1ByName(pngFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hash)


	txtFilename := "F:\\Download\\package-lock.json"
	hash, err = CalcLargeFileHashInSHA1ByName(txtFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hash)

}
