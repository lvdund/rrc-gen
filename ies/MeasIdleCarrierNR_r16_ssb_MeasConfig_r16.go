package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdleCarrierNR_r16_ssb_MeasConfig_r16 struct {
	nrofSS_BlocksToAverage_r16          *int64               `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	absThreshSS_BlocksConsolidation_r16 *ThresholdNR         `optional`
	smtc_r16                            *SSB_MTC             `optional`
	ssb_ToMeasure_r16                   *SSB_ToMeasure       `optional`
	deriveSSB_IndexFromCell_r16         bool                 `madatory`
	ss_RSSI_Measurement_r16             *SS_RSSI_Measurement `optional`
}

func (ie *MeasIdleCarrierNR_r16_ssb_MeasConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nrofSS_BlocksToAverage_r16 != nil, ie.absThreshSS_BlocksConsolidation_r16 != nil, ie.smtc_r16 != nil, ie.ssb_ToMeasure_r16 != nil, ie.ss_RSSI_Measurement_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nrofSS_BlocksToAverage_r16 != nil {
		if err = w.WriteInteger(*ie.nrofSS_BlocksToAverage_r16, &uper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Encode nrofSS_BlocksToAverage_r16", err)
		}
	}
	if ie.absThreshSS_BlocksConsolidation_r16 != nil {
		if err = ie.absThreshSS_BlocksConsolidation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode absThreshSS_BlocksConsolidation_r16", err)
		}
	}
	if ie.smtc_r16 != nil {
		if err = ie.smtc_r16.Encode(w); err != nil {
			return utils.WrapError("Encode smtc_r16", err)
		}
	}
	if ie.ssb_ToMeasure_r16 != nil {
		if err = ie.ssb_ToMeasure_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_ToMeasure_r16", err)
		}
	}
	if err = w.WriteBoolean(ie.deriveSSB_IndexFromCell_r16); err != nil {
		return utils.WrapError("WriteBoolean deriveSSB_IndexFromCell_r16", err)
	}
	if ie.ss_RSSI_Measurement_r16 != nil {
		if err = ie.ss_RSSI_Measurement_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ss_RSSI_Measurement_r16", err)
		}
	}
	return nil
}

func (ie *MeasIdleCarrierNR_r16_ssb_MeasConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var nrofSS_BlocksToAverage_r16Present bool
	if nrofSS_BlocksToAverage_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var absThreshSS_BlocksConsolidation_r16Present bool
	if absThreshSS_BlocksConsolidation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var smtc_r16Present bool
	if smtc_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_ToMeasure_r16Present bool
	if ssb_ToMeasure_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ss_RSSI_Measurement_r16Present bool
	if ss_RSSI_Measurement_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if nrofSS_BlocksToAverage_r16Present {
		var tmp_int_nrofSS_BlocksToAverage_r16 int64
		if tmp_int_nrofSS_BlocksToAverage_r16, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Decode nrofSS_BlocksToAverage_r16", err)
		}
		ie.nrofSS_BlocksToAverage_r16 = &tmp_int_nrofSS_BlocksToAverage_r16
	}
	if absThreshSS_BlocksConsolidation_r16Present {
		ie.absThreshSS_BlocksConsolidation_r16 = new(ThresholdNR)
		if err = ie.absThreshSS_BlocksConsolidation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode absThreshSS_BlocksConsolidation_r16", err)
		}
	}
	if smtc_r16Present {
		ie.smtc_r16 = new(SSB_MTC)
		if err = ie.smtc_r16.Decode(r); err != nil {
			return utils.WrapError("Decode smtc_r16", err)
		}
	}
	if ssb_ToMeasure_r16Present {
		ie.ssb_ToMeasure_r16 = new(SSB_ToMeasure)
		if err = ie.ssb_ToMeasure_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_ToMeasure_r16", err)
		}
	}
	var tmp_bool_deriveSSB_IndexFromCell_r16 bool
	if tmp_bool_deriveSSB_IndexFromCell_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean deriveSSB_IndexFromCell_r16", err)
	}
	ie.deriveSSB_IndexFromCell_r16 = tmp_bool_deriveSSB_IndexFromCell_r16
	if ss_RSSI_Measurement_r16Present {
		ie.ss_RSSI_Measurement_r16 = new(SS_RSSI_Measurement)
		if err = ie.ss_RSSI_Measurement_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ss_RSSI_Measurement_r16", err)
		}
	}
	return nil
}
