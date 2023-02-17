package dumpHL7

import (
	"testing"

	"github.com/minerofish/go-hl7"
	"github.com/minerofish/go-hl7/lib/hl7v23"
)

// not a test - just sample use case
func Test_dump(t *testing.T) {
	sample := `MSH|^~\&|SWISSLAB|FFM||FFM|20230203080903||ORM^O01|o3057937.000001|P|2.3|
PID|||01077843||CL5AVA0N7K|||U|
PV1||S|||||||||||||||||S01077843|
ORC|NW||23071012||||||20230203080800|||BSD|
OBR|1||23071012|HINAET^HIV-1/2 PCR mit erhöhter Sensitivität|||20230203080900|||||||||BSD|||||||||P|
OBR|2||23071012|HCNAET^HCV PCR mit erhöhter Sensitivität|||20230203080900|||||||||BSD|||||||||P|
OBR|3||23071012|HBNAET^HBV PCR mit erhöhter Sensitivität|||20230203080900|||||||||BSD|||||||||P|`
	var dest hl7v23.ORM_O01
	hl7.Unmarshal([]byte(sample), &dest, hl7.EncodingUTF8, hl7.TimezoneEuropeBerlin)
	dump(dest, "dest")
}
