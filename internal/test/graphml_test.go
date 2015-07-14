package test

import (
	// "fmt"
	"github.com/carneades/carneades-4/internal/engine/caes"
	// "github.com/carneades/carneades-4/internal/engine/caes/encoding/graphml"
	"github.com/carneades/carneades-4/internal/engine/caes/encoding/yaml"
	//	"log"
	// "errors"
	"os"
	"testing"
)

func gmlcheck(t *testing.T, e error) {
	if e != nil {
		t.Errorf(e.Error())
		t.Skip("Skip Test")
	}
}

func TestIOGmlTandem(t *testing.T) {
	ioGmlTest(t, "../../examples/AGs/tandem.yml", "../../examples/AGs/TempTandem.graphml")
}

func TestIOGmlBachelor(t *testing.T) {
	ioGmlTest(t, "../../examples/AGs/bachelor.yml", "../../examples/AGs/TempBachelor.graphml")
}

func TestIOGmlFrisan(t *testing.T) {
	ioGmlTest(t, "../../examples/AGs/frisian.yml", "../../examples/AGs/TempFrisian.graphml")
}

func TestIOGmlJogging(t *testing.T) {
	ioGmlTest(t, "../../examples/AGs/jogging.yml", "../../examples/AGs/TempJogging.graphml")
}

func TestIOGmlSherlock(t *testing.T) {
	ioGmlTest(t, "../../examples/AGs/sherlock.yml", "../../examples/AGs/TempSherlock.graphml")
}

func TestIOGmlVacation(t *testing.T) {
	ioGmlTest(t, "../../examples/AGs/vacation.yml", "../../examples/AGs/TempVacation.graphml")
}

func ioGmlTest(t *testing.T, filename1 string, filename2 string) {

	var ag *caes.ArgGraph
	var err error

	file, err := os.Open(filename1)
	gmlcheck(t, err)
	ag, err = yaml.Import(file)
	file.Close()
	gmlcheck(t, err)
	//	fmt.Printf("---------- WriteArgGraph %s ----------\n", filename1)
	//	yaml.ExportWithReferences(os.Stdout, ag)
	//	fmt.Printf("---------- End: WriteArgGraph %s ----------\n", filename1)
	l := ag.GroundedLabelling()
	ag.ApplyLabelling(l)
	//	fmt.Printf("---------- printLabeling %s ----------\n", filename1)
	//	printLabeling(l)
	//	fmt.Printf("---------- End: printLabeling %s ----------\n", filename1)
	//file, err = os.Create(filename2)
	//gmlcheck(t, err)
	//graphml.Export(file, *ag)

}
