package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResult2UTRA_FDD_r16 struct {
	carrierFreq_r16             ARFCN_ValueUTRA_FDD_r16    `madatory`
	measResultNeighCellList_r16 MeasResultListUTRA_FDD_r16 `madatory`
}

func (ie *MeasResult2UTRA_FDD_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.carrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq_r16", err)
	}
	if err = ie.measResultNeighCellList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResultNeighCellList_r16", err)
	}
	return nil
}

func (ie *MeasResult2UTRA_FDD_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.carrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq_r16", err)
	}
	if err = ie.measResultNeighCellList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode measResultNeighCellList_r16", err)
	}
	return nil
}
