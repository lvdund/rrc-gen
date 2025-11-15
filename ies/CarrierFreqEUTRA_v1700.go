package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CarrierFreqEUTRA_v1700 struct {
	eutra_FreqNeighHSDN_CellList_r17 *EUTRA_FreqNeighHSDN_CellList_r17 `optional`
}

func (ie *CarrierFreqEUTRA_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.eutra_FreqNeighHSDN_CellList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.eutra_FreqNeighHSDN_CellList_r17 != nil {
		if err = ie.eutra_FreqNeighHSDN_CellList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_FreqNeighHSDN_CellList_r17", err)
		}
	}
	return nil
}

func (ie *CarrierFreqEUTRA_v1700) Decode(r *uper.UperReader) error {
	var err error
	var eutra_FreqNeighHSDN_CellList_r17Present bool
	if eutra_FreqNeighHSDN_CellList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if eutra_FreqNeighHSDN_CellList_r17Present {
		ie.eutra_FreqNeighHSDN_CellList_r17 = new(EUTRA_FreqNeighHSDN_CellList_r17)
		if err = ie.eutra_FreqNeighHSDN_CellList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_FreqNeighHSDN_CellList_r17", err)
		}
	}
	return nil
}
