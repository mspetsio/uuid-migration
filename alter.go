package main

import (
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	// var s = "1, mike, den kserw"
	// var sss []string = strings.Split(s, ",")

	// m := test1(os.Args[1])
	test3(os.Args[3], test2(os.Args[2], test1(os.Args[1])))

	// for n := range m {
	// 	fmt.Printf("%s\t%s\n", n, m[n])
	// }

}

func test1(filename string) map[string]string {

	f, err := os.Create("test1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var m = make(map[string]string)
	data, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(string(data), "\n") {

		// fmt.Printf("%s\n", line)
		sss := strings.Split(line, ",")
		uiiiid := uuid.New().String()
		m[sss[0]] = uiiiid
		strings.Join(sss[1:], ",")
		_, err2 := f.WriteString(m[sss[0]] + "," + strings.Join(sss, ",") + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
		// fmt.Printf("%s,%s\n", m[sss[0]], strings.Join(sss, ","))
	}
	return m

}

func test2(filename string, m map[string]string) map[string]string {

	f, err := os.Create("test2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var n = make(map[string]string)
	data, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(string(data), "\n") {

		// fmt.Printf("%s\n", line)
		sss := strings.Split(line, ",")
		uiiiid := uuid.New().String()
		n[sss[0]] = uiiiid
		i := strings.TrimSpace(sss[5])
		_, err2 := f.WriteString(n[sss[0]] + "," + strings.Join(sss[1:5], ",") + "," + m[i] + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
		//cockroach sql --url 'postgres://root@:26257/movr?ssl=disable' --insecure
		// fmt.Printf("%s,%s,%s\n", n[sss[0]], strings.Join(sss[1:5], ","), m[i])
	}
	return n

}

func test3(filename string, m map[string]string) map[string]string {

	f, err := os.Create("test3.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var n = make(map[string]string)
	data, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(string(data), "\n") {

		// fmt.Printf("%s\n", line)
		sss := strings.Split(line, ",")
		uiiiid := uuid.New().String()
		n[sss[0]] = uiiiid
		i := strings.TrimSpace(sss[1])
		_, err2 := f.WriteString(n[sss[0]] + "," + m[i] + "," + strings.Join(sss[2:], ",") + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
		//cockroach sql --url 'postgres://root@:26257/movr?ssl=disable' --insecure
		// fmt.Printf("%s,%s,%s\n", n[sss[0]], strings.Join(sss[1:5], ","), m[i])
	}
	return n

}
