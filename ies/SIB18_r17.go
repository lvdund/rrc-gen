package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB18_r17 struct {
	gin_ElementList_r17      []GIN_Element_r17  `lb:1,ub:maxGIN_r17,optional`
	gins_PerSNPN_List_r17    []GINs_PerSNPN_r17 `lb:1,ub:maxNPN_r16,optional`
	lateNonCriticalExtension *[]byte            `optional`
}

func (ie *SIB18_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.gin_ElementList_r17) > 0, len(ie.gins_PerSNPN_List_r17) > 0, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.gin_ElementList_r17) > 0 {
		tmp_gin_ElementList_r17 := utils.NewSequence[*GIN_Element_r17]([]*GIN_Element_r17{}, uper.Constraint{Lb: 1, Ub: maxGIN_r17}, false)
		for _, i := range ie.gin_ElementList_r17 {
			tmp_gin_ElementList_r17.Value = append(tmp_gin_ElementList_r17.Value, &i)
		}
		if err = tmp_gin_ElementList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode gin_ElementList_r17", err)
		}
	}
	if len(ie.gins_PerSNPN_List_r17) > 0 {
		tmp_gins_PerSNPN_List_r17 := utils.NewSequence[*GINs_PerSNPN_r17]([]*GINs_PerSNPN_r17{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
		for _, i := range ie.gins_PerSNPN_List_r17 {
			tmp_gins_PerSNPN_List_r17.Value = append(tmp_gins_PerSNPN_List_r17.Value, &i)
		}
		if err = tmp_gins_PerSNPN_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode gins_PerSNPN_List_r17", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB18_r17) Decode(r *uper.UperReader) error {
	var err error
	var gin_ElementList_r17Present bool
	if gin_ElementList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gins_PerSNPN_List_r17Present bool
	if gins_PerSNPN_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if gin_ElementList_r17Present {
		tmp_gin_ElementList_r17 := utils.NewSequence[*GIN_Element_r17]([]*GIN_Element_r17{}, uper.Constraint{Lb: 1, Ub: maxGIN_r17}, false)
		fn_gin_ElementList_r17 := func() *GIN_Element_r17 {
			return new(GIN_Element_r17)
		}
		if err = tmp_gin_ElementList_r17.Decode(r, fn_gin_ElementList_r17); err != nil {
			return utils.WrapError("Decode gin_ElementList_r17", err)
		}
		ie.gin_ElementList_r17 = []GIN_Element_r17{}
		for _, i := range tmp_gin_ElementList_r17.Value {
			ie.gin_ElementList_r17 = append(ie.gin_ElementList_r17, *i)
		}
	}
	if gins_PerSNPN_List_r17Present {
		tmp_gins_PerSNPN_List_r17 := utils.NewSequence[*GINs_PerSNPN_r17]([]*GINs_PerSNPN_r17{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
		fn_gins_PerSNPN_List_r17 := func() *GINs_PerSNPN_r17 {
			return new(GINs_PerSNPN_r17)
		}
		if err = tmp_gins_PerSNPN_List_r17.Decode(r, fn_gins_PerSNPN_List_r17); err != nil {
			return utils.WrapError("Decode gins_PerSNPN_List_r17", err)
		}
		ie.gins_PerSNPN_List_r17 = []GINs_PerSNPN_r17{}
		for _, i := range tmp_gins_PerSNPN_List_r17.Value {
			ie.gins_PerSNPN_List_r17 = append(ie.gins_PerSNPN_List_r17, *i)
		}
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
