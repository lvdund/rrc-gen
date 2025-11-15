package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_nothing uint64 = iota
	UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_uePositioningAssistanceInfo_r17
	UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_criticalExtensionsFuture
)

type UEPositioningAssistanceInfo_r17_CriticalExtensions struct {
	Choice                          uint64
	uePositioningAssistanceInfo_r17 *UEPositioningAssistanceInfo_r17_IEs
	criticalExtensionsFuture        interface{} `madatory`
}

func (ie *UEPositioningAssistanceInfo_r17_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_uePositioningAssistanceInfo_r17:
		if err = ie.uePositioningAssistanceInfo_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode uePositioningAssistanceInfo_r17", err)
		}
	case UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UEPositioningAssistanceInfo_r17_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_uePositioningAssistanceInfo_r17:
		ie.uePositioningAssistanceInfo_r17 = new(UEPositioningAssistanceInfo_r17_IEs)
		if err = ie.uePositioningAssistanceInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode uePositioningAssistanceInfo_r17", err)
		}
	case UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
