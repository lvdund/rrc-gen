package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17 struct {
	maxNumListDL_TCI_r17 *MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17_maxNumListDL_TCI_r17 `optional`
	maxNumListUL_TCI_r17 *MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17_maxNumListUL_TCI_r17 `optional`
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxNumListDL_TCI_r17 != nil, ie.maxNumListUL_TCI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxNumListDL_TCI_r17 != nil {
		if err = ie.maxNumListDL_TCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumListDL_TCI_r17", err)
		}
	}
	if ie.maxNumListUL_TCI_r17 != nil {
		if err = ie.maxNumListUL_TCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumListUL_TCI_r17", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17) Decode(r *uper.UperReader) error {
	var err error
	var maxNumListDL_TCI_r17Present bool
	if maxNumListDL_TCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumListUL_TCI_r17Present bool
	if maxNumListUL_TCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxNumListDL_TCI_r17Present {
		ie.maxNumListDL_TCI_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17_maxNumListDL_TCI_r17)
		if err = ie.maxNumListDL_TCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumListDL_TCI_r17", err)
		}
	}
	if maxNumListUL_TCI_r17Present {
		ie.maxNumListUL_TCI_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17_maxNumListUL_TCI_r17)
		if err = ie.maxNumListUL_TCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumListUL_TCI_r17", err)
		}
	}
	return nil
}
