package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MBSBroadcastConfiguration_r17_CriticalExtensions_Choice_nothing uint64 = iota
	MBSBroadcastConfiguration_r17_CriticalExtensions_Choice_mbsBroadcastConfiguration_r17
	MBSBroadcastConfiguration_r17_CriticalExtensions_Choice_criticalExtensionsFuture
)

type MBSBroadcastConfiguration_r17_CriticalExtensions struct {
	Choice                        uint64
	mbsBroadcastConfiguration_r17 *MBSBroadcastConfiguration_r17_IEs
	criticalExtensionsFuture      interface{} `madatory`
}

func (ie *MBSBroadcastConfiguration_r17_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MBSBroadcastConfiguration_r17_CriticalExtensions_Choice_mbsBroadcastConfiguration_r17:
		if err = ie.mbsBroadcastConfiguration_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode mbsBroadcastConfiguration_r17", err)
		}
	case MBSBroadcastConfiguration_r17_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MBSBroadcastConfiguration_r17_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MBSBroadcastConfiguration_r17_CriticalExtensions_Choice_mbsBroadcastConfiguration_r17:
		ie.mbsBroadcastConfiguration_r17 = new(MBSBroadcastConfiguration_r17_IEs)
		if err = ie.mbsBroadcastConfiguration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mbsBroadcastConfiguration_r17", err)
		}
	case MBSBroadcastConfiguration_r17_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
