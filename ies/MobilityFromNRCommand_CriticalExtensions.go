package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MobilityFromNRCommand_CriticalExtensions_Choice_nothing uint64 = iota
	MobilityFromNRCommand_CriticalExtensions_Choice_mobilityFromNRCommand
	MobilityFromNRCommand_CriticalExtensions_Choice_criticalExtensionsFuture
)

type MobilityFromNRCommand_CriticalExtensions struct {
	Choice                   uint64
	mobilityFromNRCommand    *MobilityFromNRCommand_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *MobilityFromNRCommand_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MobilityFromNRCommand_CriticalExtensions_Choice_mobilityFromNRCommand:
		if err = ie.mobilityFromNRCommand.Encode(w); err != nil {
			err = utils.WrapError("Encode mobilityFromNRCommand", err)
		}
	case MobilityFromNRCommand_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MobilityFromNRCommand_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MobilityFromNRCommand_CriticalExtensions_Choice_mobilityFromNRCommand:
		ie.mobilityFromNRCommand = new(MobilityFromNRCommand_IEs)
		if err = ie.mobilityFromNRCommand.Decode(r); err != nil {
			return utils.WrapError("Decode mobilityFromNRCommand", err)
		}
	case MobilityFromNRCommand_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
