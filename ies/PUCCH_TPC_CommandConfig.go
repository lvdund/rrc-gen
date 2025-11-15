package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_TPC_CommandConfig struct {
	tpc_IndexPCell                               *int64 `lb:1,ub:15,optional`
	tpc_IndexPUCCH_SCell                         *int64 `lb:1,ub:15,optional`
	tpc_IndexPUCCH_sSCell_r17                    *int64 `lb:1,ub:15,optional,ext-1`
	tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 *int64 `lb:1,ub:15,optional,ext-1`
}

func (ie *PUCCH_TPC_CommandConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.tpc_IndexPUCCH_sSCell_r17 != nil || ie.tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 != nil
	preambleBits := []bool{hasExtensions, ie.tpc_IndexPCell != nil, ie.tpc_IndexPUCCH_SCell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.tpc_IndexPCell != nil {
		if err = w.WriteInteger(*ie.tpc_IndexPCell, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode tpc_IndexPCell", err)
		}
	}
	if ie.tpc_IndexPUCCH_SCell != nil {
		if err = w.WriteInteger(*ie.tpc_IndexPUCCH_SCell, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode tpc_IndexPUCCH_SCell", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.tpc_IndexPUCCH_sSCell_r17 != nil || ie.tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUCCH_TPC_CommandConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.tpc_IndexPUCCH_sSCell_r17 != nil, ie.tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode tpc_IndexPUCCH_sSCell_r17 optional
			if ie.tpc_IndexPUCCH_sSCell_r17 != nil {
				if err = extWriter.WriteInteger(*ie.tpc_IndexPUCCH_sSCell_r17, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode tpc_IndexPUCCH_sSCell_r17", err)
				}
			}
			// encode tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 optional
			if ie.tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 != nil {
				if err = extWriter.WriteInteger(*ie.tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17", err)
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

func (ie *PUCCH_TPC_CommandConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_IndexPCellPresent bool
	if tpc_IndexPCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_IndexPUCCH_SCellPresent bool
	if tpc_IndexPUCCH_SCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if tpc_IndexPCellPresent {
		var tmp_int_tpc_IndexPCell int64
		if tmp_int_tpc_IndexPCell, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode tpc_IndexPCell", err)
		}
		ie.tpc_IndexPCell = &tmp_int_tpc_IndexPCell
	}
	if tpc_IndexPUCCH_SCellPresent {
		var tmp_int_tpc_IndexPUCCH_SCell int64
		if tmp_int_tpc_IndexPUCCH_SCell, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode tpc_IndexPUCCH_SCell", err)
		}
		ie.tpc_IndexPUCCH_SCell = &tmp_int_tpc_IndexPUCCH_SCell
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

			tpc_IndexPUCCH_sSCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode tpc_IndexPUCCH_sSCell_r17 optional
			if tpc_IndexPUCCH_sSCell_r17Present {
				var tmp_int_tpc_IndexPUCCH_sSCell_r17 int64
				if tmp_int_tpc_IndexPUCCH_sSCell_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode tpc_IndexPUCCH_sSCell_r17", err)
				}
				ie.tpc_IndexPUCCH_sSCell_r17 = &tmp_int_tpc_IndexPUCCH_sSCell_r17
			}
			// decode tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 optional
			if tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17Present {
				var tmp_int_tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 int64
				if tmp_int_tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17", err)
				}
				ie.tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 = &tmp_int_tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17
			}
		}
	}
	return nil
}
