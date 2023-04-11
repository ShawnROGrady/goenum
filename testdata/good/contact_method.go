package contactmethod

type PreferredContactMethod int

const (
	PreferredContactMethodUnspecified PreferredContactMethod = iota
	ContactByEmail                                           //goenum:"name=EMAIL"
	ContactByCellphone                                       //goenum:"name=CELLPHONE"
	ContactByLandLine                                        //goenum:"name=LANDLINE"
)
