package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpatialRelationInfo_PDC_r17_referenceSignal_Choice_nothing uint64 = iota
	SpatialRelationInfo_PDC_r17_referenceSignal_Choice_ssb_Index
	SpatialRelationInfo_PDC_r17_referenceSignal_Choice_csi_RS_Index
	SpatialRelationInfo_PDC_r17_referenceSignal_Choice_dl_PRS_PDC
	SpatialRelationInfo_PDC_r17_referenceSignal_Choice_srs
)

type SpatialRelationInfo_PDC_r17_referenceSignal struct {
	Choice       uint64
	ssb_Index    *SSB_Index
	csi_RS_Index *NZP_CSI_RS_ResourceId
	dl_PRS_PDC   *NR_DL_PRS_ResourceID_r17
	srs          *SpatialRelationInfo_PDC_r17_referenceSignal_srs
}

func (ie *SpatialRelationInfo_PDC_r17_referenceSignal) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_ssb_Index:
		if err = ie.ssb_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb_Index", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_csi_RS_Index:
		if err = ie.csi_RS_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode csi_RS_Index", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_dl_PRS_PDC:
		if err = ie.dl_PRS_PDC.Encode(w); err != nil {
			err = utils.WrapError("Encode dl_PRS_PDC", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_srs:
		if err = ie.srs.Encode(w); err != nil {
			err = utils.WrapError("Encode srs", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SpatialRelationInfo_PDC_r17_referenceSignal) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_ssb_Index:
		ie.ssb_Index = new(SSB_Index)
		if err = ie.ssb_Index.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Index", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_csi_RS_Index:
		ie.csi_RS_Index = new(NZP_CSI_RS_ResourceId)
		if err = ie.csi_RS_Index.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_Index", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_dl_PRS_PDC:
		ie.dl_PRS_PDC = new(NR_DL_PRS_ResourceID_r17)
		if err = ie.dl_PRS_PDC.Decode(r); err != nil {
			return utils.WrapError("Decode dl_PRS_PDC", err)
		}
	case SpatialRelationInfo_PDC_r17_referenceSignal_Choice_srs:
		ie.srs = new(SpatialRelationInfo_PDC_r17_referenceSignal_srs)
		if err = ie.srs.Decode(r); err != nil {
			return utils.WrapError("Decode srs", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
