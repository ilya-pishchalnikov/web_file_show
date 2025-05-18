package util

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func GetExecDir() string {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting path to executable:", err)
		log.Fatal(err)
	}
	return filepath.Dir(execPath) + "/"
}

func PostFile(text []byte, postMethod string, username string, password string) error {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("POST", postMethod, bytes.NewBuffer(text))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	request.Header.Set("Content-Type", "text/plain")

	/*credentialsString := os.Getenv("filesender_credentials")

	if credentialsString == "" {
		return errors.New("Credentials are not set in the filesender_credentials environment variable")
	}

	loginPassword := strings.Split(credentialsString, "=")

	if len(loginPassword) != 2 {
		return errors.New("Invalid format of the filesender_credentials environment variable")
	}

	username := loginPassword[0]
	password := loginPassword[1]*/

	auth := username + ":" + password
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	request.Header.Set("Authorization", "Basic "+encoded)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error executing POST request:", err)
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading answer:", err)
		return err
	}

	fmt.Println("response:", string(body))

	return nil
}
