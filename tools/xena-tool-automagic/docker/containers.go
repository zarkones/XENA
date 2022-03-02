package docker

import (
	"errors"
	"main/shell"
)

func InitAtila(serviceSecret, dbSecret, publicKey string) error {
	if err := initAtilaDatabase(dbSecret); err != nil {
		return err
	}
	if err := initAtilaService(serviceSecret, dbSecret, publicKey); err != nil {
		return err
	}
	return nil
}

func initAtilaDatabase(secret string) error {
	if _, err := shell.Run("docker run -d --name xena-atila-postgres --net xena -e POSTGRES_DB=xena-atila -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=" + secret + " postgres"); err != nil {
		return errors.New(" " + err.Error() + ". Failed to initialize xena-atila-postgres container")
	}
	return nil
}

func initAtilaService(serviceSecret, dbSecret, publicKey string) error {
	address, err := shell.Run("docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' xena-atila-postgres")
	if err != nil {
		return errors.New(" " + err.Error() + ". Unable to get the IP address of xena-atila-postgres")
	}
	if _, err := shell.Run("cd services/xena-service-atila && docker build -t xena-service-atila . && docker run -d --net xena --name='xena-atila' -e PG_HOST='" + address + "' -e CORS_POLICY_ALLOWED_ORIGINS='http://127.0.0.1:3000' -e PG_PASSWORD='" + dbSecret + "' -e APP_KEY='" + serviceSecret + "' -e TRUSTED_PUBLIC_KEY='" + publicKey + "' -p 60666:60666 xena-service-atila"); err != nil {
		return errors.New(" " + err.Error() + ". Unable to build and run xena-atila container.")
	}
	if _, err := shell.Run("docker exec -ti xena-atila sh -c \"node build/ace migration:run --force\""); err != nil {
		return errors.New(" " + err.Error() + ". Unable to run migrations on xena-atila container.")
	}
	return nil
}

func InitPyramid(serviceSecret, dbSecret, publicKey string) error {
	err := initPyramidDatabase(dbSecret)
	if err != nil {
		return err
	}
	initPyramidService(serviceSecret, dbSecret, publicKey)
	if err != nil {
		return err
	}
	return nil
}

func initPyramidDatabase(secret string) error {
	if _, err := shell.Run("docker run -d --name xena-pyramid-postgres --net xena -e POSTGRES_DB=xena-pyramid -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=" + secret + " postgres"); err != nil {
		return errors.New(" " + err.Error() + ". Failed to initialize xena-pyramid-postgres container")
	}
	return nil
}

func initPyramidService(serviceSecret, dbSecret, publicKey string) error {
	address, err := shell.Run("docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' xena-pyramid-postgres")
	if err != nil {
		return errors.New(" " + err.Error() + ". Unable to get the IP address of xena-pyramid-postgres")
	}
	if _, err := shell.Run("cd services/xena-service-pyramid && docker build -t xena-service-pyramid . && docker run -d --net xena --name='xena-pyramid' -e XENA_GIT_BRANCH='stage' -e PG_HOST='" + address + "' -e CORS_POLICY_ALLOWED_ORIGINS='http://127.0.0.1:3000' -e PG_PASSWORD='" + dbSecret + "' -e APP_KEY='" + serviceSecret + "' -e TRUSTED_PUBLIC_KEY='" + publicKey + "' -p 60667:60667 xena-service-pyramid"); err != nil {
		return errors.New(" " + err.Error() + ". Unable to build and run xena-pyramid container.")
	}
	if _, err := shell.Run("docker exec -ti xena-pyramid sh -c \"node build/ace migration:run --force\""); err != nil {
		return errors.New(" " + err.Error() + ". Unable to run migrations on xena-pyramid container.")
	}
	return nil
}

func InitDomena(serviceSecret, dbSecret, publicKey string) error {
	err := initDomenaDatabase(dbSecret)
	if err != nil {
		return err
	}
	err = initDomenaService(serviceSecret, dbSecret, publicKey)
	if err != nil {
		return err
	}
	return nil
}

func initDomenaDatabase(secret string) error {
	if _, err := shell.Run("docker run -d --name xena-domena-postgres --net xena -e POSTGRES_DB=xena-domena -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=" + secret + " postgres"); err != nil {
		return errors.New(" " + err.Error() + ". Failed to initialize xena-domena-postgres container")
	}
	return nil
}

func initDomenaService(serviceSecret, dbSecret, publicKey string) error {
	address, err := shell.Run("docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' xena-domena-postgres")
	if err != nil {
		return errors.New(" " + err.Error() + ". Unable to get the IP address of xena-domena-postgres")
	}
	if _, err := shell.Run("cd services/xena-service-domena && docker build -t xena-service-domena . && docker run -d --net xena --name='xena-domena' -e PG_HOST='" + address + "' -e CORS_POLICY_ALLOWED_ORIGINS='http://127.0.0.1:3000' -e PG_PASSWORD='" + dbSecret + "' -e APP_KEY='" + serviceSecret + "' -e TRUSTED_PUBLIC_KEY='" + publicKey + "' -p 60798:60798 xena-service-domena"); err != nil {
		return errors.New(" " + err.Error() + ". Unable to build and run xena-domena container.")
	}
	if _, err := shell.Run("docker exec -ti xena-domena sh -c \"node build/ace migration:run --force\""); err != nil {
		return errors.New(" " + err.Error() + ". Unable to run migrations on xena-domena container.")
	}
	return nil
}

func InitGateway() error {
	err := initGatewayService()
	if err != nil {
		return err
	}
	return nil
}

func initGatewayService() error {
	domenaAddress, err := shell.Run("docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' xena-atila")
	if err != nil {
		return errors.New(" " + err.Error() + ". Unable to get the IP address of xena-atila")
	}
	atilaAddress, err := shell.Run("docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' xena-domena")
	if err != nil {
		return errors.New(" " + err.Error() + ". Unable to get the IP address of xena-domena")
	}
	if _, err := shell.Run("cd services/xena-service-gateway && docker build -t xena-service-gateway . && docker run -d --net xena --name='xena-gateway'  -e DOMENA_HOST='http://" + domenaAddress + ":60798' -e ATILA_HOST='http://" + atilaAddress + ":60666' -p 60606:60606 xena-service-gateway"); err != nil {
		return errors.New(" " + err.Error() + ". Unable to build and run xena-gateway container.")
	}
	return nil
}

func InitFace() error {
	err := initFaceService()
	if err != nil {
		return err
	}
	return nil
}

func initFaceService() error {
	if _, err := shell.Run("cd user-interfaces/xena-service-face && docker build -t xena-service-face . && docker run -d -p 3000:3000 --net xena --name='xena-face' xena-service-face"); err != nil {
		return errors.New(" " + err.Error() + ". Unable to build and run xena-face container.")
	}
	return nil
}
