package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceSet struct {
	pucch_ResourceSetId PUCCH_ResourceSetId `madatory`
	resourceList        []PUCCH_ResourceId  `lb:1,ub:maxNrofPUCCH_ResourcesPerSet,madatory`
	maxPayloadSize      *int64              `lb:4,ub:256,optional`
}

func (ie *PUCCH_ResourceSet) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxPayloadSize != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.pucch_ResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_ResourceSetId", err)
	}
	tmp_resourceList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourcesPerSet}, false)
	for _, i := range ie.resourceList {
		tmp_resourceList.Value = append(tmp_resourceList.Value, &i)
	}
	if err = tmp_resourceList.Encode(w); err != nil {
		return utils.WrapError("Encode resourceList", err)
	}
	if ie.maxPayloadSize != nil {
		if err = w.WriteInteger(*ie.maxPayloadSize, &uper.Constraint{Lb: 4, Ub: 256}, false); err != nil {
			return utils.WrapError("Encode maxPayloadSize", err)
		}
	}
	return nil
}

func (ie *PUCCH_ResourceSet) Decode(r *uper.UperReader) error {
	var err error
	var maxPayloadSizePresent bool
	if maxPayloadSizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.pucch_ResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_ResourceSetId", err)
	}
	tmp_resourceList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourcesPerSet}, false)
	fn_resourceList := func() *PUCCH_ResourceId {
		return new(PUCCH_ResourceId)
	}
	if err = tmp_resourceList.Decode(r, fn_resourceList); err != nil {
		return utils.WrapError("Decode resourceList", err)
	}
	ie.resourceList = []PUCCH_ResourceId{}
	for _, i := range tmp_resourceList.Value {
		ie.resourceList = append(ie.resourceList, *i)
	}
	if maxPayloadSizePresent {
		var tmp_int_maxPayloadSize int64
		if tmp_int_maxPayloadSize, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode maxPayloadSize", err)
		}
		ie.maxPayloadSize = &tmp_int_maxPayloadSize
	}
	return nil
}
