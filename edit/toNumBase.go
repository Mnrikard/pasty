package edit

import "strconv"

func (e *EditorArgs) ToNumBase(input string) (string, error) {
	base := e.getBase()

	//parse as base10 from input
	output, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return "", err
	}
	//output as baseX
	return strconv.FormatInt(output, int(base)), nil
}

func (e *EditorArgs) FromNumBase(input string) (string, error) {
	base := e.getBase()

	//parse the baseX string to an int
	output, err := strconv.ParseInt(input, base, 64)
	if err != nil {
		return "", err
	}
	//output as base10
	return strconv.FormatInt(output, 10), nil
}

func (e *EditorArgs) getBase() int {
	base, err := strconv.ParseInt(e.Option, 10, 0)

	if(e.Option == "0x") {
		base = 16
	} else if( err != nil || base == 0){
		base = 8;
	}

	return int(base)
}
