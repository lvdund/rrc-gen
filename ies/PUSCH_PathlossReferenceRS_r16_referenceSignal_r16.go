package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_nothing uint64 = iota
	PUSCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_ssb_Index_r16
	PUSCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_csi_RS_Index_r16
)

type PUSCH_PathlossReferenceRS_r16_referenceSignal_r16 struct {
	Choice           uint64
	ssb_Index_r16    *SSB_Index
	csi_RS_Index_r16 *NZP_CSI_RS_ResourceId
}

func (ie *PUSCH_PathlossReferenceRS_r16_referenceSignal_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUSCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_ssb_Index_r16:
		if err = ie.ssb_Index_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb_Index_r16", err)
		}
	case PUSCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_csi_RS_Index_r16:
		if err = ie.csi_RS_Index_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode csi_RS_Index_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUSCH_PathlossReferenceRS_r16_referenceSignal_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUSCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_ssb_Index_r16:
		ie.ssb_Index_r16 = new(SSB_Index)
		if err = ie.ssb_Index_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Index_r16", err)
		}
	case PUSCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_csi_RS_Index_r16:
		ie.csi_RS_Index_r16 = new(NZP_CSI_RS_ResourceId)
		if err = ie.csi_RS_Index_r16.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_Index_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
