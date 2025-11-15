package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResourceSet_r16 struct {
	srs_PosResourceSetId_r16    SRS_PosResourceSetId_r16                            `madatory`
	srs_PosResourceIdList_r16   []SRS_PosResourceId_r16                             `lb:1,ub:maxNrofSRS_ResourcesPerSet,optional`
	resourceType_r16            ResourceType_r16_SRS_PosResourceSet_r16             `madatory`
	alpha_r16                   *Alpha                                              `optional`
	p0_r16                      *int64                                              `lb:-202,ub:24,optional`
	pathlossReferenceRS_Pos_r16 *SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16 `optional`
}

func (ie *SRS_PosResourceSet_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.srs_PosResourceIdList_r16) > 0, ie.alpha_r16 != nil, ie.p0_r16 != nil, ie.pathlossReferenceRS_Pos_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.srs_PosResourceSetId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode srs_PosResourceSetId_r16", err)
	}
	if len(ie.srs_PosResourceIdList_r16) > 0 {
		tmp_srs_PosResourceIdList_r16 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourcesPerSet}, false)
		for _, i := range ie.srs_PosResourceIdList_r16 {
			tmp_srs_PosResourceIdList_r16.Value = append(tmp_srs_PosResourceIdList_r16.Value, &i)
		}
		if err = tmp_srs_PosResourceIdList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode srs_PosResourceIdList_r16", err)
		}
	}
	if err = ie.resourceType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode resourceType_r16", err)
	}
	if ie.alpha_r16 != nil {
		if err = ie.alpha_r16.Encode(w); err != nil {
			return utils.WrapError("Encode alpha_r16", err)
		}
	}
	if ie.p0_r16 != nil {
		if err = w.WriteInteger(*ie.p0_r16, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Encode p0_r16", err)
		}
	}
	if ie.pathlossReferenceRS_Pos_r16 != nil {
		if err = ie.pathlossReferenceRS_Pos_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pathlossReferenceRS_Pos_r16", err)
		}
	}
	return nil
}

func (ie *SRS_PosResourceSet_r16) Decode(r *uper.UperReader) error {
	var err error
	var srs_PosResourceIdList_r16Present bool
	if srs_PosResourceIdList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var alpha_r16Present bool
	if alpha_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var p0_r16Present bool
	if p0_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pathlossReferenceRS_Pos_r16Present bool
	if pathlossReferenceRS_Pos_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.srs_PosResourceSetId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode srs_PosResourceSetId_r16", err)
	}
	if srs_PosResourceIdList_r16Present {
		tmp_srs_PosResourceIdList_r16 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourcesPerSet}, false)
		fn_srs_PosResourceIdList_r16 := func() *SRS_PosResourceId_r16 {
			return new(SRS_PosResourceId_r16)
		}
		if err = tmp_srs_PosResourceIdList_r16.Decode(r, fn_srs_PosResourceIdList_r16); err != nil {
			return utils.WrapError("Decode srs_PosResourceIdList_r16", err)
		}
		ie.srs_PosResourceIdList_r16 = []SRS_PosResourceId_r16{}
		for _, i := range tmp_srs_PosResourceIdList_r16.Value {
			ie.srs_PosResourceIdList_r16 = append(ie.srs_PosResourceIdList_r16, *i)
		}
	}
	if err = ie.resourceType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode resourceType_r16", err)
	}
	if alpha_r16Present {
		ie.alpha_r16 = new(Alpha)
		if err = ie.alpha_r16.Decode(r); err != nil {
			return utils.WrapError("Decode alpha_r16", err)
		}
	}
	if p0_r16Present {
		var tmp_int_p0_r16 int64
		if tmp_int_p0_r16, err = r.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode p0_r16", err)
		}
		ie.p0_r16 = &tmp_int_p0_r16
	}
	if pathlossReferenceRS_Pos_r16Present {
		ie.pathlossReferenceRS_Pos_r16 = new(SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16)
		if err = ie.pathlossReferenceRS_Pos_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pathlossReferenceRS_Pos_r16", err)
		}
	}
	return nil
}
