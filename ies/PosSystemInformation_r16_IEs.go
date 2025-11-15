package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosSystemInformation_r16_IEs struct {
	posSIB_TypeAndInfo_r16   []PosSIB_TypeAndInfo_r16 `lb:1,ub:maxSIB,madatory`
	lateNonCriticalExtension *[]byte                  `optional`
	nonCriticalExtension     interface{}              `optional`
}

func (ie *PosSystemInformation_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_posSIB_TypeAndInfo_r16 := utils.NewSequence[*PosSIB_TypeAndInfo_r16]([]*PosSIB_TypeAndInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxSIB}, false)
	for _, i := range ie.posSIB_TypeAndInfo_r16 {
		tmp_posSIB_TypeAndInfo_r16.Value = append(tmp_posSIB_TypeAndInfo_r16.Value, &i)
	}
	if err = tmp_posSIB_TypeAndInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode posSIB_TypeAndInfo_r16", err)
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *PosSystemInformation_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_posSIB_TypeAndInfo_r16 := utils.NewSequence[*PosSIB_TypeAndInfo_r16]([]*PosSIB_TypeAndInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxSIB}, false)
	fn_posSIB_TypeAndInfo_r16 := func() *PosSIB_TypeAndInfo_r16 {
		return new(PosSIB_TypeAndInfo_r16)
	}
	if err = tmp_posSIB_TypeAndInfo_r16.Decode(r, fn_posSIB_TypeAndInfo_r16); err != nil {
		return utils.WrapError("Decode posSIB_TypeAndInfo_r16", err)
	}
	ie.posSIB_TypeAndInfo_r16 = []PosSIB_TypeAndInfo_r16{}
	for _, i := range tmp_posSIB_TypeAndInfo_r16.Value {
		ie.posSIB_TypeAndInfo_r16 = append(ie.posSIB_TypeAndInfo_r16, *i)
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
