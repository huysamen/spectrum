package spectrum

var (
	ReferenceACIE1931                          = [3]float64{1.09850, 1.00000, 0.35585}  //	02° (CIE 1931) - Incandescent/tungsten
	ReferenceACIE1964                          = [3]float64{1.11144, 1.00000, 0.35200}  //	10° (CIE 1964) - Incandescent/tungsten
	ReferenceBCIE1931                          = [3]float64{0.990927, 1.00000, 0.85313} //	02° (CIE 1931) - Old direct sunlight at noon
	ReferenceBCIE1964                          = [3]float64{0.99178, 1.00000, 0.843493} //	10° (CIE 1964) - Old direct sunlight at noon
	ReferenceCCIE1931                          = [3]float64{0.98074, 1.00000, 1.18232}  //	02° (CIE 1931) - Old daylight
	ReferenceCCIE1964                          = [3]float64{0.97285, 1.00000, 1.16145}  //	10° (CIE 1964) - Old daylight
	ReferenceD50CIE1931                        = [3]float64{0.96422, 1.00000, 0.82521}  //	02° (CIE 1931) - ICC profile PCS
	ReferenceD50CIE1964                        = [3]float64{0.96720, 1.00000, 0.81427}  //	10° (CIE 1964) - ICC profile PCS
	ReferenceD55CIE1931                        = [3]float64{0.95682, 1.00000, 0.92149}  //	02° (CIE 1931) - Mid-morning daylight
	ReferenceD55CIE1964                        = [3]float64{0.95799, 1.00000, 0.90926}  //	10° (CIE 1964) - Mid-morning daylight
	ReferenceD65CIE1931                        = [3]float64{0.95047, 1.00000, 1.08883}  //	02° (CIE 1931) - Daylight, sRGB, Adobe-RGB
	ReferenceD65CIE1964                        = [3]float64{0.94811, 1.00000, 1.07304}  //	10° (CIE 1964) - Daylight, sRGB, Adobe-RGB
	ReferenceD75CIE1931                        = [3]float64{0.94972, 1.00000, 1.22638}  //	02° (CIE 1931) - North sky daylight
	ReferenceD75CIE1964                        = [3]float64{0.94416, 1.00000, 1.20641}  //	10° (CIE 1964) - North sky daylight
	ReferenceECIE1931                          = [3]float64{1.00000, 1.00000, 1.00000}  //	02° (CIE 1931) - Equal energy
	ReferenceECIE1964                          = [3]float64{1.00000, 1.00000, 1.00000}  //	10° (CIE 1964) - Equal energy
	ReferenceF1CIE1931                         = [3]float64{0.92834, 1.00000, 1.03665}  //	02° (CIE 1931) - Daylight Fluorescent
	ReferenceF1CIE1964                         = [3]float64{0.94791, 1.00000, 1.03191}  //	10° (CIE 1964) - Daylight Fluorescent
	ReferenceF2CIE1931                         = [3]float64{0.99187, 1.00000, 0.67395}  //	02° (CIE 1931) - Cool fluorescent
	ReferenceF2CIE1964                         = [3]float64{1.03280, 1.00000, 0.69026}  //	10° (CIE 1964) - Cool fluorescent
	ReferenceF3CIE1931                         = [3]float64{1.03754, 1.00000, 0.49861}  //	02° (CIE 1931) - White Fluorescent
	ReferenceF3CIE1964                         = [3]float64{1.08968, 1.00000, 0.51965}  //	10° (CIE 1964) - White Fluorescent
	ReferenceF4CIE1931                         = [3]float64{1.09147, 1.00000, 0.38813}  //	02° (CIE 1931) - Warm White Fluorescent
	ReferenceF4CIE1964                         = [3]float64{1.14961, 1.00000, 0.40963}  //	10° (CIE 1964) - Warm White Fluorescent
	ReferenceF5CIE1931                         = [3]float64{0.90872, 1.00000, 0.98723}  //	02° (CIE 1931) - Daylight Fluorescent
	ReferenceF5CIE1964                         = [3]float64{0.93369, 1.00000, 0.98636}  //	10° (CIE 1964) - Daylight Fluorescent
	ReferenceF6CIE1931                         = [3]float64{0.97309, 1.00000, 0.60191}  //	02° (CIE 1931) - Lite White Fluorescent
	ReferenceF6CIE1964                         = [3]float64{1.02148, 1.00000, 0.62074}  //	10° (CIE 1964) - Lite White Fluorescent
	ReferenceF7CIE1931                         = [3]float64{0.95044, 1.00000, 1.08755}  //	02° (CIE 1931) - Daylight fluorescent, D65 simulator
	ReferenceF7CIE1964                         = [3]float64{0.95792, 1.00000, 1.07687}  //	10° (CIE 1964) - Daylight fluorescent, D65 simulator
	ReferenceF8CIE1931                         = [3]float64{0.96413, 1.00000, 0.82333}  //	02° (CIE 1931) - Sylvania F40, D50 simulator
	ReferenceF8CIE1964                         = [3]float64{0.97115, 1.00000, 0.81135}  //	10° (CIE 1964) - Sylvania F40, D50 simulator
	ReferenceF9CIE1931                         = [3]float64{1.00365, 1.00000, 0.67868}  //	02° (CIE 1931) - Cool White Fluorescent
	ReferenceF9CIE1964                         = [3]float64{1.02116, 1.00000, 0.67826}  //	10° (CIE 1964) - Cool White Fluorescent
	ReferenceF10CIE1931                        = [3]float64{0.96174, 1.00000, 0.81712}  //	02° (CIE 1931) - Ultralume 50, Philips TL85
	ReferenceF10CIE1964                        = [3]float64{0.99001, 1.00000, 0.83134}  //	10° (CIE 1964) - Ultralume 50, Philips TL85
	ReferenceF11CIE1931                        = [3]float64{1.00966, 1.00000, 0.64370}  //	02° (CIE 1931) - Ultralume 40, Philips TL84
	ReferenceF11CIE1964                        = [3]float64{1.03866, 1.00000, 0.65627}  //	10° (CIE 1964) - Ultralume 40, Philips TL84
	ReferenceF12CIE1931                        = [3]float64{1.08046, 1.00000, 0.39228}  //	02° (CIE 1931) - Ultralume 30, Philips TL83
	ReferenceF12CIE1964                        = [3]float64{1.11428, 1.00000, 0.40353}  //	10° (CIE 1964) - Ultralume 30, Philips TL83
	ReferenceIncandescentTungsten02            = ReferenceACIE1931
	ReferenceIncandescentTungsten10            = ReferenceACIE1964
	ReferenceOldDirectSunlightAtNoon02         = ReferenceBCIE1931
	ReferenceOldDirectSunlightAtNoon10         = ReferenceBCIE1964
	ReferenceOldDaylight02                     = ReferenceCCIE1931
	ReferenceOldDaylight10                     = ReferenceCCIE1964
	ReferenceICCProfilePCS02                   = ReferenceD50CIE1931
	ReferenceICCProfilePCS10                   = ReferenceD50CIE1964
	ReferenceMidMorningDaylight02              = ReferenceD55CIE1931
	ReferenceMidMorningDaylight10              = ReferenceD55CIE1964
	ReferenceDaylightRGB02                     = ReferenceD65CIE1931
	ReferenceDaylightRGB10                     = ReferenceD65CIE1964
	ReferenceNorthSkyDaylight02                = ReferenceD75CIE1931
	ReferenceNorthSkyDaylight10                = ReferenceD75CIE1964
	ReferenceEqualEnergy02                     = ReferenceECIE1931
	ReferenceEqualEnergy04                     = ReferenceECIE1964
	ReferenceDaylightFluorescent02             = ReferenceF1CIE1931
	ReferenceDaylightFluorescent10             = ReferenceF1CIE1964
	ReferenceCoolFluorescent02                 = ReferenceF2CIE1931
	ReferenceCoolFluorescent10                 = ReferenceF2CIE1964
	ReferenceWhiteFluorescent02                = ReferenceF3CIE1931
	ReferenceWhiteFluorescent10                = ReferenceF3CIE1964
	ReferenceWarmWhiteFluorescent02            = ReferenceF4CIE1931
	ReferenceWarmWhiteFluorescent10            = ReferenceF4CIE1964
	ReferenceDaylightWhiteFluorescent02        = ReferenceF5CIE1931
	ReferenceDaylightWhiteFluorescent10        = ReferenceF5CIE1964
	ReferenceLiteWhiteFluorescent02            = ReferenceF6CIE1931
	ReferenceLiteWhiteFluorescent10            = ReferenceF6CIE1964
	ReferenceDaylightFluorescentD65Simulator02 = ReferenceF7CIE1931
	ReferenceDaylightFluorescentD65Simulator10 = ReferenceF7CIE1964
	ReferenceSylvaniaF40D50Simulator02         = ReferenceF8CIE1931
	ReferenceSylvaniaF40D50Simulator10         = ReferenceF8CIE1964
	ReferenceCoolWhiteFluorescent02            = ReferenceF9CIE1931
	ReferenceCoolWhiteFluorescent10            = ReferenceF9CIE1964
	ReferenceUltralume50PhilipsTL8502          = ReferenceF10CIE1931
	ReferenceUltralume50PhilipsTL8510          = ReferenceF10CIE1964
	ReferenceUltralume40PhilipsTL8402          = ReferenceF11CIE1931
	ReferenceUltralume40PhilipsTL8410          = ReferenceF11CIE1964
	ReferenceUltralume30PhilipsTL8302          = ReferenceF12CIE1931
	ReferenceUltralume30PhilipsTL8310          = ReferenceF12CIE1964
)