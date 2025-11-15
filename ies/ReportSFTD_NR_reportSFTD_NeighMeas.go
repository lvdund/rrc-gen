package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReportSFTD_NR_reportSFTD_NeighMeas_Enum_true uper.Enumerated = 0
)

type ReportSFTD_NR_reportSFTD_NeighMeas struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *ReportSFTD_NR_reportSFTD_NeighMeas) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode ReportSFTD_NR_reportSFTD_NeighMeas", err)
	}
	return nil
}

func (ie *ReportSFTD_NR_reportSFTD_NeighMeas) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode ReportSFTD_NR_reportSFTD_NeighMeas", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
