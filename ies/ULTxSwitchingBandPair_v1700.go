package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ULTxSwitchingBandPair_v1700 struct {
	uplinkTxSwitchingPeriod2T2T_r17 *ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17 `optional`
}

func (ie *ULTxSwitchingBandPair_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uplinkTxSwitchingPeriod2T2T_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.uplinkTxSwitchingPeriod2T2T_r17 != nil {
		if err = ie.uplinkTxSwitchingPeriod2T2T_r17.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkTxSwitchingPeriod2T2T_r17", err)
		}
	}
	return nil
}

func (ie *ULTxSwitchingBandPair_v1700) Decode(r *uper.UperReader) error {
	var err error
	var uplinkTxSwitchingPeriod2T2T_r17Present bool
	if uplinkTxSwitchingPeriod2T2T_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if uplinkTxSwitchingPeriod2T2T_r17Present {
		ie.uplinkTxSwitchingPeriod2T2T_r17 = new(ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17)
		if err = ie.uplinkTxSwitchingPeriod2T2T_r17.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkTxSwitchingPeriod2T2T_r17", err)
		}
	}
	return nil
}
