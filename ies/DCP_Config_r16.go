package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DCP_Config_r16 struct {
	ps_RNTI_r16                     RNTI_Value                                      `madatory`
	ps_Offset_r16                   int64                                           `lb:1,ub:120,madatory`
	sizeDCI_2_6_r16                 int64                                           `lb:1,ub:maxDCI_2_6_Size_r16,madatory`
	ps_PositionDCI_2_6_r16          int64                                           `lb:0,ub:maxDCI_2_6_Size_1_r16,madatory`
	ps_WakeUp_r16                   *DCP_Config_r16_ps_WakeUp_r16                   `optional`
	ps_TransmitPeriodicL1_RSRP_r16  *DCP_Config_r16_ps_TransmitPeriodicL1_RSRP_r16  `optional`
	ps_TransmitOtherPeriodicCSI_r16 *DCP_Config_r16_ps_TransmitOtherPeriodicCSI_r16 `optional`
}

func (ie *DCP_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ps_WakeUp_r16 != nil, ie.ps_TransmitPeriodicL1_RSRP_r16 != nil, ie.ps_TransmitOtherPeriodicCSI_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ps_RNTI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ps_RNTI_r16", err)
	}
	if err = w.WriteInteger(ie.ps_Offset_r16, &uper.Constraint{Lb: 1, Ub: 120}, false); err != nil {
		return utils.WrapError("WriteInteger ps_Offset_r16", err)
	}
	if err = w.WriteInteger(ie.sizeDCI_2_6_r16, &uper.Constraint{Lb: 1, Ub: maxDCI_2_6_Size_r16}, false); err != nil {
		return utils.WrapError("WriteInteger sizeDCI_2_6_r16", err)
	}
	if err = w.WriteInteger(ie.ps_PositionDCI_2_6_r16, &uper.Constraint{Lb: 0, Ub: maxDCI_2_6_Size_1_r16}, false); err != nil {
		return utils.WrapError("WriteInteger ps_PositionDCI_2_6_r16", err)
	}
	if ie.ps_WakeUp_r16 != nil {
		if err = ie.ps_WakeUp_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ps_WakeUp_r16", err)
		}
	}
	if ie.ps_TransmitPeriodicL1_RSRP_r16 != nil {
		if err = ie.ps_TransmitPeriodicL1_RSRP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ps_TransmitPeriodicL1_RSRP_r16", err)
		}
	}
	if ie.ps_TransmitOtherPeriodicCSI_r16 != nil {
		if err = ie.ps_TransmitOtherPeriodicCSI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ps_TransmitOtherPeriodicCSI_r16", err)
		}
	}
	return nil
}

func (ie *DCP_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var ps_WakeUp_r16Present bool
	if ps_WakeUp_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ps_TransmitPeriodicL1_RSRP_r16Present bool
	if ps_TransmitPeriodicL1_RSRP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ps_TransmitOtherPeriodicCSI_r16Present bool
	if ps_TransmitOtherPeriodicCSI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ps_RNTI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ps_RNTI_r16", err)
	}
	var tmp_int_ps_Offset_r16 int64
	if tmp_int_ps_Offset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 120}, false); err != nil {
		return utils.WrapError("ReadInteger ps_Offset_r16", err)
	}
	ie.ps_Offset_r16 = tmp_int_ps_Offset_r16
	var tmp_int_sizeDCI_2_6_r16 int64
	if tmp_int_sizeDCI_2_6_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxDCI_2_6_Size_r16}, false); err != nil {
		return utils.WrapError("ReadInteger sizeDCI_2_6_r16", err)
	}
	ie.sizeDCI_2_6_r16 = tmp_int_sizeDCI_2_6_r16
	var tmp_int_ps_PositionDCI_2_6_r16 int64
	if tmp_int_ps_PositionDCI_2_6_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxDCI_2_6_Size_1_r16}, false); err != nil {
		return utils.WrapError("ReadInteger ps_PositionDCI_2_6_r16", err)
	}
	ie.ps_PositionDCI_2_6_r16 = tmp_int_ps_PositionDCI_2_6_r16
	if ps_WakeUp_r16Present {
		ie.ps_WakeUp_r16 = new(DCP_Config_r16_ps_WakeUp_r16)
		if err = ie.ps_WakeUp_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ps_WakeUp_r16", err)
		}
	}
	if ps_TransmitPeriodicL1_RSRP_r16Present {
		ie.ps_TransmitPeriodicL1_RSRP_r16 = new(DCP_Config_r16_ps_TransmitPeriodicL1_RSRP_r16)
		if err = ie.ps_TransmitPeriodicL1_RSRP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ps_TransmitPeriodicL1_RSRP_r16", err)
		}
	}
	if ps_TransmitOtherPeriodicCSI_r16Present {
		ie.ps_TransmitOtherPeriodicCSI_r16 = new(DCP_Config_r16_ps_TransmitOtherPeriodicCSI_r16)
		if err = ie.ps_TransmitOtherPeriodicCSI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ps_TransmitOtherPeriodicCSI_r16", err)
		}
	}
	return nil
}
