package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosConfig_r17 struct {
	srs_PosResourceSetToReleaseList_r17 []SRS_PosResourceSetId_r16 `lb:1,ub:maxNrofSRS_PosResourceSets_r16,optional`
	srs_PosResourceSetToAddModList_r17  []SRS_PosResourceSet_r16   `lb:1,ub:maxNrofSRS_PosResourceSets_r16,optional`
	srs_PosResourceToReleaseList_r17    []SRS_PosResourceId_r16    `lb:1,ub:maxNrofSRS_PosResources_r16,optional`
	srs_PosResourceToAddModList_r17     []SRS_PosResource_r16      `lb:1,ub:maxNrofSRS_PosResources_r16,optional`
}

func (ie *SRS_PosConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.srs_PosResourceSetToReleaseList_r17) > 0, len(ie.srs_PosResourceSetToAddModList_r17) > 0, len(ie.srs_PosResourceToReleaseList_r17) > 0, len(ie.srs_PosResourceToAddModList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.srs_PosResourceSetToReleaseList_r17) > 0 {
		tmp_srs_PosResourceSetToReleaseList_r17 := utils.NewSequence[*SRS_PosResourceSetId_r16]([]*SRS_PosResourceSetId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
		for _, i := range ie.srs_PosResourceSetToReleaseList_r17 {
			tmp_srs_PosResourceSetToReleaseList_r17.Value = append(tmp_srs_PosResourceSetToReleaseList_r17.Value, &i)
		}
		if err = tmp_srs_PosResourceSetToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srs_PosResourceSetToReleaseList_r17", err)
		}
	}
	if len(ie.srs_PosResourceSetToAddModList_r17) > 0 {
		tmp_srs_PosResourceSetToAddModList_r17 := utils.NewSequence[*SRS_PosResourceSet_r16]([]*SRS_PosResourceSet_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
		for _, i := range ie.srs_PosResourceSetToAddModList_r17 {
			tmp_srs_PosResourceSetToAddModList_r17.Value = append(tmp_srs_PosResourceSetToAddModList_r17.Value, &i)
		}
		if err = tmp_srs_PosResourceSetToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srs_PosResourceSetToAddModList_r17", err)
		}
	}
	if len(ie.srs_PosResourceToReleaseList_r17) > 0 {
		tmp_srs_PosResourceToReleaseList_r17 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
		for _, i := range ie.srs_PosResourceToReleaseList_r17 {
			tmp_srs_PosResourceToReleaseList_r17.Value = append(tmp_srs_PosResourceToReleaseList_r17.Value, &i)
		}
		if err = tmp_srs_PosResourceToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srs_PosResourceToReleaseList_r17", err)
		}
	}
	if len(ie.srs_PosResourceToAddModList_r17) > 0 {
		tmp_srs_PosResourceToAddModList_r17 := utils.NewSequence[*SRS_PosResource_r16]([]*SRS_PosResource_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
		for _, i := range ie.srs_PosResourceToAddModList_r17 {
			tmp_srs_PosResourceToAddModList_r17.Value = append(tmp_srs_PosResourceToAddModList_r17.Value, &i)
		}
		if err = tmp_srs_PosResourceToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srs_PosResourceToAddModList_r17", err)
		}
	}
	return nil
}

func (ie *SRS_PosConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var srs_PosResourceSetToReleaseList_r17Present bool
	if srs_PosResourceSetToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_PosResourceSetToAddModList_r17Present bool
	if srs_PosResourceSetToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_PosResourceToReleaseList_r17Present bool
	if srs_PosResourceToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_PosResourceToAddModList_r17Present bool
	if srs_PosResourceToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if srs_PosResourceSetToReleaseList_r17Present {
		tmp_srs_PosResourceSetToReleaseList_r17 := utils.NewSequence[*SRS_PosResourceSetId_r16]([]*SRS_PosResourceSetId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
		fn_srs_PosResourceSetToReleaseList_r17 := func() *SRS_PosResourceSetId_r16 {
			return new(SRS_PosResourceSetId_r16)
		}
		if err = tmp_srs_PosResourceSetToReleaseList_r17.Decode(r, fn_srs_PosResourceSetToReleaseList_r17); err != nil {
			return utils.WrapError("Decode srs_PosResourceSetToReleaseList_r17", err)
		}
		ie.srs_PosResourceSetToReleaseList_r17 = []SRS_PosResourceSetId_r16{}
		for _, i := range tmp_srs_PosResourceSetToReleaseList_r17.Value {
			ie.srs_PosResourceSetToReleaseList_r17 = append(ie.srs_PosResourceSetToReleaseList_r17, *i)
		}
	}
	if srs_PosResourceSetToAddModList_r17Present {
		tmp_srs_PosResourceSetToAddModList_r17 := utils.NewSequence[*SRS_PosResourceSet_r16]([]*SRS_PosResourceSet_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
		fn_srs_PosResourceSetToAddModList_r17 := func() *SRS_PosResourceSet_r16 {
			return new(SRS_PosResourceSet_r16)
		}
		if err = tmp_srs_PosResourceSetToAddModList_r17.Decode(r, fn_srs_PosResourceSetToAddModList_r17); err != nil {
			return utils.WrapError("Decode srs_PosResourceSetToAddModList_r17", err)
		}
		ie.srs_PosResourceSetToAddModList_r17 = []SRS_PosResourceSet_r16{}
		for _, i := range tmp_srs_PosResourceSetToAddModList_r17.Value {
			ie.srs_PosResourceSetToAddModList_r17 = append(ie.srs_PosResourceSetToAddModList_r17, *i)
		}
	}
	if srs_PosResourceToReleaseList_r17Present {
		tmp_srs_PosResourceToReleaseList_r17 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
		fn_srs_PosResourceToReleaseList_r17 := func() *SRS_PosResourceId_r16 {
			return new(SRS_PosResourceId_r16)
		}
		if err = tmp_srs_PosResourceToReleaseList_r17.Decode(r, fn_srs_PosResourceToReleaseList_r17); err != nil {
			return utils.WrapError("Decode srs_PosResourceToReleaseList_r17", err)
		}
		ie.srs_PosResourceToReleaseList_r17 = []SRS_PosResourceId_r16{}
		for _, i := range tmp_srs_PosResourceToReleaseList_r17.Value {
			ie.srs_PosResourceToReleaseList_r17 = append(ie.srs_PosResourceToReleaseList_r17, *i)
		}
	}
	if srs_PosResourceToAddModList_r17Present {
		tmp_srs_PosResourceToAddModList_r17 := utils.NewSequence[*SRS_PosResource_r16]([]*SRS_PosResource_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
		fn_srs_PosResourceToAddModList_r17 := func() *SRS_PosResource_r16 {
			return new(SRS_PosResource_r16)
		}
		if err = tmp_srs_PosResourceToAddModList_r17.Decode(r, fn_srs_PosResourceToAddModList_r17); err != nil {
			return utils.WrapError("Decode srs_PosResourceToAddModList_r17", err)
		}
		ie.srs_PosResourceToAddModList_r17 = []SRS_PosResource_r16{}
		for _, i := range tmp_srs_PosResourceToAddModList_r17.Value {
			ie.srs_PosResourceToAddModList_r17 = append(ie.srs_PosResourceToAddModList_r17, *i)
		}
	}
	return nil
}
