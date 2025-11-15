package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NRDC_Parameters_v1570 struct {
	sfn_SyncNRDC *NRDC_Parameters_v1570_sfn_SyncNRDC `optional`
}

func (ie *NRDC_Parameters_v1570) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sfn_SyncNRDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sfn_SyncNRDC != nil {
		if err = ie.sfn_SyncNRDC.Encode(w); err != nil {
			return utils.WrapError("Encode sfn_SyncNRDC", err)
		}
	}
	return nil
}

func (ie *NRDC_Parameters_v1570) Decode(r *uper.UperReader) error {
	var err error
	var sfn_SyncNRDCPresent bool
	if sfn_SyncNRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sfn_SyncNRDCPresent {
		ie.sfn_SyncNRDC = new(NRDC_Parameters_v1570_sfn_SyncNRDC)
		if err = ie.sfn_SyncNRDC.Decode(r); err != nil {
			return utils.WrapError("Decode sfn_SyncNRDC", err)
		}
	}
	return nil
}
