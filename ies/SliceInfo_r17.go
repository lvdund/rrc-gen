package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SliceInfo_r17 struct {
	nsag_IdentityInfo_r17               NSAG_IdentityInfo_r17              `madatory`
	nsag_CellReselectionPriority_r17    *CellReselectionPriority           `optional`
	nsag_CellReselectionSubPriority_r17 *CellReselectionSubPriority        `optional`
	sliceCellListNR_r17                 *SliceInfo_r17_sliceCellListNR_r17 `optional`
}

func (ie *SliceInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nsag_CellReselectionPriority_r17 != nil, ie.nsag_CellReselectionSubPriority_r17 != nil, ie.sliceCellListNR_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.nsag_IdentityInfo_r17.Encode(w); err != nil {
		return utils.WrapError("Encode nsag_IdentityInfo_r17", err)
	}
	if ie.nsag_CellReselectionPriority_r17 != nil {
		if err = ie.nsag_CellReselectionPriority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nsag_CellReselectionPriority_r17", err)
		}
	}
	if ie.nsag_CellReselectionSubPriority_r17 != nil {
		if err = ie.nsag_CellReselectionSubPriority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nsag_CellReselectionSubPriority_r17", err)
		}
	}
	if ie.sliceCellListNR_r17 != nil {
		if err = ie.sliceCellListNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sliceCellListNR_r17", err)
		}
	}
	return nil
}

func (ie *SliceInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var nsag_CellReselectionPriority_r17Present bool
	if nsag_CellReselectionPriority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nsag_CellReselectionSubPriority_r17Present bool
	if nsag_CellReselectionSubPriority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sliceCellListNR_r17Present bool
	if sliceCellListNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.nsag_IdentityInfo_r17.Decode(r); err != nil {
		return utils.WrapError("Decode nsag_IdentityInfo_r17", err)
	}
	if nsag_CellReselectionPriority_r17Present {
		ie.nsag_CellReselectionPriority_r17 = new(CellReselectionPriority)
		if err = ie.nsag_CellReselectionPriority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode nsag_CellReselectionPriority_r17", err)
		}
	}
	if nsag_CellReselectionSubPriority_r17Present {
		ie.nsag_CellReselectionSubPriority_r17 = new(CellReselectionSubPriority)
		if err = ie.nsag_CellReselectionSubPriority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode nsag_CellReselectionSubPriority_r17", err)
		}
	}
	if sliceCellListNR_r17Present {
		ie.sliceCellListNR_r17 = new(SliceInfo_r17_sliceCellListNR_r17)
		if err = ie.sliceCellListNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sliceCellListNR_r17", err)
		}
	}
	return nil
}
