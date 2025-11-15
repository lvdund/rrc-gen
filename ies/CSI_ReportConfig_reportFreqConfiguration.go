package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfig_reportFreqConfiguration struct {
	cqi_FormatIndicator *CSI_ReportConfig_reportFreqConfiguration_cqi_FormatIndicator `optional`
	pmi_FormatIndicator *CSI_ReportConfig_reportFreqConfiguration_pmi_FormatIndicator `optional`
	csi_ReportingBand   *CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand   `lb:3,ub:3,optional`
}

func (ie *CSI_ReportConfig_reportFreqConfiguration) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cqi_FormatIndicator != nil, ie.pmi_FormatIndicator != nil, ie.csi_ReportingBand != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cqi_FormatIndicator != nil {
		if err = ie.cqi_FormatIndicator.Encode(w); err != nil {
			return utils.WrapError("Encode cqi_FormatIndicator", err)
		}
	}
	if ie.pmi_FormatIndicator != nil {
		if err = ie.pmi_FormatIndicator.Encode(w); err != nil {
			return utils.WrapError("Encode pmi_FormatIndicator", err)
		}
	}
	if ie.csi_ReportingBand != nil {
		if err = ie.csi_ReportingBand.Encode(w); err != nil {
			return utils.WrapError("Encode csi_ReportingBand", err)
		}
	}
	return nil
}

func (ie *CSI_ReportConfig_reportFreqConfiguration) Decode(r *uper.UperReader) error {
	var err error
	var cqi_FormatIndicatorPresent bool
	if cqi_FormatIndicatorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pmi_FormatIndicatorPresent bool
	if pmi_FormatIndicatorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_ReportingBandPresent bool
	if csi_ReportingBandPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if cqi_FormatIndicatorPresent {
		ie.cqi_FormatIndicator = new(CSI_ReportConfig_reportFreqConfiguration_cqi_FormatIndicator)
		if err = ie.cqi_FormatIndicator.Decode(r); err != nil {
			return utils.WrapError("Decode cqi_FormatIndicator", err)
		}
	}
	if pmi_FormatIndicatorPresent {
		ie.pmi_FormatIndicator = new(CSI_ReportConfig_reportFreqConfiguration_pmi_FormatIndicator)
		if err = ie.pmi_FormatIndicator.Decode(r); err != nil {
			return utils.WrapError("Decode pmi_FormatIndicator", err)
		}
	}
	if csi_ReportingBandPresent {
		ie.csi_ReportingBand = new(CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand)
		if err = ie.csi_ReportingBand.Decode(r); err != nil {
			return utils.WrapError("Decode csi_ReportingBand", err)
		}
	}
	return nil
}
