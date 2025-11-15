package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqInfo_v1700 struct {
	interFreqNeighHSDN_CellList_r17 *InterFreqNeighHSDN_CellList_r17                           `optional`
	highSpeedMeasInterFreq_r17      *InterFreqCarrierFreqInfo_v1700_highSpeedMeasInterFreq_r17 `optional`
	redCapAccessAllowed_r17         *InterFreqCarrierFreqInfo_v1700_redCapAccessAllowed_r17    `optional`
	ssb_PositionQCL_Common_r17      *SSB_PositionQCL_Relation_r17                              `optional`
	interFreqNeighCellList_v1710    *InterFreqNeighCellList_v1710                              `optional`
}

func (ie *InterFreqCarrierFreqInfo_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.interFreqNeighHSDN_CellList_r17 != nil, ie.highSpeedMeasInterFreq_r17 != nil, ie.redCapAccessAllowed_r17 != nil, ie.ssb_PositionQCL_Common_r17 != nil, ie.interFreqNeighCellList_v1710 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.interFreqNeighHSDN_CellList_r17 != nil {
		if err = ie.interFreqNeighHSDN_CellList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqNeighHSDN_CellList_r17", err)
		}
	}
	if ie.highSpeedMeasInterFreq_r17 != nil {
		if err = ie.highSpeedMeasInterFreq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedMeasInterFreq_r17", err)
		}
	}
	if ie.redCapAccessAllowed_r17 != nil {
		if err = ie.redCapAccessAllowed_r17.Encode(w); err != nil {
			return utils.WrapError("Encode redCapAccessAllowed_r17", err)
		}
	}
	if ie.ssb_PositionQCL_Common_r17 != nil {
		if err = ie.ssb_PositionQCL_Common_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_PositionQCL_Common_r17", err)
		}
	}
	if ie.interFreqNeighCellList_v1710 != nil {
		if err = ie.interFreqNeighCellList_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqNeighCellList_v1710", err)
		}
	}
	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1700) Decode(r *uper.UperReader) error {
	var err error
	var interFreqNeighHSDN_CellList_r17Present bool
	if interFreqNeighHSDN_CellList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var highSpeedMeasInterFreq_r17Present bool
	if highSpeedMeasInterFreq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var redCapAccessAllowed_r17Present bool
	if redCapAccessAllowed_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_PositionQCL_Common_r17Present bool
	if ssb_PositionQCL_Common_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqNeighCellList_v1710Present bool
	if interFreqNeighCellList_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	if interFreqNeighHSDN_CellList_r17Present {
		ie.interFreqNeighHSDN_CellList_r17 = new(InterFreqNeighHSDN_CellList_r17)
		if err = ie.interFreqNeighHSDN_CellList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqNeighHSDN_CellList_r17", err)
		}
	}
	if highSpeedMeasInterFreq_r17Present {
		ie.highSpeedMeasInterFreq_r17 = new(InterFreqCarrierFreqInfo_v1700_highSpeedMeasInterFreq_r17)
		if err = ie.highSpeedMeasInterFreq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedMeasInterFreq_r17", err)
		}
	}
	if redCapAccessAllowed_r17Present {
		ie.redCapAccessAllowed_r17 = new(InterFreqCarrierFreqInfo_v1700_redCapAccessAllowed_r17)
		if err = ie.redCapAccessAllowed_r17.Decode(r); err != nil {
			return utils.WrapError("Decode redCapAccessAllowed_r17", err)
		}
	}
	if ssb_PositionQCL_Common_r17Present {
		ie.ssb_PositionQCL_Common_r17 = new(SSB_PositionQCL_Relation_r17)
		if err = ie.ssb_PositionQCL_Common_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_PositionQCL_Common_r17", err)
		}
	}
	if interFreqNeighCellList_v1710Present {
		ie.interFreqNeighCellList_v1710 = new(InterFreqNeighCellList_v1710)
		if err = ie.interFreqNeighCellList_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqNeighCellList_v1710", err)
		}
	}
	return nil
}
