package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Q_OffsetRangeList struct {
	rsrpOffsetSSB    Q_OffsetRange `madatory`
	rsrqOffsetSSB    Q_OffsetRange `madatory`
	sinrOffsetSSB    Q_OffsetRange `madatory`
	rsrpOffsetCSI_RS Q_OffsetRange `madatory`
	rsrqOffsetCSI_RS Q_OffsetRange `madatory`
	sinrOffsetCSI_RS Q_OffsetRange `madatory`
}

func (ie *Q_OffsetRangeList) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rsrpOffsetSSB.Encode(w); err != nil {
		return utils.WrapError("Encode rsrpOffsetSSB", err)
	}
	if err = ie.rsrqOffsetSSB.Encode(w); err != nil {
		return utils.WrapError("Encode rsrqOffsetSSB", err)
	}
	if err = ie.sinrOffsetSSB.Encode(w); err != nil {
		return utils.WrapError("Encode sinrOffsetSSB", err)
	}
	if err = ie.rsrpOffsetCSI_RS.Encode(w); err != nil {
		return utils.WrapError("Encode rsrpOffsetCSI_RS", err)
	}
	if err = ie.rsrqOffsetCSI_RS.Encode(w); err != nil {
		return utils.WrapError("Encode rsrqOffsetCSI_RS", err)
	}
	if err = ie.sinrOffsetCSI_RS.Encode(w); err != nil {
		return utils.WrapError("Encode sinrOffsetCSI_RS", err)
	}
	return nil
}

func (ie *Q_OffsetRangeList) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rsrpOffsetSSB.Decode(r); err != nil {
		return utils.WrapError("Decode rsrpOffsetSSB", err)
	}
	if err = ie.rsrqOffsetSSB.Decode(r); err != nil {
		return utils.WrapError("Decode rsrqOffsetSSB", err)
	}
	if err = ie.sinrOffsetSSB.Decode(r); err != nil {
		return utils.WrapError("Decode sinrOffsetSSB", err)
	}
	if err = ie.rsrpOffsetCSI_RS.Decode(r); err != nil {
		return utils.WrapError("Decode rsrpOffsetCSI_RS", err)
	}
	if err = ie.rsrqOffsetCSI_RS.Decode(r); err != nil {
		return utils.WrapError("Decode rsrqOffsetCSI_RS", err)
	}
	if err = ie.sinrOffsetCSI_RS.Decode(r); err != nil {
		return utils.WrapError("Decode sinrOffsetCSI_RS", err)
	}
	return nil
}
