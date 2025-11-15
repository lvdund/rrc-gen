package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ControlResourceSet struct {
	controlResourceSetId          ControlResourceSetId                           `madatory`
	frequencyDomainResources      uper.BitString                                 `lb:45,ub:45,madatory`
	duration                      int64                                          `lb:1,ub:maxCoReSetDuration,madatory`
	cce_REG_MappingType           *ControlResourceSet_cce_REG_MappingType        `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional`
	precoderGranularity           ControlResourceSet_precoderGranularity         `madatory`
	tci_StatesPDCCH_ToAddList     []TCI_StateId                                  `lb:1,ub:maxNrofTCI_StatesPDCCH,optional`
	tci_StatesPDCCH_ToReleaseList []TCI_StateId                                  `lb:1,ub:maxNrofTCI_StatesPDCCH,optional`
	tci_PresentInDCI              *ControlResourceSet_tci_PresentInDCI           `optional`
	pdcch_DMRS_ScramblingID       *int64                                         `lb:0,ub:65535,optional`
	rb_Offset_r16                 *int64                                         `lb:0,ub:5,optional,ext-1`
	tci_PresentDCI_1_2_r16        *int64                                         `lb:1,ub:3,optional,ext-1`
	coresetPoolIndex_r16          *int64                                         `lb:0,ub:1,optional,ext-1`
	controlResourceSetId_v1610    *ControlResourceSetId_v1610                    `optional,ext-1`
	followUnifiedTCI_State_r17    *ControlResourceSet_followUnifiedTCI_State_r17 `optional,ext-2`
}

func (ie *ControlResourceSet) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.rb_Offset_r16 != nil || ie.tci_PresentDCI_1_2_r16 != nil || ie.coresetPoolIndex_r16 != nil || ie.controlResourceSetId_v1610 != nil || ie.followUnifiedTCI_State_r17 != nil
	preambleBits := []bool{hasExtensions, ie.cce_REG_MappingType != nil, len(ie.tci_StatesPDCCH_ToAddList) > 0, len(ie.tci_StatesPDCCH_ToReleaseList) > 0, ie.tci_PresentInDCI != nil, ie.pdcch_DMRS_ScramblingID != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.controlResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode controlResourceSetId", err)
	}
	if err = w.WriteBitString(ie.frequencyDomainResources.Bytes, uint(ie.frequencyDomainResources.NumBits), &uper.Constraint{Lb: 45, Ub: 45}, false); err != nil {
		return utils.WrapError("WriteBitString frequencyDomainResources", err)
	}
	if err = w.WriteInteger(ie.duration, &uper.Constraint{Lb: 1, Ub: maxCoReSetDuration}, false); err != nil {
		return utils.WrapError("WriteInteger duration", err)
	}
	if ie.cce_REG_MappingType != nil {
		if err = ie.cce_REG_MappingType.Encode(w); err != nil {
			return utils.WrapError("Encode cce_REG_MappingType", err)
		}
	}
	if err = ie.precoderGranularity.Encode(w); err != nil {
		return utils.WrapError("Encode precoderGranularity", err)
	}
	if len(ie.tci_StatesPDCCH_ToAddList) > 0 {
		tmp_tci_StatesPDCCH_ToAddList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_StatesPDCCH}, false)
		for _, i := range ie.tci_StatesPDCCH_ToAddList {
			tmp_tci_StatesPDCCH_ToAddList.Value = append(tmp_tci_StatesPDCCH_ToAddList.Value, &i)
		}
		if err = tmp_tci_StatesPDCCH_ToAddList.Encode(w); err != nil {
			return utils.WrapError("Encode tci_StatesPDCCH_ToAddList", err)
		}
	}
	if len(ie.tci_StatesPDCCH_ToReleaseList) > 0 {
		tmp_tci_StatesPDCCH_ToReleaseList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_StatesPDCCH}, false)
		for _, i := range ie.tci_StatesPDCCH_ToReleaseList {
			tmp_tci_StatesPDCCH_ToReleaseList.Value = append(tmp_tci_StatesPDCCH_ToReleaseList.Value, &i)
		}
		if err = tmp_tci_StatesPDCCH_ToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode tci_StatesPDCCH_ToReleaseList", err)
		}
	}
	if ie.tci_PresentInDCI != nil {
		if err = ie.tci_PresentInDCI.Encode(w); err != nil {
			return utils.WrapError("Encode tci_PresentInDCI", err)
		}
	}
	if ie.pdcch_DMRS_ScramblingID != nil {
		if err = w.WriteInteger(*ie.pdcch_DMRS_ScramblingID, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode pdcch_DMRS_ScramblingID", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.rb_Offset_r16 != nil || ie.tci_PresentDCI_1_2_r16 != nil || ie.coresetPoolIndex_r16 != nil || ie.controlResourceSetId_v1610 != nil, ie.followUnifiedTCI_State_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ControlResourceSet", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.rb_Offset_r16 != nil, ie.tci_PresentDCI_1_2_r16 != nil, ie.coresetPoolIndex_r16 != nil, ie.controlResourceSetId_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode rb_Offset_r16 optional
			if ie.rb_Offset_r16 != nil {
				if err = extWriter.WriteInteger(*ie.rb_Offset_r16, &uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode rb_Offset_r16", err)
				}
			}
			// encode tci_PresentDCI_1_2_r16 optional
			if ie.tci_PresentDCI_1_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.tci_PresentDCI_1_2_r16, &uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
					return utils.WrapError("Encode tci_PresentDCI_1_2_r16", err)
				}
			}
			// encode coresetPoolIndex_r16 optional
			if ie.coresetPoolIndex_r16 != nil {
				if err = extWriter.WriteInteger(*ie.coresetPoolIndex_r16, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
					return utils.WrapError("Encode coresetPoolIndex_r16", err)
				}
			}
			// encode controlResourceSetId_v1610 optional
			if ie.controlResourceSetId_v1610 != nil {
				if err = ie.controlResourceSetId_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode controlResourceSetId_v1610", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.followUnifiedTCI_State_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode followUnifiedTCI_State_r17 optional
			if ie.followUnifiedTCI_State_r17 != nil {
				if err = ie.followUnifiedTCI_State_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode followUnifiedTCI_State_r17", err)
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

func (ie *ControlResourceSet) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var cce_REG_MappingTypePresent bool
	if cce_REG_MappingTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tci_StatesPDCCH_ToAddListPresent bool
	if tci_StatesPDCCH_ToAddListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tci_StatesPDCCH_ToReleaseListPresent bool
	if tci_StatesPDCCH_ToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tci_PresentInDCIPresent bool
	if tci_PresentInDCIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_DMRS_ScramblingIDPresent bool
	if pdcch_DMRS_ScramblingIDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.controlResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode controlResourceSetId", err)
	}
	var tmp_bs_frequencyDomainResources []byte
	var n_frequencyDomainResources uint
	if tmp_bs_frequencyDomainResources, n_frequencyDomainResources, err = r.ReadBitString(&uper.Constraint{Lb: 45, Ub: 45}, false); err != nil {
		return utils.WrapError("ReadBitString frequencyDomainResources", err)
	}
	ie.frequencyDomainResources = uper.BitString{
		Bytes:   tmp_bs_frequencyDomainResources,
		NumBits: uint64(n_frequencyDomainResources),
	}
	var tmp_int_duration int64
	if tmp_int_duration, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxCoReSetDuration}, false); err != nil {
		return utils.WrapError("ReadInteger duration", err)
	}
	ie.duration = tmp_int_duration
	if cce_REG_MappingTypePresent {
		ie.cce_REG_MappingType = new(ControlResourceSet_cce_REG_MappingType)
		if err = ie.cce_REG_MappingType.Decode(r); err != nil {
			return utils.WrapError("Decode cce_REG_MappingType", err)
		}
	}
	if err = ie.precoderGranularity.Decode(r); err != nil {
		return utils.WrapError("Decode precoderGranularity", err)
	}
	if tci_StatesPDCCH_ToAddListPresent {
		tmp_tci_StatesPDCCH_ToAddList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_StatesPDCCH}, false)
		fn_tci_StatesPDCCH_ToAddList := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_tci_StatesPDCCH_ToAddList.Decode(r, fn_tci_StatesPDCCH_ToAddList); err != nil {
			return utils.WrapError("Decode tci_StatesPDCCH_ToAddList", err)
		}
		ie.tci_StatesPDCCH_ToAddList = []TCI_StateId{}
		for _, i := range tmp_tci_StatesPDCCH_ToAddList.Value {
			ie.tci_StatesPDCCH_ToAddList = append(ie.tci_StatesPDCCH_ToAddList, *i)
		}
	}
	if tci_StatesPDCCH_ToReleaseListPresent {
		tmp_tci_StatesPDCCH_ToReleaseList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_StatesPDCCH}, false)
		fn_tci_StatesPDCCH_ToReleaseList := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_tci_StatesPDCCH_ToReleaseList.Decode(r, fn_tci_StatesPDCCH_ToReleaseList); err != nil {
			return utils.WrapError("Decode tci_StatesPDCCH_ToReleaseList", err)
		}
		ie.tci_StatesPDCCH_ToReleaseList = []TCI_StateId{}
		for _, i := range tmp_tci_StatesPDCCH_ToReleaseList.Value {
			ie.tci_StatesPDCCH_ToReleaseList = append(ie.tci_StatesPDCCH_ToReleaseList, *i)
		}
	}
	if tci_PresentInDCIPresent {
		ie.tci_PresentInDCI = new(ControlResourceSet_tci_PresentInDCI)
		if err = ie.tci_PresentInDCI.Decode(r); err != nil {
			return utils.WrapError("Decode tci_PresentInDCI", err)
		}
	}
	if pdcch_DMRS_ScramblingIDPresent {
		var tmp_int_pdcch_DMRS_ScramblingID int64
		if tmp_int_pdcch_DMRS_ScramblingID, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode pdcch_DMRS_ScramblingID", err)
		}
		ie.pdcch_DMRS_ScramblingID = &tmp_int_pdcch_DMRS_ScramblingID
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			rb_Offset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			tci_PresentDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			coresetPoolIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			controlResourceSetId_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode rb_Offset_r16 optional
			if rb_Offset_r16Present {
				var tmp_int_rb_Offset_r16 int64
				if tmp_int_rb_Offset_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode rb_Offset_r16", err)
				}
				ie.rb_Offset_r16 = &tmp_int_rb_Offset_r16
			}
			// decode tci_PresentDCI_1_2_r16 optional
			if tci_PresentDCI_1_2_r16Present {
				var tmp_int_tci_PresentDCI_1_2_r16 int64
				if tmp_int_tci_PresentDCI_1_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
					return utils.WrapError("Decode tci_PresentDCI_1_2_r16", err)
				}
				ie.tci_PresentDCI_1_2_r16 = &tmp_int_tci_PresentDCI_1_2_r16
			}
			// decode coresetPoolIndex_r16 optional
			if coresetPoolIndex_r16Present {
				var tmp_int_coresetPoolIndex_r16 int64
				if tmp_int_coresetPoolIndex_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
					return utils.WrapError("Decode coresetPoolIndex_r16", err)
				}
				ie.coresetPoolIndex_r16 = &tmp_int_coresetPoolIndex_r16
			}
			// decode controlResourceSetId_v1610 optional
			if controlResourceSetId_v1610Present {
				ie.controlResourceSetId_v1610 = new(ControlResourceSetId_v1610)
				if err = ie.controlResourceSetId_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode controlResourceSetId_v1610", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			followUnifiedTCI_State_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode followUnifiedTCI_State_r17 optional
			if followUnifiedTCI_State_r17Present {
				ie.followUnifiedTCI_State_r17 = new(ControlResourceSet_followUnifiedTCI_State_r17)
				if err = ie.followUnifiedTCI_State_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode followUnifiedTCI_State_r17", err)
				}
			}
		}
	}
	return nil
}
