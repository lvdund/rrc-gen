package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxSwitchingBandParameters_v1700 struct {
	bandIndex_r17                                  int64                                                                                 `lb:1,ub:maxSimultaneousBands,madatory`
	uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17 *UplinkTxSwitchingBandParameters_v1700_uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17 `optional`
}

func (ie *UplinkTxSwitchingBandParameters_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.bandIndex_r17, &uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
		return utils.WrapError("WriteInteger bandIndex_r17", err)
	}
	if ie.uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17 != nil {
		if err = ie.uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17", err)
		}
	}
	return nil
}

func (ie *UplinkTxSwitchingBandParameters_v1700) Decode(r *uper.UperReader) error {
	var err error
	var uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17Present bool
	if uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_bandIndex_r17 int64
	if tmp_int_bandIndex_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
		return utils.WrapError("ReadInteger bandIndex_r17", err)
	}
	ie.bandIndex_r17 = tmp_int_bandIndex_r17
	if uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17Present {
		ie.uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17 = new(UplinkTxSwitchingBandParameters_v1700_uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17)
		if err = ie.uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkTxSwitching2T2T_PUSCH_TransCoherence_r17", err)
		}
	}
	return nil
}
