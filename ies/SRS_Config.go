package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Config struct {
	srs_ResourceSetToReleaseList            []SRS_ResourceSetId          `lb:1,ub:maxNrofSRS_ResourceSets,optional`
	srs_ResourceSetToAddModList             []SRS_ResourceSet            `lb:1,ub:maxNrofSRS_ResourceSets,optional`
	srs_ResourceToReleaseList               []SRS_ResourceId             `lb:1,ub:maxNrofSRS_Resources,optional`
	srs_ResourceToAddModList                []SRS_Resource               `lb:1,ub:maxNrofSRS_Resources,optional`
	tpc_Accumulation                        *SRS_Config_tpc_Accumulation `optional`
	srs_RequestDCI_1_2_r16                  *int64                       `lb:1,ub:2,optional,ext-1`
	srs_RequestDCI_0_2_r16                  *int64                       `lb:1,ub:2,optional,ext-1`
	srs_ResourceSetToAddModListDCI_0_2_r16  []SRS_ResourceSet            `lb:1,ub:maxNrofSRS_ResourceSets,optional,ext-1`
	srs_ResourceSetToReleaseListDCI_0_2_r16 []SRS_ResourceSetId          `lb:1,ub:maxNrofSRS_ResourceSets,optional,ext-1`
	srs_PosResourceSetToReleaseList_r16     []SRS_PosResourceSetId_r16   `lb:1,ub:maxNrofSRS_PosResourceSets_r16,optional,ext-1`
	srs_PosResourceSetToAddModList_r16      []SRS_PosResourceSet_r16     `lb:1,ub:maxNrofSRS_PosResourceSets_r16,optional,ext-1`
	srs_PosResourceToReleaseList_r16        []SRS_PosResourceId_r16      `lb:1,ub:maxNrofSRS_PosResources_r16,optional,ext-1`
	srs_PosResourceToAddModList_r16         []SRS_PosResource_r16        `lb:1,ub:maxNrofSRS_PosResources_r16,optional,ext-1`
}

func (ie *SRS_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.srs_RequestDCI_1_2_r16 != nil || ie.srs_RequestDCI_0_2_r16 != nil || len(ie.srs_ResourceSetToAddModListDCI_0_2_r16) > 0 || len(ie.srs_ResourceSetToReleaseListDCI_0_2_r16) > 0 || len(ie.srs_PosResourceSetToReleaseList_r16) > 0 || len(ie.srs_PosResourceSetToAddModList_r16) > 0 || len(ie.srs_PosResourceToReleaseList_r16) > 0 || len(ie.srs_PosResourceToAddModList_r16) > 0
	preambleBits := []bool{hasExtensions, len(ie.srs_ResourceSetToReleaseList) > 0, len(ie.srs_ResourceSetToAddModList) > 0, len(ie.srs_ResourceToReleaseList) > 0, len(ie.srs_ResourceToAddModList) > 0, ie.tpc_Accumulation != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.srs_ResourceSetToReleaseList) > 0 {
		tmp_srs_ResourceSetToReleaseList := utils.NewSequence[*SRS_ResourceSetId]([]*SRS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
		for _, i := range ie.srs_ResourceSetToReleaseList {
			tmp_srs_ResourceSetToReleaseList.Value = append(tmp_srs_ResourceSetToReleaseList.Value, &i)
		}
		if err = tmp_srs_ResourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode srs_ResourceSetToReleaseList", err)
		}
	}
	if len(ie.srs_ResourceSetToAddModList) > 0 {
		tmp_srs_ResourceSetToAddModList := utils.NewSequence[*SRS_ResourceSet]([]*SRS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
		for _, i := range ie.srs_ResourceSetToAddModList {
			tmp_srs_ResourceSetToAddModList.Value = append(tmp_srs_ResourceSetToAddModList.Value, &i)
		}
		if err = tmp_srs_ResourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode srs_ResourceSetToAddModList", err)
		}
	}
	if len(ie.srs_ResourceToReleaseList) > 0 {
		tmp_srs_ResourceToReleaseList := utils.NewSequence[*SRS_ResourceId]([]*SRS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_Resources}, false)
		for _, i := range ie.srs_ResourceToReleaseList {
			tmp_srs_ResourceToReleaseList.Value = append(tmp_srs_ResourceToReleaseList.Value, &i)
		}
		if err = tmp_srs_ResourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode srs_ResourceToReleaseList", err)
		}
	}
	if len(ie.srs_ResourceToAddModList) > 0 {
		tmp_srs_ResourceToAddModList := utils.NewSequence[*SRS_Resource]([]*SRS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_Resources}, false)
		for _, i := range ie.srs_ResourceToAddModList {
			tmp_srs_ResourceToAddModList.Value = append(tmp_srs_ResourceToAddModList.Value, &i)
		}
		if err = tmp_srs_ResourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode srs_ResourceToAddModList", err)
		}
	}
	if ie.tpc_Accumulation != nil {
		if err = ie.tpc_Accumulation.Encode(w); err != nil {
			return utils.WrapError("Encode tpc_Accumulation", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.srs_RequestDCI_1_2_r16 != nil || ie.srs_RequestDCI_0_2_r16 != nil || len(ie.srs_ResourceSetToAddModListDCI_0_2_r16) > 0 || len(ie.srs_ResourceSetToReleaseListDCI_0_2_r16) > 0 || len(ie.srs_PosResourceSetToReleaseList_r16) > 0 || len(ie.srs_PosResourceSetToAddModList_r16) > 0 || len(ie.srs_PosResourceToReleaseList_r16) > 0 || len(ie.srs_PosResourceToAddModList_r16) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SRS_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.srs_RequestDCI_1_2_r16 != nil, ie.srs_RequestDCI_0_2_r16 != nil, len(ie.srs_ResourceSetToAddModListDCI_0_2_r16) > 0, len(ie.srs_ResourceSetToReleaseListDCI_0_2_r16) > 0, len(ie.srs_PosResourceSetToReleaseList_r16) > 0, len(ie.srs_PosResourceSetToAddModList_r16) > 0, len(ie.srs_PosResourceToReleaseList_r16) > 0, len(ie.srs_PosResourceToAddModList_r16) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode srs_RequestDCI_1_2_r16 optional
			if ie.srs_RequestDCI_1_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.srs_RequestDCI_1_2_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode srs_RequestDCI_1_2_r16", err)
				}
			}
			// encode srs_RequestDCI_0_2_r16 optional
			if ie.srs_RequestDCI_0_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.srs_RequestDCI_0_2_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode srs_RequestDCI_0_2_r16", err)
				}
			}
			// encode srs_ResourceSetToAddModListDCI_0_2_r16 optional
			if len(ie.srs_ResourceSetToAddModListDCI_0_2_r16) > 0 {
				tmp_srs_ResourceSetToAddModListDCI_0_2_r16 := utils.NewSequence[*SRS_ResourceSet]([]*SRS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
				for _, i := range ie.srs_ResourceSetToAddModListDCI_0_2_r16 {
					tmp_srs_ResourceSetToAddModListDCI_0_2_r16.Value = append(tmp_srs_ResourceSetToAddModListDCI_0_2_r16.Value, &i)
				}
				if err = tmp_srs_ResourceSetToAddModListDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_ResourceSetToAddModListDCI_0_2_r16", err)
				}
			}
			// encode srs_ResourceSetToReleaseListDCI_0_2_r16 optional
			if len(ie.srs_ResourceSetToReleaseListDCI_0_2_r16) > 0 {
				tmp_srs_ResourceSetToReleaseListDCI_0_2_r16 := utils.NewSequence[*SRS_ResourceSetId]([]*SRS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
				for _, i := range ie.srs_ResourceSetToReleaseListDCI_0_2_r16 {
					tmp_srs_ResourceSetToReleaseListDCI_0_2_r16.Value = append(tmp_srs_ResourceSetToReleaseListDCI_0_2_r16.Value, &i)
				}
				if err = tmp_srs_ResourceSetToReleaseListDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_ResourceSetToReleaseListDCI_0_2_r16", err)
				}
			}
			// encode srs_PosResourceSetToReleaseList_r16 optional
			if len(ie.srs_PosResourceSetToReleaseList_r16) > 0 {
				tmp_srs_PosResourceSetToReleaseList_r16 := utils.NewSequence[*SRS_PosResourceSetId_r16]([]*SRS_PosResourceSetId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
				for _, i := range ie.srs_PosResourceSetToReleaseList_r16 {
					tmp_srs_PosResourceSetToReleaseList_r16.Value = append(tmp_srs_PosResourceSetToReleaseList_r16.Value, &i)
				}
				if err = tmp_srs_PosResourceSetToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_PosResourceSetToReleaseList_r16", err)
				}
			}
			// encode srs_PosResourceSetToAddModList_r16 optional
			if len(ie.srs_PosResourceSetToAddModList_r16) > 0 {
				tmp_srs_PosResourceSetToAddModList_r16 := utils.NewSequence[*SRS_PosResourceSet_r16]([]*SRS_PosResourceSet_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
				for _, i := range ie.srs_PosResourceSetToAddModList_r16 {
					tmp_srs_PosResourceSetToAddModList_r16.Value = append(tmp_srs_PosResourceSetToAddModList_r16.Value, &i)
				}
				if err = tmp_srs_PosResourceSetToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_PosResourceSetToAddModList_r16", err)
				}
			}
			// encode srs_PosResourceToReleaseList_r16 optional
			if len(ie.srs_PosResourceToReleaseList_r16) > 0 {
				tmp_srs_PosResourceToReleaseList_r16 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
				for _, i := range ie.srs_PosResourceToReleaseList_r16 {
					tmp_srs_PosResourceToReleaseList_r16.Value = append(tmp_srs_PosResourceToReleaseList_r16.Value, &i)
				}
				if err = tmp_srs_PosResourceToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_PosResourceToReleaseList_r16", err)
				}
			}
			// encode srs_PosResourceToAddModList_r16 optional
			if len(ie.srs_PosResourceToAddModList_r16) > 0 {
				tmp_srs_PosResourceToAddModList_r16 := utils.NewSequence[*SRS_PosResource_r16]([]*SRS_PosResource_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
				for _, i := range ie.srs_PosResourceToAddModList_r16 {
					tmp_srs_PosResourceToAddModList_r16.Value = append(tmp_srs_PosResourceToAddModList_r16.Value, &i)
				}
				if err = tmp_srs_PosResourceToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_PosResourceToAddModList_r16", err)
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

func (ie *SRS_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_ResourceSetToReleaseListPresent bool
	if srs_ResourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_ResourceSetToAddModListPresent bool
	if srs_ResourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_ResourceToReleaseListPresent bool
	if srs_ResourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_ResourceToAddModListPresent bool
	if srs_ResourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_AccumulationPresent bool
	if tpc_AccumulationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if srs_ResourceSetToReleaseListPresent {
		tmp_srs_ResourceSetToReleaseList := utils.NewSequence[*SRS_ResourceSetId]([]*SRS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
		fn_srs_ResourceSetToReleaseList := func() *SRS_ResourceSetId {
			return new(SRS_ResourceSetId)
		}
		if err = tmp_srs_ResourceSetToReleaseList.Decode(r, fn_srs_ResourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode srs_ResourceSetToReleaseList", err)
		}
		ie.srs_ResourceSetToReleaseList = []SRS_ResourceSetId{}
		for _, i := range tmp_srs_ResourceSetToReleaseList.Value {
			ie.srs_ResourceSetToReleaseList = append(ie.srs_ResourceSetToReleaseList, *i)
		}
	}
	if srs_ResourceSetToAddModListPresent {
		tmp_srs_ResourceSetToAddModList := utils.NewSequence[*SRS_ResourceSet]([]*SRS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
		fn_srs_ResourceSetToAddModList := func() *SRS_ResourceSet {
			return new(SRS_ResourceSet)
		}
		if err = tmp_srs_ResourceSetToAddModList.Decode(r, fn_srs_ResourceSetToAddModList); err != nil {
			return utils.WrapError("Decode srs_ResourceSetToAddModList", err)
		}
		ie.srs_ResourceSetToAddModList = []SRS_ResourceSet{}
		for _, i := range tmp_srs_ResourceSetToAddModList.Value {
			ie.srs_ResourceSetToAddModList = append(ie.srs_ResourceSetToAddModList, *i)
		}
	}
	if srs_ResourceToReleaseListPresent {
		tmp_srs_ResourceToReleaseList := utils.NewSequence[*SRS_ResourceId]([]*SRS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_Resources}, false)
		fn_srs_ResourceToReleaseList := func() *SRS_ResourceId {
			return new(SRS_ResourceId)
		}
		if err = tmp_srs_ResourceToReleaseList.Decode(r, fn_srs_ResourceToReleaseList); err != nil {
			return utils.WrapError("Decode srs_ResourceToReleaseList", err)
		}
		ie.srs_ResourceToReleaseList = []SRS_ResourceId{}
		for _, i := range tmp_srs_ResourceToReleaseList.Value {
			ie.srs_ResourceToReleaseList = append(ie.srs_ResourceToReleaseList, *i)
		}
	}
	if srs_ResourceToAddModListPresent {
		tmp_srs_ResourceToAddModList := utils.NewSequence[*SRS_Resource]([]*SRS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_Resources}, false)
		fn_srs_ResourceToAddModList := func() *SRS_Resource {
			return new(SRS_Resource)
		}
		if err = tmp_srs_ResourceToAddModList.Decode(r, fn_srs_ResourceToAddModList); err != nil {
			return utils.WrapError("Decode srs_ResourceToAddModList", err)
		}
		ie.srs_ResourceToAddModList = []SRS_Resource{}
		for _, i := range tmp_srs_ResourceToAddModList.Value {
			ie.srs_ResourceToAddModList = append(ie.srs_ResourceToAddModList, *i)
		}
	}
	if tpc_AccumulationPresent {
		ie.tpc_Accumulation = new(SRS_Config_tpc_Accumulation)
		if err = ie.tpc_Accumulation.Decode(r); err != nil {
			return utils.WrapError("Decode tpc_Accumulation", err)
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

			srs_RequestDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_RequestDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_ResourceSetToAddModListDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_ResourceSetToReleaseListDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_PosResourceSetToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_PosResourceSetToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_PosResourceToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_PosResourceToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode srs_RequestDCI_1_2_r16 optional
			if srs_RequestDCI_1_2_r16Present {
				var tmp_int_srs_RequestDCI_1_2_r16 int64
				if tmp_int_srs_RequestDCI_1_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode srs_RequestDCI_1_2_r16", err)
				}
				ie.srs_RequestDCI_1_2_r16 = &tmp_int_srs_RequestDCI_1_2_r16
			}
			// decode srs_RequestDCI_0_2_r16 optional
			if srs_RequestDCI_0_2_r16Present {
				var tmp_int_srs_RequestDCI_0_2_r16 int64
				if tmp_int_srs_RequestDCI_0_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode srs_RequestDCI_0_2_r16", err)
				}
				ie.srs_RequestDCI_0_2_r16 = &tmp_int_srs_RequestDCI_0_2_r16
			}
			// decode srs_ResourceSetToAddModListDCI_0_2_r16 optional
			if srs_ResourceSetToAddModListDCI_0_2_r16Present {
				tmp_srs_ResourceSetToAddModListDCI_0_2_r16 := utils.NewSequence[*SRS_ResourceSet]([]*SRS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
				fn_srs_ResourceSetToAddModListDCI_0_2_r16 := func() *SRS_ResourceSet {
					return new(SRS_ResourceSet)
				}
				if err = tmp_srs_ResourceSetToAddModListDCI_0_2_r16.Decode(extReader, fn_srs_ResourceSetToAddModListDCI_0_2_r16); err != nil {
					return utils.WrapError("Decode srs_ResourceSetToAddModListDCI_0_2_r16", err)
				}
				ie.srs_ResourceSetToAddModListDCI_0_2_r16 = []SRS_ResourceSet{}
				for _, i := range tmp_srs_ResourceSetToAddModListDCI_0_2_r16.Value {
					ie.srs_ResourceSetToAddModListDCI_0_2_r16 = append(ie.srs_ResourceSetToAddModListDCI_0_2_r16, *i)
				}
			}
			// decode srs_ResourceSetToReleaseListDCI_0_2_r16 optional
			if srs_ResourceSetToReleaseListDCI_0_2_r16Present {
				tmp_srs_ResourceSetToReleaseListDCI_0_2_r16 := utils.NewSequence[*SRS_ResourceSetId]([]*SRS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
				fn_srs_ResourceSetToReleaseListDCI_0_2_r16 := func() *SRS_ResourceSetId {
					return new(SRS_ResourceSetId)
				}
				if err = tmp_srs_ResourceSetToReleaseListDCI_0_2_r16.Decode(extReader, fn_srs_ResourceSetToReleaseListDCI_0_2_r16); err != nil {
					return utils.WrapError("Decode srs_ResourceSetToReleaseListDCI_0_2_r16", err)
				}
				ie.srs_ResourceSetToReleaseListDCI_0_2_r16 = []SRS_ResourceSetId{}
				for _, i := range tmp_srs_ResourceSetToReleaseListDCI_0_2_r16.Value {
					ie.srs_ResourceSetToReleaseListDCI_0_2_r16 = append(ie.srs_ResourceSetToReleaseListDCI_0_2_r16, *i)
				}
			}
			// decode srs_PosResourceSetToReleaseList_r16 optional
			if srs_PosResourceSetToReleaseList_r16Present {
				tmp_srs_PosResourceSetToReleaseList_r16 := utils.NewSequence[*SRS_PosResourceSetId_r16]([]*SRS_PosResourceSetId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
				fn_srs_PosResourceSetToReleaseList_r16 := func() *SRS_PosResourceSetId_r16 {
					return new(SRS_PosResourceSetId_r16)
				}
				if err = tmp_srs_PosResourceSetToReleaseList_r16.Decode(extReader, fn_srs_PosResourceSetToReleaseList_r16); err != nil {
					return utils.WrapError("Decode srs_PosResourceSetToReleaseList_r16", err)
				}
				ie.srs_PosResourceSetToReleaseList_r16 = []SRS_PosResourceSetId_r16{}
				for _, i := range tmp_srs_PosResourceSetToReleaseList_r16.Value {
					ie.srs_PosResourceSetToReleaseList_r16 = append(ie.srs_PosResourceSetToReleaseList_r16, *i)
				}
			}
			// decode srs_PosResourceSetToAddModList_r16 optional
			if srs_PosResourceSetToAddModList_r16Present {
				tmp_srs_PosResourceSetToAddModList_r16 := utils.NewSequence[*SRS_PosResourceSet_r16]([]*SRS_PosResourceSet_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
				fn_srs_PosResourceSetToAddModList_r16 := func() *SRS_PosResourceSet_r16 {
					return new(SRS_PosResourceSet_r16)
				}
				if err = tmp_srs_PosResourceSetToAddModList_r16.Decode(extReader, fn_srs_PosResourceSetToAddModList_r16); err != nil {
					return utils.WrapError("Decode srs_PosResourceSetToAddModList_r16", err)
				}
				ie.srs_PosResourceSetToAddModList_r16 = []SRS_PosResourceSet_r16{}
				for _, i := range tmp_srs_PosResourceSetToAddModList_r16.Value {
					ie.srs_PosResourceSetToAddModList_r16 = append(ie.srs_PosResourceSetToAddModList_r16, *i)
				}
			}
			// decode srs_PosResourceToReleaseList_r16 optional
			if srs_PosResourceToReleaseList_r16Present {
				tmp_srs_PosResourceToReleaseList_r16 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
				fn_srs_PosResourceToReleaseList_r16 := func() *SRS_PosResourceId_r16 {
					return new(SRS_PosResourceId_r16)
				}
				if err = tmp_srs_PosResourceToReleaseList_r16.Decode(extReader, fn_srs_PosResourceToReleaseList_r16); err != nil {
					return utils.WrapError("Decode srs_PosResourceToReleaseList_r16", err)
				}
				ie.srs_PosResourceToReleaseList_r16 = []SRS_PosResourceId_r16{}
				for _, i := range tmp_srs_PosResourceToReleaseList_r16.Value {
					ie.srs_PosResourceToReleaseList_r16 = append(ie.srs_PosResourceToReleaseList_r16, *i)
				}
			}
			// decode srs_PosResourceToAddModList_r16 optional
			if srs_PosResourceToAddModList_r16Present {
				tmp_srs_PosResourceToAddModList_r16 := utils.NewSequence[*SRS_PosResource_r16]([]*SRS_PosResource_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
				fn_srs_PosResourceToAddModList_r16 := func() *SRS_PosResource_r16 {
					return new(SRS_PosResource_r16)
				}
				if err = tmp_srs_PosResourceToAddModList_r16.Decode(extReader, fn_srs_PosResourceToAddModList_r16); err != nil {
					return utils.WrapError("Decode srs_PosResourceToAddModList_r16", err)
				}
				ie.srs_PosResourceToAddModList_r16 = []SRS_PosResource_r16{}
				for _, i := range tmp_srs_PosResourceToAddModList_r16.Value {
					ie.srs_PosResourceToAddModList_r16 = append(ie.srs_PosResourceToAddModList_r16, *i)
				}
			}
		}
	}
	return nil
}
