package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_SpatialRelationInfo_referenceSignal_Choice_nothing uint64 = iota
	PUCCH_SpatialRelationInfo_referenceSignal_Choice_ssb_Index
	PUCCH_SpatialRelationInfo_referenceSignal_Choice_csi_RS_Index
	PUCCH_SpatialRelationInfo_referenceSignal_Choice_srs
)

type PUCCH_SpatialRelationInfo_referenceSignal struct {
	Choice       uint64
	ssb_Index    *SSB_Index
	csi_RS_Index *NZP_CSI_RS_ResourceId
	srs          *PUCCH_SRS
}

func (ie *PUCCH_SpatialRelationInfo_referenceSignal) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_SpatialRelationInfo_referenceSignal_Choice_ssb_Index:
		if err = ie.ssb_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb_Index", err)
		}
	case PUCCH_SpatialRelationInfo_referenceSignal_Choice_csi_RS_Index:
		if err = ie.csi_RS_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode csi_RS_Index", err)
		}
	case PUCCH_SpatialRelationInfo_referenceSignal_Choice_srs:
		if err = ie.srs.Encode(w); err != nil {
			err = utils.WrapError("Encode srs", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUCCH_SpatialRelationInfo_referenceSignal) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_SpatialRelationInfo_referenceSignal_Choice_ssb_Index:
		ie.ssb_Index = new(SSB_Index)
		if err = ie.ssb_Index.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Index", err)
		}
	case PUCCH_SpatialRelationInfo_referenceSignal_Choice_csi_RS_Index:
		ie.csi_RS_Index = new(NZP_CSI_RS_ResourceId)
		if err = ie.csi_RS_Index.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_Index", err)
		}
	case PUCCH_SpatialRelationInfo_referenceSignal_Choice_srs:
		ie.srs = new(PUCCH_SRS)
		if err = ie.srs.Decode(r); err != nil {
			return utils.WrapError("Decode srs", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
