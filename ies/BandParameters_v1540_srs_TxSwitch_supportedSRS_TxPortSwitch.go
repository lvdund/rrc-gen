package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch_Enum_t1r2         uper.Enumerated = 0
	BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch_Enum_t1r4         uper.Enumerated = 1
	BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch_Enum_t2r4         uper.Enumerated = 2
	BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch_Enum_t1r4_t2r4    uper.Enumerated = 3
	BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch_Enum_t1r1         uper.Enumerated = 4
	BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch_Enum_t2r2         uper.Enumerated = 5
	BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch_Enum_t4r4         uper.Enumerated = 6
	BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch_Enum_notSupported uper.Enumerated = 7
)

type BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch", err)
	}
	return nil
}

func (ie *BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
