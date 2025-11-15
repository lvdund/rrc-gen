package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	TCI_UL_State_r17_referenceSignal_r17_Choice_nothing uint64 = iota
	TCI_UL_State_r17_referenceSignal_r17_Choice_ssb_Index_r17
	TCI_UL_State_r17_referenceSignal_r17_Choice_csi_RS_Index_r17
	TCI_UL_State_r17_referenceSignal_r17_Choice_srs_r17
)

type TCI_UL_State_r17_referenceSignal_r17 struct {
	Choice           uint64
	ssb_Index_r17    *SSB_Index
	csi_RS_Index_r17 *NZP_CSI_RS_ResourceId
	srs_r17          *SRS_ResourceId
}

func (ie *TCI_UL_State_r17_referenceSignal_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TCI_UL_State_r17_referenceSignal_r17_Choice_ssb_Index_r17:
		if err = ie.ssb_Index_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb_Index_r17", err)
		}
	case TCI_UL_State_r17_referenceSignal_r17_Choice_csi_RS_Index_r17:
		if err = ie.csi_RS_Index_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode csi_RS_Index_r17", err)
		}
	case TCI_UL_State_r17_referenceSignal_r17_Choice_srs_r17:
		if err = ie.srs_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode srs_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *TCI_UL_State_r17_referenceSignal_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TCI_UL_State_r17_referenceSignal_r17_Choice_ssb_Index_r17:
		ie.ssb_Index_r17 = new(SSB_Index)
		if err = ie.ssb_Index_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Index_r17", err)
		}
	case TCI_UL_State_r17_referenceSignal_r17_Choice_csi_RS_Index_r17:
		ie.csi_RS_Index_r17 = new(NZP_CSI_RS_ResourceId)
		if err = ie.csi_RS_Index_r17.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_Index_r17", err)
		}
	case TCI_UL_State_r17_referenceSignal_r17_Choice_srs_r17:
		ie.srs_r17 = new(SRS_ResourceId)
		if err = ie.srs_r17.Decode(r); err != nil {
			return utils.WrapError("Decode srs_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
