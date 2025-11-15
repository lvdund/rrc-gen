package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_ResourceSet struct {
	srs_ResourceSetId                SRS_ResourceSetId                                 `madatory`
	srs_ResourceIdList               []SRS_ResourceId                                  `lb:1,ub:maxNrofSRS_ResourcesPerSet,optional`
	resourceType                     []int64                                           `lb:1,ub:maxNrofSRS_TriggerStates_2,e_lb:1,e_ub:maxNrofSRS_TriggerStates_1,optional`
	usage                            SRS_ResourceSet_usage                             `madatory,ext`
	alpha                            *Alpha                                            `optional,ext`
	p0                               *int64                                            `lb:-202,ub:24,optional,ext`
	pathlossReferenceRS              *PathlossReferenceRS_Config                       `optional,ext`
	srs_PowerControlAdjustmentStates *SRS_ResourceSet_srs_PowerControlAdjustmentStates `optional,ext`
	pathlossReferenceRSList_r16      *PathlossReferenceRSList_r16                      `optional,ext-2,setuprelease`
	usagePDC_r17                     *SRS_ResourceSet_usagePDC_r17                     `optional,ext-3`
	availableSlotOffsetList_r17      []AvailableSlotOffset_r17                         `lb:1,ub:4,optional,ext-3`
	followUnifiedTCI_StateSRS_r17    *SRS_ResourceSet_followUnifiedTCI_StateSRS_r17    `optional,ext-3`
}

func (ie *SRS_ResourceSet) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.pathlossReferenceRSList_r16 != nil || ie.usagePDC_r17 != nil || len(ie.availableSlotOffsetList_r17) > 0 || ie.followUnifiedTCI_StateSRS_r17 != nil
	preambleBits := []bool{hasExtensions, len(ie.srs_ResourceIdList) > 0, len(ie.resourceType) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.srs_ResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode srs_ResourceSetId", err)
	}
	if len(ie.srs_ResourceIdList) > 0 {
		tmp_srs_ResourceIdList := utils.NewSequence[*SRS_ResourceId]([]*SRS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourcesPerSet}, false)
		for _, i := range ie.srs_ResourceIdList {
			tmp_srs_ResourceIdList.Value = append(tmp_srs_ResourceIdList.Value, &i)
		}
		if err = tmp_srs_ResourceIdList.Encode(w); err != nil {
			return utils.WrapError("Encode srs_ResourceIdList", err)
		}
	}
	if len(ie.resourceType) > 0 {
		tmp_resourceType := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_2}, false)
		for _, i := range ie.resourceType {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_1}, false)
			tmp_resourceType.Value = append(tmp_resourceType.Value, &tmp_ie)
		}
		if err = tmp_resourceType.Encode(w); err != nil {
			return utils.WrapError("Encode resourceType", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.pathlossReferenceRSList_r16 != nil, ie.usagePDC_r17 != nil || len(ie.availableSlotOffsetList_r17) > 0 || ie.followUnifiedTCI_StateSRS_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SRS_ResourceSet", err)
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.pathlossReferenceRSList_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pathlossReferenceRSList_r16 optional
			if ie.pathlossReferenceRSList_r16 != nil {
				tmp_pathlossReferenceRSList_r16 := utils.SetupRelease[*PathlossReferenceRSList_r16]{
					Setup: ie.pathlossReferenceRSList_r16,
				}
				if err = tmp_pathlossReferenceRSList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pathlossReferenceRSList_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.usagePDC_r17 != nil, len(ie.availableSlotOffsetList_r17) > 0, ie.followUnifiedTCI_StateSRS_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode usagePDC_r17 optional
			if ie.usagePDC_r17 != nil {
				if err = ie.usagePDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode usagePDC_r17", err)
				}
			}
			// encode availableSlotOffsetList_r17 optional
			if len(ie.availableSlotOffsetList_r17) > 0 {
				tmp_availableSlotOffsetList_r17 := utils.NewSequence[*AvailableSlotOffset_r17]([]*AvailableSlotOffset_r17{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				for _, i := range ie.availableSlotOffsetList_r17 {
					tmp_availableSlotOffsetList_r17.Value = append(tmp_availableSlotOffsetList_r17.Value, &i)
				}
				if err = tmp_availableSlotOffsetList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode availableSlotOffsetList_r17", err)
				}
			}
			// encode followUnifiedTCI_StateSRS_r17 optional
			if ie.followUnifiedTCI_StateSRS_r17 != nil {
				if err = ie.followUnifiedTCI_StateSRS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode followUnifiedTCI_StateSRS_r17", err)
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

func (ie *SRS_ResourceSet) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_ResourceIdListPresent bool
	if srs_ResourceIdListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var resourceTypePresent bool
	if resourceTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.srs_ResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode srs_ResourceSetId", err)
	}
	if srs_ResourceIdListPresent {
		tmp_srs_ResourceIdList := utils.NewSequence[*SRS_ResourceId]([]*SRS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourcesPerSet}, false)
		fn_srs_ResourceIdList := func() *SRS_ResourceId {
			return new(SRS_ResourceId)
		}
		if err = tmp_srs_ResourceIdList.Decode(r, fn_srs_ResourceIdList); err != nil {
			return utils.WrapError("Decode srs_ResourceIdList", err)
		}
		ie.srs_ResourceIdList = []SRS_ResourceId{}
		for _, i := range tmp_srs_ResourceIdList.Value {
			ie.srs_ResourceIdList = append(ie.srs_ResourceIdList, *i)
		}
	}
	if resourceTypePresent {
		tmp_resourceType := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_2}, false)
		fn_resourceType := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_1}, false)
			return &ie
		}
		if err = tmp_resourceType.Decode(r, fn_resourceType); err != nil {
			return utils.WrapError("Decode resourceType", err)
		}
		ie.resourceType = []int64{}
		for _, i := range tmp_resourceType.Value {
			ie.resourceType = append(ie.resourceType, int64(i.Value))
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			pathlossReferenceRSList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pathlossReferenceRSList_r16 optional
			if pathlossReferenceRSList_r16Present {
				tmp_pathlossReferenceRSList_r16 := utils.SetupRelease[*PathlossReferenceRSList_r16]{}
				if err = tmp_pathlossReferenceRSList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pathlossReferenceRSList_r16", err)
				}
				ie.pathlossReferenceRSList_r16 = tmp_pathlossReferenceRSList_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			usagePDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			availableSlotOffsetList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			followUnifiedTCI_StateSRS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode usagePDC_r17 optional
			if usagePDC_r17Present {
				ie.usagePDC_r17 = new(SRS_ResourceSet_usagePDC_r17)
				if err = ie.usagePDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode usagePDC_r17", err)
				}
			}
			// decode availableSlotOffsetList_r17 optional
			if availableSlotOffsetList_r17Present {
				tmp_availableSlotOffsetList_r17 := utils.NewSequence[*AvailableSlotOffset_r17]([]*AvailableSlotOffset_r17{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				fn_availableSlotOffsetList_r17 := func() *AvailableSlotOffset_r17 {
					return new(AvailableSlotOffset_r17)
				}
				if err = tmp_availableSlotOffsetList_r17.Decode(extReader, fn_availableSlotOffsetList_r17); err != nil {
					return utils.WrapError("Decode availableSlotOffsetList_r17", err)
				}
				ie.availableSlotOffsetList_r17 = []AvailableSlotOffset_r17{}
				for _, i := range tmp_availableSlotOffsetList_r17.Value {
					ie.availableSlotOffsetList_r17 = append(ie.availableSlotOffsetList_r17, *i)
				}
			}
			// decode followUnifiedTCI_StateSRS_r17 optional
			if followUnifiedTCI_StateSRS_r17Present {
				ie.followUnifiedTCI_StateSRS_r17 = new(SRS_ResourceSet_followUnifiedTCI_StateSRS_r17)
				if err = ie.followUnifiedTCI_StateSRS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode followUnifiedTCI_StateSRS_r17", err)
				}
			}
		}
	}
	return nil
}
