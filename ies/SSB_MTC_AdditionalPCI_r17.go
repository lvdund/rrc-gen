package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_MTC_AdditionalPCI_r17 struct {
	additionalPCIIndex_r17   AdditionalPCIIndex_r17                             `madatory`
	additionalPCI_r17        PhysCellId                                         `madatory`
	periodicity_r17          SSB_MTC_AdditionalPCI_r17_periodicity_r17          `madatory`
	ssb_PositionsInBurst_r17 SSB_MTC_AdditionalPCI_r17_ssb_PositionsInBurst_r17 `lb:4,ub:4,madatory`
	ss_PBCH_BlockPower_r17   int64                                              `lb:-60,ub:50,madatory`
}

func (ie *SSB_MTC_AdditionalPCI_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.additionalPCIIndex_r17.Encode(w); err != nil {
		return utils.WrapError("Encode additionalPCIIndex_r17", err)
	}
	if err = ie.additionalPCI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode additionalPCI_r17", err)
	}
	if err = ie.periodicity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode periodicity_r17", err)
	}
	if err = ie.ssb_PositionsInBurst_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_PositionsInBurst_r17", err)
	}
	if err = w.WriteInteger(ie.ss_PBCH_BlockPower_r17, &uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("WriteInteger ss_PBCH_BlockPower_r17", err)
	}
	return nil
}

func (ie *SSB_MTC_AdditionalPCI_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.additionalPCIIndex_r17.Decode(r); err != nil {
		return utils.WrapError("Decode additionalPCIIndex_r17", err)
	}
	if err = ie.additionalPCI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode additionalPCI_r17", err)
	}
	if err = ie.periodicity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode periodicity_r17", err)
	}
	if err = ie.ssb_PositionsInBurst_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_PositionsInBurst_r17", err)
	}
	var tmp_int_ss_PBCH_BlockPower_r17 int64
	if tmp_int_ss_PBCH_BlockPower_r17, err = r.ReadInteger(&uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("ReadInteger ss_PBCH_BlockPower_r17", err)
	}
	ie.ss_PBCH_BlockPower_r17 = tmp_int_ss_PBCH_BlockPower_r17
	return nil
}
