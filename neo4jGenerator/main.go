package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gregpechiro/neo4jGenerator"
)

var (
	out = flag.String("o", "", "what file to write")
	pkg = flag.String("pkg", ".", "what package to get the interface from")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	st := flag.Arg(0)

	if st == "" {
		log.Fatal("need to specify a struct name")
	}

	gen, err := neo4jGenerator.NewNeo4jGenerator(*pkg, st)
	if err != nil {
		log.Fatalf("main.go >> neo4jGenerator.NewNeo4jGenerator() >> %v\n", err)
	}

	var buf bytes.Buffer

	log.Printf("Generating Neo4j Service Layer for %s\n", st)
	err = gen.Write(&buf)
	if err != nil {
		log.Fatalf("main.go >> gen.Write() >> %v\n", err)
	}

	if *out == "" {
		*out = ToLowerFirst(st) + "Neo4jService_generated.go"

	}
	if err := ioutil.WriteFile(*out, buf.Bytes(), 0666); err != nil {
		log.Fatalf("main.go >> ioutil.WriteFile() >> %v\n", err)
	}
}

func ToLowerFirst(s string) string {
	if len(s) > 1 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	return strings.ToLower(s)
}
