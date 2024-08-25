package main

import (
	"flag"
	"log"
	"os"
	"tests/cases"
	"tests/common"
)

func main() {
	build := flag.Bool("build", false, "build the release version of XENA")
	test := flag.Bool("test", false, "build and test the release version of XENA")

	flag.Parse()

	if err := os.Chdir(common.ROOT_PATH_RELATIVE_TO_TESTER); err != nil {
		log.Fatalln("failed to set root of the project as working directory:", err)
	}

	// Pre-compilation tests.
	if *test {
		log.Println("pre-compilation tests... IN PROGRESS")

		if err := cases.AgentMainFileValidation(); err != nil {
			log.Fatalln("cases.AgentMainFileValidation failed:", err)
		}

		// TODO: Add a test case that if --prod is passed then it would check that DEBUG mode is not enabled in .go file.

		log.Println("pre-compilation tests... DONE")
	}

	if *build {
		if err := buildXENA(); err != nil {
			log.Fatalln("building of the project failed with generic error")
		}
	}

	// Post-compilation tests.
	if *test {
		log.Println("post-compilation tests...")

		if err := cases.C2(); err != nil {
			log.Fatalln("cases.C2():", err)
		}

		log.Println("post-compilation tests... DONE")
	}
}
