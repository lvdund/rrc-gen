package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PathlossReferenceRS_r17_referenceSignal_r17_Choice_nothing uint64 = iota
	PathlossReferenceRS_r17_referenceSignal_r17_Choice_ssb_Index
	PathlossReferenceRS_r17_referenceSignal_r17_Choice_csi_RS_Index
)

type PathlossReferenceRS_r17_referenceSignal_r17 struct {
	Choice       uint64
	ssb_Index    *SSB_Index
	csi_RS_Index *NZP_CSI_RS_ResourceId
}

func (ie *PathlossReferenceRS_r17_referenceSignal_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PathlossReferenceRS_r17_referenceSignal_r17_Choice_ssb_Index:
		if err = ie.ssb_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb_Index", err)
		}
	case PathlossReferenceRS_r17_referenceSignal_r17_Choice_csi_RS_Index:
		if err = ie.csi_RS_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode csi_RS_Index", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PathlossReferenceRS_r17_referenceSignal_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PathlossReferenceRS_r17_referenceSignal_r17_Choice_ssb_Index:
		ie.ssb_Index = new(SSB_Index)
		if err = ie.ssb_Index.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Index", err)
		}
	case PathlossReferenceRS_r17_referenceSignal_r17_Choice_csi_RS_Index:
		ie.csi_RS_Index = new(NZP_CSI_RS_ResourceId)
		if err = ie.csi_RS_Index.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_Index", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
