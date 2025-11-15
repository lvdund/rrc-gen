package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCGFailureInformation_CriticalExtensions_Choice_nothing uint64 = iota
	SCGFailureInformation_CriticalExtensions_Choice_scgFailureInformation
	SCGFailureInformation_CriticalExtensions_Choice_criticalExtensionsFuture
)

type SCGFailureInformation_CriticalExtensions struct {
	Choice                   uint64
	scgFailureInformation    *SCGFailureInformation_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *SCGFailureInformation_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCGFailureInformation_CriticalExtensions_Choice_scgFailureInformation:
		if err = ie.scgFailureInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode scgFailureInformation", err)
		}
	case SCGFailureInformation_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SCGFailureInformation_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCGFailureInformation_CriticalExtensions_Choice_scgFailureInformation:
		ie.scgFailureInformation = new(SCGFailureInformation_IEs)
		if err = ie.scgFailureInformation.Decode(r); err != nil {
			return utils.WrapError("Decode scgFailureInformation", err)
		}
	case SCGFailureInformation_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
