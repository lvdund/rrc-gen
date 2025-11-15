package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s5   uper.Enumerated = 0
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s10  uper.Enumerated = 1
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s15  uper.Enumerated = 2
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s20  uper.Enumerated = 3
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s25  uper.Enumerated = 4
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s30  uper.Enumerated = 5
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s35  uper.Enumerated = 6
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s40  uper.Enumerated = 7
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s45  uper.Enumerated = 8
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s50  uper.Enumerated = 9
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s55  uper.Enumerated = 10
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s60  uper.Enumerated = 11
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s120 uper.Enumerated = 12
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s180 uper.Enumerated = 13
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s240 uper.Enumerated = 14
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s900 uper.Enumerated = 15
)

type NTN_Config_r17_ntn_UlSyncValidityDuration_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *NTN_Config_r17_ntn_UlSyncValidityDuration_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode NTN_Config_r17_ntn_UlSyncValidityDuration_r17", err)
	}
	return nil
}

func (ie *NTN_Config_r17_ntn_UlSyncValidityDuration_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode NTN_Config_r17_ntn_UlSyncValidityDuration_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
