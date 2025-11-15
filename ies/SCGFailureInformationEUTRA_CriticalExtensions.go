package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCGFailureInformationEUTRA_CriticalExtensions_Choice_nothing uint64 = iota
	SCGFailureInformationEUTRA_CriticalExtensions_Choice_scgFailureInformationEUTRA
	SCGFailureInformationEUTRA_CriticalExtensions_Choice_criticalExtensionsFuture
)

type SCGFailureInformationEUTRA_CriticalExtensions struct {
	Choice                     uint64
	scgFailureInformationEUTRA *SCGFailureInformationEUTRA_IEs
	criticalExtensionsFuture   interface{} `madatory`
}

func (ie *SCGFailureInformationEUTRA_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCGFailureInformationEUTRA_CriticalExtensions_Choice_scgFailureInformationEUTRA:
		if err = ie.scgFailureInformationEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode scgFailureInformationEUTRA", err)
		}
	case SCGFailureInformationEUTRA_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SCGFailureInformationEUTRA_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCGFailureInformationEUTRA_CriticalExtensions_Choice_scgFailureInformationEUTRA:
		ie.scgFailureInformationEUTRA = new(SCGFailureInformationEUTRA_IEs)
		if err = ie.scgFailureInformationEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode scgFailureInformationEUTRA", err)
		}
	case SCGFailureInformationEUTRA_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
