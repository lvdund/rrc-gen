package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RxTxReportInterval_r17_Enum_ms80   uper.Enumerated = 0
	RxTxReportInterval_r17_Enum_ms120  uper.Enumerated = 1
	RxTxReportInterval_r17_Enum_ms160  uper.Enumerated = 2
	RxTxReportInterval_r17_Enum_ms240  uper.Enumerated = 3
	RxTxReportInterval_r17_Enum_ms320  uper.Enumerated = 4
	RxTxReportInterval_r17_Enum_ms480  uper.Enumerated = 5
	RxTxReportInterval_r17_Enum_ms640  uper.Enumerated = 6
	RxTxReportInterval_r17_Enum_ms1024 uper.Enumerated = 7
	RxTxReportInterval_r17_Enum_ms1280 uper.Enumerated = 8
	RxTxReportInterval_r17_Enum_ms2048 uper.Enumerated = 9
	RxTxReportInterval_r17_Enum_ms2560 uper.Enumerated = 10
	RxTxReportInterval_r17_Enum_ms5120 uper.Enumerated = 11
	RxTxReportInterval_r17_Enum_spare4 uper.Enumerated = 12
	RxTxReportInterval_r17_Enum_spare3 uper.Enumerated = 13
	RxTxReportInterval_r17_Enum_spare2 uper.Enumerated = 14
	RxTxReportInterval_r17_Enum_spare1 uper.Enumerated = 15
)

type RxTxReportInterval_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RxTxReportInterval_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RxTxReportInterval_r17", err)
	}
	return nil
}

func (ie *RxTxReportInterval_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RxTxReportInterval_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
