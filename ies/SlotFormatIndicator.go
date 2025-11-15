package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SlotFormatIndicator struct {
	sfi_RNTI                              RNTI_Value                      `madatory`
	dci_PayloadSize                       int64                           `lb:1,ub:maxSFI_DCI_PayloadSize,madatory`
	slotFormatCombToAddModList            []SlotFormatCombinationsPerCell `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional`
	slotFormatCombToReleaseList           []ServCellIndex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional`
	availableRB_SetsToAddModList_r16      []AvailableRB_SetsPerCell_r16   `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional,ext-1`
	availableRB_SetsToReleaseList_r16     []ServCellIndex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional,ext-1`
	switchTriggerToAddModList_r16         []SearchSpaceSwitchTrigger_r16  `lb:1,ub:4,optional,ext-1`
	switchTriggerToReleaseList_r16        []ServCellIndex                 `lb:1,ub:4,optional,ext-1`
	co_DurationsPerCellToAddModList_r16   []CO_DurationsPerCell_r16       `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional,ext-1`
	co_DurationsPerCellToReleaseList_r16  []ServCellIndex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional,ext-1`
	switchTriggerToAddModListSizeExt_r16  []SearchSpaceSwitchTrigger_r16  `lb:1,ub:maxNrofAggregatedCellsPerCellGroupMinus4_r16,optional,ext-2`
	switchTriggerToReleaseListSizeExt_r16 []ServCellIndex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroupMinus4_r16,optional,ext-2`
	co_DurationsPerCellToAddModList_r17   []CO_DurationsPerCell_r17       `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional,ext-3`
}

func (ie *SlotFormatIndicator) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.availableRB_SetsToAddModList_r16) > 0 || len(ie.availableRB_SetsToReleaseList_r16) > 0 || len(ie.switchTriggerToAddModList_r16) > 0 || len(ie.switchTriggerToReleaseList_r16) > 0 || len(ie.co_DurationsPerCellToAddModList_r16) > 0 || len(ie.co_DurationsPerCellToReleaseList_r16) > 0 || len(ie.switchTriggerToAddModListSizeExt_r16) > 0 || len(ie.switchTriggerToReleaseListSizeExt_r16) > 0 || len(ie.co_DurationsPerCellToAddModList_r17) > 0
	preambleBits := []bool{hasExtensions, len(ie.slotFormatCombToAddModList) > 0, len(ie.slotFormatCombToReleaseList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sfi_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode sfi_RNTI", err)
	}
	if err = w.WriteInteger(ie.dci_PayloadSize, &uper.Constraint{Lb: 1, Ub: maxSFI_DCI_PayloadSize}, false); err != nil {
		return utils.WrapError("WriteInteger dci_PayloadSize", err)
	}
	if len(ie.slotFormatCombToAddModList) > 0 {
		tmp_slotFormatCombToAddModList := utils.NewSequence[*SlotFormatCombinationsPerCell]([]*SlotFormatCombinationsPerCell{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
		for _, i := range ie.slotFormatCombToAddModList {
			tmp_slotFormatCombToAddModList.Value = append(tmp_slotFormatCombToAddModList.Value, &i)
		}
		if err = tmp_slotFormatCombToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode slotFormatCombToAddModList", err)
		}
	}
	if len(ie.slotFormatCombToReleaseList) > 0 {
		tmp_slotFormatCombToReleaseList := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
		for _, i := range ie.slotFormatCombToReleaseList {
			tmp_slotFormatCombToReleaseList.Value = append(tmp_slotFormatCombToReleaseList.Value, &i)
		}
		if err = tmp_slotFormatCombToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode slotFormatCombToReleaseList", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{len(ie.availableRB_SetsToAddModList_r16) > 0 || len(ie.availableRB_SetsToReleaseList_r16) > 0 || len(ie.switchTriggerToAddModList_r16) > 0 || len(ie.switchTriggerToReleaseList_r16) > 0 || len(ie.co_DurationsPerCellToAddModList_r16) > 0 || len(ie.co_DurationsPerCellToReleaseList_r16) > 0, len(ie.switchTriggerToAddModListSizeExt_r16) > 0 || len(ie.switchTriggerToReleaseListSizeExt_r16) > 0, len(ie.co_DurationsPerCellToAddModList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SlotFormatIndicator", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.availableRB_SetsToAddModList_r16) > 0, len(ie.availableRB_SetsToReleaseList_r16) > 0, len(ie.switchTriggerToAddModList_r16) > 0, len(ie.switchTriggerToReleaseList_r16) > 0, len(ie.co_DurationsPerCellToAddModList_r16) > 0, len(ie.co_DurationsPerCellToReleaseList_r16) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode availableRB_SetsToAddModList_r16 optional
			if len(ie.availableRB_SetsToAddModList_r16) > 0 {
				tmp_availableRB_SetsToAddModList_r16 := utils.NewSequence[*AvailableRB_SetsPerCell_r16]([]*AvailableRB_SetsPerCell_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				for _, i := range ie.availableRB_SetsToAddModList_r16 {
					tmp_availableRB_SetsToAddModList_r16.Value = append(tmp_availableRB_SetsToAddModList_r16.Value, &i)
				}
				if err = tmp_availableRB_SetsToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode availableRB_SetsToAddModList_r16", err)
				}
			}
			// encode availableRB_SetsToReleaseList_r16 optional
			if len(ie.availableRB_SetsToReleaseList_r16) > 0 {
				tmp_availableRB_SetsToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				for _, i := range ie.availableRB_SetsToReleaseList_r16 {
					tmp_availableRB_SetsToReleaseList_r16.Value = append(tmp_availableRB_SetsToReleaseList_r16.Value, &i)
				}
				if err = tmp_availableRB_SetsToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode availableRB_SetsToReleaseList_r16", err)
				}
			}
			// encode switchTriggerToAddModList_r16 optional
			if len(ie.switchTriggerToAddModList_r16) > 0 {
				tmp_switchTriggerToAddModList_r16 := utils.NewSequence[*SearchSpaceSwitchTrigger_r16]([]*SearchSpaceSwitchTrigger_r16{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				for _, i := range ie.switchTriggerToAddModList_r16 {
					tmp_switchTriggerToAddModList_r16.Value = append(tmp_switchTriggerToAddModList_r16.Value, &i)
				}
				if err = tmp_switchTriggerToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode switchTriggerToAddModList_r16", err)
				}
			}
			// encode switchTriggerToReleaseList_r16 optional
			if len(ie.switchTriggerToReleaseList_r16) > 0 {
				tmp_switchTriggerToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				for _, i := range ie.switchTriggerToReleaseList_r16 {
					tmp_switchTriggerToReleaseList_r16.Value = append(tmp_switchTriggerToReleaseList_r16.Value, &i)
				}
				if err = tmp_switchTriggerToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode switchTriggerToReleaseList_r16", err)
				}
			}
			// encode co_DurationsPerCellToAddModList_r16 optional
			if len(ie.co_DurationsPerCellToAddModList_r16) > 0 {
				tmp_co_DurationsPerCellToAddModList_r16 := utils.NewSequence[*CO_DurationsPerCell_r16]([]*CO_DurationsPerCell_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				for _, i := range ie.co_DurationsPerCellToAddModList_r16 {
					tmp_co_DurationsPerCellToAddModList_r16.Value = append(tmp_co_DurationsPerCellToAddModList_r16.Value, &i)
				}
				if err = tmp_co_DurationsPerCellToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode co_DurationsPerCellToAddModList_r16", err)
				}
			}
			// encode co_DurationsPerCellToReleaseList_r16 optional
			if len(ie.co_DurationsPerCellToReleaseList_r16) > 0 {
				tmp_co_DurationsPerCellToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				for _, i := range ie.co_DurationsPerCellToReleaseList_r16 {
					tmp_co_DurationsPerCellToReleaseList_r16.Value = append(tmp_co_DurationsPerCellToReleaseList_r16.Value, &i)
				}
				if err = tmp_co_DurationsPerCellToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode co_DurationsPerCellToReleaseList_r16", err)
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
			optionals_ext_2 := []bool{len(ie.switchTriggerToAddModListSizeExt_r16) > 0, len(ie.switchTriggerToReleaseListSizeExt_r16) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode switchTriggerToAddModListSizeExt_r16 optional
			if len(ie.switchTriggerToAddModListSizeExt_r16) > 0 {
				tmp_switchTriggerToAddModListSizeExt_r16 := utils.NewSequence[*SearchSpaceSwitchTrigger_r16]([]*SearchSpaceSwitchTrigger_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroupMinus4_r16}, false)
				for _, i := range ie.switchTriggerToAddModListSizeExt_r16 {
					tmp_switchTriggerToAddModListSizeExt_r16.Value = append(tmp_switchTriggerToAddModListSizeExt_r16.Value, &i)
				}
				if err = tmp_switchTriggerToAddModListSizeExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode switchTriggerToAddModListSizeExt_r16", err)
				}
			}
			// encode switchTriggerToReleaseListSizeExt_r16 optional
			if len(ie.switchTriggerToReleaseListSizeExt_r16) > 0 {
				tmp_switchTriggerToReleaseListSizeExt_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroupMinus4_r16}, false)
				for _, i := range ie.switchTriggerToReleaseListSizeExt_r16 {
					tmp_switchTriggerToReleaseListSizeExt_r16.Value = append(tmp_switchTriggerToReleaseListSizeExt_r16.Value, &i)
				}
				if err = tmp_switchTriggerToReleaseListSizeExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode switchTriggerToReleaseListSizeExt_r16", err)
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
			optionals_ext_3 := []bool{len(ie.co_DurationsPerCellToAddModList_r17) > 0}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode co_DurationsPerCellToAddModList_r17 optional
			if len(ie.co_DurationsPerCellToAddModList_r17) > 0 {
				tmp_co_DurationsPerCellToAddModList_r17 := utils.NewSequence[*CO_DurationsPerCell_r17]([]*CO_DurationsPerCell_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				for _, i := range ie.co_DurationsPerCellToAddModList_r17 {
					tmp_co_DurationsPerCellToAddModList_r17.Value = append(tmp_co_DurationsPerCellToAddModList_r17.Value, &i)
				}
				if err = tmp_co_DurationsPerCellToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode co_DurationsPerCellToAddModList_r17", err)
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

func (ie *SlotFormatIndicator) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var slotFormatCombToAddModListPresent bool
	if slotFormatCombToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var slotFormatCombToReleaseListPresent bool
	if slotFormatCombToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sfi_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode sfi_RNTI", err)
	}
	var tmp_int_dci_PayloadSize int64
	if tmp_int_dci_PayloadSize, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxSFI_DCI_PayloadSize}, false); err != nil {
		return utils.WrapError("ReadInteger dci_PayloadSize", err)
	}
	ie.dci_PayloadSize = tmp_int_dci_PayloadSize
	if slotFormatCombToAddModListPresent {
		tmp_slotFormatCombToAddModList := utils.NewSequence[*SlotFormatCombinationsPerCell]([]*SlotFormatCombinationsPerCell{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
		fn_slotFormatCombToAddModList := func() *SlotFormatCombinationsPerCell {
			return new(SlotFormatCombinationsPerCell)
		}
		if err = tmp_slotFormatCombToAddModList.Decode(r, fn_slotFormatCombToAddModList); err != nil {
			return utils.WrapError("Decode slotFormatCombToAddModList", err)
		}
		ie.slotFormatCombToAddModList = []SlotFormatCombinationsPerCell{}
		for _, i := range tmp_slotFormatCombToAddModList.Value {
			ie.slotFormatCombToAddModList = append(ie.slotFormatCombToAddModList, *i)
		}
	}
	if slotFormatCombToReleaseListPresent {
		tmp_slotFormatCombToReleaseList := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
		fn_slotFormatCombToReleaseList := func() *ServCellIndex {
			return new(ServCellIndex)
		}
		if err = tmp_slotFormatCombToReleaseList.Decode(r, fn_slotFormatCombToReleaseList); err != nil {
			return utils.WrapError("Decode slotFormatCombToReleaseList", err)
		}
		ie.slotFormatCombToReleaseList = []ServCellIndex{}
		for _, i := range tmp_slotFormatCombToReleaseList.Value {
			ie.slotFormatCombToReleaseList = append(ie.slotFormatCombToReleaseList, *i)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			availableRB_SetsToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			availableRB_SetsToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			switchTriggerToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			switchTriggerToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			co_DurationsPerCellToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			co_DurationsPerCellToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode availableRB_SetsToAddModList_r16 optional
			if availableRB_SetsToAddModList_r16Present {
				tmp_availableRB_SetsToAddModList_r16 := utils.NewSequence[*AvailableRB_SetsPerCell_r16]([]*AvailableRB_SetsPerCell_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				fn_availableRB_SetsToAddModList_r16 := func() *AvailableRB_SetsPerCell_r16 {
					return new(AvailableRB_SetsPerCell_r16)
				}
				if err = tmp_availableRB_SetsToAddModList_r16.Decode(extReader, fn_availableRB_SetsToAddModList_r16); err != nil {
					return utils.WrapError("Decode availableRB_SetsToAddModList_r16", err)
				}
				ie.availableRB_SetsToAddModList_r16 = []AvailableRB_SetsPerCell_r16{}
				for _, i := range tmp_availableRB_SetsToAddModList_r16.Value {
					ie.availableRB_SetsToAddModList_r16 = append(ie.availableRB_SetsToAddModList_r16, *i)
				}
			}
			// decode availableRB_SetsToReleaseList_r16 optional
			if availableRB_SetsToReleaseList_r16Present {
				tmp_availableRB_SetsToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				fn_availableRB_SetsToReleaseList_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_availableRB_SetsToReleaseList_r16.Decode(extReader, fn_availableRB_SetsToReleaseList_r16); err != nil {
					return utils.WrapError("Decode availableRB_SetsToReleaseList_r16", err)
				}
				ie.availableRB_SetsToReleaseList_r16 = []ServCellIndex{}
				for _, i := range tmp_availableRB_SetsToReleaseList_r16.Value {
					ie.availableRB_SetsToReleaseList_r16 = append(ie.availableRB_SetsToReleaseList_r16, *i)
				}
			}
			// decode switchTriggerToAddModList_r16 optional
			if switchTriggerToAddModList_r16Present {
				tmp_switchTriggerToAddModList_r16 := utils.NewSequence[*SearchSpaceSwitchTrigger_r16]([]*SearchSpaceSwitchTrigger_r16{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				fn_switchTriggerToAddModList_r16 := func() *SearchSpaceSwitchTrigger_r16 {
					return new(SearchSpaceSwitchTrigger_r16)
				}
				if err = tmp_switchTriggerToAddModList_r16.Decode(extReader, fn_switchTriggerToAddModList_r16); err != nil {
					return utils.WrapError("Decode switchTriggerToAddModList_r16", err)
				}
				ie.switchTriggerToAddModList_r16 = []SearchSpaceSwitchTrigger_r16{}
				for _, i := range tmp_switchTriggerToAddModList_r16.Value {
					ie.switchTriggerToAddModList_r16 = append(ie.switchTriggerToAddModList_r16, *i)
				}
			}
			// decode switchTriggerToReleaseList_r16 optional
			if switchTriggerToReleaseList_r16Present {
				tmp_switchTriggerToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				fn_switchTriggerToReleaseList_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_switchTriggerToReleaseList_r16.Decode(extReader, fn_switchTriggerToReleaseList_r16); err != nil {
					return utils.WrapError("Decode switchTriggerToReleaseList_r16", err)
				}
				ie.switchTriggerToReleaseList_r16 = []ServCellIndex{}
				for _, i := range tmp_switchTriggerToReleaseList_r16.Value {
					ie.switchTriggerToReleaseList_r16 = append(ie.switchTriggerToReleaseList_r16, *i)
				}
			}
			// decode co_DurationsPerCellToAddModList_r16 optional
			if co_DurationsPerCellToAddModList_r16Present {
				tmp_co_DurationsPerCellToAddModList_r16 := utils.NewSequence[*CO_DurationsPerCell_r16]([]*CO_DurationsPerCell_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				fn_co_DurationsPerCellToAddModList_r16 := func() *CO_DurationsPerCell_r16 {
					return new(CO_DurationsPerCell_r16)
				}
				if err = tmp_co_DurationsPerCellToAddModList_r16.Decode(extReader, fn_co_DurationsPerCellToAddModList_r16); err != nil {
					return utils.WrapError("Decode co_DurationsPerCellToAddModList_r16", err)
				}
				ie.co_DurationsPerCellToAddModList_r16 = []CO_DurationsPerCell_r16{}
				for _, i := range tmp_co_DurationsPerCellToAddModList_r16.Value {
					ie.co_DurationsPerCellToAddModList_r16 = append(ie.co_DurationsPerCellToAddModList_r16, *i)
				}
			}
			// decode co_DurationsPerCellToReleaseList_r16 optional
			if co_DurationsPerCellToReleaseList_r16Present {
				tmp_co_DurationsPerCellToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				fn_co_DurationsPerCellToReleaseList_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_co_DurationsPerCellToReleaseList_r16.Decode(extReader, fn_co_DurationsPerCellToReleaseList_r16); err != nil {
					return utils.WrapError("Decode co_DurationsPerCellToReleaseList_r16", err)
				}
				ie.co_DurationsPerCellToReleaseList_r16 = []ServCellIndex{}
				for _, i := range tmp_co_DurationsPerCellToReleaseList_r16.Value {
					ie.co_DurationsPerCellToReleaseList_r16 = append(ie.co_DurationsPerCellToReleaseList_r16, *i)
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

			switchTriggerToAddModListSizeExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			switchTriggerToReleaseListSizeExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode switchTriggerToAddModListSizeExt_r16 optional
			if switchTriggerToAddModListSizeExt_r16Present {
				tmp_switchTriggerToAddModListSizeExt_r16 := utils.NewSequence[*SearchSpaceSwitchTrigger_r16]([]*SearchSpaceSwitchTrigger_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroupMinus4_r16}, false)
				fn_switchTriggerToAddModListSizeExt_r16 := func() *SearchSpaceSwitchTrigger_r16 {
					return new(SearchSpaceSwitchTrigger_r16)
				}
				if err = tmp_switchTriggerToAddModListSizeExt_r16.Decode(extReader, fn_switchTriggerToAddModListSizeExt_r16); err != nil {
					return utils.WrapError("Decode switchTriggerToAddModListSizeExt_r16", err)
				}
				ie.switchTriggerToAddModListSizeExt_r16 = []SearchSpaceSwitchTrigger_r16{}
				for _, i := range tmp_switchTriggerToAddModListSizeExt_r16.Value {
					ie.switchTriggerToAddModListSizeExt_r16 = append(ie.switchTriggerToAddModListSizeExt_r16, *i)
				}
			}
			// decode switchTriggerToReleaseListSizeExt_r16 optional
			if switchTriggerToReleaseListSizeExt_r16Present {
				tmp_switchTriggerToReleaseListSizeExt_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroupMinus4_r16}, false)
				fn_switchTriggerToReleaseListSizeExt_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_switchTriggerToReleaseListSizeExt_r16.Decode(extReader, fn_switchTriggerToReleaseListSizeExt_r16); err != nil {
					return utils.WrapError("Decode switchTriggerToReleaseListSizeExt_r16", err)
				}
				ie.switchTriggerToReleaseListSizeExt_r16 = []ServCellIndex{}
				for _, i := range tmp_switchTriggerToReleaseListSizeExt_r16.Value {
					ie.switchTriggerToReleaseListSizeExt_r16 = append(ie.switchTriggerToReleaseListSizeExt_r16, *i)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			co_DurationsPerCellToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode co_DurationsPerCellToAddModList_r17 optional
			if co_DurationsPerCellToAddModList_r17Present {
				tmp_co_DurationsPerCellToAddModList_r17 := utils.NewSequence[*CO_DurationsPerCell_r17]([]*CO_DurationsPerCell_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				fn_co_DurationsPerCellToAddModList_r17 := func() *CO_DurationsPerCell_r17 {
					return new(CO_DurationsPerCell_r17)
				}
				if err = tmp_co_DurationsPerCellToAddModList_r17.Decode(extReader, fn_co_DurationsPerCellToAddModList_r17); err != nil {
					return utils.WrapError("Decode co_DurationsPerCellToAddModList_r17", err)
				}
				ie.co_DurationsPerCellToAddModList_r17 = []CO_DurationsPerCell_r17{}
				for _, i := range tmp_co_DurationsPerCellToAddModList_r17.Value {
					ie.co_DurationsPerCellToAddModList_r17 = append(ie.co_DurationsPerCellToAddModList_r17, *i)
				}
			}
		}
	}
	return nil
}
