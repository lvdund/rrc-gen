package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_nothing uint64 = iota
	MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_csi_RS_Resource_r17
	MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_ssb_Resource_r17
)

type MPE_Resource_r17_mpe_ReferenceSignal_r17 struct {
	Choice              uint64
	csi_RS_Resource_r17 *NZP_CSI_RS_ResourceId
	ssb_Resource_r17    *SSB_Index
}

func (ie *MPE_Resource_r17_mpe_ReferenceSignal_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_csi_RS_Resource_r17:
		if err = ie.csi_RS_Resource_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode csi_RS_Resource_r17", err)
		}
	case MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_ssb_Resource_r17:
		if err = ie.ssb_Resource_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb_Resource_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MPE_Resource_r17_mpe_ReferenceSignal_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_csi_RS_Resource_r17:
		ie.csi_RS_Resource_r17 = new(NZP_CSI_RS_ResourceId)
		if err = ie.csi_RS_Resource_r17.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_Resource_r17", err)
		}
	case MPE_Resource_r17_mpe_ReferenceSignal_r17_Choice_ssb_Resource_r17:
		ie.ssb_Resource_r17 = new(SSB_Index)
		if err = ie.ssb_Resource_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Resource_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
