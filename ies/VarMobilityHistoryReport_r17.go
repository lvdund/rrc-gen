package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarMobilityHistoryReport_r17 struct {
	visitedCellInfoList_r16         VisitedCellInfoList_r16    `madatory`
	visitedPSCellInfoListReport_r17 *VisitedPSCellInfoList_r17 `optional`
}

func (ie *VarMobilityHistoryReport_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.visitedPSCellInfoListReport_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.visitedCellInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode visitedCellInfoList_r16", err)
	}
	if ie.visitedPSCellInfoListReport_r17 != nil {
		if err = ie.visitedPSCellInfoListReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode visitedPSCellInfoListReport_r17", err)
		}
	}
	return nil
}

func (ie *VarMobilityHistoryReport_r17) Decode(r *uper.UperReader) error {
	var err error
	var visitedPSCellInfoListReport_r17Present bool
	if visitedPSCellInfoListReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.visitedCellInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode visitedCellInfoList_r16", err)
	}
	if visitedPSCellInfoListReport_r17Present {
		ie.visitedPSCellInfoListReport_r17 = new(VisitedPSCellInfoList_r17)
		if err = ie.visitedPSCellInfoListReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode visitedPSCellInfoListReport_r17", err)
		}
	}
	return nil
}
