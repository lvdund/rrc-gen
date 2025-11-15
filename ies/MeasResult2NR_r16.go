package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResult2NR_r16 struct {
	ssbFrequency_r16   *ARFCN_ValueNR   `optional`
	refFreqCSI_RS_r16  *ARFCN_ValueNR   `optional`
	measResultList_r16 MeasResultListNR `madatory`
}

func (ie *MeasResult2NR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssbFrequency_r16 != nil, ie.refFreqCSI_RS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ssbFrequency_r16 != nil {
		if err = ie.ssbFrequency_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssbFrequency_r16", err)
		}
	}
	if ie.refFreqCSI_RS_r16 != nil {
		if err = ie.refFreqCSI_RS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode refFreqCSI_RS_r16", err)
		}
	}
	if err = ie.measResultList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResultList_r16", err)
	}
	return nil
}

func (ie *MeasResult2NR_r16) Decode(r *uper.UperReader) error {
	var err error
	var ssbFrequency_r16Present bool
	if ssbFrequency_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var refFreqCSI_RS_r16Present bool
	if refFreqCSI_RS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ssbFrequency_r16Present {
		ie.ssbFrequency_r16 = new(ARFCN_ValueNR)
		if err = ie.ssbFrequency_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssbFrequency_r16", err)
		}
	}
	if refFreqCSI_RS_r16Present {
		ie.refFreqCSI_RS_r16 = new(ARFCN_ValueNR)
		if err = ie.refFreqCSI_RS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode refFreqCSI_RS_r16", err)
		}
	}
	if err = ie.measResultList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode measResultList_r16", err)
	}
	return nil
}
