package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_Configuration_r16 struct {
	ssb_Freq_r16             ARFCN_ValueNR                              `madatory`
	halfFrameIndex_r16       SSB_Configuration_r16_halfFrameIndex_r16   `madatory`
	ssbSubcarrierSpacing_r16 SubcarrierSpacing                          `madatory`
	ssb_Periodicity_r16      *SSB_Configuration_r16_ssb_Periodicity_r16 `optional`
	sfn0_Offset_r16          *SSB_Configuration_r16_sfn0_Offset_r16     `lb:0,ub:1023,optional`
	sfn_SSB_Offset_r16       int64                                      `lb:0,ub:15,madatory`
	ss_PBCH_BlockPower_r16   *int64                                     `lb:-60,ub:50,optional`
}

func (ie *SSB_Configuration_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssb_Periodicity_r16 != nil, ie.sfn0_Offset_r16 != nil, ie.ss_PBCH_BlockPower_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ssb_Freq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_Freq_r16", err)
	}
	if err = ie.halfFrameIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode halfFrameIndex_r16", err)
	}
	if err = ie.ssbSubcarrierSpacing_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ssbSubcarrierSpacing_r16", err)
	}
	if ie.ssb_Periodicity_r16 != nil {
		if err = ie.ssb_Periodicity_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_Periodicity_r16", err)
		}
	}
	if ie.sfn0_Offset_r16 != nil {
		if err = ie.sfn0_Offset_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sfn0_Offset_r16", err)
		}
	}
	if err = w.WriteInteger(ie.sfn_SSB_Offset_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger sfn_SSB_Offset_r16", err)
	}
	if ie.ss_PBCH_BlockPower_r16 != nil {
		if err = w.WriteInteger(*ie.ss_PBCH_BlockPower_r16, &uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
			return utils.WrapError("Encode ss_PBCH_BlockPower_r16", err)
		}
	}
	return nil
}

func (ie *SSB_Configuration_r16) Decode(r *uper.UperReader) error {
	var err error
	var ssb_Periodicity_r16Present bool
	if ssb_Periodicity_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sfn0_Offset_r16Present bool
	if sfn0_Offset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ss_PBCH_BlockPower_r16Present bool
	if ss_PBCH_BlockPower_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ssb_Freq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_Freq_r16", err)
	}
	if err = ie.halfFrameIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode halfFrameIndex_r16", err)
	}
	if err = ie.ssbSubcarrierSpacing_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ssbSubcarrierSpacing_r16", err)
	}
	if ssb_Periodicity_r16Present {
		ie.ssb_Periodicity_r16 = new(SSB_Configuration_r16_ssb_Periodicity_r16)
		if err = ie.ssb_Periodicity_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Periodicity_r16", err)
		}
	}
	if sfn0_Offset_r16Present {
		ie.sfn0_Offset_r16 = new(SSB_Configuration_r16_sfn0_Offset_r16)
		if err = ie.sfn0_Offset_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sfn0_Offset_r16", err)
		}
	}
	var tmp_int_sfn_SSB_Offset_r16 int64
	if tmp_int_sfn_SSB_Offset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger sfn_SSB_Offset_r16", err)
	}
	ie.sfn_SSB_Offset_r16 = tmp_int_sfn_SSB_Offset_r16
	if ss_PBCH_BlockPower_r16Present {
		var tmp_int_ss_PBCH_BlockPower_r16 int64
		if tmp_int_ss_PBCH_BlockPower_r16, err = r.ReadInteger(&uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
			return utils.WrapError("Decode ss_PBCH_BlockPower_r16", err)
		}
		ie.ss_PBCH_BlockPower_r16 = &tmp_int_ss_PBCH_BlockPower_r16
	}
	return nil
}
