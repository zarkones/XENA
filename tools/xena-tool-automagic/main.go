package main

import (
	"fmt"
	"main/crypto"
	"main/docker"
	"main/shell"
	"strings"
)

func issue() {
	fmt.Println()
	fmt.Println()
	fmt.Println("- - - LOGS START - - -")

	fmt.Println("- - - - - - - XENA_ATILA - - - - - - -")
	logs, err := shell.Run("docker container logs xena-atila")
	if err != nil {
		fmt.Print("An error happened during the printing of container logs: ")
		fmt.Println(err)
	}
	fmt.Println(logs)
	fmt.Println("- - - - - - - XENA_ATILA - - - - - - -")
	fmt.Println("- - - - - - - XENA_ATILA_POSTGRES - - - - - - -")
	logs, err = shell.Run("docker container logs xena-atila-postgres")
	if err != nil {
		fmt.Print("An error happened during the printing of container logs: ")
		fmt.Println(err)
	}
	fmt.Println(logs)

	fmt.Println("- - - - - - - XENA_DOMENA - - - - - - -")
	logs, err = shell.Run("docker container logs xena-domena")
	if err != nil {
		fmt.Print("An error happened during the printing of container logs: ")
		fmt.Println(err)
	}
	fmt.Println(logs)
	fmt.Println("- - - - - - - XENA_DOMENA_POSTGRES - - - - - - -")
	logs, err = shell.Run("docker container logs xena-domena-postgres")
	if err != nil {
		fmt.Print("An error happened during the printing of container logs: ")
		fmt.Println(err)
	}
	fmt.Println(logs)

	fmt.Println("- - - - - - - XENA_PYRAMID - - - - - - -")
	logs, err = shell.Run("docker container logs xena-pyramid")
	if err != nil {
		fmt.Print("An error happened during the printing of container logs: ")
		fmt.Println(err)
	}
	fmt.Println(logs)
	fmt.Println("- - - - - - - XENA_PYRAMID_POSTGRES - - - - - - -")
	logs, err = shell.Run("docker container logs xena-pyramid-postgres")
	if err != nil {
		fmt.Print("An error happened during the printing of container logs: ")
		fmt.Println(err)
	}
	fmt.Println(logs)

	fmt.Println("- - - - - - - XENA_FACE - - - - - - -")
	logs, err = shell.Run("docker container logs xena-face")
	if err != nil {
		fmt.Print("An error happened during the printing of container logs: ")
		fmt.Println(err)
	}
	fmt.Println(logs)

	fmt.Println("- - - - - - - XENA_GATEWAY - - - - - - -")
	logs, err = shell.Run("docker container logs xena-gateway")
	if err != nil {
		fmt.Print("An error happened during the printing of container logs: ")
		fmt.Println(err)
	}
	fmt.Println(logs)

	fmt.Println("- - - DOCKER PS - - -")
	logs, err = shell.Run("docker ps -a | grep xena")
	if err != nil {
		fmt.Print("An error happened during the printing of container logs: ")
		fmt.Println(err)
	}
	fmt.Println(logs)

	fmt.Println("- - - LOGS END - - -")
	fmt.Println()
	fmt.Println()

	fmt.Println("Seems like something went wrong.")
	fmt.Println("Please, submit a ticket at: https://github.com/zarkones/XENA/issues")

	fmt.Println()
	fmt.Println("Provide the logs when submiting a ticket.")
	fmt.Println()

	fmt.Println("FIX TIPS:")
	fmt.Println(`
	1. "exit status 125" indicates a conflict in container names. Delete all containers with the prefix 'xena' and run setup process again.
	`)
}

func main() {
	fmt.Println("Welcome to XENA!")
	fmt.Println("Setup process initiated. This may take some time.")

	fmt.Print("Creating a key pair. ")
	privateKeyOrg, publicKeyOrg, err := crypto.KeyPair()
	if err != nil {
		issue()
		panic("CRYPTO_PRIVATE_KEY_GENERATION_FAILED")
	}
	fmt.Println("Created.")

	publicKey := strings.ReplaceAll(publicKeyOrg, "\n", "\\n")

	atilaDbSecret := crypto.UniqueSecret()
	atilaKeySecret := crypto.UniqueSecret()

	pyramidDbSecret := crypto.UniqueSecret()
	pyramidKeySecret := crypto.UniqueSecret()

	domenaDbSecret := crypto.UniqueSecret()
	domenaKeySecret := crypto.UniqueSecret()

	fmt.Print("Checking Docker. ")
	if err = docker.Download(); err != nil {
		issue()
		fmt.Println(err)
		panic("DOCKER_DOWNLOAD_FAILED")
	}
	fmt.Println("Installed.")

	fmt.Print("Creating a docker network. ")
	if err = docker.CreateNetwork(); err != nil {
		issue()
		fmt.Println(err)
		panic("DOCKER_NETWORK_CREATION_FAILED")
	}
	fmt.Println("Created.")

	fmt.Print("Creating two containers: xena-atila, xena-atila-postgres. ")
	if err = docker.InitAtila(atilaKeySecret, atilaDbSecret, publicKey); err != nil {
		issue()
		fmt.Println(err)
		panic("DOCKER_INIT_ATILA_FAILED")
	}
	fmt.Println("Success.")

	fmt.Print("Creating two containers: xena-domena, xena-domena-postgres. ")
	if err = docker.InitDomena(domenaKeySecret, domenaDbSecret, publicKey); err != nil {
		issue()
		fmt.Println(err)
		panic("DOCKER_INIT_DOMENA_FAILED")
	}
	fmt.Println("Success.")

	fmt.Print("Creating two containers: xena-pyramid, xena-pyramid-postgres. ")
	if err = docker.InitPyramid(pyramidKeySecret, pyramidDbSecret, publicKey); err != nil {
		issue()
		fmt.Println(err)
		panic("DOCKER_INIT_PYRAMID_FAILED")
	}
	fmt.Println("Success.")

	fmt.Print("Creating one container: xena-gateway. ")
	if err = docker.InitGateway(); err != nil {
		issue()
		fmt.Println(err)
		panic("DOCKER_INIT_GATEWAY_FAILED")
	}
	fmt.Println("Success.")

	fmt.Print("Creating one container: xena-Face. ")
	if err = docker.InitFace(); err != nil {
		issue()
		fmt.Println(err)
		panic("DOCKER_INIT_FACE_FAILED")
	}
	fmt.Println("Success.")
	fmt.Println()
	fmt.Println("Installation process complete!")
	fmt.Println()

	fmt.Println(privateKeyOrg)

	fmt.Println(publicKeyOrg)

	fmt.Println("\n 1. SAVE YOUR KEYS")
	fmt.Println(" 2. Login using any username and private key at http://127.0.0.1:3000")
	fmt.Println(" 3. Using the menu bar on the left navigate to Settings, or go to http://127.0.0.1:3000/settings")
	fmt.Println(" 4. Switch to 'IDENTITY' tab, then click on 'Generate A Token' button.")
	fmt.Println(" 5. Copy the generated JSON Web Token.")
	fmt.Println(" 6. Go back to 'CONNECTIONS' tab.")
	fmt.Println(" 7. Paste that token in the text-fields which have prefix 'Token for'.")

	fmt.Println("\nDONATIONS:")
	fmt.Println(`
	https://www.patreon.com/zarkones
	
	Monero: 87RZRSuASU4cdVXWSXvxLUQ84MpnbpxNHfkucSP9gjNE4TzCUSWT5H7fYunk7JLGPpE9QwHXjZQtaNpeyuKsB8WWLGz13ZJ

	Ethereum: 0x787Ba8EF8d75489160C6687296839F947DC62736
	`)

	fmt.Println("\nDISCLAIMER:")
	fmt.Println(`
	You bear the full responsibility of your actions. This is an open-source project.
	This is not a commercial project. This project is not related to my current nor
	past employers. This is my hobby and I'm developing this tool for the sake of
	learning and understanding, meaning educational purpose. You shall not hold viable
	its creator/s nor any contributor to the project for any damage you may have done.
	If you contribute to the project bare in mind that your code or data may be changed
	in the future without a notice and that it belongs to the project, not you. This
	software is not allowed to be used for political purpose. This software is not
	allowed to be used for commercial purpose of any kind, shape or form without a
	written permission. Selling and redistributing this software is not permitted
	without a written permission. This software is not allowed to be used for
	training of algorithms without a written permission.
	`)

}
