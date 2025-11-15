package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_nr struct {
	bandNR                 FreqBandIndicatorNR  `madatory`
	ca_BandwidthClassDL_NR *CA_BandwidthClassNR `optional`
	ca_BandwidthClassUL_NR *CA_BandwidthClassNR `optional`
}

func (ie *BandParameters_nr) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_BandwidthClassDL_NR != nil, ie.ca_BandwidthClassUL_NR != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.bandNR.Encode(w); err != nil {
		return utils.WrapError("Encode bandNR", err)
	}
	if ie.ca_BandwidthClassDL_NR != nil {
		if err = ie.ca_BandwidthClassDL_NR.Encode(w); err != nil {
			return utils.WrapError("Encode ca_BandwidthClassDL_NR", err)
		}
	}
	if ie.ca_BandwidthClassUL_NR != nil {
		if err = ie.ca_BandwidthClassUL_NR.Encode(w); err != nil {
			return utils.WrapError("Encode ca_BandwidthClassUL_NR", err)
		}
	}
	return nil
}

func (ie *BandParameters_nr) Decode(r *uper.UperReader) error {
	var err error
	var ca_BandwidthClassDL_NRPresent bool
	if ca_BandwidthClassDL_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_BandwidthClassUL_NRPresent bool
	if ca_BandwidthClassUL_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.bandNR.Decode(r); err != nil {
		return utils.WrapError("Decode bandNR", err)
	}
	if ca_BandwidthClassDL_NRPresent {
		ie.ca_BandwidthClassDL_NR = new(CA_BandwidthClassNR)
		if err = ie.ca_BandwidthClassDL_NR.Decode(r); err != nil {
			return utils.WrapError("Decode ca_BandwidthClassDL_NR", err)
		}
	}
	if ca_BandwidthClassUL_NRPresent {
		ie.ca_BandwidthClassUL_NR = new(CA_BandwidthClassNR)
		if err = ie.ca_BandwidthClassUL_NR.Decode(r); err != nil {
			return utils.WrapError("Decode ca_BandwidthClassUL_NR", err)
		}
	}
	return nil
}
