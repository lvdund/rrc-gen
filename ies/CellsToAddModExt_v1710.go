package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellsToAddModExt_v1710 struct {
	ntn_PolarizationDL_r17 *CellsToAddModExt_v1710_ntn_PolarizationDL_r17 `optional`
	ntn_PolarizationUL_r17 *CellsToAddModExt_v1710_ntn_PolarizationUL_r17 `optional`
}

func (ie *CellsToAddModExt_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ntn_PolarizationDL_r17 != nil, ie.ntn_PolarizationUL_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ntn_PolarizationDL_r17 != nil {
		if err = ie.ntn_PolarizationDL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ntn_PolarizationDL_r17", err)
		}
	}
	if ie.ntn_PolarizationUL_r17 != nil {
		if err = ie.ntn_PolarizationUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ntn_PolarizationUL_r17", err)
		}
	}
	return nil
}

func (ie *CellsToAddModExt_v1710) Decode(r *uper.UperReader) error {
	var err error
	var ntn_PolarizationDL_r17Present bool
	if ntn_PolarizationDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ntn_PolarizationUL_r17Present bool
	if ntn_PolarizationUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ntn_PolarizationDL_r17Present {
		ie.ntn_PolarizationDL_r17 = new(CellsToAddModExt_v1710_ntn_PolarizationDL_r17)
		if err = ie.ntn_PolarizationDL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ntn_PolarizationDL_r17", err)
		}
	}
	if ntn_PolarizationUL_r17Present {
		ie.ntn_PolarizationUL_r17 = new(CellsToAddModExt_v1710_ntn_PolarizationUL_r17)
		if err = ie.ntn_PolarizationUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ntn_PolarizationUL_r17", err)
		}
	}
	return nil
}
