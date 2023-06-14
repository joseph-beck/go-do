package todo

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeDeadline(t *testing.T) {
	d := makeDeadline("00/00/0000-00:00")
	assert.NotNil(t, d)

	if d.Date.Str() != "00/00/0000" {
		log.Fatalln("Bad date conversion for deadline")
	}

	if d.Time.Str() != "00:00" {
		log.Fatalln("Bad time conversion for deadline")
	}

	if d.Str() != "00/00/0000-00:00" {
		log.Fatalln("Bad deadline conversion for deadline")
	}
}

func TestVerifyDeadline(t *testing.T) {

}
