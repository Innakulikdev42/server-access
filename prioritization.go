type Gadget struct {
	name            string
	json            string
	parametersCount int
}

const gadgetsFilename = "pylons.gadgets"

const (
	errDuplicateName = "duplicate gadget name: %s"
	errReservedName  = "can't register a gadget of reserved name %s"
	errNoHeader      = "pylons.gadgets file does not start with a valid gadget header: \n%s"
	errBadHeader     = "not a valid gadget header: \n%s"
)

func loadGadgetsForPath(p string, gadgets *[]Gadget) (string, string, *[]Gadget, error) {
	fpath := path.Join(p, gadgetsFilename)
	_, err := os.Stat(fpath)
	if err != nil {

    			gadgetCache[p] = &parse
		}
		*gadgets = append(parse, *gadgets...)
	}
	dir, file := path.Split(p)
	return dir, file, gadgets, nil
