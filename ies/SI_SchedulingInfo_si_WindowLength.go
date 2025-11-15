package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SI_SchedulingInfo_si_WindowLength_Enum_s5          uper.Enumerated = 0
	SI_SchedulingInfo_si_WindowLength_Enum_s10         uper.Enumerated = 1
	SI_SchedulingInfo_si_WindowLength_Enum_s20         uper.Enumerated = 2
	SI_SchedulingInfo_si_WindowLength_Enum_s40         uper.Enumerated = 3
	SI_SchedulingInfo_si_WindowLength_Enum_s80         uper.Enumerated = 4
	SI_SchedulingInfo_si_WindowLength_Enum_s160        uper.Enumerated = 5
	SI_SchedulingInfo_si_WindowLength_Enum_s320        uper.Enumerated = 6
	SI_SchedulingInfo_si_WindowLength_Enum_s640        uper.Enumerated = 7
	SI_SchedulingInfo_si_WindowLength_Enum_s1280       uper.Enumerated = 8
	SI_SchedulingInfo_si_WindowLength_Enum_s2560_v1710 uper.Enumerated = 9
	SI_SchedulingInfo_si_WindowLength_Enum_s5120_v1710 uper.Enumerated = 10
)

type SI_SchedulingInfo_si_WindowLength struct {
	Value uper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *SI_SchedulingInfo_si_WindowLength) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode SI_SchedulingInfo_si_WindowLength", err)
	}
	return nil
}

func (ie *SI_SchedulingInfo_si_WindowLength) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode SI_SchedulingInfo_si_WindowLength", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
