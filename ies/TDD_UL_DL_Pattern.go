package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_Pattern struct {
	dl_UL_TransmissionPeriodicity       TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity        `madatory`
	nrofDownlinkSlots                   int64                                                  `lb:0,ub:maxNrofSlots,madatory`
	nrofDownlinkSymbols                 int64                                                  `lb:0,ub:maxNrofSymbols_1,madatory`
	nrofUplinkSlots                     int64                                                  `lb:0,ub:maxNrofSlots,madatory`
	nrofUplinkSymbols                   int64                                                  `lb:0,ub:maxNrofSymbols_1,madatory`
	dl_UL_TransmissionPeriodicity_v1530 *TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530 `optional,ext-1`
}

func (ie *TDD_UL_DL_Pattern) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.dl_UL_TransmissionPeriodicity_v1530 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.dl_UL_TransmissionPeriodicity.Encode(w); err != nil {
		return utils.WrapError("Encode dl_UL_TransmissionPeriodicity", err)
	}
	if err = w.WriteInteger(ie.nrofDownlinkSlots, &uper.Constraint{Lb: 0, Ub: maxNrofSlots}, false); err != nil {
		return utils.WrapError("WriteInteger nrofDownlinkSlots", err)
	}
	if err = w.WriteInteger(ie.nrofDownlinkSymbols, &uper.Constraint{Lb: 0, Ub: maxNrofSymbols_1}, false); err != nil {
		return utils.WrapError("WriteInteger nrofDownlinkSymbols", err)
	}
	if err = w.WriteInteger(ie.nrofUplinkSlots, &uper.Constraint{Lb: 0, Ub: maxNrofSlots}, false); err != nil {
		return utils.WrapError("WriteInteger nrofUplinkSlots", err)
	}
	if err = w.WriteInteger(ie.nrofUplinkSymbols, &uper.Constraint{Lb: 0, Ub: maxNrofSymbols_1}, false); err != nil {
		return utils.WrapError("WriteInteger nrofUplinkSymbols", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.dl_UL_TransmissionPeriodicity_v1530 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap TDD_UL_DL_Pattern", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.dl_UL_TransmissionPeriodicity_v1530 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode dl_UL_TransmissionPeriodicity_v1530 optional
			if ie.dl_UL_TransmissionPeriodicity_v1530 != nil {
				if err = ie.dl_UL_TransmissionPeriodicity_v1530.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_UL_TransmissionPeriodicity_v1530", err)
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

func (ie *TDD_UL_DL_Pattern) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.dl_UL_TransmissionPeriodicity.Decode(r); err != nil {
		return utils.WrapError("Decode dl_UL_TransmissionPeriodicity", err)
	}
	var tmp_int_nrofDownlinkSlots int64
	if tmp_int_nrofDownlinkSlots, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSlots}, false); err != nil {
		return utils.WrapError("ReadInteger nrofDownlinkSlots", err)
	}
	ie.nrofDownlinkSlots = tmp_int_nrofDownlinkSlots
	var tmp_int_nrofDownlinkSymbols int64
	if tmp_int_nrofDownlinkSymbols, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSymbols_1}, false); err != nil {
		return utils.WrapError("ReadInteger nrofDownlinkSymbols", err)
	}
	ie.nrofDownlinkSymbols = tmp_int_nrofDownlinkSymbols
	var tmp_int_nrofUplinkSlots int64
	if tmp_int_nrofUplinkSlots, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSlots}, false); err != nil {
		return utils.WrapError("ReadInteger nrofUplinkSlots", err)
	}
	ie.nrofUplinkSlots = tmp_int_nrofUplinkSlots
	var tmp_int_nrofUplinkSymbols int64
	if tmp_int_nrofUplinkSymbols, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSymbols_1}, false); err != nil {
		return utils.WrapError("ReadInteger nrofUplinkSymbols", err)
	}
	ie.nrofUplinkSymbols = tmp_int_nrofUplinkSymbols

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
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

			dl_UL_TransmissionPeriodicity_v1530Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode dl_UL_TransmissionPeriodicity_v1530 optional
			if dl_UL_TransmissionPeriodicity_v1530Present {
				ie.dl_UL_TransmissionPeriodicity_v1530 = new(TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530)
				if err = ie.dl_UL_TransmissionPeriodicity_v1530.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_UL_TransmissionPeriodicity_v1530", err)
				}
			}
		}
	}
	return nil
}
