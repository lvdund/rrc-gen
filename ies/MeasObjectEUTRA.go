package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectEUTRA struct {
	carrierFreq                     ARFCN_ValueEUTRA           `madatory`
	allowedMeasBandwidth            EUTRA_AllowedMeasBandwidth `madatory`
	cellsToRemoveListEUTRAN         *EUTRA_CellIndexList       `optional`
	cellsToAddModListEUTRAN         []EUTRA_Cell               `lb:1,ub:maxCellMeasEUTRA,optional`
	excludedCellsToRemoveListEUTRAN *EUTRA_CellIndexList       `optional`
	excludedCellsToAddModListEUTRAN []EUTRA_ExcludedCell       `lb:1,ub:maxCellMeasEUTRA,optional`
	eutra_PresenceAntennaPort1      EUTRA_PresenceAntennaPort1 `madatory`
	eutra_Q_OffsetRange             *EUTRA_Q_OffsetRange       `optional`
	widebandRSRQ_Meas               bool                       `madatory`
	associatedMeasGap_r17           *MeasGapId_r17             `optional,ext-1`
}

func (ie *MeasObjectEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.associatedMeasGap_r17 != nil
	preambleBits := []bool{hasExtensions, ie.cellsToRemoveListEUTRAN != nil, len(ie.cellsToAddModListEUTRAN) > 0, ie.excludedCellsToRemoveListEUTRAN != nil, len(ie.excludedCellsToAddModListEUTRAN) > 0, ie.eutra_Q_OffsetRange != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.carrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq", err)
	}
	if err = ie.allowedMeasBandwidth.Encode(w); err != nil {
		return utils.WrapError("Encode allowedMeasBandwidth", err)
	}
	if ie.cellsToRemoveListEUTRAN != nil {
		if err = ie.cellsToRemoveListEUTRAN.Encode(w); err != nil {
			return utils.WrapError("Encode cellsToRemoveListEUTRAN", err)
		}
	}
	if len(ie.cellsToAddModListEUTRAN) > 0 {
		tmp_cellsToAddModListEUTRAN := utils.NewSequence[*EUTRA_Cell]([]*EUTRA_Cell{}, uper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false)
		for _, i := range ie.cellsToAddModListEUTRAN {
			tmp_cellsToAddModListEUTRAN.Value = append(tmp_cellsToAddModListEUTRAN.Value, &i)
		}
		if err = tmp_cellsToAddModListEUTRAN.Encode(w); err != nil {
			return utils.WrapError("Encode cellsToAddModListEUTRAN", err)
		}
	}
	if ie.excludedCellsToRemoveListEUTRAN != nil {
		if err = ie.excludedCellsToRemoveListEUTRAN.Encode(w); err != nil {
			return utils.WrapError("Encode excludedCellsToRemoveListEUTRAN", err)
		}
	}
	if len(ie.excludedCellsToAddModListEUTRAN) > 0 {
		tmp_excludedCellsToAddModListEUTRAN := utils.NewSequence[*EUTRA_ExcludedCell]([]*EUTRA_ExcludedCell{}, uper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false)
		for _, i := range ie.excludedCellsToAddModListEUTRAN {
			tmp_excludedCellsToAddModListEUTRAN.Value = append(tmp_excludedCellsToAddModListEUTRAN.Value, &i)
		}
		if err = tmp_excludedCellsToAddModListEUTRAN.Encode(w); err != nil {
			return utils.WrapError("Encode excludedCellsToAddModListEUTRAN", err)
		}
	}
	if err = ie.eutra_PresenceAntennaPort1.Encode(w); err != nil {
		return utils.WrapError("Encode eutra_PresenceAntennaPort1", err)
	}
	if ie.eutra_Q_OffsetRange != nil {
		if err = ie.eutra_Q_OffsetRange.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_Q_OffsetRange", err)
		}
	}
	if err = w.WriteBoolean(ie.widebandRSRQ_Meas); err != nil {
		return utils.WrapError("WriteBoolean widebandRSRQ_Meas", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.associatedMeasGap_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasObjectEUTRA", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.associatedMeasGap_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode associatedMeasGap_r17 optional
			if ie.associatedMeasGap_r17 != nil {
				if err = ie.associatedMeasGap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode associatedMeasGap_r17", err)
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

func (ie *MeasObjectEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var cellsToRemoveListEUTRANPresent bool
	if cellsToRemoveListEUTRANPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cellsToAddModListEUTRANPresent bool
	if cellsToAddModListEUTRANPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var excludedCellsToRemoveListEUTRANPresent bool
	if excludedCellsToRemoveListEUTRANPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var excludedCellsToAddModListEUTRANPresent bool
	if excludedCellsToAddModListEUTRANPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var eutra_Q_OffsetRangePresent bool
	if eutra_Q_OffsetRangePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.carrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq", err)
	}
	if err = ie.allowedMeasBandwidth.Decode(r); err != nil {
		return utils.WrapError("Decode allowedMeasBandwidth", err)
	}
	if cellsToRemoveListEUTRANPresent {
		ie.cellsToRemoveListEUTRAN = new(EUTRA_CellIndexList)
		if err = ie.cellsToRemoveListEUTRAN.Decode(r); err != nil {
			return utils.WrapError("Decode cellsToRemoveListEUTRAN", err)
		}
	}
	if cellsToAddModListEUTRANPresent {
		tmp_cellsToAddModListEUTRAN := utils.NewSequence[*EUTRA_Cell]([]*EUTRA_Cell{}, uper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false)
		fn_cellsToAddModListEUTRAN := func() *EUTRA_Cell {
			return new(EUTRA_Cell)
		}
		if err = tmp_cellsToAddModListEUTRAN.Decode(r, fn_cellsToAddModListEUTRAN); err != nil {
			return utils.WrapError("Decode cellsToAddModListEUTRAN", err)
		}
		ie.cellsToAddModListEUTRAN = []EUTRA_Cell{}
		for _, i := range tmp_cellsToAddModListEUTRAN.Value {
			ie.cellsToAddModListEUTRAN = append(ie.cellsToAddModListEUTRAN, *i)
		}
	}
	if excludedCellsToRemoveListEUTRANPresent {
		ie.excludedCellsToRemoveListEUTRAN = new(EUTRA_CellIndexList)
		if err = ie.excludedCellsToRemoveListEUTRAN.Decode(r); err != nil {
			return utils.WrapError("Decode excludedCellsToRemoveListEUTRAN", err)
		}
	}
	if excludedCellsToAddModListEUTRANPresent {
		tmp_excludedCellsToAddModListEUTRAN := utils.NewSequence[*EUTRA_ExcludedCell]([]*EUTRA_ExcludedCell{}, uper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false)
		fn_excludedCellsToAddModListEUTRAN := func() *EUTRA_ExcludedCell {
			return new(EUTRA_ExcludedCell)
		}
		if err = tmp_excludedCellsToAddModListEUTRAN.Decode(r, fn_excludedCellsToAddModListEUTRAN); err != nil {
			return utils.WrapError("Decode excludedCellsToAddModListEUTRAN", err)
		}
		ie.excludedCellsToAddModListEUTRAN = []EUTRA_ExcludedCell{}
		for _, i := range tmp_excludedCellsToAddModListEUTRAN.Value {
			ie.excludedCellsToAddModListEUTRAN = append(ie.excludedCellsToAddModListEUTRAN, *i)
		}
	}
	if err = ie.eutra_PresenceAntennaPort1.Decode(r); err != nil {
		return utils.WrapError("Decode eutra_PresenceAntennaPort1", err)
	}
	if eutra_Q_OffsetRangePresent {
		ie.eutra_Q_OffsetRange = new(EUTRA_Q_OffsetRange)
		if err = ie.eutra_Q_OffsetRange.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_Q_OffsetRange", err)
		}
	}
	var tmp_bool_widebandRSRQ_Meas bool
	if tmp_bool_widebandRSRQ_Meas, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean widebandRSRQ_Meas", err)
	}
	ie.widebandRSRQ_Meas = tmp_bool_widebandRSRQ_Meas

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

			associatedMeasGap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode associatedMeasGap_r17 optional
			if associatedMeasGap_r17Present {
				ie.associatedMeasGap_r17 = new(MeasGapId_r17)
				if err = ie.associatedMeasGap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode associatedMeasGap_r17", err)
				}
			}
		}
	}
	return nil
}
