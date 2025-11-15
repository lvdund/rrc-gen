package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersXDD_Diff struct {
	intraAndInterF_MeasAndReport *MeasAndMobParametersXDD_Diff_intraAndInterF_MeasAndReport `optional`
	eventA_MeasAndReport         *MeasAndMobParametersXDD_Diff_eventA_MeasAndReport         `optional`
	handoverInterF               *MeasAndMobParametersXDD_Diff_handoverInterF               `optional,ext-1`
	handoverLTE_EPC              *MeasAndMobParametersXDD_Diff_handoverLTE_EPC              `optional,ext-1`
	handoverLTE_5GC              *MeasAndMobParametersXDD_Diff_handoverLTE_5GC              `optional,ext-1`
	sftd_MeasNR_Neigh            *MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh            `optional,ext-2`
	sftd_MeasNR_Neigh_DRX        *MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX        `optional,ext-2`
	dummy                        *MeasAndMobParametersXDD_Diff_dummy                        `optional,ext-3`
}

func (ie *MeasAndMobParametersXDD_Diff) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.handoverInterF != nil || ie.handoverLTE_EPC != nil || ie.handoverLTE_5GC != nil || ie.sftd_MeasNR_Neigh != nil || ie.sftd_MeasNR_Neigh_DRX != nil || ie.dummy != nil
	preambleBits := []bool{hasExtensions, ie.intraAndInterF_MeasAndReport != nil, ie.eventA_MeasAndReport != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.intraAndInterF_MeasAndReport != nil {
		if err = ie.intraAndInterF_MeasAndReport.Encode(w); err != nil {
			return utils.WrapError("Encode intraAndInterF_MeasAndReport", err)
		}
	}
	if ie.eventA_MeasAndReport != nil {
		if err = ie.eventA_MeasAndReport.Encode(w); err != nil {
			return utils.WrapError("Encode eventA_MeasAndReport", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.handoverInterF != nil || ie.handoverLTE_EPC != nil || ie.handoverLTE_5GC != nil, ie.sftd_MeasNR_Neigh != nil || ie.sftd_MeasNR_Neigh_DRX != nil, ie.dummy != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasAndMobParametersXDD_Diff", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.handoverInterF != nil, ie.handoverLTE_EPC != nil, ie.handoverLTE_5GC != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode handoverInterF optional
			if ie.handoverInterF != nil {
				if err = ie.handoverInterF.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverInterF", err)
				}
			}
			// encode handoverLTE_EPC optional
			if ie.handoverLTE_EPC != nil {
				if err = ie.handoverLTE_EPC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverLTE_EPC", err)
				}
			}
			// encode handoverLTE_5GC optional
			if ie.handoverLTE_5GC != nil {
				if err = ie.handoverLTE_5GC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverLTE_5GC", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.sftd_MeasNR_Neigh != nil, ie.sftd_MeasNR_Neigh_DRX != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sftd_MeasNR_Neigh optional
			if ie.sftd_MeasNR_Neigh != nil {
				if err = ie.sftd_MeasNR_Neigh.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sftd_MeasNR_Neigh", err)
				}
			}
			// encode sftd_MeasNR_Neigh_DRX optional
			if ie.sftd_MeasNR_Neigh_DRX != nil {
				if err = ie.sftd_MeasNR_Neigh_DRX.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sftd_MeasNR_Neigh_DRX", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.dummy != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode dummy optional
			if ie.dummy != nil {
				if err = ie.dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dummy", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *MeasAndMobParametersXDD_Diff) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var intraAndInterF_MeasAndReportPresent bool
	if intraAndInterF_MeasAndReportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var eventA_MeasAndReportPresent bool
	if eventA_MeasAndReportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if intraAndInterF_MeasAndReportPresent {
		ie.intraAndInterF_MeasAndReport = new(MeasAndMobParametersXDD_Diff_intraAndInterF_MeasAndReport)
		if err = ie.intraAndInterF_MeasAndReport.Decode(r); err != nil {
			return utils.WrapError("Decode intraAndInterF_MeasAndReport", err)
		}
	}
	if eventA_MeasAndReportPresent {
		ie.eventA_MeasAndReport = new(MeasAndMobParametersXDD_Diff_eventA_MeasAndReport)
		if err = ie.eventA_MeasAndReport.Decode(r); err != nil {
			return utils.WrapError("Decode eventA_MeasAndReport", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			handoverInterFPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			handoverLTE_EPCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			handoverLTE_5GCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode handoverInterF optional
			if handoverInterFPresent {
				ie.handoverInterF = new(MeasAndMobParametersXDD_Diff_handoverInterF)
				if err = ie.handoverInterF.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverInterF", err)
				}
			}
			// decode handoverLTE_EPC optional
			if handoverLTE_EPCPresent {
				ie.handoverLTE_EPC = new(MeasAndMobParametersXDD_Diff_handoverLTE_EPC)
				if err = ie.handoverLTE_EPC.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverLTE_EPC", err)
				}
			}
			// decode handoverLTE_5GC optional
			if handoverLTE_5GCPresent {
				ie.handoverLTE_5GC = new(MeasAndMobParametersXDD_Diff_handoverLTE_5GC)
				if err = ie.handoverLTE_5GC.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverLTE_5GC", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sftd_MeasNR_NeighPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sftd_MeasNR_Neigh_DRXPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sftd_MeasNR_Neigh optional
			if sftd_MeasNR_NeighPresent {
				ie.sftd_MeasNR_Neigh = new(MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh)
				if err = ie.sftd_MeasNR_Neigh.Decode(extReader); err != nil {
					return utils.WrapError("Decode sftd_MeasNR_Neigh", err)
				}
			}
			// decode sftd_MeasNR_Neigh_DRX optional
			if sftd_MeasNR_Neigh_DRXPresent {
				ie.sftd_MeasNR_Neigh_DRX = new(MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX)
				if err = ie.sftd_MeasNR_Neigh_DRX.Decode(extReader); err != nil {
					return utils.WrapError("Decode sftd_MeasNR_Neigh_DRX", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			dummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode dummy optional
			if dummyPresent {
				ie.dummy = new(MeasAndMobParametersXDD_Diff_dummy)
				if err = ie.dummy.Decode(extReader); err != nil {
					return utils.WrapError("Decode dummy", err)
				}
			}
		}
	}
	return nil
}
