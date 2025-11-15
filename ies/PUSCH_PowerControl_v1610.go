package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_PowerControl_v1610 struct {
	pathlossReferenceRSToAddModListSizeExt_v1610  []PUSCH_PathlossReferenceRS_r16             `lb:1,ub:maxNrofPUSCH_PathlossReferenceRSsDiff_r16,optional`
	pathlossReferenceRSToReleaseListSizeExt_v1610 []PUSCH_PathlossReferenceRS_Id_v1610        `lb:1,ub:maxNrofPUSCH_PathlossReferenceRSsDiff_r16,optional`
	p0_PUSCH_SetList_r16                          []P0_PUSCH_Set_r16                          `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional`
	olpc_ParameterSet                             *PUSCH_PowerControl_v1610_olpc_ParameterSet `lb:1,ub:2,optional`
	sri_PUSCH_MappingToAddModList2_r17            []SRI_PUSCH_PowerControl                    `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional,ext-1`
	sri_PUSCH_MappingToReleaseList2_r17           []SRI_PUSCH_PowerControlId                  `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional,ext-1`
	p0_PUSCH_SetList2_r17                         []P0_PUSCH_Set_r16                          `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional,ext-1`
	dummy                                         []DummyPathlossReferenceRS_v1710            `lb:1,ub:maxNrofPUSCH_PathlossReferenceRSs_r16,optional,ext-1`
}

func (ie *PUSCH_PowerControl_v1610) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.sri_PUSCH_MappingToAddModList2_r17) > 0 || len(ie.sri_PUSCH_MappingToReleaseList2_r17) > 0 || len(ie.p0_PUSCH_SetList2_r17) > 0 || len(ie.dummy) > 0
	preambleBits := []bool{hasExtensions, len(ie.pathlossReferenceRSToAddModListSizeExt_v1610) > 0, len(ie.pathlossReferenceRSToReleaseListSizeExt_v1610) > 0, len(ie.p0_PUSCH_SetList_r16) > 0, ie.olpc_ParameterSet != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.pathlossReferenceRSToAddModListSizeExt_v1610) > 0 {
		tmp_pathlossReferenceRSToAddModListSizeExt_v1610 := utils.NewSequence[*PUSCH_PathlossReferenceRS_r16]([]*PUSCH_PathlossReferenceRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSsDiff_r16}, false)
		for _, i := range ie.pathlossReferenceRSToAddModListSizeExt_v1610 {
			tmp_pathlossReferenceRSToAddModListSizeExt_v1610.Value = append(tmp_pathlossReferenceRSToAddModListSizeExt_v1610.Value, &i)
		}
		if err = tmp_pathlossReferenceRSToAddModListSizeExt_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode pathlossReferenceRSToAddModListSizeExt_v1610", err)
		}
	}
	if len(ie.pathlossReferenceRSToReleaseListSizeExt_v1610) > 0 {
		tmp_pathlossReferenceRSToReleaseListSizeExt_v1610 := utils.NewSequence[*PUSCH_PathlossReferenceRS_Id_v1610]([]*PUSCH_PathlossReferenceRS_Id_v1610{}, uper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSsDiff_r16}, false)
		for _, i := range ie.pathlossReferenceRSToReleaseListSizeExt_v1610 {
			tmp_pathlossReferenceRSToReleaseListSizeExt_v1610.Value = append(tmp_pathlossReferenceRSToReleaseListSizeExt_v1610.Value, &i)
		}
		if err = tmp_pathlossReferenceRSToReleaseListSizeExt_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode pathlossReferenceRSToReleaseListSizeExt_v1610", err)
		}
	}
	if len(ie.p0_PUSCH_SetList_r16) > 0 {
		tmp_p0_PUSCH_SetList_r16 := utils.NewSequence[*P0_PUSCH_Set_r16]([]*P0_PUSCH_Set_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		for _, i := range ie.p0_PUSCH_SetList_r16 {
			tmp_p0_PUSCH_SetList_r16.Value = append(tmp_p0_PUSCH_SetList_r16.Value, &i)
		}
		if err = tmp_p0_PUSCH_SetList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode p0_PUSCH_SetList_r16", err)
		}
	}
	if ie.olpc_ParameterSet != nil {
		if err = ie.olpc_ParameterSet.Encode(w); err != nil {
			return utils.WrapError("Encode olpc_ParameterSet", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{len(ie.sri_PUSCH_MappingToAddModList2_r17) > 0 || len(ie.sri_PUSCH_MappingToReleaseList2_r17) > 0 || len(ie.p0_PUSCH_SetList2_r17) > 0 || len(ie.dummy) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUSCH_PowerControl_v1610", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.sri_PUSCH_MappingToAddModList2_r17) > 0, len(ie.sri_PUSCH_MappingToReleaseList2_r17) > 0, len(ie.p0_PUSCH_SetList2_r17) > 0, len(ie.dummy) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sri_PUSCH_MappingToAddModList2_r17 optional
			if len(ie.sri_PUSCH_MappingToAddModList2_r17) > 0 {
				tmp_sri_PUSCH_MappingToAddModList2_r17 := utils.NewSequence[*SRI_PUSCH_PowerControl]([]*SRI_PUSCH_PowerControl{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				for _, i := range ie.sri_PUSCH_MappingToAddModList2_r17 {
					tmp_sri_PUSCH_MappingToAddModList2_r17.Value = append(tmp_sri_PUSCH_MappingToAddModList2_r17.Value, &i)
				}
				if err = tmp_sri_PUSCH_MappingToAddModList2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sri_PUSCH_MappingToAddModList2_r17", err)
				}
			}
			// encode sri_PUSCH_MappingToReleaseList2_r17 optional
			if len(ie.sri_PUSCH_MappingToReleaseList2_r17) > 0 {
				tmp_sri_PUSCH_MappingToReleaseList2_r17 := utils.NewSequence[*SRI_PUSCH_PowerControlId]([]*SRI_PUSCH_PowerControlId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				for _, i := range ie.sri_PUSCH_MappingToReleaseList2_r17 {
					tmp_sri_PUSCH_MappingToReleaseList2_r17.Value = append(tmp_sri_PUSCH_MappingToReleaseList2_r17.Value, &i)
				}
				if err = tmp_sri_PUSCH_MappingToReleaseList2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sri_PUSCH_MappingToReleaseList2_r17", err)
				}
			}
			// encode p0_PUSCH_SetList2_r17 optional
			if len(ie.p0_PUSCH_SetList2_r17) > 0 {
				tmp_p0_PUSCH_SetList2_r17 := utils.NewSequence[*P0_PUSCH_Set_r16]([]*P0_PUSCH_Set_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				for _, i := range ie.p0_PUSCH_SetList2_r17 {
					tmp_p0_PUSCH_SetList2_r17.Value = append(tmp_p0_PUSCH_SetList2_r17.Value, &i)
				}
				if err = tmp_p0_PUSCH_SetList2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode p0_PUSCH_SetList2_r17", err)
				}
			}
			// encode dummy optional
			if len(ie.dummy) > 0 {
				tmp_dummy := utils.NewSequence[*DummyPathlossReferenceRS_v1710]([]*DummyPathlossReferenceRS_v1710{}, uper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs_r16}, false)
				for _, i := range ie.dummy {
					tmp_dummy.Value = append(tmp_dummy.Value, &i)
				}
				if err = tmp_dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dummy", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *PUSCH_PowerControl_v1610) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var pathlossReferenceRSToAddModListSizeExt_v1610Present bool
	if pathlossReferenceRSToAddModListSizeExt_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pathlossReferenceRSToReleaseListSizeExt_v1610Present bool
	if pathlossReferenceRSToReleaseListSizeExt_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var p0_PUSCH_SetList_r16Present bool
	if p0_PUSCH_SetList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var olpc_ParameterSetPresent bool
	if olpc_ParameterSetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if pathlossReferenceRSToAddModListSizeExt_v1610Present {
		tmp_pathlossReferenceRSToAddModListSizeExt_v1610 := utils.NewSequence[*PUSCH_PathlossReferenceRS_r16]([]*PUSCH_PathlossReferenceRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSsDiff_r16}, false)
		fn_pathlossReferenceRSToAddModListSizeExt_v1610 := func() *PUSCH_PathlossReferenceRS_r16 {
			return new(PUSCH_PathlossReferenceRS_r16)
		}
		if err = tmp_pathlossReferenceRSToAddModListSizeExt_v1610.Decode(r, fn_pathlossReferenceRSToAddModListSizeExt_v1610); err != nil {
			return utils.WrapError("Decode pathlossReferenceRSToAddModListSizeExt_v1610", err)
		}
		ie.pathlossReferenceRSToAddModListSizeExt_v1610 = []PUSCH_PathlossReferenceRS_r16{}
		for _, i := range tmp_pathlossReferenceRSToAddModListSizeExt_v1610.Value {
			ie.pathlossReferenceRSToAddModListSizeExt_v1610 = append(ie.pathlossReferenceRSToAddModListSizeExt_v1610, *i)
		}
	}
	if pathlossReferenceRSToReleaseListSizeExt_v1610Present {
		tmp_pathlossReferenceRSToReleaseListSizeExt_v1610 := utils.NewSequence[*PUSCH_PathlossReferenceRS_Id_v1610]([]*PUSCH_PathlossReferenceRS_Id_v1610{}, uper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSsDiff_r16}, false)
		fn_pathlossReferenceRSToReleaseListSizeExt_v1610 := func() *PUSCH_PathlossReferenceRS_Id_v1610 {
			return new(PUSCH_PathlossReferenceRS_Id_v1610)
		}
		if err = tmp_pathlossReferenceRSToReleaseListSizeExt_v1610.Decode(r, fn_pathlossReferenceRSToReleaseListSizeExt_v1610); err != nil {
			return utils.WrapError("Decode pathlossReferenceRSToReleaseListSizeExt_v1610", err)
		}
		ie.pathlossReferenceRSToReleaseListSizeExt_v1610 = []PUSCH_PathlossReferenceRS_Id_v1610{}
		for _, i := range tmp_pathlossReferenceRSToReleaseListSizeExt_v1610.Value {
			ie.pathlossReferenceRSToReleaseListSizeExt_v1610 = append(ie.pathlossReferenceRSToReleaseListSizeExt_v1610, *i)
		}
	}
	if p0_PUSCH_SetList_r16Present {
		tmp_p0_PUSCH_SetList_r16 := utils.NewSequence[*P0_PUSCH_Set_r16]([]*P0_PUSCH_Set_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		fn_p0_PUSCH_SetList_r16 := func() *P0_PUSCH_Set_r16 {
			return new(P0_PUSCH_Set_r16)
		}
		if err = tmp_p0_PUSCH_SetList_r16.Decode(r, fn_p0_PUSCH_SetList_r16); err != nil {
			return utils.WrapError("Decode p0_PUSCH_SetList_r16", err)
		}
		ie.p0_PUSCH_SetList_r16 = []P0_PUSCH_Set_r16{}
		for _, i := range tmp_p0_PUSCH_SetList_r16.Value {
			ie.p0_PUSCH_SetList_r16 = append(ie.p0_PUSCH_SetList_r16, *i)
		}
	}
	if olpc_ParameterSetPresent {
		ie.olpc_ParameterSet = new(PUSCH_PowerControl_v1610_olpc_ParameterSet)
		if err = ie.olpc_ParameterSet.Decode(r); err != nil {
			return utils.WrapError("Decode olpc_ParameterSet", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sri_PUSCH_MappingToAddModList2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sri_PUSCH_MappingToReleaseList2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			p0_PUSCH_SetList2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sri_PUSCH_MappingToAddModList2_r17 optional
			if sri_PUSCH_MappingToAddModList2_r17Present {
				tmp_sri_PUSCH_MappingToAddModList2_r17 := utils.NewSequence[*SRI_PUSCH_PowerControl]([]*SRI_PUSCH_PowerControl{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				fn_sri_PUSCH_MappingToAddModList2_r17 := func() *SRI_PUSCH_PowerControl {
					return new(SRI_PUSCH_PowerControl)
				}
				if err = tmp_sri_PUSCH_MappingToAddModList2_r17.Decode(extReader, fn_sri_PUSCH_MappingToAddModList2_r17); err != nil {
					return utils.WrapError("Decode sri_PUSCH_MappingToAddModList2_r17", err)
				}
				ie.sri_PUSCH_MappingToAddModList2_r17 = []SRI_PUSCH_PowerControl{}
				for _, i := range tmp_sri_PUSCH_MappingToAddModList2_r17.Value {
					ie.sri_PUSCH_MappingToAddModList2_r17 = append(ie.sri_PUSCH_MappingToAddModList2_r17, *i)
				}
			}
			// decode sri_PUSCH_MappingToReleaseList2_r17 optional
			if sri_PUSCH_MappingToReleaseList2_r17Present {
				tmp_sri_PUSCH_MappingToReleaseList2_r17 := utils.NewSequence[*SRI_PUSCH_PowerControlId]([]*SRI_PUSCH_PowerControlId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				fn_sri_PUSCH_MappingToReleaseList2_r17 := func() *SRI_PUSCH_PowerControlId {
					return new(SRI_PUSCH_PowerControlId)
				}
				if err = tmp_sri_PUSCH_MappingToReleaseList2_r17.Decode(extReader, fn_sri_PUSCH_MappingToReleaseList2_r17); err != nil {
					return utils.WrapError("Decode sri_PUSCH_MappingToReleaseList2_r17", err)
				}
				ie.sri_PUSCH_MappingToReleaseList2_r17 = []SRI_PUSCH_PowerControlId{}
				for _, i := range tmp_sri_PUSCH_MappingToReleaseList2_r17.Value {
					ie.sri_PUSCH_MappingToReleaseList2_r17 = append(ie.sri_PUSCH_MappingToReleaseList2_r17, *i)
				}
			}
			// decode p0_PUSCH_SetList2_r17 optional
			if p0_PUSCH_SetList2_r17Present {
				tmp_p0_PUSCH_SetList2_r17 := utils.NewSequence[*P0_PUSCH_Set_r16]([]*P0_PUSCH_Set_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				fn_p0_PUSCH_SetList2_r17 := func() *P0_PUSCH_Set_r16 {
					return new(P0_PUSCH_Set_r16)
				}
				if err = tmp_p0_PUSCH_SetList2_r17.Decode(extReader, fn_p0_PUSCH_SetList2_r17); err != nil {
					return utils.WrapError("Decode p0_PUSCH_SetList2_r17", err)
				}
				ie.p0_PUSCH_SetList2_r17 = []P0_PUSCH_Set_r16{}
				for _, i := range tmp_p0_PUSCH_SetList2_r17.Value {
					ie.p0_PUSCH_SetList2_r17 = append(ie.p0_PUSCH_SetList2_r17, *i)
				}
			}
			// decode dummy optional
			if dummyPresent {
				tmp_dummy := utils.NewSequence[*DummyPathlossReferenceRS_v1710]([]*DummyPathlossReferenceRS_v1710{}, uper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs_r16}, false)
				fn_dummy := func() *DummyPathlossReferenceRS_v1710 {
					return new(DummyPathlossReferenceRS_v1710)
				}
				if err = tmp_dummy.Decode(extReader, fn_dummy); err != nil {
					return utils.WrapError("Decode dummy", err)
				}
				ie.dummy = []DummyPathlossReferenceRS_v1710{}
				for _, i := range tmp_dummy.Value {
					ie.dummy = append(ie.dummy, *i)
				}
			}
		}
	}
	return nil
}
