package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config_drb struct {
	discardTimer         *PDCP_Config_drb_discardTimer         `optional`
	pdcp_SN_SizeUL       *PDCP_Config_drb_pdcp_SN_SizeUL       `optional`
	pdcp_SN_SizeDL       *PDCP_Config_drb_pdcp_SN_SizeDL       `optional`
	headerCompression    *PDCP_Config_drb_headerCompression    `lb:1,ub:16383,optional`
	integrityProtection  *PDCP_Config_drb_integrityProtection  `optional`
	statusReportRequired *PDCP_Config_drb_statusReportRequired `optional`
	outOfOrderDelivery   *PDCP_Config_drb_outOfOrderDelivery   `optional`
}

func (ie *PDCP_Config_drb) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.discardTimer != nil, ie.pdcp_SN_SizeUL != nil, ie.pdcp_SN_SizeDL != nil, ie.headerCompression != nil, ie.integrityProtection != nil, ie.statusReportRequired != nil, ie.outOfOrderDelivery != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.discardTimer != nil {
		if err = ie.discardTimer.Encode(w); err != nil {
			return utils.WrapError("Encode discardTimer", err)
		}
	}
	if ie.pdcp_SN_SizeUL != nil {
		if err = ie.pdcp_SN_SizeUL.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_SN_SizeUL", err)
		}
	}
	if ie.pdcp_SN_SizeDL != nil {
		if err = ie.pdcp_SN_SizeDL.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_SN_SizeDL", err)
		}
	}
	if ie.headerCompression != nil {
		if err = ie.headerCompression.Encode(w); err != nil {
			return utils.WrapError("Encode headerCompression", err)
		}
	}
	if ie.integrityProtection != nil {
		if err = ie.integrityProtection.Encode(w); err != nil {
			return utils.WrapError("Encode integrityProtection", err)
		}
	}
	if ie.statusReportRequired != nil {
		if err = ie.statusReportRequired.Encode(w); err != nil {
			return utils.WrapError("Encode statusReportRequired", err)
		}
	}
	if ie.outOfOrderDelivery != nil {
		if err = ie.outOfOrderDelivery.Encode(w); err != nil {
			return utils.WrapError("Encode outOfOrderDelivery", err)
		}
	}
	return nil
}

func (ie *PDCP_Config_drb) Decode(r *uper.UperReader) error {
	var err error
	var discardTimerPresent bool
	if discardTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcp_SN_SizeULPresent bool
	if pdcp_SN_SizeULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcp_SN_SizeDLPresent bool
	if pdcp_SN_SizeDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var headerCompressionPresent bool
	if headerCompressionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var integrityProtectionPresent bool
	if integrityProtectionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var statusReportRequiredPresent bool
	if statusReportRequiredPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var outOfOrderDeliveryPresent bool
	if outOfOrderDeliveryPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if discardTimerPresent {
		ie.discardTimer = new(PDCP_Config_drb_discardTimer)
		if err = ie.discardTimer.Decode(r); err != nil {
			return utils.WrapError("Decode discardTimer", err)
		}
	}
	if pdcp_SN_SizeULPresent {
		ie.pdcp_SN_SizeUL = new(PDCP_Config_drb_pdcp_SN_SizeUL)
		if err = ie.pdcp_SN_SizeUL.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_SN_SizeUL", err)
		}
	}
	if pdcp_SN_SizeDLPresent {
		ie.pdcp_SN_SizeDL = new(PDCP_Config_drb_pdcp_SN_SizeDL)
		if err = ie.pdcp_SN_SizeDL.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_SN_SizeDL", err)
		}
	}
	if headerCompressionPresent {
		ie.headerCompression = new(PDCP_Config_drb_headerCompression)
		if err = ie.headerCompression.Decode(r); err != nil {
			return utils.WrapError("Decode headerCompression", err)
		}
	}
	if integrityProtectionPresent {
		ie.integrityProtection = new(PDCP_Config_drb_integrityProtection)
		if err = ie.integrityProtection.Decode(r); err != nil {
			return utils.WrapError("Decode integrityProtection", err)
		}
	}
	if statusReportRequiredPresent {
		ie.statusReportRequired = new(PDCP_Config_drb_statusReportRequired)
		if err = ie.statusReportRequired.Decode(r); err != nil {
			return utils.WrapError("Decode statusReportRequired", err)
		}
	}
	if outOfOrderDeliveryPresent {
		ie.outOfOrderDelivery = new(PDCP_Config_drb_outOfOrderDelivery)
		if err = ie.outOfOrderDelivery.Decode(r); err != nil {
			return utils.WrapError("Decode outOfOrderDelivery", err)
		}
	}
	return nil
}
