package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_SpatialRelationInfoPos_r16_Choice_nothing uint64 = iota
	SRS_SpatialRelationInfoPos_r16_Choice_servingRS_r16
	SRS_SpatialRelationInfoPos_r16_Choice_ssb_Ncell_r16
	SRS_SpatialRelationInfoPos_r16_Choice_dl_PRS_r16
)

type SRS_SpatialRelationInfoPos_r16 struct {
	Choice        uint64
	servingRS_r16 *SRS_SpatialRelationInfoPos_r16_servingRS_r16
	ssb_Ncell_r16 *SSB_InfoNcell_r16
	dl_PRS_r16    *DL_PRS_Info_r16
}

func (ie *SRS_SpatialRelationInfoPos_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_SpatialRelationInfoPos_r16_Choice_servingRS_r16:
		if err = ie.servingRS_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode servingRS_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_Choice_ssb_Ncell_r16:
		if err = ie.ssb_Ncell_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb_Ncell_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_Choice_dl_PRS_r16:
		if err = ie.dl_PRS_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode dl_PRS_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_SpatialRelationInfoPos_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_SpatialRelationInfoPos_r16_Choice_servingRS_r16:
		ie.servingRS_r16 = new(SRS_SpatialRelationInfoPos_r16_servingRS_r16)
		if err = ie.servingRS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode servingRS_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_Choice_ssb_Ncell_r16:
		ie.ssb_Ncell_r16 = new(SSB_InfoNcell_r16)
		if err = ie.ssb_Ncell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Ncell_r16", err)
		}
	case SRS_SpatialRelationInfoPos_r16_Choice_dl_PRS_r16:
		ie.dl_PRS_r16 = new(DL_PRS_Info_r16)
		if err = ie.dl_PRS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dl_PRS_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
