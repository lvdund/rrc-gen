package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RateMatchPatternLTE_CRS struct {
	carrierFreqDL            int64                                      `lb:0,ub:16383,madatory`
	carrierBandwidthDL       RateMatchPatternLTE_CRS_carrierBandwidthDL `madatory`
	mbsfn_SubframeConfigList *EUTRA_MBSFN_SubframeConfigList            `optional`
	nrofCRS_Ports            RateMatchPatternLTE_CRS_nrofCRS_Ports      `madatory`
	v_Shift                  RateMatchPatternLTE_CRS_v_Shift            `madatory`
}

func (ie *RateMatchPatternLTE_CRS) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mbsfn_SubframeConfigList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.carrierFreqDL, &uper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
		return utils.WrapError("WriteInteger carrierFreqDL", err)
	}
	if err = ie.carrierBandwidthDL.Encode(w); err != nil {
		return utils.WrapError("Encode carrierBandwidthDL", err)
	}
	if ie.mbsfn_SubframeConfigList != nil {
		if err = ie.mbsfn_SubframeConfigList.Encode(w); err != nil {
			return utils.WrapError("Encode mbsfn_SubframeConfigList", err)
		}
	}
	if err = ie.nrofCRS_Ports.Encode(w); err != nil {
		return utils.WrapError("Encode nrofCRS_Ports", err)
	}
	if err = ie.v_Shift.Encode(w); err != nil {
		return utils.WrapError("Encode v_Shift", err)
	}
	return nil
}

func (ie *RateMatchPatternLTE_CRS) Decode(r *uper.UperReader) error {
	var err error
	var mbsfn_SubframeConfigListPresent bool
	if mbsfn_SubframeConfigListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_carrierFreqDL int64
	if tmp_int_carrierFreqDL, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
		return utils.WrapError("ReadInteger carrierFreqDL", err)
	}
	ie.carrierFreqDL = tmp_int_carrierFreqDL
	if err = ie.carrierBandwidthDL.Decode(r); err != nil {
		return utils.WrapError("Decode carrierBandwidthDL", err)
	}
	if mbsfn_SubframeConfigListPresent {
		ie.mbsfn_SubframeConfigList = new(EUTRA_MBSFN_SubframeConfigList)
		if err = ie.mbsfn_SubframeConfigList.Decode(r); err != nil {
			return utils.WrapError("Decode mbsfn_SubframeConfigList", err)
		}
	}
	if err = ie.nrofCRS_Ports.Decode(r); err != nil {
		return utils.WrapError("Decode nrofCRS_Ports", err)
	}
	if err = ie.v_Shift.Decode(r); err != nil {
		return utils.WrapError("Decode v_Shift", err)
	}
	return nil
}
